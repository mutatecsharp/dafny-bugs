// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm0(p0, p1, p2, p3, globalState) {
      return (((true) ? (_dafny.Set.fromElements(true, false)) : (_dafny.Set.fromElements(false)))).Union(_dafny.Set.fromElements(!(true), true, false));
    };
    static fm3(globalState) {
      return !(!(((new BigNumber((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(true, false))).cardinality()),new BigNumber((_dafny.Seq.of(!(false))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(679), new BigNumber(-336))).cardinality()),!(false))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(793),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(8))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, false))).cardinality()),new BigNumber(76)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),new BigNumber((_dafny.Seq.of(false)).length)))).Elements) {
          let _0_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(true, false))).cardinality()),new BigNumber((_dafny.Seq.of(!(false))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(679), new BigNumber(-336))).cardinality()),!(false))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(793),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(8))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, false))).cardinality()),new BigNumber(76)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),new BigNumber((_dafny.Seq.of(false)).length))), _0_v0)) {
            _coll0.add(_0_v0);
          }
        }
        return _coll0;
      }()).length)).plus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true)).length))).cardinality()))).isEqualTo(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-574), new BigNumber((_dafny.Seq.of(false, !(true))).length), new BigNumber(-776))).cardinality()))));
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return _dafny.MultiSet.fromElements((new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())).isLessThan(new BigNumber(402)), true, !((_dafny.MultiSet.fromElements(!(!(false)), true, true)).contains(false)), true, false);
    };
    static fm7(p0, p1, p2, globalState) {
      return true;
    };
    static fm8(p0, p1, p2, globalState) {
      return _module.__default.safeModuloInt(new BigNumber(544), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("agf")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("ufyke")).length)))).length));
    };
    static fm10(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(new BigNumber(324), new BigNumber(155));
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return _module.D2.create_DC7(false);
    };
    static fm12(p0, p1, p2, globalState) {
      if (false) {
        return _module.D3.create_DC12(_module.D3.create_DC12(_module.D3.create_DC12(_module.D3.create_DC12(_module.D3.create_DC11(false, new BigNumber(707), _dafny.Set.fromElements(false))))));
      } else {
        return _module.D3.create_DC12(_module.D3.create_DC11(false, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-504)), function (_1_i0) {
  return new _dafny.CodePoint('h'.codePointAt(0));
})).length), _dafny.Set.fromElements(false)));
      }
    };
    static fm13(globalState) {
      if ((!(false)) || (false)) {
        if (true) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('x'.codePointAt(0));
        }
      } else if (!(!(true))) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('o'.codePointAt(0));
      }
    };
    static fm14(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.UnicodeFromString("nxqveefwj");
    };
    static fm15(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(!(true)))).length),new BigNumber((_dafny.Seq.of(true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("xw")).length),new BigNumber((_dafny.Set.fromElements(false)).length)));
    };
    static fm16(globalState) {
      return _dafny.Seq.of(_module.D1.create_DC5(new BigNumber(710)), _module.D1.create_DC5(new BigNumber(611)), _module.D1.create_DC5(new BigNumber(-796)));
    };
    static fm17(p0, globalState) {
      if ((_dafny.Set.fromElements(false, true, !(false))).contains(false)) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),true));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),true);
      }
    };
    static fm18(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false, false, true, false), _dafny.Seq.Concat(_dafny.Seq.of(true, (_module.D12.create_DC34(!(true))).dtor_cf65, true), _dafny.Seq.of(false)));
    };
    static fm21(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC2((new BigNumber(687)).plus(new BigNumber((_dafny.Seq.of(true)).length)), (new BigNumber(818)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-321)), function (_2_i0) {
  return _dafny.Seq.UnicodeFromString("qyetl");
})).length)), ((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(614)), function (_3_i1) {
  return new BigNumber(8);
})) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(818)), function (_4_i2) {
  return new BigNumber(-302);
}))));
    };
    static fm22(globalState) {
      return (function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (function () {
          let _coll2 = new _dafny.Map();
          for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,true),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-951)), function (_5_i0) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          })).length),new BigNumber(168)))).length))).Keys.Elements) {
            let _6_v0 = _compr_2;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,true),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-951)), function (_5_i0) {
              return new _dafny.CodePoint('v'.codePointAt(0));
            })).length),new BigNumber(168)))).length))).contains(_6_v0)) {
              _coll2.push([_6_v0,!(true)]);
            }
          }
          return _coll2;
        }()).Keys.Elements) {
          let _7_v1 = _compr_1;
          if ((function () {
            let _coll3 = new _dafny.Map();
            for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,true),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-951)), function (_5_i0) {
              return new _dafny.CodePoint('v'.codePointAt(0));
            })).length),new BigNumber(168)))).length))).Keys.Elements) {
              let _6_v0 = _compr_3;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,true),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-951)), function (_5_i0) {
                return new _dafny.CodePoint('v'.codePointAt(0));
              })).length),new BigNumber(168)))).length))).contains(_6_v0)) {
                _coll3.push([_6_v0,!(true)]);
              }
            }
            return _coll3;
          }()).contains(_7_v1)) {
            _coll1.add(_7_v1);
          }
        }
        return _coll1;
      }()).Union(((true) ? (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,true))) : (function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(true),true),(_dafny.ZERO).minus(new BigNumber(-400)))).Keys.Elements) {
          let _8_v2 = _compr_4;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(true),true),(_dafny.ZERO).minus(new BigNumber(-400)))).contains(_8_v2)) {
            _coll4.add(_8_v2);
          }
        }
        return _coll4;
      }())));
    };
    static fm23(p0, p1, p2, p3, globalState) {
      return (((true) ? (_dafny.Map.Empty.slice().updateUnsafe(true,false)) : (_dafny.Map.Empty.slice().updateUnsafe(false,false)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),true)));
    };
    static fm24(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(!(false),(_dafny.ZERO).minus(new BigNumber(-285)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(!(false))),new BigNumber(-46)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(289))).cardinality()))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-844))));
    };
    static fm25(globalState) {
      let _source0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(381)), function (_9_i0) {
        return _module.D1.create_DC5(new BigNumber(928));
      });
      let _10___mcc_h0 = _source0;
      let _11_cf24 = _10___mcc_h0;
      return _module.D1.create_DC5(new BigNumber(752));
    };
    static fm26(p0, p1, p2, globalState) {
      return _module.D1.create_DC4(new _dafny.CodePoint('a'.codePointAt(0)), (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.of(true)).length)).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("bfmrgm")).length)))), new _dafny.CodePoint('q'.codePointAt(0)), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length))).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-71)), function (_12_i0) {
  return new _dafny.CodePoint('a'.codePointAt(0));
}))).length)).multipliedBy(new BigNumber(961)), (_dafny.MultiSet.fromElements(!(true), true)).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true))));
    };
    static fm27(p0, p1, p2, globalState) {
      return (new BigNumber((_dafny.Seq.UnicodeFromString("nxgruga")).length)).isEqualTo(new BigNumber((_dafny.Seq.of(_module.D14.create_DC40((_dafny.ZERO).minus(new BigNumber(-75)), true, _dafny.Seq.UnicodeFromString("dnsuin")))).length));
    };
    static fm28(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.UnicodeFromString("pbthlure");
    };
    static fm29(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of((_module.D0.create_DC2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_module.D12.create_DC33(false, _dafny.Seq.UnicodeFromString("yuplewtm")))).length), true, _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("bensvahfw")).length)))).dtor_cf5)), _dafny.Seq.of(true, true, false));
    };
    static fm30(p0, globalState) {
      return ((_module.D16.create_DC45(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())), new BigNumber(292)))).dtor_cf80).Difference(function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of _dafny.IntegerRange(new BigNumber(363), new BigNumber(186))) {
          let _13_v0 = _compr_5;
          if (((new BigNumber(363)).isLessThanOrEqualTo(_13_v0)) && ((_13_v0).isLessThan(new BigNumber(186)))) {
            _coll5.add((_13_v0).multipliedBy(new BigNumber(-949)));
          }
        }
        return _coll5;
      }());
    };
    static fm31(globalState) {
      return _module.D6.create_DC17(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(992)), function (_14_i0) {
  return new _dafny.CodePoint('d'.codePointAt(0));
}), _dafny.Seq.UnicodeFromString("eanbpr")), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(156)), function (_15_i1) {
  return new _dafny.CodePoint('s'.codePointAt(0));
}), _dafny.Seq.UnicodeFromString("xm")));
    };
    static fm32(p0, p1, globalState) {
      return (((false) ? (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(!(false))).length))).cardinality()))).length)),_dafny.Seq.UnicodeFromString("sorttx"))) : (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.Create(_module.__default.abs(new BigNumber(599)), function (_16_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      }))).length), new BigNumber((_dafny.Seq.UnicodeFromString("rkbghw")).length))).length), new BigNumber(559)),_dafny.Seq.UnicodeFromString("yhnnprjd"))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(!(true), true)).length), new BigNumber(67)),_dafny.Seq.UnicodeFromString("ug")));
    };
    static fm33(globalState) {
      return _dafny.MultiSet.fromElements(((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).minus(new BigNumber(77)));
    };
    static fm34(p0, p1, p2, globalState) {
      if (!(_dafny.Set.fromElements(new BigNumber(919), new BigNumber((_dafny.Set.fromElements(!(false))).length), new BigNumber((_dafny.Seq.of(true)).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("sswah")).length))))).contains(new BigNumber(788))) {
        return _dafny.Seq.of(_dafny.Seq.of(true, true, false));
      } else {
        return _dafny.Seq.of(_dafny.Seq.of(true, true));
      }
    };
    static m0(globalState) {
      let _17_v0;
      _17_v0 = true;
      if (_17_v0) {
        let _18_v1;
        _18_v1 = new BigNumber(146);
        (globalState).f1 = (_dafny.ZERO).minus(_18_v1);
        let _19_v2;
        _19_v2 = _dafny.MultiSet.fromElements(_18_v1);
        let _20_v3;
        let _nw0 = Array((new BigNumber(19)).toNumber());
        _20_v3 = _nw0;
        let _21_v4;
        _21_v4 = _dafny.Map.Empty.slice().updateUnsafe((_18_v1).multipliedBy(_18_v1),_20_v3);
        let _22_v5;
        _22_v5 = _dafny.Seq.of(_module.__default.fm13(globalState));
        let _23_v6;
        _23_v6 = new _dafny.CodePoint('f'.codePointAt(0));
        let _24_v7;
        let _nw1 = new _module.C4();
        _nw1.__ctor(new BigNumber((_22_v5).length), _17_v0, _17_v0, _17_v0, _23_v6, _22_v5);
        _24_v7 = _nw1;
        let _25_v8;
        _25_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_22_v5).length),_24_v7);
        let _rhs0 = _18_v1;
        let _rhs1 = (_19_v2).update(new BigNumber(82), _module.__default.abs(new BigNumber(397)));
        let _rhs2 = (_21_v4).Merge(_21_v4);
        let _rhs3 = (((_19_v2).contains(_module.__default.safeDivisionInt(new BigNumber((_22_v5).length), _18_v1))) ? ((_19_v2).get(_module.__default.safeDivisionInt(new BigNumber((_22_v5).length), _18_v1))) : (_18_v1));
        let _rhs4 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(new BigNumber((_25_v8).length))).length), _module.__default.safeModuloInt(_18_v1, _18_v1));
        let _lhs0 = globalState;
        let _lhs1 = globalState;
        _18_v1 = _rhs0;
        _19_v2 = _rhs1;
        _21_v4 = _rhs2;
        _lhs0.f1 = _rhs3;
        _lhs1.f1 = _rhs4;
        if (!(_17_v0)) {
          (globalState).f1 = (_18_v1).multipliedBy(_18_v1);
          let _26_v9;
          _26_v9 = _module.D6.create_DC16(_19_v2);
          let _27_v10;
          _27_v10 = _dafny.Seq.of(_26_v9, _26_v9, _26_v9);
          let _28_v11;
          _28_v11 = _dafny.Seq.of(_17_v0, true, _17_v0);
          _27_v10 = _dafny.Seq.Concat(_dafny.Seq.update(_27_v10, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_28_v11)).cardinality()), new BigNumber((_27_v10).length)), _26_v9), _27_v10);
          _22_v5 = ((_17_v0) ? ((_24_v7).f5) : ((_24_v7).f5));
          let _29_v12;
          let _nw2 = Array((new BigNumber(18)).toNumber()).fill(_dafny.MultiSet.Empty);
          _29_v12 = _nw2;
          let _30_v13;
          _30_v13 = _dafny.MultiSet.fromElements(_17_v0);
          let _index0 = _module.__default.safeIndex(new BigNumber(430), new BigNumber((_29_v12).length));
          (_29_v12)[_index0] = _30_v13;
          let _index1 = _module.__default.safeIndex(new BigNumber(430), new BigNumber((_29_v12).length));
          (_29_v12)[_index1] = _30_v13;
          let _31_v14;
          let _nw3 = new _module.C1();
          _nw3.__ctor(_18_v1, _17_v0, _17_v0, _17_v0, _23_v6, _22_v5);
          _31_v14 = _nw3;
          let _32_v15;
          _32_v15 = _dafny.Seq.of(_31_v14, _31_v14, _31_v14, _31_v14);
          let _33_v16;
          _33_v16 = _dafny.Map.Empty.slice().updateUnsafe(_32_v15,_18_v1);
          let _34_v17;
          _34_v17 = _dafny.Map.Empty.slice().updateUnsafe((((_33_v16).contains(_dafny.Seq.of(_31_v14))) ? ((_33_v16).get(_dafny.Seq.of(_31_v14))) : (_18_v1)),_18_v1);
          let _35_v18;
          _35_v18 = _dafny.Seq.of(_18_v1);
          _34_v17 = (_34_v17).update(_18_v1, new BigNumber((_35_v18).length));
        } else {
          let _36_v20;
          let _init0 = function (_37_i0) {
            return _module.__default.safeDivisionInt(_37_i0, new BigNumber((function () {
              let _coll6 = new _dafny.Set();
              for (const _compr_6 of _dafny.IntegerRange(new BigNumber(-674), new BigNumber(-208))) {
                let _38_v19 = _compr_6;
                if (((new BigNumber(-674)).isLessThanOrEqualTo(_38_v19)) && ((_38_v19).isLessThan(new BigNumber(-208)))) {
                  _coll6.add((_38_v19).plus(new BigNumber(382)));
                }
              }
              return _coll6;
            }()).length));
          };
          let _nw4 = Array((new BigNumber(18)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw4.length); _i0_0++) {
            _nw4[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _36_v20 = _nw4;
          let _index2 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_36_v20).length));
          (_36_v20)[_index2] = ((_17_v0) ? (_18_v1) : (_18_v1));
          let _39_v21;
          let _nw5 = Array((new BigNumber(8)).toNumber()).fill([]);
          _39_v21 = _nw5;
          let _40_v22;
          let _nw6 = Array((new BigNumber(17)).toNumber()).fill(false);
          _40_v22 = _nw6;
          let _index3 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_40_v22).length));
          (_40_v22)[_index3] = (_19_v2).equals(_19_v2);
          let _index4 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_36_v20).length));
          let _index5 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_40_v22).length));
          let _rhs5 = _24_v7.f4;
          let _rhs6 = new BigNumber(-977);
          let _rhs7 = ((((_17_v0) ? (_module.__default.fm7(_24_v7.f4, _18_v1, _17_v0, globalState)) : (_17_v0))) ? (_39_v21) : (_39_v21));
          let _rhs8 = !(false) || (!(!(_17_v0) || (_17_v0)));
          let _lhs2 = _36_v20;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_36_v20).length));
          let _lhs4 = _40_v22;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_40_v22).length));
          _23_v6 = _rhs5;
          _lhs2[_lhs3] = _rhs6;
          _39_v21 = _rhs7;
          _lhs4[_lhs5] = _rhs8;
          _23_v6 = _24_v7.f4;
          _17_v0 = _17_v0;
          let _41_v23;
          let _init1 = ((_42_v7) => function (_43_i1) {
            return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("itjlfvghe"),_42_v7.f4);
          })(_24_v7);
          let _nw7 = Array((new BigNumber(3)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw7.length); _i0_1++) {
            _nw7[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _41_v23 = _nw7;
          let _44_v24;
          _44_v24 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("dtgg"),_24_v7.f4);
          let _index6 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_41_v23).length));
          (_41_v23)[_index6] = (_44_v24).Merge(_44_v24);
          let _index7 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_41_v23).length));
          (_41_v23)[_index7] = _44_v24;
          _40_v22 = _40_v22;
        }
        _17_v0 = _module.__default.fm3(globalState);
        (globalState).f1 = (_18_v1).minus(_18_v1);
      } else {
        _17_v0 = _17_v0;
        let _45_v25;
        let _init2 = ((_46_v0) => function (_47_i2) {
          return _46_v0;
        })(_17_v0);
        let _nw8 = Array((new BigNumber(11)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw8.length); _i0_2++) {
          _nw8[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _45_v25 = _nw8;
        let _48_v26;
        let _init3 = ((_49_v0) => function (_50_i3) {
          return _dafny.Seq.of(_49_v0, _49_v0);
        })(_17_v0);
        let _nw9 = Array((new BigNumber(6)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw9.length); _i0_3++) {
          _nw9[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _48_v26 = _nw9;
        let _51_v27;
        _51_v27 = _module.D1.create_DC3(_48_v26);
        let _52_v28;
        _52_v28 = _dafny.Seq.of(_48_v26);
        let _53_v29;
        _53_v29 = new BigNumber(80);
        let _pat_let_tv0 = _48_v26;
        let _54_v30;
        let _nw10 = Array((new BigNumber(24)).toNumber());
        _nw10[(_dafny.ZERO).toNumber()] = _51_v27;
        _nw10[(_dafny.ONE).toNumber()] = _51_v27;
        _nw10[(new BigNumber(2)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(3)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(4)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(5)).toNumber()] = _module.D1.create_DC3((_52_v28)[_module.__default.safeIndex(_53_v29, new BigNumber((_52_v28).length))]);
        _nw10[(new BigNumber(6)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(7)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(8)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(9)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(10)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(11)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(12)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(13)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(14)).toNumber()] = _module.D1.create_DC3(_48_v26);
        _nw10[(new BigNumber(15)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(16)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(17)).toNumber()] = function (_pat_let0_0) {
          return function (_55_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_56_dt__update_hcf7_h0) {
                return _module.D1.create_DC3(_56_dt__update_hcf7_h0);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_51_v27);
        _nw10[(new BigNumber(18)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(19)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(20)).toNumber()] = _module.D1.create_DC3(_48_v26);
        _nw10[(new BigNumber(21)).toNumber()] = _51_v27;
        _nw10[(new BigNumber(22)).toNumber()] = _module.D1.create_DC3(_48_v26);
        _nw10[(new BigNumber(23)).toNumber()] = _51_v27;
        _54_v30 = _nw10;
        let _57_v31;
        let _nw11 = new _module.C2();
        _nw11.__ctor(_45_v25, _54_v30);
        _57_v31 = _nw11;
        (globalState).f1 = _53_v29;
        let _58_v32;
        _58_v32 = _dafny.MultiSet.fromElements(_53_v29, new BigNumber(536), _53_v29, _module.__default.fm8(_53_v29, _17_v0, _17_v0, globalState));
        let _59_v33;
        let _init4 = ((_60_v29, _61_v0) => function (_62_i4) {
          return (_62_i4).multipliedBy((_module.D14.create_DC40(_60_v29, _61_v0, _dafny.Seq.UnicodeFromString("ywpyyikps"))).dtor_cf73);
        })(_53_v29, _17_v0);
        let _nw12 = Array((new BigNumber(19)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw12.length); _i0_4++) {
          _nw12[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _59_v33 = _nw12;
        let _index8 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_59_v33).length));
        (_59_v33)[_index8] = _53_v29;
        let _index9 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_59_v33).length));
        let _rhs9 = (_53_v29).plus(_53_v29);
        let _rhs10 = ((_58_v32).Union(_58_v32)).Intersect((_58_v32).Union(_module.__default.fm33(globalState)));
        let _rhs11 = _53_v29;
        let _lhs6 = _59_v33;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_59_v33).length));
        _53_v29 = _rhs9;
        _58_v32 = _rhs10;
        _lhs6[_lhs7] = _rhs11;
        let _63_v34;
        _63_v34 = _dafny.Seq.UnicodeFromString("vejr");
        _63_v34 = _63_v34;
      }
      let _64_v35;
      let _nw13 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _64_v35 = _nw13;
      let _65_v36;
      _65_v36 = new BigNumber(-214);
      let _index10 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_64_v35).length));
      (_64_v35)[_index10] = _65_v36;
      let _index11 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_64_v35).length));
      (_64_v35)[_index11] = _65_v36;
      let _index12 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_64_v35).length));
      (_64_v35)[_index12] = (_65_v36).plus(_65_v36);
      _65_v36 = (_64_v35)[_module.__default.safeIndex(new BigNumber(91), new BigNumber((_64_v35).length))];
      let _index13 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_64_v35).length));
      (_64_v35)[_index13] = _65_v36;
      let _66_v37;
      _66_v37 = _dafny.Seq.of(_65_v36);
      let _67_v38;
      _67_v38 = new _dafny.CodePoint('x'.codePointAt(0));
      let _68_v39;
      let _nw14 = new _module.C0();
      _nw14.__ctor(_17_v0, _17_v0, _67_v38, _dafny.Seq.Create(_module.__default.abs(new BigNumber(407)), ((_69_v38) => function (_70_i5) {
        return _69_v38;
      })(_67_v38)));
      _68_v39 = _nw14;
      let _71_v40;
      _71_v40 = _dafny.MultiSet.fromElements(_68_v39, _68_v39);
      let _72_v41;
      _72_v41 = _dafny.Seq.of(_17_v0);
      let _73_v42;
      _73_v42 = _dafny.Set.fromElements(new BigNumber((_71_v40).cardinality()), new BigNumber(733), new BigNumber((_dafny.Seq.update(_72_v41, _module.__default.safeIndex(new BigNumber(688), new BigNumber((_72_v41).length)), _17_v0)).length));
      let _74_v43;
      _74_v43 = _dafny.Map.Empty.slice().updateUnsafe(false,(_68_v39).f7);
      let _75_v44;
      let _nw15 = new _module.C1();
      _nw15.__ctor(new BigNumber((_73_v42).length), (_68_v39).f6, (_68_v39).f6, (((_74_v43).contains((_68_v39).f7)) ? ((_74_v43).get((_68_v39).f7)) : (_17_v0)), _67_v38, (_68_v39).f5);
      _75_v44 = _nw15;
      let _76_v45;
      _76_v45 = _dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC21(_66_v37, _75_v44),_65_v36);
      let _77_v46;
      _77_v46 = _module.D7.create_DC21(_66_v37, _75_v44);
      let _pat_let_tv1 = _66_v37;
      let _pat_let_tv2 = _75_v44;
      let _pat_let_tv3 = _66_v37;
      let _pat_let_tv4 = _65_v36;
      _76_v45 = (_76_v45).update(function (_pat_let2_0) {
        return function (_78_dt__update__tmp_h1) {
          return function (_pat_let3_0) {
            return function (_79_dt__update_hcf39_h0) {
              return _module.D7.create_DC21(_79_dt__update_hcf39_h0, (_78_dt__update__tmp_h1).dtor_cf40);
            }(_pat_let3_0);
          }(_dafny.Seq.update(_pat_let_tv1, _module.__default.safeIndex((_pat_let_tv2).f8, new BigNumber((_pat_let_tv3).length)), _pat_let_tv4));
        }(_pat_let2_0);
      }(_77_v46), new BigNumber(957));
      return;
    }
    static Main(__noArgsParameter) {
      let _80_v0;
      _80_v0 = true;
      let _81_v1;
      let _nw16 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
      _81_v1 = _nw16;
      let _82_v2;
      _82_v2 = _dafny.Map.Empty.slice().updateUnsafe(_80_v0,_81_v1);
      let _83_globalState;
      let _nw17 = new _module.GlobalState();
      _nw17.__ctor(false, new BigNumber(130), false, (((_82_v2).contains(_80_v0)) ? ((_82_v2).get(_80_v0)) : (_81_v1)));
      _83_globalState = _nw17;
      let _84_v3;
      _84_v3 = new BigNumber(287);
      _80_v0 = ((_dafny.ZERO).minus(_84_v3)).isLessThan(new BigNumber(398));
      let _85_v4;
      _85_v4 = _dafny.Seq.UnicodeFromString("ikmk");
      let _86_v5;
      _86_v5 = _dafny.MultiSet.fromElements(_module.__default.fm0(_80_v0, new BigNumber((_85_v4).length), _84_v3, _80_v0, _83_globalState));
      _86_v5 = _86_v5;
      _80_v0 = _80_v0;
      _module.__default.m0(_83_globalState);
      let _87_v6;
      _87_v6 = _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bmln"), _dafny.Seq.UnicodeFromString("bkkejxop")), _dafny.Seq.UnicodeFromString("vnqpsmb"), _85_v4);
      let _88_v7;
      _88_v7 = _dafny.Seq.of(_85_v4);
      _87_v6 = function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of (((_80_v0) ? (_88_v7) : (_88_v7))).Elements) {
          let _89_v8 = _compr_7;
          if (_dafny.Seq.contains(((_80_v0) ? (_88_v7) : (_88_v7)), _89_v8)) {
            _coll7.add(_89_v8);
          }
        }
        return _coll7;
      }();
      let _90_i0;
      _90_i0 = _dafny.ZERO;
      L0: {
        while (_80_v0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_90_i0)) {
              break L0;
            }
            _90_i0 = (_90_i0).plus(_dafny.ONE);
            let _91_v9;
            _91_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(349),_85_v4);
            let _92_v10;
            _92_v10 = new _dafny.CodePoint('h'.codePointAt(0));
            let _93_v11;
            _93_v11 = _dafny.Map.Empty.slice().updateUnsafe(_80_v0,_module.__default.fm7(_92_v10, _84_v3, true, _83_globalState));
            let _94_v12;
            let _nw18 = new _module.C3();
            _nw18.__ctor(false, _80_v0, new BigNumber((_91_v9).length), _80_v0, true, (new BigNumber((_93_v11).length)).isLessThanOrEqualTo(_84_v3), _92_v10, _dafny.Seq.Concat(_85_v4, _85_v4));
            _94_v12 = _nw18;
            let _95_v13;
            _95_v13 = _dafny.Seq.of(_94_v12.f12, _94_v12.f12);
            let _96_v14;
            let _nw19 = Array((new BigNumber(29)).toNumber());
            _nw19[(_dafny.ZERO).toNumber()] = _80_v0;
            _nw19[(_dafny.ONE).toNumber()] = (_94_v12).f13;
            _nw19[(new BigNumber(2)).toNumber()] = (_94_v12).f13;
            _nw19[(new BigNumber(3)).toNumber()] = _94_v12.f12;
            _nw19[(new BigNumber(4)).toNumber()] = (_94_v12).f13;
            _nw19[(new BigNumber(5)).toNumber()] = _80_v0;
            _nw19[(new BigNumber(6)).toNumber()] = (_94_v12).f13;
            _nw19[(new BigNumber(7)).toNumber()] = _80_v0;
            _nw19[(new BigNumber(8)).toNumber()] = (_94_v12).f13;
            _nw19[(new BigNumber(9)).toNumber()] = (_94_v12).f13;
            _nw19[(new BigNumber(10)).toNumber()] = _module.__default.fm3(_83_globalState);
            _nw19[(new BigNumber(11)).toNumber()] = (_94_v12).f13;
            _nw19[(new BigNumber(12)).toNumber()] = (_95_v13)[_module.__default.safeIndex(_84_v3, new BigNumber((_95_v13).length))];
            _nw19[(new BigNumber(13)).toNumber()] = !(_80_v0);
            _nw19[(new BigNumber(14)).toNumber()] = _80_v0;
            _nw19[(new BigNumber(15)).toNumber()] = _80_v0;
            _nw19[(new BigNumber(16)).toNumber()] = (_94_v12).f13;
            _nw19[(new BigNumber(17)).toNumber()] = !((_94_v12).f13);
            _nw19[(new BigNumber(18)).toNumber()] = _94_v12.f12;
            _nw19[(new BigNumber(19)).toNumber()] = _80_v0;
            _nw19[(new BigNumber(20)).toNumber()] = _94_v12.f12;
            _nw19[(new BigNumber(21)).toNumber()] = _80_v0;
            _nw19[(new BigNumber(22)).toNumber()] = (_94_v12).f13;
            _nw19[(new BigNumber(23)).toNumber()] = _94_v12.f12;
            _nw19[(new BigNumber(24)).toNumber()] = _94_v12.f12;
            _nw19[(new BigNumber(25)).toNumber()] = !(_94_v12.f12);
            _nw19[(new BigNumber(26)).toNumber()] = _80_v0;
            _nw19[(new BigNumber(27)).toNumber()] = _94_v12.f12;
            _nw19[(new BigNumber(28)).toNumber()] = !(_94_v12.f12);
            _96_v14 = _nw19;
            let _97_v15;
            _97_v15 = _dafny.Map.Empty.slice().updateUnsafe(_96_v14,(((_93_v11).contains((_94_v12).f13)) ? ((_93_v11).get((_94_v12).f13)) : ((_94_v12).f13)));
            _97_v15 = (_97_v15).update(_96_v14, (_94_v12).f13);
            let _98_v16;
            let _init5 = ((_99_v3, _100_v13) => function (_101_i1) {
              return (_dafny.MultiSet.fromElements(_99_v3, new BigNumber((_100_v13).length))).Difference(_dafny.MultiSet.fromElements(_99_v3, _99_v3, _99_v3));
            })(_84_v3, _95_v13);
            let _nw20 = Array((new BigNumber(19)).toNumber());
            for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw20.length); _i0_5++) {
              _nw20[_i0_5] = _init5(new BigNumber(_i0_5));
            }
            _98_v16 = _nw20;
            _98_v16 = _98_v16;
            let _102_v17;
            _102_v17 = _dafny.Map.Empty.slice().updateUnsafe((_94_v12).f13,_95_v13);
            let _103_v18;
            let _nw21 = Array((new BigNumber(29)).toNumber());
            _nw21[(_dafny.ZERO).toNumber()] = _95_v13;
            _nw21[(_dafny.ONE).toNumber()] = _95_v13;
            _nw21[(new BigNumber(2)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(3)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(4)).toNumber()] = (((_102_v17).contains(false)) ? ((_102_v17).get(false)) : (_95_v13));
            _nw21[(new BigNumber(5)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(6)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(7)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(8)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(9)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(10)).toNumber()] = _dafny.Seq.of(true, _94_v12.f12);
            _nw21[(new BigNumber(11)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(12)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(13)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(14)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(15)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(16)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(17)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(18)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(19)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(20)).toNumber()] = _module.__default.fm18(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(640)), ((_104_v10) => function (_105_i2) {
              return _104_v10;
            })(_92_v10)), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("vqt")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(640)), ((_106_v10) => function (_107_i2) {
              return _106_v10;
            })(_92_v10))).length)), _92_v10)).length), _94_v12.f12, _83_globalState);
            _nw21[(new BigNumber(21)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(22)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(23)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(24)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(25)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(26)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(27)).toNumber()] = _95_v13;
            _nw21[(new BigNumber(28)).toNumber()] = _95_v13;
            _103_v18 = _nw21;
            let _108_v19;
            _108_v19 = _module.D1.create_DC3(_103_v18);
            let _109_v20;
            _109_v20 = _dafny.Map.Empty.slice().updateUnsafe(_84_v3,_108_v19);
            let _110_v21;
            _110_v21 = _module.D2.create_DC6(_109_v20);
            let _111_v22;
            _111_v22 = _dafny.Set.fromElements(_110_v21);
            let _112_v23;
            _112_v23 = _module.D13.create_DC35(_111_v22);
            _111_v22 = (_112_v23).dtor_cf66;
          }
        }
      }
      _80_v0 = ((true) ? (_module.__default.fm3(_83_globalState)) : (false));
      _80_v0 = _80_v0;
      let _113_v24;
      _113_v24 = _dafny.MultiSet.fromElements(_84_v3);
      _113_v24 = _113_v24;
      _84_v3 = _84_v3;
      let _114_v25;
      _114_v25 = _dafny.Seq.of(new BigNumber(372), _84_v3);
      let _115_v26;
      _115_v26 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_114_v25));
      let _116_v27;
      _116_v27 = new _dafny.CodePoint('b'.codePointAt(0));
      let _117_v28;
      let _nw22 = new _module.C1();
      _nw22.__ctor(_84_v3, _80_v0, _80_v0, _80_v0, _116_v27, _85_v4);
      _117_v28 = _nw22;
      let _118_v29;
      let _nw23 = Array((new BigNumber(2)).toNumber()).fill(false);
      _118_v29 = _nw23;
      let _119_v30;
      let _nw24 = new _module.C5();
      _nw24.__ctor(_118_v29, _81_v1, !(_80_v0), true, (_85_v4)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-922)), function (_120_i3) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })).length), new BigNumber((_85_v4).length))], _85_v4, _84_v3, _80_v0);
      _119_v30 = _nw24;
      let _121_v31;
      _121_v31 = _dafny.Map.Empty.slice().updateUnsafe(_80_v0,_119_v30);
      let _122_v32;
      let _nw25 = new _module.C4();
      _nw25.__ctor(_84_v3, _80_v0, _module.__default.fm27(_84_v3, _80_v0, _module.__default.fm27(_84_v3, !(!(_module.__default.fm27(_84_v3, _80_v0, _80_v0, _83_globalState))), _80_v0, _83_globalState), _83_globalState), _80_v0, new _dafny.CodePoint('m'.codePointAt(0)), _85_v4);
      _122_v32 = _nw25;
      let _123_v33;
      _123_v33 = _dafny.Map.Empty.slice().updateUnsafe(_122_v32,(_122_v32).f7);
      let _124_v34;
      _124_v34 = _module.D0.create_DC2(_84_v3, _80_v0, _114_v25);
      let _125_v35;
      let _nw26 = Array((new BigNumber(28)).toNumber());
      _nw26[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.FromArray(_114_v25);
      _nw26[(_dafny.ONE).toNumber()] = (_113_v24).Intersect(_dafny.MultiSet.fromElements(_84_v3));
      _nw26[(new BigNumber(2)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(3)).toNumber()] = ((_115_v26)[_module.__default.safeIndex(_84_v3, new BigNumber((_115_v26).length))]).Intersect(_113_v24);
      _nw26[(new BigNumber(4)).toNumber()] = (_113_v24).update(_84_v3, _module.__default.abs(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_117_v28,new BigNumber(((_113_v24).update(_84_v3, _module.__default.abs(_84_v3))).cardinality()))).update(_117_v28, _84_v3)).length)));
      _nw26[(new BigNumber(5)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(6)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(7)).toNumber()] = (_113_v24).Difference(_113_v24);
      _nw26[(new BigNumber(8)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(9)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(10)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(11)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(12)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(13)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(14)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(15)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(16)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(17)).toNumber()] = _module.__default.fm33(_83_globalState);
      _nw26[(new BigNumber(18)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(19)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber((_121_v31).length), _84_v3)).Difference(_113_v24);
      _nw26[(new BigNumber(20)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(608)), ((_126_v3) => function (_127_i4) {
        return (_dafny.ZERO).minus(_126_v3);
      })(_84_v3)), _dafny.Seq.of(_84_v3, new BigNumber((_114_v25).length), _84_v3)));
      _nw26[(new BigNumber(21)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(22)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(23)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber((_85_v4).length))).Union(_113_v24);
      _nw26[(new BigNumber(24)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(25)).toNumber()] = _113_v24;
      _nw26[(new BigNumber(26)).toNumber()] = (_dafny.MultiSet.fromElements(_84_v3)).Intersect(_113_v24);
      _nw26[(new BigNumber(27)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_123_v33).length), new BigNumber((_dafny.Seq.of((_122_v32).f6, (_122_v32).f7, (_122_v32).f7, (_124_v34).dtor_cf5, _80_v0)).length), new BigNumber(-415), _84_v3, _84_v3);
      _125_v35 = _nw26;
      let _index14 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_125_v35).length));
      (_125_v35)[_index14] = _113_v24;
      let _index15 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_125_v35).length));
      (_125_v35)[_index15] = (_113_v24).Intersect(_113_v24);
      let _source1 = _module.__default.fm26((_122_v32).f7, new BigNumber(147), _122_v32.f4, _83_globalState);
      if (_source1.is_DC4) {
        let _128___mcc_h0 = (_source1).cf8;
        let _129___mcc_h1 = (_source1).cf9;
        let _130___mcc_h2 = (_source1).cf10;
        let _131___mcc_h3 = (_source1).cf11;
        let _132___mcc_h4 = (_source1).cf12;
        let _133_cf12 = _132___mcc_h4;
        let _134_cf11 = _131___mcc_h3;
        let _135_cf10 = _130___mcc_h2;
        let _136_cf9 = _129___mcc_h1;
        let _137_cf8 = _128___mcc_h0;
        let _138_v36;
        _138_v36 = _dafny.MultiSet.fromElements(_80_v0, (_122_v32).f6);
        let _139_v37;
        let _140_v38;
        let _141_v39;
        let _out0;
        let _out1;
        let _out2;
        let _outcollector0 = (_119_v30).m1((_122_v32).f6, (_122_v32).f6, (_138_v36).update((_122_v32).f6, _module.__default.abs(new BigNumber(71))), (_122_v32).f6, _83_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _139_v37 = _out0;
        _140_v38 = _out1;
        _141_v39 = _out2;
        let _142_v40;
        _142_v40 = _dafny.Seq.of((_122_v32).f9);
        let _143_v41;
        _143_v41 = _dafny.Map.Empty.slice().updateUnsafe(_80_v0,new BigNumber((_142_v40).length));
        let _144_v42;
        _144_v42 = _module.D10.create_DC28(_133_cf12, new BigNumber((_143_v41).length), (_122_v32).f9, (_122_v32).f9, _117_v28);
        let _145_v43;
        let _nw27 = new _module.C1();
        _nw27.__ctor((_122_v32).f8, (_122_v32).f6, (_122_v32).f6, (_144_v42).dtor_cf50, new _dafny.CodePoint('t'.codePointAt(0)), (_122_v32).f5);
        _145_v43 = _nw27;
        let _146_v44;
        _146_v44 = _dafny.Map.Empty.slice().updateUnsafe(_145_v43,_133_cf12);
        let _147_v45;
        _147_v45 = _module.D5.create_DC15((_119_v30).f10, (_122_v32).f9, _143_v41, new BigNumber((_146_v44).length));
        let _148_v46;
        let _nw28 = new _module.C3();
        _nw28.__ctor(!((_122_v32).f7), ((_122_v32).f7) && (_80_v0), _136_cf9, (_147_v45).dtor_cf27, false, false, _137_cf8, _dafny.Seq.Concat((_122_v32).f5, _dafny.Seq.Create(_module.__default.abs(new BigNumber(519)), ((_149_v32) => function (_150_i5) {
          return _149_v32.f4;
        })(_122_v32))));
        _148_v46 = _nw28;
        _148_v46 = _145_v43;
        _133_cf12 = _dafny.Seq.IsPrefixOf((_145_v43).f5, (_145_v43).f5);
        (_148_v46).f4 = _135_cf10;
      } else if (_source1.is_DC5) {
        let _151___mcc_h5 = (_source1).cf13;
        let _152_cf13 = _151___mcc_h5;
        (_83_globalState).f1 = _84_v3;
        let _153_v47;
        _153_v47 = _dafny.Map.Empty.slice().updateUnsafe((_84_v3).minus(_152_cf13),(_122_v32).f6);
        _153_v47 = _153_v47;
        let _154_v48;
        _154_v48 = _dafny.Seq.of(false);
        let _155_v49;
        _155_v49 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_152_cf13, (_122_v32).f8),_154_v48);
        _155_v49 = _155_v49;
        let _156_v50;
        _156_v50 = _module.D2.create_DC7((_122_v32).f9);
        _80_v0 = (_117_v28).fm19((_122_v32).f6, _80_v0, _dafny.Set.fromElements(_156_v50, _156_v50, _156_v50, _156_v50, _156_v50), _83_globalState);
      } else {
        let _157___mcc_h6 = (_source1).cf7;
        let _158_cf7 = _157___mcc_h6;
        let _159_v51;
        let _nw29 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
        _159_v51 = _nw29;
        let _160_v52;
        _160_v52 = _module.D14.create_DC39(_159_v51);
        let _161_v53;
        _161_v53 = _dafny.Map.Empty.slice().updateUnsafe((_160_v52).dtor_cf72,_dafny.Seq.IsPrefixOf((_122_v32).f5, _85_v4));
        if (!(!((((_161_v53).contains(_159_v51)) ? ((_161_v53).get(_159_v51)) : ((_122_v32).f9))))) {
          let _162_v54;
          let _nw30 = new _module.C3();
          _nw30.__ctor((_122_v32).f9, (_122_v32).f7, (_122_v32).f8, (_122_v32).f6, (_122_v32).f6, _80_v0, _122_v32.f4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(424)), ((_163_v32) => function (_164_i6) {
            return _163_v32.f4;
          })(_122_v32)));
          _162_v54 = _nw30;
          let _165_v55;
          _165_v55 = _module.D7.create_DC21(_114_v25, _162_v54);
          _80_v0 = (_module.__default.safeModuloInt(new BigNumber(((_165_v55).dtor_cf39).length), _84_v3)).isLessThanOrEqualTo((_122_v32).f8);
          let _166_v56;
          _166_v56 = _dafny.Set.fromElements(new BigNumber(353));
          let _167_v57;
          _167_v57 = _dafny.Seq.of(_166_v56);
          let _168_v58;
          _168_v58 = _dafny.Map.Empty.slice().updateUnsafe((((_122_v32).f7) ? (_166_v56) : (_166_v56)),new BigNumber((_167_v57).length));
          _168_v58 = _168_v58;
          _80_v0 = (_122_v32).f7;
          let _169_v59;
          _169_v59 = _dafny.Map.Empty.slice().updateUnsafe((_122_v32).f9,(_122_v32).f9);
          let _170_v60;
          _170_v60 = _dafny.Seq.of((_122_v32).f7);
          _169_v59 = (_169_v59).update(!(!(_dafny.areEqual(_dafny.Seq.of((_162_v54).f8, (_dafny.ZERO).minus(new BigNumber((_170_v60).length))), _114_v25))), (_162_v54).f7);
          let _171_v61;
          _171_v61 = _module.D5.create_DC14((_119_v30).f11);
          let _172_v62;
          _172_v62 = _dafny.Set.fromElements(_module.D5.create_DC14(_81_v1), _171_v61, _171_v61, _module.D5.create_DC14(_81_v1), _171_v61);
          let _173_v63;
          _173_v63 = _dafny.Map.Empty.slice().updateUnsafe(_116_v27,((_122_v32).f8).multipliedBy(new BigNumber((_172_v62).length)));
          _173_v63 = (_173_v63).update(_162_v54.f4, _84_v3);
        } else {
          let _index16 = _module.__default.safeIndex(new BigNumber(265), new BigNumber(((_119_v30).f10).length));
          ((_119_v30).f10)[_index16] = (_122_v32).f7;
          let _index17 = _module.__default.safeIndex(new BigNumber(265), new BigNumber(((_119_v30).f10).length));
          ((_119_v30).f10)[_index17] = !(!((_122_v32).f7));
          let _174_v64;
          _174_v64 = _dafny.Map.Empty.slice().updateUnsafe((_122_v32).f8,(_122_v32).f7);
          let _175_v65;
          let _nw31 = new _module.C0();
          _nw31.__ctor((((_174_v64).contains((_122_v32).f8)) ? ((_174_v64).get((_122_v32).f8)) : (!((_122_v32).f6))), ((_119_v30).f10)[_module.__default.safeIndex(new BigNumber(265), new BigNumber(((_119_v30).f10).length))], _122_v32.f4, (_122_v32).f5);
          _175_v65 = _nw31;
          _175_v65 = _175_v65;
          let _176_v66;
          _176_v66 = _dafny.Map.Empty.slice().updateUnsafe(_119_v30,(_122_v32).f8);
          let _177_v67;
          _177_v67 = _dafny.Seq.of(_119_v30);
          (_83_globalState).f1 = (_114_v25)[_module.__default.safeIndex((_dafny.ZERO).minus(((((_176_v66).contains((_177_v67)[_module.__default.safeIndex(_84_v3, new BigNumber((_177_v67).length))])) ? ((_176_v66).get((_177_v67)[_module.__default.safeIndex(_84_v3, new BigNumber((_177_v67).length))])) : (_84_v3))).minus((_122_v32).f8)), new BigNumber((_114_v25).length))];
          let _178_v68;
          _178_v68 = _module.D12.create_DC34((_122_v32).f6);
          let _179_v69;
          let _nw32 = new _module.C4();
          _nw32.__ctor((_122_v32).f8, (_122_v32).f7, false, (_178_v68).dtor_cf65, _122_v32.f4, _dafny.Seq.Concat((_122_v32).f5, _dafny.Seq.Create(_module.__default.abs(new BigNumber(32)), ((_180_v32) => function (_181_i7) {
            return _180_v32.f4;
          })(_122_v32))));
          _179_v69 = _nw32;
          let _182_v70;
          let _init6 = ((_183_v32) => function (_184_i8) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(971)), ((_185_v32) => function (_186_i9) {
              return _185_v32.f4;
            })(_183_v32));
          })(_122_v32);
          let _nw33 = Array((new BigNumber(20)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw33.length); _i0_6++) {
            _nw33[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _182_v70 = _nw33;
          let _index18 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_182_v70).length));
          (_182_v70)[_index18] = _dafny.Seq.UnicodeFromString("dlgmycrna");
          let _index19 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_182_v70).length));
          (_182_v70)[_index19] = _dafny.Seq.Concat((_122_v32).f5, _dafny.Seq.Create(_module.__default.abs(new BigNumber(211)), function (_187_i10) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          }));
        }
        let _188_v71;
        _188_v71 = _dafny.Seq.of((_122_v32).f7);
        let _189_v72;
        _189_v72 = _dafny.Set.fromElements(_188_v71);
        let _190_v73;
        let _nw34 = new _module.C4();
        _nw34.__ctor(new BigNumber((_189_v72).length), _80_v0, (_122_v32).f6, !(!(false)), new _dafny.CodePoint('g'.codePointAt(0)), _dafny.Seq.Concat((_122_v32).f5, _dafny.Seq.UnicodeFromString("lctrtkorn")));
        _190_v73 = _nw34;
        _190_v73 = _190_v73;
        _84_v3 = (_122_v32).f8;
        let _191_v74;
        _191_v74 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(162),_84_v3);
        let _192_v75;
        _192_v75 = _dafny.MultiSet.fromElements((_122_v32).f9, true, _module.__default.fm3(_83_globalState));
        let _193_v76;
        let _194_v77;
        let _195_v78;
        let _out3;
        let _out4;
        let _out5;
        let _outcollector1 = (_190_v73).m1(_80_v0, !(!(_191_v74).contains(new BigNumber(((_122_v32).f5).length))), _192_v75, (_122_v32).f6, _83_globalState);
        _out3 = _outcollector1[0];
        _out4 = _outcollector1[1];
        _out5 = _outcollector1[2];
        _193_v76 = _out3;
        _194_v77 = _out4;
        _195_v78 = _out5;
      }
      let _196_v79;
      _196_v79 = _dafny.Set.fromElements((_122_v32).f9);
      let _197_v80;
      _197_v80 = _dafny.Seq.of((_196_v79).IsSubsetOf(_196_v79), (_122_v32).f9, (_122_v32).f9);
      _197_v80 = _dafny.Seq.Concat(_197_v80, _197_v80);
      let _198_v81;
      let _nw35 = new _module.C4();
      _nw35.__ctor(_84_v3, _80_v0, _80_v0, true, _116_v27, _85_v4);
      _198_v81 = _nw35;
      let _199_v82;
      _199_v82 = _dafny.Set.fromElements(_198_v81, ((_module.__default.fm27((_dafny.ZERO).minus(new BigNumber(-56)), true, !((_122_v32).f7), _83_globalState)) ? (_198_v81) : (_198_v81)));
      let _200_v83;
      let _nw36 = new _module.C4();
      _nw36.__ctor(_84_v3, (_122_v32).f7, (_122_v32).f9, _80_v0, _122_v32.f4, _85_v4);
      _200_v83 = _nw36;
      let _201_v84;
      _201_v84 = _dafny.Map.Empty.slice().updateUnsafe(_80_v0,(_80_v0) || ((_122_v32).f9));
      let _202_v85;
      let _nw37 = Array((new BigNumber(5)).toNumber()).fill(_module.D11.Default());
      _202_v85 = _nw37;
      let _rhs12 = _199_v82;
      let _rhs13 = (_122_v32).f7;
      let _rhs14 = _200_v83;
      let _rhs15 = _201_v84;
      let _rhs16 = _202_v85;
      _199_v82 = _rhs12;
      _80_v0 = _rhs13;
      _200_v83 = _rhs14;
      _201_v84 = _rhs15;
      _202_v85 = _rhs16;
      let _203_v86;
      _203_v86 = _dafny.Map.Empty.slice().updateUnsafe((_122_v32).f8,true);
      if ((((_203_v86).contains(_module.__default.safeModuloInt((_122_v32).f8, new BigNumber(((_122_v32).f5).length)))) ? ((_203_v86).get(_module.__default.safeModuloInt((_122_v32).f8, new BigNumber(((_122_v32).f5).length)))) : ((_197_v80)[_module.__default.safeIndex(new BigNumber((_85_v4).length), new BigNumber((_197_v80).length))]))) {
        if (_80_v0) {
          let _204_v87;
          let _nw38 = new _module.C0();
          _nw38.__ctor((((_122_v32).f9) ? ((_122_v32).f7) : ((_122_v32).f9)), (_122_v32).f6, ((_200_v83).f5)[_module.__default.safeIndex(_84_v3, new BigNumber(((_200_v83).f5).length))], (_122_v32).f5);
          _204_v87 = _nw38;
          _204_v87 = _204_v87;
          (_122_v32).f4 = _200_v83.f4;
          let _205_v88;
          let _nw39 = Array((new BigNumber(22)).toNumber()).fill(_module.D7.Default());
          _205_v88 = _nw39;
          let _206_v89;
          _206_v89 = _dafny.Map.Empty.slice().updateUnsafe((_122_v32).f6,_205_v88);
          let _207_v90;
          _207_v90 = _dafny.Map.Empty.slice().updateUnsafe((_206_v89).Merge(_206_v89),((_117_v28).fm1((_122_v32).f7, _201_v84, (_122_v32).f6, _83_globalState)).isLessThan((_122_v32).f8));
          let _208_v91;
          _208_v91 = _dafny.Seq.of((_119_v30).f10, (_119_v30).f10);
          let _209_v92;
          _209_v92 = _dafny.Seq.of(_208_v91, _208_v91);
          _207_v90 = (_207_v90).update(_206_v89, ((_125_v35)[_module.__default.safeIndex(new BigNumber(124), new BigNumber((_125_v35).length))]).IsProperSubsetOf(_dafny.MultiSet.fromElements((_122_v32).f8, new BigNumber(((_209_v92)[_module.__default.safeIndex((_122_v32).f8, new BigNumber((_209_v92).length))]).length), _84_v3)));
          let _210_v93;
          _210_v93 = _dafny.MultiSet.fromElements((_122_v32).f6);
          let _211_v94;
          let _212_v95;
          let _213_v96;
          let _out6;
          let _out7;
          let _out8;
          let _outcollector2 = (_122_v32).m1((_122_v32).f6, _80_v0, _210_v93, !((_122_v32).f9), _83_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _out8 = _outcollector2[2];
          _211_v94 = _out6;
          _212_v95 = _out7;
          _213_v96 = _out8;
          let _214_v97;
          _214_v97 = _module.D0.create_DC1(false, _211_v94, (_122_v32).f7, _84_v3);
          _80_v0 = !((_214_v97).dtor_cf2);
        } else {
          let _215_v98;
          _215_v98 = _dafny.Seq.of((_module.__default.fm15(_83_globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_122_v32).f8,_84_v3)));
          _215_v98 = _215_v98;
          let _216_v99;
          _216_v99 = _module.D12.create_DC34(false);
          _80_v0 = (_216_v99).dtor_cf65;
          _80_v0 = (_122_v32).f9;
          let _index20 = _module.__default.safeIndex(new BigNumber(738), new BigNumber(((_119_v30).f11).length));
          ((_119_v30).f11)[_index20] = (new BigNumber(85)).multipliedBy(_84_v3);
          let _index21 = _module.__default.safeIndex(new BigNumber(738), new BigNumber(((_119_v30).f11).length));
          ((_119_v30).f11)[_index21] = new BigNumber(766);
          let _217_v100;
          let _nw40 = new _module.C3();
          _nw40.__ctor(false, (_122_v32).f7, ((_119_v30).f11)[_module.__default.safeIndex(new BigNumber(738), new BigNumber(((_119_v30).f11).length))], true, _80_v0, !((_122_v32).f7) || (true), _122_v32.f4, (_200_v83).f5);
          _217_v100 = _nw40;
        }
        let _218_v101;
        _218_v101 = _module.D7.create_DC20();
        let _source2 = _218_v101;
        if (_source2.is_DC19) {
          let _219___mcc_h7 = (_source2).cf34;
          let _220___mcc_h8 = (_source2).cf35;
          let _221___mcc_h9 = (_source2).cf36;
          let _222___mcc_h10 = (_source2).cf37;
          let _223___mcc_h11 = (_source2).cf38;
          let _224_cf38 = _223___mcc_h11;
          let _225_cf37 = _222___mcc_h10;
          let _226_cf36 = _221___mcc_h9;
          let _227_cf35 = _220___mcc_h8;
          let _228_cf34 = _219___mcc_h7;
          let _index22 = _module.__default.safeIndex(new BigNumber(147), new BigNumber(((_119_v30).f10).length));
          ((_119_v30).f10)[_index22] = _80_v0;
          let _229_v102;
          _229_v102 = _dafny.Map.Empty.slice().updateUnsafe(_197_v80,(_122_v32).f9);
          let _index23 = _module.__default.safeIndex(new BigNumber(147), new BigNumber(((_119_v30).f10).length));
          ((_119_v30).f10)[_index23] = !(!((_122_v32).f6) || (_228_cf34)) || (_dafny.areEqual((_122_v32).f5, _dafny.Seq.update((_200_v83).f5, _module.__default.safeIndex(new BigNumber((_229_v102).length), new BigNumber(((_200_v83).f5).length)), _200_v83.f4)));
          _198_v81 = _198_v81;
          let _230_v103;
          _230_v103 = _dafny.Map.Empty.slice().updateUnsafe((_122_v32).fm1((_122_v32).f6, _201_v84, (_122_v32).f6, _83_globalState),_84_v3);
          let _231_v104;
          _231_v104 = _dafny.MultiSet.fromElements(_198_v81);
          _230_v103 = (_230_v103).update((_dafny.ZERO).minus((_122_v32).f8), (((_231_v104).contains(_198_v81)) ? ((_231_v104).get(_198_v81)) : ((_122_v32).f8)));
          let _232_v105;
          _232_v105 = _dafny.Map.Empty.slice().updateUnsafe(_200_v83,(_122_v32).f8);
          _232_v105 = (_232_v105).update(_200_v83, new BigNumber(732));
        } else if (_source2.is_DC20) {
          let _233_v106;
          let _nw41 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _233_v106 = _nw41;
          let _index24 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_233_v106).length));
          (_233_v106)[_index24] = _85_v4;
          let _index25 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_233_v106).length));
          (_233_v106)[_index25] = (_200_v83).f5;
          let _234_v107;
          let _nw42 = new _module.C5();
          _nw42.__ctor(_118_v29, (_119_v30).f11, (_122_v32).f7, (_122_v32).f9, _200_v83.f4, _85_v4, (_84_v3).minus((_dafny.ZERO).minus((_122_v32).f8)), (_122_v32).f6);
          _234_v107 = _nw42;
          let _235_v108;
          _235_v108 = _dafny.Map.Empty.slice().updateUnsafe(_84_v3,new BigNumber(956));
          let _236_v109;
          _236_v109 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(904)).isLessThan(new BigNumber(((_235_v108).update(new BigNumber((_114_v25).length), (_122_v32).f8)).length)),(_117_v28).fm1(!((_122_v32).f7), _module.__default.fm23(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(991)), function (_237_i11) {
            return new BigNumber(827);
          })).length), (_dafny.ZERO).minus((_122_v32).f8), _122_v32.f4, (_122_v32).f9, _83_globalState), (_122_v32).f7, _83_globalState));
          let _rhs17 = ((!((_122_v32).f9)) ? (new BigNumber((_203_v86).length)) : (new BigNumber(((_dafny.Set.fromElements(false)).Intersect(_196_v79)).length)));
          let _rhs18 = _236_v109;
          let _lhs8 = _83_globalState;
          _lhs8.f1 = _rhs17;
          _236_v109 = _rhs18;
          (_119_v30).m4(new BigNumber((_203_v86).length), _203_v86, _83_globalState);
        } else if (_source2.is_DC21) {
          let _238___mcc_h12 = (_source2).cf39;
          let _239___mcc_h13 = (_source2).cf40;
          let _240_cf40 = _239___mcc_h13;
          let _241_cf39 = _238___mcc_h12;
          let _242_v110;
          _242_v110 = _dafny.Set.fromElements((_122_v32).f8);
          _242_v110 = _dafny.Set.fromElements(_84_v3, new BigNumber(((_122_v32).f5).length), (_122_v32).f8, (_240_cf40).fm1((_122_v32).f9, _201_v84, false, _83_globalState));
          let _243_v111;
          let _nw43 = new _module.C4();
          _nw43.__ctor(new BigNumber((_113_v24).cardinality()), !_dafny.areEqual(_197_v80, _197_v80), !((_122_v32).f9), _module.__default.fm3(_83_globalState), _240_cf40.f4, _dafny.Seq.update((_module.D3.create_DC10((_122_v32).f5)).dtor_cf19, _module.__default.safeIndex(_84_v3, new BigNumber(((_module.D3.create_DC10((_122_v32).f5)).dtor_cf19).length)), _116_v27));
          _243_v111 = _nw43;
          _197_v80 = _dafny.Seq.Concat(_197_v80, _dafny.Seq.update(_dafny.Seq.Concat(_197_v80, _197_v80), _module.__default.safeIndex(new BigNumber((_196_v79).length), new BigNumber((_dafny.Seq.Concat(_197_v80, _197_v80)).length)), false));
          let _244_v112;
          let _out9;
          _out9 = (_119_v30).m2((_122_v32).f8, _83_globalState);
          _244_v112 = _out9;
        } else {
          let _245___mcc_h14 = (_source2).cf33;
          let _246_cf33 = _245___mcc_h14;
          let _247_v113;
          _247_v113 = _dafny.Map.Empty.slice().updateUnsafe((_122_v32).f8,(_dafny.ZERO).minus((_122_v32).f8));
          let _248_v114;
          _248_v114 = _dafny.Map.Empty.slice().updateUnsafe((_122_v32).f7,new BigNumber((_247_v113).length));
          let _249_v115;
          _249_v115 = _module.D12.create_DC31(_201_v84);
          let _250_v116;
          let _nw44 = new _module.C3();
          _nw44.__ctor((_122_v32).f6, (_122_v32).f7, (((_248_v114).contains((_122_v32).f9)) ? ((_248_v114).get((_122_v32).f9)) : ((_122_v32).f8)), !((_122_v32).f9), (_122_v32).f7, !(((_198_v81).fm1(!((_122_v32).f6), (_249_v115).dtor_cf60, true, _83_globalState)).isLessThanOrEqualTo(new BigNumber(575))), _200_v83.f4, (_122_v32).f5);
          _250_v116 = _nw44;
          let _251_v117;
          _251_v117 = _module.D15.create_DC42(_250_v116);
          let _rhs19 = _84_v3;
          let _rhs20 = (((_122_v32).f9) ? (((_122_v32).f8).minus((_dafny.ZERO).minus((_122_v32).f8))) : ((_dafny.ZERO).minus(((_122_v32).f8).minus((_122_v32).f8))));
          let _rhs21 = (_251_v117).dtor_cf77;
          let _lhs9 = _83_globalState;
          let _lhs10 = _83_globalState;
          _lhs9.f1 = _rhs19;
          _lhs10.f1 = _rhs20;
          _250_v116 = _rhs21;
          let _252_v118;
          _252_v118 = _dafny.MultiSet.fromElements((_122_v32).f9);
          let _253_v119;
          let _254_v120;
          let _255_v121;
          let _out10;
          let _out11;
          let _out12;
          let _outcollector3 = (_250_v116).m1((_122_v32).f6, !((_122_v32).f6), (_252_v118).Intersect(_252_v118), true, _83_globalState);
          _out10 = _outcollector3[0];
          _out11 = _outcollector3[1];
          _out12 = _outcollector3[2];
          _253_v119 = _out10;
          _254_v120 = _out11;
          _255_v121 = _out12;
          _254_v120 = ((_dafny.ZERO).minus(_84_v3)).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ubva")).length));
          _84_v3 = (_122_v32).f8;
        }
        let _256_v122;
        let _257_v123;
        let _258_v124;
        let _259_v125;
        let _out13;
        let _out14;
        let _out15;
        let _out16;
        let _outcollector4 = (_117_v28).m8((_122_v32).f7, (_122_v32).f9, _83_globalState);
        _out13 = _outcollector4[0];
        _out14 = _outcollector4[1];
        _out15 = _outcollector4[2];
        _out16 = _outcollector4[3];
        _256_v122 = _out13;
        _257_v123 = _out14;
        _258_v124 = _out15;
        _259_v125 = _out16;
        let _260_v126;
        _260_v126 = _dafny.Map.Empty.slice().updateUnsafe((_122_v32).f6,new BigNumber(-203));
        let _261_v127;
        let _nw45 = new _module.C3();
        _nw45.__ctor(_80_v0, (_122_v32).f9, _258_v124, (_122_v32).f7, (_122_v32).f6, false, new _dafny.CodePoint('u'.codePointAt(0)), (_122_v32).f5);
        _261_v127 = _nw45;
        let _262_v128;
        _262_v128 = _dafny.Map.Empty.slice().updateUnsafe(_261_v127,_258_v124);
        let _263_v129;
        _263_v129 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(420),(_dafny.ZERO).minus((_122_v32).f8));
        let _rhs22 = (_122_v32).f6;
        let _rhs23 = _module.__default.safeDivisionInt((_84_v3).plus(new BigNumber(((_260_v126).update(_80_v0, new BigNumber((_262_v128).length))).length)), _module.__default.safeModuloInt(_84_v3, (_dafny.ZERO).minus((((_263_v129).contains(_258_v124)) ? ((_263_v129).get(_258_v124)) : ((_122_v32).f8)))));
        _80_v0 = _rhs22;
        _84_v3 = _rhs23;
        let _264_v130;
        _264_v130 = _module.D3.create_DC11(!(_261_v127.f12), (_122_v32).f8, _dafny.Set.fromElements(_module.__default.fm7(_116_v27, _258_v124, (_122_v32).f6, _83_globalState), true));
        let _265_v131;
        _265_v131 = _dafny.Map.Empty.slice().updateUnsafe(_264_v130,_84_v3);
        let _266_v133;
        _266_v133 = _dafny.MultiSet.fromElements(_module.D3.create_DC11((_122_v32).f6, _258_v124, _196_v79));
        let _267_v134;
        _267_v134 = _dafny.Map.Empty.slice().updateUnsafe((_122_v32).f7,function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of (_266_v133).Elements) {
            let _268_v132 = _compr_8;
            if ((_266_v133).contains(_268_v132)) {
              _coll8.push([_268_v132,_84_v3]);
            }
          }
          return _coll8;
        }());
        _265_v131 = ((_dafny.Map.Empty.slice().updateUnsafe(_264_v130,(_122_v32).f8)).Merge((((_267_v134).contains((_122_v32).f9)) ? ((_267_v134).get((_122_v32).f9)) : (_dafny.Map.Empty.slice().updateUnsafe(_264_v130,(_122_v32).f8))))).Merge(_265_v131);
      } else {
        if ((_122_v32).f9) {
          _80_v0 = ((_122_v32).f8).isLessThanOrEqualTo(_84_v3);
          _80_v0 = (((_203_v86).contains(_84_v3)) ? ((_203_v86).get(_84_v3)) : (!((new BigNumber(((_122_v32).f5).length)).isEqualTo(_84_v3))));
          let _269_v135;
          _269_v135 = _dafny.Map.Empty.slice().updateUnsafe((_122_v32).f6,_196_v79);
          let _270_v136;
          _270_v136 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_200_v83).f5).length),_269_v135);
          _203_v86 = (_203_v86).update(_module.__default.safeDivisionInt(_84_v3, new BigNumber(((((_270_v136).contains(new BigNumber((_197_v80).length))) ? ((_270_v136).get(new BigNumber((_197_v80).length))) : (_269_v135))).length)), (_122_v32).f7);
          let _271_v137;
          let _nw46 = new _module.C3();
          _nw46.__ctor((_122_v32).f9, (_122_v32).f6, (new BigNumber((_197_v80).length)).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(370)), ((_272_v27) => function (_273_i12) {
            return _272_v27;
          })(_116_v27))).length)), !((_197_v80)[_module.__default.safeIndex((_122_v32).f8, new BigNumber((_197_v80).length))]) || ((_122_v32).f6), (_122_v32).f9, false, _116_v27, _dafny.Seq.Concat((_200_v83).f5, (_122_v32).f5));
          _271_v137 = _nw46;
          let _274_v138;
          _274_v138 = _dafny.Map.Empty.slice().updateUnsafe(!(!((_122_v32).f6)),_84_v3);
          let _275_v139;
          _275_v139 = _module.D5.create_DC15((_119_v30).f10, _80_v0, _274_v138, new BigNumber(((_122_v32).f5).length));
          let _276_v140;
          _276_v140 = _dafny.Set.fromElements(_122_v32.f4);
          let _277_v141;
          _277_v141 = _dafny.Set.fromElements(new BigNumber((_276_v140).length));
          let _278_v142;
          let _nw47 = Array((new BigNumber(19)).toNumber());
          _nw47[(_dafny.ZERO).toNumber()] = _277_v141;
          _nw47[(_dafny.ONE).toNumber()] = _277_v141;
          _nw47[(new BigNumber(2)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(3)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(4)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(5)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(6)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(7)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(8)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(9)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(10)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(11)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements((_122_v32).f8);
          _nw47[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements((_122_v32).f8);
          _nw47[(new BigNumber(14)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(15)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(16)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(17)).toNumber()] = _277_v141;
          _nw47[(new BigNumber(18)).toNumber()] = _dafny.Set.fromElements((_122_v32).f8, _84_v3, (_122_v32).f8);
          _278_v142 = _nw47;
          let _279_v143;
          _279_v143 = _module.D14.create_DC39(_278_v142);
          let _pat_let_tv5 = _278_v142;
          let _280_v144;
          _280_v144 = _dafny.Map.Empty.slice().updateUnsafe((_275_v139).dtor_cf29,function (_pat_let4_0) {
            return function (_281_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_282_dt__update_hcf72_h0) {
                  return _module.D14.create_DC39(_282_dt__update_hcf72_h0);
                }(_pat_let5_0);
              }(_pat_let_tv5);
            }(_pat_let4_0);
          }(_279_v143));
          _280_v144 = (_280_v144).update((_84_v3).minus((_122_v32).f8), _module.D14.create_DC39(_278_v142));
        } else {
          _80_v0 = _80_v0;
          let _283_v145;
          let _nw48 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _283_v145 = _nw48;
          let _index26 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_283_v145).length));
          (_283_v145)[_index26] = (_200_v83).f5;
          let _index27 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_283_v145).length));
          (_283_v145)[_index27] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(183)), ((_284_v32) => function (_285_i13) {
            return _284_v32.f4;
          })(_122_v32)), _85_v4);
          let _286_v146;
          let _nw49 = new _module.C0();
          _nw49.__ctor((_122_v32).f6, _80_v0, _122_v32.f4, _dafny.Seq.UnicodeFromString("igmblqb"));
          _286_v146 = _nw49;
          let _287_v147;
          _287_v147 = _dafny.Map.Empty.slice().updateUnsafe((_122_v32).f9,_286_v146);
          (_83_globalState).f1 = ((new BigNumber((_287_v147).length)).plus((_122_v32).f8)).minus((_122_v32).f8);
          _201_v84 = (_201_v84).update(false, (_122_v32).f9);
          let _288_v148;
          _288_v148 = _module.D12.create_DC33(!(false), (_200_v83).f5);
          let _289_v149;
          _289_v149 = _dafny.Map.Empty.slice().updateUnsafe(_288_v148,(_122_v32).f7);
          _289_v149 = (_289_v149).update(_288_v148, (_122_v32).f6);
        }
        let _290_v151;
        _290_v151 = _dafny.Set.fromElements(new BigNumber(271));
        let _291_v152;
        _291_v152 = _dafny.Map.Empty.slice().updateUnsafe(!((_122_v32).f9),(_122_v32).f8);
        let _292_v153;
        _292_v153 = _dafny.Map.Empty.slice().updateUnsafe(_84_v3,new BigNumber((_291_v152).length));
        let _293_v155;
        let _nw50 = new _module.C0();
        _nw50.__ctor((_122_v32).f9, (_122_v32).f7, _116_v27, (_122_v32).f5);
        _293_v155 = _nw50;
        let _294_v156;
        _294_v156 = _dafny.Seq.of(_293_v155);
        let _295_v157;
        _295_v157 = _dafny.Seq.of(_294_v156);
        let _296_v158;
        _296_v158 = _dafny.MultiSet.fromElements((_122_v32).f9, (_122_v32).f7, _80_v0);
        let _297_v160;
        _297_v160 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_295_v157)[_module.__default.safeIndex((((_296_v158).contains(_80_v0)) ? ((_296_v158).get(_80_v0)) : ((_122_v32).f8)), new BigNumber((_295_v157).length))]).length),function () {
          let _coll9 = new _dafny.Set();
          for (const _compr_9 of ((_125_v35)[_module.__default.safeIndex(new BigNumber(124), new BigNumber((_125_v35).length))]).Elements) {
            let _298_v159 = _compr_9;
            if (((_125_v35)[_module.__default.safeIndex(new BigNumber(124), new BigNumber((_125_v35).length))]).contains(_298_v159)) {
              _coll9.add((_298_v159).multipliedBy(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(852)), function (_299_i14) {
                return new BigNumber(-752);
              }))).cardinality())));
            }
          }
          return _coll9;
        }());
        let _300_v161;
        let _nw51 = Array((new BigNumber(15)).toNumber());
        _nw51[(_dafny.ZERO).toNumber()] = function () {
          let _coll10 = new _dafny.Set();
          for (const _compr_10 of _dafny.IntegerRange(new BigNumber(381), new BigNumber(-403))) {
            let _301_v150 = _compr_10;
            if (((new BigNumber(381)).isLessThanOrEqualTo(_301_v150)) && ((_301_v150).isLessThan(new BigNumber(-403)))) {
              _coll10.add(_module.__default.safeModuloInt(_301_v150, (_dafny.ZERO).minus((_122_v32).f8)));
            }
          }
          return _coll10;
        }();
        _nw51[(_dafny.ONE).toNumber()] = _290_v151;
        _nw51[(new BigNumber(2)).toNumber()] = _290_v151;
        _nw51[(new BigNumber(3)).toNumber()] = _290_v151;
        _nw51[(new BigNumber(4)).toNumber()] = _290_v151;
        _nw51[(new BigNumber(5)).toNumber()] = function () {
          let _coll11 = new _dafny.Set();
          for (const _compr_11 of (_292_v153).Keys.Elements) {
            let _302_v154 = _compr_11;
            if ((_292_v153).contains(_302_v154)) {
              _coll11.add(_module.__default.safeModuloInt(_302_v154, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("ucmoqf")).length))).length)));
            }
          }
          return _coll11;
        }();
        _nw51[(new BigNumber(6)).toNumber()] = _290_v151;
        _nw51[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements((_122_v32).f8);
        _nw51[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements((_122_v32).f8);
        _nw51[(new BigNumber(9)).toNumber()] = _290_v151;
        _nw51[(new BigNumber(10)).toNumber()] = _290_v151;
        _nw51[(new BigNumber(11)).toNumber()] = _290_v151;
        _nw51[(new BigNumber(12)).toNumber()] = _290_v151;
        _nw51[(new BigNumber(13)).toNumber()] = (((_297_v160).contains((_122_v32).f8)) ? ((_297_v160).get((_122_v32).f8)) : (_290_v151));
        _nw51[(new BigNumber(14)).toNumber()] = _290_v151;
        _300_v161 = _nw51;
        let _303_v162;
        _303_v162 = _module.D14.create_DC39(_300_v161);
        let _source3 = _303_v162;
        if (_source3.is_DC40) {
          let _304___mcc_h15 = (_source3).cf73;
          let _305___mcc_h16 = (_source3).cf74;
          let _306___mcc_h17 = (_source3).cf75;
          let _307_cf75 = _306___mcc_h17;
          let _308_cf74 = _305___mcc_h16;
          let _309_cf73 = _304___mcc_h15;
          _module.__default.m0(_83_globalState);
          _80_v0 = !((_122_v32).f8).isEqualTo(_module.__default.safeModuloInt(_84_v3, _309_cf73));
          _308_cf74 = (_122_v32).f9;
          let _rhs24 = (_122_v32).f6;
          let _rhs25 = (_122_v32).f8;
          let _rhs26 = (_308_cf74) || (!((_122_v32).f6));
          let _rhs27 = ((_122_v32).f8).minus(_309_cf73);
          let _lhs11 = _83_globalState;
          _80_v0 = _rhs24;
          _lhs11.f1 = _rhs25;
          _308_cf74 = _rhs26;
          _84_v3 = _rhs27;
        } else if (_source3.is_DC39) {
          let _310___mcc_h18 = (_source3).cf72;
          let _311_cf72 = _310___mcc_h18;
          _module.__default.m0(_83_globalState);
          let _312_v163;
          _312_v163 = _dafny.Seq.of(_117_v28);
          let _313_v164;
          let _nw52 = Array((new BigNumber(14)).toNumber());
          _nw52[(_dafny.ZERO).toNumber()] = _117_v28;
          _nw52[(_dafny.ONE).toNumber()] = _117_v28;
          _nw52[(new BigNumber(2)).toNumber()] = _117_v28;
          _nw52[(new BigNumber(3)).toNumber()] = _117_v28;
          _nw52[(new BigNumber(4)).toNumber()] = (_312_v163)[_module.__default.safeIndex((_122_v32).f8, new BigNumber((_312_v163).length))];
          _nw52[(new BigNumber(5)).toNumber()] = _117_v28;
          _nw52[(new BigNumber(6)).toNumber()] = _117_v28;
          _nw52[(new BigNumber(7)).toNumber()] = _117_v28;
          _nw52[(new BigNumber(8)).toNumber()] = (((_122_v32).f9) ? (_117_v28) : (_117_v28));
          _nw52[(new BigNumber(9)).toNumber()] = _117_v28;
          _nw52[(new BigNumber(10)).toNumber()] = _117_v28;
          _nw52[(new BigNumber(11)).toNumber()] = _117_v28;
          _nw52[(new BigNumber(12)).toNumber()] = (_module.D10.create_DC28((_122_v32).f9, (_122_v32).f8, _80_v0, (_122_v32).f7, _117_v28)).dtor_cf54;
          _nw52[(new BigNumber(13)).toNumber()] = _117_v28;
          _313_v164 = _nw52;
          let _index28 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_313_v164).length));
          (_313_v164)[_index28] = _117_v28;
          let _index29 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_313_v164).length));
          (_313_v164)[_index29] = _117_v28;
          (_83_globalState).f1 = _84_v3;
          _303_v162 = _303_v162;
        } else {
          let _314___mcc_h19 = (_source3).cf76;
          let _315_cf76 = _314___mcc_h19;
          let _316_v165;
          let _init7 = ((_317_v80) => function (_318_i15) {
            return _dafny.Seq.Concat(_317_v80, _317_v80);
          })(_197_v80);
          let _nw53 = Array((new BigNumber(12)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw53.length); _i0_7++) {
            _nw53[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _316_v165 = _nw53;
          _316_v165 = _316_v165;
          _80_v0 = (_122_v32).f7;
          let _319_v166;
          _319_v166 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_122_v32).f8)).length),_197_v80);
          _197_v80 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update((((_319_v166).contains((_122_v32).f8)) ? ((_319_v166).get((_122_v32).f8)) : (_197_v80)), _module.__default.safeIndex(new BigNumber(((_200_v83).f5).length), new BigNumber(((((_319_v166).contains((_122_v32).f8)) ? ((_319_v166).get((_122_v32).f8)) : (_197_v80))).length)), (_122_v32).f9), _197_v80), _197_v80);
          _80_v0 = (_122_v32).f9;
        }
        let _320_v167;
        _320_v167 = _module.D2.create_DC8(_84_v3, (_122_v32).f8);
        let _321_v168;
        _321_v168 = _module.D2.create_DC9(_320_v167);
        _321_v168 = _321_v168;
        let _322_v169;
        _322_v169 = _module.D5.create_DC14(_81_v1);
        let _323_v170;
        _323_v170 = _dafny.MultiSet.fromElements(_322_v169);
        _323_v170 = _323_v170;
        _80_v0 = (_122_v32).f9;
      }
      let _324_v171;
      _324_v171 = _module.D0.create_DC0();
      let _325_v172;
      _325_v172 = _dafny.Map.Empty.slice().updateUnsafe(_324_v171,(_122_v32).fm1(true, _201_v84, (_122_v32).f9, _83_globalState));
      _325_v172 = (_325_v172).Merge(_dafny.Map.Empty.slice().updateUnsafe(_324_v171,new BigNumber(763)));
      process.stdout.write(_dafny.toString(_80_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_81_v1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_81_v1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_82_v2).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_83_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_83_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_83_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_83_globalState).f3)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_83_globalState).f3)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_84_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_85_v4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v5).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(true, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_87_v6).equals(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ikmk")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_88_v7, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ikmk")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_90_i0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_113_v24).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_114_v25, _dafny.Seq.of(new BigNumber(372), new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_115_v26, _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(372), new BigNumber(287))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_116_v27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_118_v29)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_119_v30).f10)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_119_v30).f11)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_119_v30).f11)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_121_v31).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v32).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v32).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v32).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v32).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_122_v32.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_122_v32).f5).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_123_v33).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v34).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v34).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_124_v34).dtor_cf6, _dafny.Seq.of(new BigNumber(372), new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(new BigNumber(372), new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(12)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(13)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(14)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(15)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(16)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(17)]).equals(_dafny.MultiSet.fromElements(new BigNumber(-78)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(18)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(19)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(20)]).equals(_dafny.MultiSet.fromElements(new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(-287), new BigNumber(287), new BigNumber(287), new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(21)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(22)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(23)]).equals(_dafny.MultiSet.fromElements(new BigNumber(4), new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(24)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(25)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(26)]).equals(_dafny.MultiSet.fromElements(new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_125_v35)[new BigNumber(27)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE, new BigNumber(5), new BigNumber(-415), new BigNumber(287), new BigNumber(287)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v79).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_197_v80, _dafny.Seq.of(true, false, false, true, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_199_v82).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_200_v83.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_200_v83).f5).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_201_v84).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v86).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(287),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_325_v172).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(),new BigNumber(763)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0() {
      let $dt = new D0(0);
      return $dt;
    }
    static create_DC1(cf0, cf1, cf2, cf3) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC2(cf4, cf5, cf6) {
      let $dt = new D0(2);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf0) + ", " + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0 && _dafny.areEqual(this.cf1, other.cf1) && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4(cf8, cf9, cf10, cf11, cf12) {
      let $dt = new D1(0);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC5(cf13) {
      let $dt = new D1(1);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC3(cf7) {
      let $dt = new D1(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11) && this.cf12 === other.cf12;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf7 === other.cf7;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf15) {
      let $dt = new D2(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC8(cf16, cf17) {
      let $dt = new D2(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC6(cf14) {
      let $dt = new D2(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC9(cf18) {
      let $dt = new D2(3);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get is_DC9() { return this.$tag === 3; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC11(cf20, cf21, cf22) {
      let $dt = new D3(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC10(cf19) {
      let $dt = new D3(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC12(cf23) {
      let $dt = new D3(2);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + this.cf19.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC11(false, _dafny.ZERO, _dafny.Set.Empty);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC13(cf24) {
      let $dt = new D4(0);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC15(cf26, cf27, cf28, cf29) {
      let $dt = new D5(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC14(cf25) {
      let $dt = new D5(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf26 === other.cf26 && this.cf27 === other.cf27 && _dafny.areEqual(this.cf28, other.cf28) && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf25 === other.cf25;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC15([], false, _dafny.Map.Empty, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf31, cf32) {
      let $dt = new D6(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC16(cf30) {
      let $dt = new D6(1);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + this.cf31.toVerbatimString(true) + ", " + this.cf32.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31) && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(_dafny.Seq.UnicodeFromString(""), _dafny.Seq.UnicodeFromString(""));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC19(cf34, cf35, cf36, cf37, cf38) {
      let $dt = new D7(0);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC20() {
      let $dt = new D7(1);
      return $dt;
    }
    static create_DC21(cf39, cf40) {
      let $dt = new D7(2);
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC18(cf33) {
      let $dt = new D7(3);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get is_DC18() { return this.$tag === 3; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf34) + ", " + this.cf35.toVerbatimString(true) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC20";
      } else if (this.$tag === 2) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 3) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf34 === other.cf34 && _dafny.areEqual(this.cf35, other.cf35) && this.cf36 === other.cf36 && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf39, other.cf39) && this.cf40 === other.cf40;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf33, other.cf33);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC19(false, _dafny.Seq.UnicodeFromString(""), false, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC23(cf42, cf43) {
      let $dt = new D8(0);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC22(cf41) {
      let $dt = new D8(1);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf42, other.cf42) && this.cf43 === other.cf43;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC23(_dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC25(cf45, cf46, cf47) {
      let $dt = new D9(0);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC24(cf44) {
      let $dt = new D9(1);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf45 === other.cf45 && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC25(false, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC27(cf49) {
      let $dt = new D10(0);
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC28(cf50, cf51, cf52, cf53, cf54) {
      let $dt = new D10(1);
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC26(cf48) {
      let $dt = new D10(2);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf49, other.cf49);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf50 === other.cf50 && _dafny.areEqual(this.cf51, other.cf51) && this.cf52 === other.cf52 && this.cf53 === other.cf53 && this.cf54 === other.cf54;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf48 === other.cf48;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC27(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC30(cf56, cf57, cf58, cf59) {
      let $dt = new D11(0);
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC29(cf55) {
      let $dt = new D11(1);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC29" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf56, other.cf56) && this.cf57 === other.cf57 && this.cf58 === other.cf58 && this.cf59 === other.cf59;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf55, other.cf55);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC30(_module.D7.Default(), false, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC32(cf61, cf62) {
      let $dt = new D12(0);
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC33(cf63, cf64) {
      let $dt = new D12(1);
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC34(cf65) {
      let $dt = new D12(2);
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC31(cf60) {
      let $dt = new D12(3);
      $dt.cf60 = cf60;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get is_DC31() { return this.$tag === 3; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf60() { return this.cf60; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf63) + ", " + this.cf64.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf60) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf61 === other.cf61 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf63 === other.cf63 && _dafny.areEqual(this.cf64, other.cf64);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf65 === other.cf65;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf60, other.cf60);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC32(false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC36(cf67) {
      let $dt = new D13(0);
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC37(cf68, cf69, cf70) {
      let $dt = new D13(1);
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC35(cf66) {
      let $dt = new D13(2);
      $dt.cf66 = cf66;
      return $dt;
    }
    static create_DC38(cf71) {
      let $dt = new D13(3);
      $dt.cf71 = cf71;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get is_DC35() { return this.$tag === 2; }
    get is_DC38() { return this.$tag === 3; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf71() { return this.cf71; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC36" + "(" + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC37" + "(" + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC35" + "(" + _dafny.toString(this.cf66) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC38" + "(" + _dafny.toString(this.cf71) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf68, other.cf68) && _dafny.areEqual(this.cf69, other.cf69) && this.cf70 === other.cf70;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf66, other.cf66);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf71, other.cf71);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC36(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC40(cf73, cf74, cf75) {
      let $dt = new D14(0);
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      return $dt;
    }
    static create_DC39(cf72) {
      let $dt = new D14(1);
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC41(cf76) {
      let $dt = new D14(2);
      $dt.cf76 = cf76;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get is_DC39() { return this.$tag === 1; }
    get is_DC41() { return this.$tag === 2; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf76() { return this.cf76; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC40" + "(" + _dafny.toString(this.cf73) + ", " + _dafny.toString(this.cf74) + ", " + this.cf75.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC39" + "(" + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC41" + "(" + _dafny.toString(this.cf76) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf73, other.cf73) && this.cf74 === other.cf74 && _dafny.areEqual(this.cf75, other.cf75);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf72 === other.cf72;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf76, other.cf76);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC40(_dafny.ZERO, false, _dafny.Seq.UnicodeFromString(""));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC43(cf78) {
      let $dt = new D15(0);
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC42(cf77) {
      let $dt = new D15(1);
      $dt.cf77 = cf77;
      return $dt;
    }
    static create_DC44(cf79) {
      let $dt = new D15(2);
      $dt.cf79 = cf79;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get is_DC44() { return this.$tag === 2; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf79() { return this.cf79; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC43" + "(" + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC42" + "(" + _dafny.toString(this.cf77) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC44" + "(" + _dafny.toString(this.cf79) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf78 === other.cf78;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf77 === other.cf77;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf79, other.cf79);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC43(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC46(cf81) {
      let $dt = new D16(0);
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC45(cf80) {
      let $dt = new D16(1);
      $dt.cf80 = cf80;
      return $dt;
    }
    static create_DC47(cf82) {
      let $dt = new D16(2);
      $dt.cf82 = cf82;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC45() { return this.$tag === 1; }
    get is_DC47() { return this.$tag === 2; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf82() { return this.cf82; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC46" + "(" + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC45" + "(" + _dafny.toString(this.cf80) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC47" + "(" + _dafny.toString(this.cf82) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf81, other.cf81);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf80, other.cf80);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf82, other.cf82);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC46(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.T2 = class T2 {
  };

  $module.T3 = class T3 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = _dafny.ZERO;
      this._f0 = false;
      this._f2 = false;
      this._f3 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f4 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f6 = false;
      this._f7 = false;
      this._f5 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    set f4(value) {
      let _this = this;
      _this._f4 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f6, f7, f4, f5) {
      let _this = this;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm9(p0, p1, globalState) {
      let _this = this;
      return new BigNumber((function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(759)), function (_326_i0) {
          return _dafny.MultiSet.fromElements((_module.D0.create_DC2(new BigNumber(821), (_this).f6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(755)), function (_327_i1) {
  return new BigNumber(318);
}))).dtor_cf5, (_this).f6);
        })).Elements) {
          let _328_v0 = _compr_12;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(759)), function (_326_i0) {
            return _dafny.MultiSet.fromElements((_module.D0.create_DC2(new BigNumber(821), (_this).f6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(755)), function (_327_i1) {
  return new BigNumber(318);
}))).dtor_cf5, (_this).f6);
          }), _328_v0)) {
            _coll12.push([_328_v0,_dafny.Seq.of(_module.D2.create_DC7((_this).f7), _module.D2.create_DC7((_this).f6))]);
          }
        }
        return _coll12;
      }()).length);
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f4 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f8 = _dafny.ZERO;
      this._f9 = false;
      this._f6 = false;
      this._f7 = false;
      this._f5 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    set f4(value) {
      let _this = this;
      _this._f4 = value;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f8, f9, f6, f7, f4, f5) {
      let _this = this;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber(-485);
    };
    fm2(globalState) {
      let _this = this;
      let _source4 = _module.D3.create_DC12(_module.D3.create_DC12(_module.D3.create_DC11((_this).f9, (_this).f8, _dafny.Set.fromElements((_this).f9))));
      if (_source4.is_DC11) {
        let _329___mcc_h0 = (_source4).cf20;
        let _330___mcc_h1 = (_source4).cf21;
        let _331___mcc_h2 = (_source4).cf22;
        let _332_cf22 = _331___mcc_h2;
        let _333_cf21 = _330___mcc_h1;
        let _334_cf20 = _329___mcc_h0;
        return _dafny.Seq.of(_333_cf21);
      } else if (_source4.is_DC10) {
        let _335___mcc_h3 = (_source4).cf19;
        let _336_cf19 = _335___mcc_h3;
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-920)), function (_337_i0) {
          return (_this).f8;
        }), _dafny.Seq.of(new BigNumber(537)));
      } else {
        let _338___mcc_h4 = (_source4).cf23;
        let _339_cf23 = _338___mcc_h4;
        return _dafny.Seq.of(new BigNumber((function () {
          let _coll13 = new _dafny.Set();
          for (const _compr_13 of _dafny.IntegerRange(new BigNumber(372), new BigNumber(170))) {
            let _340_v0 = _compr_13;
            if (((new BigNumber(372)).isLessThanOrEqualTo(_340_v0)) && ((_340_v0).isLessThan(new BigNumber(170)))) {
              _coll13.add((_340_v0).plus((_this).f8));
            }
          }
          return _coll13;
        }()).length), new BigNumber((_dafny.MultiSet.fromElements(_this.f4)).cardinality()));
      }
    };
    fm19(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f7;
    };
    fm20(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(639)), function (_341_i0) {
        return _module.D1.create_DC5((_this).f8);
      }), _dafny.Seq.of(_module.D1.create_DC5((_this).f8), _module.D1.create_DC5((_this).f8), _module.D1.create_DC5((_this).f8)));
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = [];
      let _342_v0;
      _342_v0 = _module.D2.create_DC7(p1);
      let _343_v1;
      _343_v1 = _dafny.Seq.of((_342_v0).dtor_cf15);
      let _344_v2;
      _344_v2 = _dafny.Set.fromElements((_this).f7, false);
      let _345_v3;
      _345_v3 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_344_v2).length)));
      let _346_v4;
      _346_v4 = _dafny.Seq.of(new BigNumber((_345_v3).length));
      let _347_v5;
      _347_v5 = _module.D0.create_DC2((_this).f8, (_343_v1)[_module.__default.safeIndex((_this).f8, new BigNumber((_343_v1).length))], _345_v3);
      let _348_v6;
      _348_v6 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_347_v5).dtor_cf5);
      let _349_v7;
      _349_v7 = _dafny.Set.fromElements(function (_pat_let6_0) {
        return function (_350_dt__update__tmp_h0) {
          return function (_pat_let7_0) {
            return function (_351_dt__update_hcf15_h0) {
              return _module.D2.create_DC7(_351_dt__update_hcf15_h0);
            }(_pat_let7_0);
          }(true);
        }(_pat_let6_0);
      }(_module.D2.create_DC7(p0)), _342_v0, _342_v0);
      let _352_v8;
      _352_v8 = _module.D3.create_DC12(_module.D3.create_DC10((_this).f5));
      let _353_v9;
      _353_v9 = _module.D3.create_DC12(_352_v8);
      let _354_v10;
      _354_v10 = _module.D3.create_DC12(_352_v8);
      let _355_v11;
      _355_v11 = _dafny.MultiSet.fromElements(_354_v10);
      let _356_v12;
      let _nw54 = Array((new BigNumber(9)).toNumber());
      _nw54[(_dafny.ZERO).toNumber()] = (_this).fm1((_this).f6, _348_v6, (_this).f6, globalState);
      _nw54[(_dafny.ONE).toNumber()] = (_this).f8;
      _nw54[(new BigNumber(2)).toNumber()] = (((_this).f6) ? ((_this).f8) : (new BigNumber(31)));
      _nw54[(new BigNumber(3)).toNumber()] = (_this).f8;
      _nw54[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.update(_343_v1, _module.__default.safeIndex(new BigNumber(531), new BigNumber((_343_v1).length)), (_this).f7)).length);
      _nw54[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((((_this).fm19(false, p0, _349_v7, globalState)) ? (new BigNumber(593)) : ((_this).f8)));
      _nw54[(new BigNumber(6)).toNumber()] = (_this).f8;
      _nw54[(new BigNumber(7)).toNumber()] = new BigNumber(((_355_v11).update(_354_v10, _module.__default.abs((_this).f8))).cardinality());
      _nw54[(new BigNumber(8)).toNumber()] = (_this).f8;
      _356_v12 = _nw54;
      let _index30 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_356_v12).length));
      (_356_v12)[_index30] = ((_this).f8).multipliedBy((_this).f8);
      let _357_v13;
      let _nw55 = Array((new BigNumber(16)).toNumber()).fill(false);
      _357_v13 = _nw55;
      let _index31 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_357_v13).length));
      (_357_v13)[_index31] = ((p0) ? ((_this).f7) : (p3));
      let _pat_let_tv6 = p3;
      let _pat_let_tv7 = _346_v4;
      let _index32 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_356_v12).length));
      let _index33 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_357_v13).length));
      let _rhs28 = function (_source5) {
        if (_source5.is_DC0) {
          return (_this).f8;
        } else if (_source5.is_DC1) {
          let _358___mcc_h0 = (_source5).cf0;
          let _359___mcc_h1 = (_source5).cf1;
          let _360___mcc_h2 = (_source5).cf2;
          let _361___mcc_h3 = (_source5).cf3;
          let _362_cf3 = _361___mcc_h3;
          let _363_cf2 = _360___mcc_h2;
          let _364_cf1 = _359___mcc_h1;
          let _365_cf0 = _358___mcc_h0;
          return new BigNumber(((_this).f5).length);
        } else {
          let _366___mcc_h4 = (_source5).cf4;
          let _367___mcc_h5 = (_source5).cf5;
          let _368___mcc_h6 = (_source5).cf6;
          let _369_cf6 = _368___mcc_h6;
          let _370_cf5 = _367___mcc_h5;
          let _371_cf4 = _366___mcc_h4;
          return (_this).f8;
        }
      }(function (_pat_let8_0) {
        return function (_372_dt__update__tmp_h1) {
          return function (_pat_let9_0) {
            return function (_373_dt__update_hcf5_h0) {
              return function (_pat_let10_0) {
                return function (_374_dt__update_hcf6_h0) {
                  return _module.D0.create_DC2((_372_dt__update__tmp_h1).dtor_cf4, _373_dt__update_hcf5_h0, _374_dt__update_hcf6_h0);
                }(_pat_let10_0);
              }(_pat_let_tv7);
            }(_pat_let9_0);
          }(_pat_let_tv6);
        }(_pat_let8_0);
      }(_module.__default.fm21((_this).f6, true, (_this).f9, p0, globalState)));
      let _rhs29 = true;
      let _rhs30 = _module.__default.safeDivisionInt((_this).f8, (_this).f8);
      let _lhs12 = _356_v12;
      let _lhs13 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_356_v12).length));
      let _lhs14 = _357_v13;
      let _lhs15 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_357_v13).length));
      _lhs12[_lhs13] = _rhs28;
      _lhs14[_lhs15] = _rhs29;
      r1 = _rhs30;
      let _index34 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_357_v13).length));
      (_357_v13)[_index34] = false;
      let _375_v14;
      let _nw56 = new _module.C0();
      _nw56.__ctor(!(p3), p0, _this.f4, (((_357_v13)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_357_v13).length))]) ? ((_this).f5) : ((_this).f5)));
      _375_v14 = _nw56;
      let _376_v15;
      _376_v15 = _dafny.Set.fromElements((_this).f5, (_this).f5);
      let _377_v16;
      _377_v16 = _dafny.Set.fromElements((_356_v12)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_356_v12).length))], (_this).f8);
      let _378_v17;
      _378_v17 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements((_this).f5, _dafny.Seq.update((_this).f5, _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(668)), function (_379_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), _module.__default.safeIndex(new BigNumber(((_this).f5).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(668)), function (_380_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length)), _this.f4)).length), new BigNumber(((_this).f5).length)), _this.f4), _dafny.Seq.of(_this.f4), (_this).f5, (_this).f5)).Intersect(_376_v15),(_377_v16).IsDisjointFrom(_dafny.Set.fromElements((_this).f8, (_this).f8)));
      _378_v17 = (_378_v17).update((_376_v15).Difference(_dafny.Set.fromElements((_this).f5)), (_this).f7);
      let _381_v18;
      _381_v18 = _dafny.Set.fromElements(_357_v13, _357_v13);
      let _index35 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_356_v12).length));
      (_356_v12)[_index35] = new BigNumber((_381_v18).length);
      let _382_v19;
      let _nw57 = new _module.C0();
      _nw57.__ctor(p0, (_this).f7, _this.f4, (_this).f5);
      _382_v19 = _nw57;
      r0 = (_this).f8;
      r1 = (_dafny.ZERO).minus(new BigNumber((_345_v3).length));
      let _383_v20;
      let _nw58 = Array((new BigNumber(14)).toNumber()).fill([]);
      _383_v20 = _nw58;
      r2 = _383_v20;
      return [r0, r1, r2];
    }
    m8(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = [];
      let r2 = _dafny.ZERO;
      let r3 = [];
      (globalState).f1 = (_this).f8;
      let _384_v0;
      let _nw59 = Array((new BigNumber(7)).toNumber()).fill(_dafny.MultiSet.Empty);
      _384_v0 = _nw59;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_384_v0).length))) {
        let _385_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_385_i0)) && ((_385_i0).isLessThan(new BigNumber((_384_v0).length))))) {
          (_384_v0)[(_385_i0)] = (_dafny.MultiSet.fromElements((_this).f8, (_this).f8, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_this).f5).length))), new BigNumber(-67))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(p1, (_this).f9, (_this).f6)).length)), _dafny.Seq.of(new BigNumber(173)))));
        }
      }
      let _386_i1;
      _386_i1 = _dafny.ZERO;
      L1: {
        while (false) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_386_i1)) {
              break L1;
            }
            _386_i1 = (_386_i1).plus(_dafny.ONE);
            let _387_v1;
            _387_v1 = _module.D1.create_DC4(_this.f4, (_this).f8, _this.f4, (_this).f8, true);
            let _388_v2;
            _388_v2 = _dafny.Seq.of(_this.f4, (_387_v1).dtor_cf10);
            _388_v2 = _dafny.Seq.update((_this).f5, _module.__default.safeIndex(_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f8), (_this).f8), new BigNumber(((_this).f5).length)), _this.f4);
            let _389_v3;
            _389_v3 = false;
            _389_v3 = (_this).f6;
            (globalState).f1 = new BigNumber(76);
            let _390_v4;
            _390_v4 = _module.D3.create_DC10(_388_v2);
            let _391_v5;
            _391_v5 = _module.D3.create_DC12(_390_v4);
            _391_v5 = _module.D3.create_DC12(_390_v4);
          }
        }
      }
      _module.__default.m0(globalState);
      let _392_v6;
      _392_v6 = false;
      _392_v6 = (_this).f7;
      let _393_v7;
      _393_v7 = _dafny.Seq.of((_this).f8);
      let _394_v8;
      _394_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).fm2(globalState));
      let _395_v9;
      _395_v9 = _dafny.Set.fromElements(p0, (_this).f7, (_this).f9, _392_v6);
      let _396_v10;
      _396_v10 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_395_v9).length)), _393_v7, (_this).fm2(globalState));
      let _397_v11;
      let _nw60 = Array((new BigNumber(23)).toNumber());
      _nw60[(_dafny.ZERO).toNumber()] = _dafny.Seq.of((_this).f8);
      _nw60[(_dafny.ONE).toNumber()] = _393_v7;
      _nw60[(new BigNumber(2)).toNumber()] = _393_v7;
      _nw60[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_393_v7, _393_v7), _module.__default.safeIndex((_this).f8, new BigNumber((_dafny.Seq.Concat(_393_v7, _393_v7)).length)), (_this).f8);
      _nw60[(new BigNumber(4)).toNumber()] = _393_v7;
      _nw60[(new BigNumber(5)).toNumber()] = _393_v7;
      _nw60[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(((false) ? ((((_394_v8).contains(p1)) ? ((_394_v8).get(p1)) : ((_396_v10)[_module.__default.safeIndex(new BigNumber(((_this).f5).length), new BigNumber((_396_v10).length))]))) : (_393_v7)), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f8), new BigNumber((((false) ? ((((_394_v8).contains(p1)) ? ((_394_v8).get(p1)) : ((_396_v10)[_module.__default.safeIndex(new BigNumber(((_this).f5).length), new BigNumber((_396_v10).length))]))) : (_393_v7))).length)), (_this).f8);
      _nw60[(new BigNumber(7)).toNumber()] = _dafny.Seq.of((_this).f8, (_this).f8, (_this).f8, (_this).f8, (_this).f8);
      _nw60[(new BigNumber(8)).toNumber()] = _393_v7;
      _nw60[(new BigNumber(9)).toNumber()] = _393_v7;
      _nw60[(new BigNumber(10)).toNumber()] = _393_v7;
      _nw60[(new BigNumber(11)).toNumber()] = _393_v7;
      _nw60[(new BigNumber(12)).toNumber()] = _dafny.Seq.of((_this).f8, (_this).f8, (_this).f8, (_this).f8);
      _nw60[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_393_v7, _393_v7);
      _nw60[(new BigNumber(14)).toNumber()] = _393_v7;
      _nw60[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_393_v7, _393_v7);
      _nw60[(new BigNumber(16)).toNumber()] = _393_v7;
      _nw60[(new BigNumber(17)).toNumber()] = _393_v7;
      _nw60[(new BigNumber(18)).toNumber()] = (_this).fm2(globalState);
      _nw60[(new BigNumber(19)).toNumber()] = _393_v7;
      _nw60[(new BigNumber(20)).toNumber()] = _393_v7;
      _nw60[(new BigNumber(21)).toNumber()] = _393_v7;
      _nw60[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_393_v7, _393_v7);
      _397_v11 = _nw60;
      let _398_v12;
      _398_v12 = _dafny.Seq.of(_392_v6);
      let _index36 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_397_v11).length));
      (_397_v11)[_index36] = _dafny.Seq.Concat((_this).fm2(globalState), _dafny.Seq.update(_393_v7, _module.__default.safeIndex((_393_v7)[_module.__default.safeIndex((_this).f8, new BigNumber((_393_v7).length))], new BigNumber((_393_v7).length)), new BigNumber((_dafny.Seq.update(_398_v12, _module.__default.safeIndex(new BigNumber((_393_v7).length), new BigNumber((_398_v12).length)), p1)).length)));
      let _399_v13;
      _399_v13 = _module.D0.create_DC2(new BigNumber(981), (_this).f7, _393_v7);
      let _index37 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_397_v11).length));
      (_397_v11)[_index37] = (_399_v13).dtor_cf6;
      let _400_v14;
      _400_v14 = _dafny.Map.Empty.slice().updateUnsafe(_392_v6,_dafny.Seq.Create(_module.__default.abs(new BigNumber(697)), function (_401_i2) {
        return _this.f4;
      }));
      let _402_v15;
      _402_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _403_v16;
      _403_v16 = _dafny.Set.fromElements(_402_v15);
      let _404_v17;
      _404_v17 = _dafny.Seq.of(_403_v16);
      let _405_v18;
      _405_v18 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f7), _module.__default.fm23((_this).f8, (_this).f8, _this.f4, (_this).f6, globalState));
      r0 = (((_400_v14).contains(_392_v6)) ? ((((_this).f6) ? ((_404_v17)[_module.__default.safeIndex((_this).f8, new BigNumber((_404_v17).length))]) : (_module.__default.fm22(globalState)))) : ((_403_v16).Intersect(function () {
        let _coll14 = new _dafny.Set();
        for (const _compr_14 of (_405_v18).Elements) {
          let _406_v19 = _compr_14;
          if (_dafny.Seq.contains(_405_v18, _406_v19)) {
            _coll14.add(_406_v19);
          }
        }
        return _coll14;
      }())));
      let _407_v20;
      let _init8 = function (_408_i3) {
        return (_this).f7;
      };
      let _nw61 = Array((new BigNumber(2)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw61.length); _i0_8++) {
        _nw61[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _407_v20 = _nw61;
      r1 = _407_v20;
      r2 = (_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_this.f4, _this.f4, _this.f4)).length))).minus((_this).f8));
      r3 = _407_v20;
      return [r0, r1, r2, r3];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this.f15 = [];
      this._f14 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f14, f15) {
      let _this = this;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      return;
    }
    fm6(globalState) {
      let _this = this;
      if ((_dafny.Map.Empty.slice().updateUnsafe(!(true),false)).equals(_dafny.Map.Empty.slice().updateUnsafe(!(false),true))) {
        return _dafny.Seq.of(_dafny.Set.fromElements(true));
      } else {
        return _dafny.Seq.of(_dafny.Set.fromElements(false), _dafny.Set.fromElements(false, true, true));
      }
    };
    m6(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _module.D1.Default();
      let r2 = [];
      let r3 = _module.D1.Default();
      let _409_v0;
      _409_v0 = _dafny.Seq.UnicodeFromString("svbxxdv");
      let _410_i0;
      _410_i0 = _dafny.ZERO;
      L2: {
        while (_dafny.areEqual(_dafny.Seq.UnicodeFromString("hfrnqnere"), _409_v0)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_410_i0)) {
              break L2;
            }
            _410_i0 = (_410_i0).plus(_dafny.ONE);
            let _411_v1;
            _411_v1 = true;
            let _412_v2;
            _412_v2 = new BigNumber(-884);
            let _413_v3;
            let _init9 = ((_414_v1) => function (_415_i1) {
              return _dafny.Seq.of(!(_414_v1));
            })(_411_v1);
            let _nw62 = Array((new BigNumber(2)).toNumber());
            for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw62.length); _i0_9++) {
              _nw62[_i0_9] = _init9(new BigNumber(_i0_9));
            }
            _413_v3 = _nw62;
            let _416_v4;
            _416_v4 = _dafny.Map.Empty.slice().updateUnsafe(_412_v2,_module.D1.create_DC3(_413_v3));
            let _417_v5;
            _417_v5 = _dafny.Map.Empty.slice().updateUnsafe(_411_v1,_416_v4);
            let _418_v6;
            _418_v6 = _module.D2.create_DC6(_416_v4);
            _417_v5 = (_417_v5).update((_411_v1) === (_411_v1), (_416_v4).Merge((_418_v6).dtor_cf14));
            let _419_v7;
            _419_v7 = new _dafny.CodePoint('y'.codePointAt(0));
            let _420_v8;
            _420_v8 = _dafny.Map.Empty.slice().updateUnsafe(_411_v1,_419_v7);
            let _421_v9;
            _421_v9 = _module.D1.create_DC3(_413_v3);
            let _422_v10;
            _422_v10 = _dafny.MultiSet.fromElements(_421_v9);
            let _423_v11;
            _423_v11 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm7(_419_v7, new BigNumber((_422_v10).cardinality()), _411_v1, globalState),_411_v1);
            let _424_v12;
            _424_v12 = _dafny.Map.Empty.slice().updateUnsafe(_420_v8,_423_v11);
            _424_v12 = _424_v12;
            _409_v0 = _dafny.Seq.UnicodeFromString("d");
            let _425_v13;
            _425_v13 = _dafny.Seq.of(!(_411_v1));
            let _426_v14;
            _426_v14 = _dafny.Set.fromElements(_411_v1, _411_v1);
            let _427_v15;
            _427_v15 = _module.D0.create_DC2(_412_v2, _411_v1, _dafny.Seq.of(new BigNumber((_426_v14).length), new BigNumber((_409_v0).length)));
            _409_v0 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("aetllib"), _dafny.Seq.update(_409_v0, _module.__default.safeIndex(_module.__default.fm8(new BigNumber((_dafny.Seq.of(_419_v7)).length), _411_v1, _411_v1, globalState), new BigNumber((_409_v0).length)), new _dafny.CodePoint('o'.codePointAt(0)))), _dafny.Seq.update(_409_v0, _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_425_v13, _module.__default.safeIndex(_412_v2, new BigNumber((_425_v13).length)), (_427_v15).dtor_cf5)).length), new BigNumber((_409_v0).length)), _419_v7));
          }
        }
      }
      if (true) {
        let _428_v16;
        _428_v16 = false;
        let _429_v17;
        _429_v17 = new _dafny.CodePoint('j'.codePointAt(0));
        let _430_v18;
        let _nw63 = new _module.C0();
        _nw63.__ctor(_428_v16, _428_v16, _429_v17, _409_v0);
        _430_v18 = _nw63;
        let _431_v19;
        _431_v19 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("gmxt"),_428_v16);
        _431_v19 = (_431_v19).Merge((_431_v19).update(_409_v0, _428_v16));
        let _432_v20;
        _432_v20 = new BigNumber(475);
        let _433_v21;
        _433_v21 = _dafny.Set.fromElements(_432_v20, _432_v20);
        (globalState).f1 = new BigNumber(((((_433_v21).IsDisjointFrom(_433_v21)) ? (_433_v21) : (_module.__default.fm10(true, _432_v20, _module.__default.fm8(_432_v20, _428_v16, _428_v16, globalState), globalState)))).length);
        let _434_v22;
        _434_v22 = _module.D1.create_DC5(_432_v20);
        _434_v22 = ((_428_v16) ? (_434_v22) : (_434_v22));
        let _435_v23;
        _435_v23 = _dafny.Seq.of(_432_v20, (_dafny.ZERO).minus(_432_v20));
        let _436_v24;
        _436_v24 = _dafny.Map.Empty.slice().updateUnsafe(_409_v0,_432_v20);
        let _437_v25;
        let _nw64 = Array((new BigNumber(4)).toNumber());
        _nw64[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_435_v23, _435_v23);
        _nw64[(_dafny.ONE).toNumber()] = _435_v23;
        _nw64[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_435_v23, _dafny.Seq.of(_432_v20, new BigNumber((_436_v24).length), _432_v20, new BigNumber((_dafny.MultiSet.fromElements(_428_v16, _428_v16)).cardinality()), new BigNumber(782)));
        _nw64[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_435_v23, _module.__default.safeIndex(_432_v20, new BigNumber((_435_v23).length)), _432_v20);
        _437_v25 = _nw64;
        _437_v25 = _437_v25;
      } else {
        let _438_v26;
        _438_v26 = new BigNumber(650);
        (globalState).f1 = (_438_v26).minus(_438_v26);
        let _439_v27;
        _439_v27 = false;
        _439_v27 = _439_v27;
        let _440_v28;
        _440_v28 = _dafny.Seq.of(_438_v26, _438_v26);
        let _441_v29;
        _441_v29 = _module.D0.create_DC2(_438_v26, false, _440_v28);
        if (!((_441_v29).dtor_cf5)) {
          let _442_v30;
          let _nw65 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _442_v30 = _nw65;
          let _443_v31;
          _443_v31 = _dafny.Seq.of(_439_v27, _439_v27, _439_v27);
          let _444_v32;
          _444_v32 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_439_v27), _443_v31);
          let _index38 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_442_v30).length));
          (_442_v30)[_index38] = (new BigNumber((_444_v32).cardinality())).multipliedBy(_438_v26);
          let _445_v33;
          let _nw66 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
          _445_v33 = _nw66;
          let _446_v34;
          _446_v34 = _module.D1.create_DC3(_445_v33);
          let _447_v35;
          _447_v35 = _dafny.Map.Empty.slice().updateUnsafe(_438_v26,_446_v34);
          let _448_v36;
          _448_v36 = _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(_438_v26,_446_v34));
          let _449_v37;
          _449_v37 = _dafny.MultiSet.fromElements(((_439_v27) ? (_module.D2.create_DC6(_447_v35)) : (_448_v36)), _448_v36);
          let _index39 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_442_v30).length));
          (_442_v30)[_index39] = new BigNumber((_449_v37).cardinality());
          let _index40 = _module.__default.safeIndex(new BigNumber(310), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index40] = true;
          let _450_v38;
          _450_v38 = _dafny.MultiSet.fromElements(_439_v27, (_441_v29).dtor_cf5, _439_v27);
          let _index41 = _module.__default.safeIndex(new BigNumber(310), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index41] = ((((_450_v38).contains(_439_v27)) ? ((_450_v38).get(_439_v27)) : ((_442_v30)[_module.__default.safeIndex(new BigNumber(333), new BigNumber((_442_v30).length))]))).isLessThan(new BigNumber(181));
          let _451_v39;
          _451_v39 = new _dafny.CodePoint('h'.codePointAt(0));
          let _452_v40;
          let _nw67 = new _module.C0();
          _nw67.__ctor(((_this).f14)[_module.__default.safeIndex(new BigNumber(310), new BigNumber(((_this).f14).length))], ((_439_v27) ? (_439_v27) : (((_this).f14)[_module.__default.safeIndex(new BigNumber(310), new BigNumber(((_this).f14).length))])), _451_v39, _409_v0);
          _452_v40 = _nw67;
          _452_v40 = _452_v40;
          let _index42 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_442_v30).length));
          let _index43 = _module.__default.safeIndex(new BigNumber(310), new BigNumber(((_this).f14).length));
          let _rhs31 = ((_this).f14)[_module.__default.safeIndex(new BigNumber(310), new BigNumber(((_this).f14).length))];
          let _rhs32 = _438_v26;
          let _rhs33 = ((_this).f14)[_module.__default.safeIndex(new BigNumber(310), new BigNumber(((_this).f14).length))];
          let _rhs34 = _451_v39;
          let _lhs16 = _442_v30;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_442_v30).length));
          let _lhs18 = (_this).f14;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(310), new BigNumber(((_this).f14).length));
          _439_v27 = _rhs31;
          _lhs16[_lhs17] = _rhs32;
          _lhs18[_lhs19] = _rhs33;
          _451_v39 = _rhs34;
          r0 = _438_v26;
        } else {
          let _453_v41;
          _453_v41 = _dafny.Set.fromElements(_439_v27);
          let _454_v42;
          _454_v42 = _module.D2.create_DC8((_438_v26).plus(_438_v26), _438_v26);
          let _455_v43;
          _455_v43 = _dafny.Seq.of(_439_v27, _439_v27, _439_v27, false);
          let _456_v44;
          _456_v44 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(((!(_439_v27)) ? (_455_v43) : (_455_v43)), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("ibbeoyt")).length), new BigNumber((((!(_439_v27)) ? (_455_v43) : (_455_v43))).length)), _439_v27)).length), new BigNumber((_dafny.Seq.UnicodeFromString("vjhcwwr")).length), _438_v26);
          let _rhs35 = (((_456_v44).contains(_438_v26)) ? ((_456_v44).get(_438_v26)) : ((_dafny.ZERO).minus(_438_v26)));
          let _rhs36 = (_453_v41).Union(_453_v41);
          let _rhs37 = _454_v42;
          let _rhs38 = _dafny.Seq.Create(_module.__default.abs(_dafny.ZERO), ((_457_v26) => function (_458_i2) {
            return _457_v26;
          })(_438_v26));
          let _rhs39 = _dafny.Seq.Concat(_455_v43, _455_v43);
          let _lhs20 = globalState;
          _lhs20.f1 = _rhs35;
          _453_v41 = _rhs36;
          _454_v42 = _rhs37;
          _440_v28 = _rhs38;
          _455_v43 = _rhs39;
          let _459_v45;
          let _nw68 = Array((new BigNumber(15)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _459_v45 = _nw68;
          _459_v45 = _459_v45;
          let _index44 = _module.__default.safeIndex(new BigNumber(550), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index44] = _439_v27;
          let _index45 = _module.__default.safeIndex(new BigNumber(550), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index45] = _439_v27;
          (globalState).f1 = _438_v26;
          let _460_v46;
          _460_v46 = _dafny.Set.fromElements(_438_v26, new BigNumber(666));
          let _461_v47;
          _461_v47 = _dafny.Seq.of(_460_v46, _dafny.Set.fromElements(new BigNumber((_409_v0).length), _module.__default.fm8(new BigNumber(688), _439_v27, !(_439_v27), globalState)));
          _439_v27 = ((_461_v47)[_module.__default.safeIndex(_438_v26, new BigNumber((_461_v47).length))]).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber((_456_v44).cardinality())));
        }
        let _462_v48;
        let _nw69 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _462_v48 = _nw69;
        let _463_v49;
        _463_v49 = _dafny.Seq.of(_462_v48, _462_v48, _462_v48, _462_v48, _462_v48);
        let _rhs40 = _439_v27;
        let _rhs41 = _438_v26;
        let _rhs42 = !(_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of(_462_v48, _462_v48), _463_v49), _462_v48));
        let _rhs43 = _439_v27;
        let _rhs44 = _439_v27;
        _439_v27 = _rhs40;
        _438_v26 = _rhs41;
        _439_v27 = _rhs42;
        _439_v27 = _rhs43;
        _439_v27 = _rhs44;
        _439_v27 = !(!(_439_v27));
      }
      let _464_v50;
      _464_v50 = new BigNumber(-178);
      (globalState).f1 = _464_v50;
      let _465_v51;
      let _nw70 = Array((new BigNumber(20)).toNumber()).fill(false);
      _465_v51 = _nw70;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_465_v51).length))) {
        let _466_i3 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_466_i3)) && ((_466_i3).isLessThan(new BigNumber((_465_v51).length))))) {
          (_465_v51)[(_466_i3)] = _dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.of(_464_v50), _dafny.Seq.of(_464_v50)), _dafny.Seq.Concat(_dafny.Seq.of(_464_v50, _464_v50, new BigNumber((_dafny.Set.fromElements(_464_v50, _464_v50)).length), new BigNumber(150)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(397)), ((_467_v50) => function (_468_i4) {
            return _467_v50;
          })(_464_v50))));
        }
      }
      let _469_v52;
      let _nw71 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _469_v52 = _nw71;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_469_v52).length))) {
        let _470_i5 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_470_i5)) && ((_470_i5).isLessThan(new BigNumber((_469_v52).length))))) {
          (_469_v52)[(_470_i5)] = _dafny.Seq.update(_409_v0, _module.__default.safeIndex(_464_v50, new BigNumber((_409_v0).length)), new _dafny.CodePoint('r'.codePointAt(0)));
        }
      }
      let _471_v53;
      _471_v53 = false;
      if (((_471_v53) ? (_471_v53) : (_dafny.Seq.IsPrefixOf(_409_v0, _409_v0)))) {
        let _472_v54;
        _472_v54 = _dafny.Seq.of(_471_v53, _471_v53);
        let _473_v55;
        _473_v55 = _dafny.Set.fromElements(_471_v53, (_472_v54)[_module.__default.safeIndex(new BigNumber((_472_v54).length), new BigNumber((_472_v54).length))]);
        let _index46 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_469_v52).length));
        (_469_v52)[_index46] = _dafny.Seq.UnicodeFromString("ptockx");
        let _index47 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_469_v52).length));
        let _rhs45 = _473_v55;
        let _rhs46 = _409_v0;
        let _lhs21 = _469_v52;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_469_v52).length));
        _473_v55 = _rhs45;
        _lhs21[_lhs22] = _rhs46;
        let _474_v56;
        let _nw72 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _474_v56 = _nw72;
        let _index48 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_474_v56).length));
        (_474_v56)[_index48] = (_464_v50).plus(new BigNumber(((_469_v52)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_469_v52).length))]).length));
        let _475_v57;
        _475_v57 = _dafny.Map.Empty.slice().updateUnsafe(_471_v53,new BigNumber((_473_v55).length));
        let _index49 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_474_v56).length));
        (_474_v56)[_index49] = ((_471_v53) ? (_module.__default.safeModuloInt(new BigNumber((_475_v57).length), _464_v50)) : ((_464_v50).multipliedBy(_464_v50)));
        let _476_v58;
        _476_v58 = _dafny.Map.Empty.slice().updateUnsafe((_474_v56)[_module.__default.safeIndex(new BigNumber(760), new BigNumber((_474_v56).length))],!(_471_v53));
        let _477_v59;
        _477_v59 = _dafny.Seq.of((_476_v58).update(_464_v50, _471_v53));
        let _478_v60;
        _478_v60 = _dafny.Seq.of(_477_v59);
        let _479_v61;
        _479_v61 = _dafny.Seq.of(new BigNumber(-869), (_474_v56)[_module.__default.safeIndex(new BigNumber(760), new BigNumber((_474_v56).length))]);
        let _480_v62;
        _480_v62 = new _dafny.CodePoint('n'.codePointAt(0));
        let _index50 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_474_v56).length));
        let _index51 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_474_v56).length));
        let _rhs47 = _471_v53;
        let _rhs48 = (_474_v56)[_module.__default.safeIndex(new BigNumber(760), new BigNumber((_474_v56).length))];
        let _rhs49 = (_478_v60)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("jmbaqsn"), _module.__default.safeIndex((_479_v61)[_module.__default.safeIndex(new BigNumber(576), new BigNumber((_479_v61).length))], new BigNumber((_dafny.Seq.UnicodeFromString("jmbaqsn")).length)), _480_v62)).length), new BigNumber((_478_v60).length))];
        let _rhs50 = _471_v53;
        let _rhs51 = new BigNumber(-502);
        let _lhs23 = _474_v56;
        let _lhs24 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_474_v56).length));
        let _lhs25 = _474_v56;
        let _lhs26 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_474_v56).length));
        _471_v53 = _rhs47;
        _lhs23[_lhs24] = _rhs48;
        _477_v59 = _rhs49;
        _471_v53 = _rhs50;
        _lhs25[_lhs26] = _rhs51;
        let _481_v63;
        let _nw73 = new _module.C0();
        _nw73.__ctor(_471_v53, ((_471_v53) ? (_471_v53) : (false)), new _dafny.CodePoint('c'.codePointAt(0)), _dafny.Seq.Concat(_409_v0, _409_v0));
        _481_v63 = _nw73;
        _480_v62 = _480_v62;
      } else {
        let _482_v64;
        let _nw74 = Array((new BigNumber(2)).toNumber());
        _nw74[(_dafny.ZERO).toNumber()] = _465_v51;
        _nw74[(_dafny.ONE).toNumber()] = (_this).f14;
        _482_v64 = _nw74;
        let _index52 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_482_v64).length));
        (_482_v64)[_index52] = (_this).f14;
        let _index53 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_482_v64).length));
        (_482_v64)[_index53] = _465_v51;
        let _483_v65;
        _483_v65 = new _dafny.CodePoint('p'.codePointAt(0));
        let _484_v66;
        let _nw75 = new _module.C0();
        _nw75.__ctor(!(true), _471_v53, _483_v65, _409_v0);
        _484_v66 = _nw75;
        let _485_v67;
        _485_v67 = _dafny.Map.Empty.slice().updateUnsafe((_471_v53) && (_471_v53),_471_v53);
        _485_v67 = (_485_v67).Merge(_485_v67);
        _471_v53 = ((!(_471_v53)) ? (_module.__default.fm7(_483_v65, _464_v50, _471_v53, globalState)) : (_471_v53));
        _483_v65 = _483_v65;
      }
      r0 = _464_v50;
      let _486_v68;
      let _nw76 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
      _486_v68 = _nw76;
      let _487_v69;
      _487_v69 = _module.D1.create_DC3(_486_v68);
      r1 = _487_v69;
      let _488_v71;
      _488_v71 = _module.D3.create_DC10(_409_v0);
      let _489_v72;
      _489_v72 = _dafny.Map.Empty.slice().updateUnsafe((_488_v71).dtor_cf19,_471_v53);
      let _490_v73;
      _490_v73 = _dafny.Seq.of(_464_v50);
      let _491_v74;
      let _nw77 = Array((new BigNumber(16)).toNumber());
      _nw77[(_dafny.ZERO).toNumber()] = (_464_v50).minus(_464_v50);
      _nw77[(_dafny.ONE).toNumber()] = _464_v50;
      _nw77[(new BigNumber(2)).toNumber()] = _464_v50;
      _nw77[(new BigNumber(3)).toNumber()] = new BigNumber((function () {
        let _coll15 = new _dafny.Map();
        for (const _compr_15 of (_489_v72).Keys.Elements) {
          let _492_v70 = _compr_15;
          if ((_489_v72).contains(_492_v70)) {
            _coll15.push([_492_v70,_464_v50]);
          }
        }
        return _coll15;
      }()).length);
      _nw77[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(_464_v50, new BigNumber(719));
      _nw77[(new BigNumber(5)).toNumber()] = _464_v50;
      _nw77[(new BigNumber(6)).toNumber()] = ((_471_v53) ? (new BigNumber(233)) : (_464_v50));
      _nw77[(new BigNumber(7)).toNumber()] = _464_v50;
      _nw77[(new BigNumber(8)).toNumber()] = (_490_v73)[_module.__default.safeIndex(_464_v50, new BigNumber((_490_v73).length))];
      _nw77[(new BigNumber(9)).toNumber()] = _464_v50;
      _nw77[(new BigNumber(10)).toNumber()] = ((_471_v53) ? (_464_v50) : (_464_v50));
      _nw77[(new BigNumber(11)).toNumber()] = new BigNumber((_409_v0).length);
      _nw77[(new BigNumber(12)).toNumber()] = _464_v50;
      _nw77[(new BigNumber(13)).toNumber()] = (_490_v73)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_490_v73).length))];
      _nw77[(new BigNumber(14)).toNumber()] = _464_v50;
      _nw77[(new BigNumber(15)).toNumber()] = new BigNumber((_409_v0).length);
      _491_v74 = _nw77;
      r2 = _491_v74;
      let _pat_let_tv8 = _486_v68;
      r3 = function (_pat_let11_0) {
        return function (_493_dt__update__tmp_h0) {
          return function (_pat_let12_0) {
            return function (_494_dt__update_hcf7_h0) {
              return _module.D1.create_DC3(_494_dt__update_hcf7_h0);
            }(_pat_let12_0);
          }(_pat_let_tv8);
        }(_pat_let11_0);
      }(_487_v69);
      return [r0, r1, r2, r3];
    }
    m7(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _495_v0;
      _495_v0 = false;
      let _496_v1;
      _496_v1 = new BigNumber(64);
      let _497_v2;
      _497_v2 = _module.D0.create_DC1(_495_v0, ((_495_v0) ? (_496_v1) : (_496_v1)), _495_v0, _496_v1);
      let _source6 = _497_v2;
      if (_source6.is_DC0) {
        if ((_496_v1).isEqualTo(_496_v1)) {
          let _498_v3;
          _498_v3 = _dafny.MultiSet.fromElements(_496_v1);
          let _499_v4;
          _499_v4 = _dafny.MultiSet.fromElements(_496_v1, new BigNumber((_498_v3).cardinality()), _module.__default.safeModuloInt(_496_v1, _module.__default.fm8(new BigNumber(-351), _495_v0, _495_v0, globalState)));
          _498_v3 = _498_v3;
          r2 = _496_v1;
          let _500_v5;
          _500_v5 = _dafny.Seq.UnicodeFromString("vcmororjs");
          let _501_v6;
          let _nw78 = new _module.C0();
          _nw78.__ctor(_495_v0, _495_v0, new _dafny.CodePoint('q'.codePointAt(0)), _500_v5);
          _501_v6 = _nw78;
          r1 = ((true) ? (_495_v0) : (true));
          _501_v6 = _501_v6;
        } else {
          r2 = _module.__default.safeModuloInt(_496_v1, _module.__default.fm8(_496_v1, _495_v0, _495_v0, globalState));
          let _502_v7;
          let _nw79 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _502_v7 = _nw79;
          let _init10 = function (_503_i0) {
            return _module.__default.safeDivisionInt(_503_i0, new BigNumber(696));
          };
          let _nw80 = Array((new BigNumber(20)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw80.length); _i0_10++) {
            _nw80[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _502_v7 = _nw80;
          let _504_v8;
          _504_v8 = _dafny.MultiSet.fromElements(!(_495_v0) || (_495_v0));
          _504_v8 = _504_v8;
          let _505_v9;
          let _506_v10;
          let _507_v11;
          let _508_v12;
          let _out17;
          let _out18;
          let _out19;
          let _out20;
          let _outcollector5 = (_this).m6(globalState);
          _out17 = _outcollector5[0];
          _out18 = _outcollector5[1];
          _out19 = _outcollector5[2];
          _out20 = _outcollector5[3];
          _505_v9 = _out17;
          _506_v10 = _out18;
          _507_v11 = _out19;
          _508_v12 = _out20;
          _495_v0 = (_505_v9).isEqualTo(_496_v1);
        }
        let _509_v13;
        _509_v13 = _dafny.Seq.of((p0).dtor_cf8);
        let _510_v14;
        _510_v14 = _dafny.Seq.of(_495_v0, _dafny.Seq.contains(_509_v13, (_509_v13)[_module.__default.safeIndex(_496_v1, new BigNumber((_509_v13).length))]));
        _495_v0 = (_510_v14)[_module.__default.safeIndex(_496_v1, new BigNumber((_510_v14).length))];
        (globalState).f1 = _496_v1;
        let _511_v15;
        _511_v15 = _module.D3.create_DC10(_dafny.Seq.Concat(_509_v13, _509_v13));
        _511_v15 = _module.D3.create_DC10(_509_v13);
      } else if (_source6.is_DC1) {
        let _512___mcc_h0 = (_source6).cf0;
        let _513___mcc_h1 = (_source6).cf1;
        let _514___mcc_h2 = (_source6).cf2;
        let _515___mcc_h3 = (_source6).cf3;
        let _516_cf3 = _515___mcc_h3;
        let _517_cf2 = _514___mcc_h2;
        let _518_cf1 = _513___mcc_h1;
        let _519_cf0 = _512___mcc_h0;
        let _520_v16;
        _520_v16 = _dafny.Set.fromElements(_495_v0);
        let _521_v17;
        _521_v17 = new _dafny.CodePoint('i'.codePointAt(0));
        let _522_v18;
        _522_v18 = _dafny.Map.Empty.slice().updateUnsafe(_496_v1,_516_cf3);
        let _523_v19;
        _523_v19 = _dafny.Set.fromElements(_518_cf1, new BigNumber((_522_v18).length));
        r0 = new BigNumber(((((_517_cf2) ? (_520_v16) : (_520_v16))).Intersect((_dafny.Set.fromElements(_module.__default.fm7(_521_v17, new BigNumber((_523_v19).length), _519_cf0, globalState), _517_cf2)).Union(_520_v16))).length);
        r1 = (_518_cf1).isEqualTo((_dafny.ZERO).minus(_496_v1));
        let _524_v20;
        let _nw81 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _524_v20 = _nw81;
        let _index54 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_524_v20).length));
        (_524_v20)[_index54] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(961)), ((_525_v17) => function (_526_i1) {
          return _525_v17;
        })(_521_v17));
        let _527_v21;
        _527_v21 = _dafny.Seq.UnicodeFromString("elp");
        let _index55 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_524_v20).length));
        (_524_v20)[_index55] = _dafny.Seq.Concat(_dafny.Seq.Concat(_527_v21, _527_v21), _527_v21);
        if (_517_cf2) {
          let _528_v22;
          let _nw82 = new _module.C0();
          _nw82.__ctor(_519_cf0, _495_v0, _521_v17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(792)), ((_529_v17) => function (_530_i2) {
            return _529_v17;
          })(_521_v17)));
          _528_v22 = _nw82;
          let _531_v23;
          _531_v23 = _module.D2.create_DC7(_517_cf2);
          let _532_v24;
          _532_v24 = _dafny.Map.Empty.slice().updateUnsafe(_528_v22,(_this).f14);
          let _533_v25;
          let _nw83 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.of());
          _533_v25 = _nw83;
          let _534_v26;
          _534_v26 = _dafny.Seq.of(_519_cf0);
          let _index56 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_533_v25).length));
          (_533_v25)[_index56] = _534_v26;
          let _index57 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_533_v25).length));
          let _rhs52 = (((new BigNumber(966)).isEqualTo(_496_v1)) ? (_528_v22) : (_528_v22));
          let _rhs53 = _module.__default.fm11(_module.__default.fm12(_519_cf0, _527_v21, _517_cf2, globalState), _517_cf2, ((_dafny.ZERO).minus(_496_v1)).minus(new BigNumber(810)), _495_v0, globalState);
          let _rhs54 = _532_v24;
          let _rhs55 = _dafny.Seq.update(_534_v26, _module.__default.safeIndex(_516_cf3, new BigNumber((_534_v26).length)), _519_cf0);
          let _lhs27 = _533_v25;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_533_v25).length));
          _528_v22 = _rhs52;
          _531_v23 = _rhs53;
          _532_v24 = _rhs54;
          _lhs27[_lhs28] = _rhs55;
          _519_cf0 = _495_v0;
          let _535_v27;
          _535_v27 = _dafny.Seq.of(_518_cf1);
          let _536_v28;
          _536_v28 = _dafny.Map.Empty.slice().updateUnsafe(_495_v0,_519_cf0);
          let _537_v29;
          _537_v29 = _dafny.Set.fromElements((_524_v20)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_524_v20).length))]);
          let _538_v30;
          let _nw84 = Array((new BigNumber(21)).toNumber());
          _nw84[(_dafny.ZERO).toNumber()] = _518_cf1;
          _nw84[(_dafny.ONE).toNumber()] = ((_535_v27)[_module.__default.safeIndex(_516_cf3, new BigNumber((_535_v27).length))]).plus(_516_cf3);
          _nw84[(new BigNumber(2)).toNumber()] = _518_cf1;
          _nw84[(new BigNumber(3)).toNumber()] = _516_cf3;
          _nw84[(new BigNumber(4)).toNumber()] = _516_cf3;
          _nw84[(new BigNumber(5)).toNumber()] = new BigNumber((_527_v21).length);
          _nw84[(new BigNumber(6)).toNumber()] = _496_v1;
          _nw84[(new BigNumber(7)).toNumber()] = (((((_536_v28).contains(_495_v0)) ? ((_536_v28).get(_495_v0)) : (_519_cf0))) ? (_496_v1) : (_518_cf1));
          _nw84[(new BigNumber(8)).toNumber()] = new BigNumber((_537_v29).length);
          _nw84[(new BigNumber(9)).toNumber()] = (_518_cf1).multipliedBy(_518_cf1);
          _nw84[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt(_516_cf3, _516_cf3);
          _nw84[(new BigNumber(11)).toNumber()] = _518_cf1;
          _nw84[(new BigNumber(12)).toNumber()] = ((_517_cf2) ? (_496_v1) : (new BigNumber((_520_v16).length)));
          _nw84[(new BigNumber(13)).toNumber()] = _518_cf1;
          _nw84[(new BigNumber(14)).toNumber()] = _516_cf3;
          _nw84[(new BigNumber(15)).toNumber()] = (_518_cf1).minus(_496_v1);
          _nw84[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(_518_cf1);
          _nw84[(new BigNumber(17)).toNumber()] = _module.__default.fm8((_dafny.ZERO).minus(_518_cf1), _495_v0, _495_v0, globalState);
          _nw84[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality());
          _nw84[(new BigNumber(19)).toNumber()] = _518_cf1;
          _nw84[(new BigNumber(20)).toNumber()] = _516_cf3;
          _538_v30 = _nw84;
          let _index58 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_538_v30).length));
          (_538_v30)[_index58] = _516_cf3;
          let _539_v31;
          _539_v31 = _dafny.MultiSet.fromElements(_521_v17);
          let _540_v32;
          _540_v32 = _dafny.Map.Empty.slice().updateUnsafe(_521_v17,_517_cf2);
          let _index59 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_538_v30).length));
          let _rhs56 = _module.__default.fm7(_521_v17, (((_539_v31).contains((_module.D1.create_DC4(new _dafny.CodePoint('j'.codePointAt(0)), _496_v1, _521_v17, _516_cf3, true)).dtor_cf10)) ? ((_539_v31).get((_module.D1.create_DC4(new _dafny.CodePoint('j'.codePointAt(0)), _496_v1, _521_v17, _516_cf3, true)).dtor_cf10)) : (new BigNumber(204))), !((_518_cf1).isLessThan(_496_v1)), globalState);
          let _rhs57 = (_516_cf3).isLessThanOrEqualTo((new BigNumber((_540_v32).length)).plus(_516_cf3));
          let _rhs58 = (_dafny.ZERO).minus((_535_v27)[_module.__default.safeIndex(_496_v1, new BigNumber((_535_v27).length))]);
          let _rhs59 = _516_cf3;
          let _lhs29 = _538_v30;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_538_v30).length));
          r1 = _rhs56;
          _495_v0 = _rhs57;
          _lhs29[_lhs30] = _rhs58;
          r2 = _rhs59;
          let _index60 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_538_v30).length));
          (_538_v30)[_index60] = _518_cf1;
          _496_v1 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat((_524_v20)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_524_v20).length))], _527_v21)).length), _496_v1);
        } else {
          let _index61 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_524_v20).length));
          (_524_v20)[_index61] = (_524_v20)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_524_v20).length))];
          r1 = _495_v0;
          _527_v21 = _dafny.Seq.Concat(_527_v21, _dafny.Seq.update(_527_v21, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_516_cf3)).length), new BigNumber((_527_v21).length)), _521_v17));
          let _541_v33;
          let _nw85 = new _module.C0();
          _nw85.__ctor(_519_cf0, (_516_cf3).isEqualTo(_module.__default.fm8(_516_cf3, _517_cf2, !(_517_cf2), globalState)), _521_v17, (_524_v20)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_524_v20).length))]);
          _541_v33 = _nw85;
          let _nw86 = new _module.C0();
          _nw86.__ctor(_495_v0, _495_v0, _module.__default.fm13(globalState), _module.__default.fm14(_496_v1, _516_cf3, _519_cf0, true, globalState));
          _541_v33 = _nw86;
          let _542_v34;
          _542_v34 = _module.D3.create_DC10(_527_v21);
          let _543_v35;
          _543_v35 = _module.D3.create_DC12(_542_v34);
          let _544_v36;
          _544_v36 = _dafny.MultiSet.fromElements(_516_cf3);
          let _545_v37;
          _545_v37 = _dafny.Seq.of(_517_cf2);
          let _546_v38;
          _546_v38 = _dafny.Map.Empty.slice().updateUnsafe(_544_v36,_545_v37);
          let _547_v39;
          _547_v39 = _dafny.Map.Empty.slice().updateUnsafe(_543_v35,_546_v38);
          _547_v39 = (_547_v39).update(_module.D3.create_DC12(_542_v34), _546_v38);
        }
      } else {
        let _548___mcc_h4 = (_source6).cf4;
        let _549___mcc_h5 = (_source6).cf5;
        let _550___mcc_h6 = (_source6).cf6;
        let _551_cf6 = _550___mcc_h6;
        let _552_cf5 = _549___mcc_h5;
        let _553_cf4 = _548___mcc_h4;
        _497_v2 = _497_v2;
        _495_v0 = (_553_cf4).isLessThan(new BigNumber(619));
        (globalState).f1 = (_dafny.ZERO).minus(_553_cf4);
        _module.__default.m0(globalState);
      }
      r1 = _495_v0;
      let _554_v40;
      _554_v40 = new _dafny.CodePoint('f'.codePointAt(0));
      let _555_v41;
      _555_v41 = _dafny.Map.Empty.slice().updateUnsafe(_495_v0,_554_v40);
      let _556_v42;
      let _nw87 = new _module.C0();
      _nw87.__ctor(true, !(_495_v0), (((_555_v41).contains(_495_v0)) ? ((_555_v41).get(_495_v0)) : (_554_v40)), _module.__default.fm14(new BigNumber(-567), _496_v1, _module.__default.fm7(_554_v40, _496_v1, _495_v0, globalState), _495_v0, globalState));
      _556_v42 = _nw87;
      _496_v1 = _496_v1;
      if ((_495_v0) && (_495_v0)) {
        let _557_v43;
        _557_v43 = _dafny.Map.Empty.slice().updateUnsafe(!(_495_v0),_module.__default.fm15(globalState));
        _557_v43 = (_557_v43).Merge(_557_v43);
        if (_495_v0) {
          let _558_v44;
          _558_v44 = _dafny.Set.fromElements(true, true);
          r2 = (_496_v1).multipliedBy(_module.__default.safeModuloInt(new BigNumber((_558_v44).length), _496_v1));
          r0 = _496_v1;
          (globalState).f1 = ((_dafny.ZERO).minus(_496_v1)).minus(_496_v1);
          let _559_v45;
          _559_v45 = _module.D1.create_DC5(_496_v1);
          let _pat_let_tv9 = _496_v1;
          let _pat_let_tv10 = _496_v1;
          let _560_v46;
          _560_v46 = _dafny.Seq.of(_559_v45, function (_pat_let13_0) {
            return function (_561_dt__update__tmp_h0) {
              return function (_pat_let14_0) {
                return function (_562_dt__update_hcf13_h0) {
                  return _module.D1.create_DC5(_562_dt__update_hcf13_h0);
                }(_pat_let14_0);
              }(_pat_let_tv9);
            }(_pat_let13_0);
          }(_559_v45), function (_pat_let15_0) {
            return function (_563_dt__update__tmp_h1) {
              return function (_pat_let16_0) {
                return function (_564_dt__update_hcf13_h1) {
                  return _module.D1.create_DC5(_564_dt__update_hcf13_h1);
                }(_pat_let16_0);
              }(_pat_let_tv10);
            }(_pat_let15_0);
          }(_559_v45));
          let _565_v47;
          _565_v47 = _module.__default.fm16(globalState);
          let _566_v48;
          _566_v48 = _dafny.Map.Empty.slice().updateUnsafe(_496_v1,(_565_v47));
          let _567_v49;
          _567_v49 = _dafny.Seq.of(_module.__default.fm16(globalState), _560_v46);
          let _568_v50;
          _568_v50 = _dafny.Seq.of(_dafny.Seq.Concat(_560_v46, _dafny.Seq.of(_559_v45)), (((_566_v48).contains(_496_v1)) ? ((_566_v48).get(_496_v1)) : (_560_v46)), (_567_v49)[_module.__default.safeIndex(_496_v1, new BigNumber((_567_v49).length))]);
          _568_v50 = _567_v49;
          let _569_v51;
          let _nw88 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _569_v51 = _nw88;
          let _index62 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_569_v51).length));
          (_569_v51)[_index62] = _496_v1;
          let _570_v52;
          _570_v52 = _dafny.Map.Empty.slice().updateUnsafe(_554_v40,_495_v0);
          let _571_v53;
          _571_v53 = _dafny.Map.Empty.slice().updateUnsafe(_495_v0,_495_v0);
          let _index63 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_569_v51).length));
          let _rhs60 = _496_v1;
          let _rhs61 = _module.__default.fm17(_496_v1, globalState);
          let _rhs62 = new _dafny.CodePoint('j'.codePointAt(0));
          let _rhs63 = ((_496_v1).minus(new BigNumber((_571_v53).length))).isEqualTo(_496_v1);
          let _rhs64 = _496_v1;
          let _lhs31 = _569_v51;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_569_v51).length));
          let _lhs33 = globalState;
          _lhs31[_lhs32] = _rhs60;
          _570_v52 = _rhs61;
          _554_v40 = _rhs62;
          _495_v0 = _rhs63;
          _lhs33.f1 = _rhs64;
        } else {
          r1 = _495_v0;
          let _rhs65 = _496_v1;
          let _rhs66 = _496_v1;
          let _lhs34 = globalState;
          let _lhs35 = globalState;
          _lhs34.f1 = _rhs65;
          _lhs35.f1 = _rhs66;
          let _572_v54;
          _572_v54 = _dafny.Seq.of(_495_v0);
          let _573_v55;
          _573_v55 = _dafny.Seq.UnicodeFromString("kossprq");
          let _574_v56;
          _574_v56 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm7(_554_v40, new BigNumber((_572_v54).length), !(_495_v0), globalState),_573_v55);
          _574_v56 = (_574_v56).update(_495_v0, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fhcrvh"), _573_v55));
          r2 = _496_v1;
          let _575_v57;
          let _nw89 = Array((new BigNumber(5)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = new BigNumber(-814);
          _nw89[(_dafny.ONE).toNumber()] = (_496_v1).minus(_496_v1);
          _nw89[(new BigNumber(2)).toNumber()] = _496_v1;
          _nw89[(new BigNumber(3)).toNumber()] = _496_v1;
          _nw89[(new BigNumber(4)).toNumber()] = new BigNumber(208);
          _575_v57 = _nw89;
          let _576_v58;
          _576_v58 = _module.D5.create_DC14(_575_v57);
          _575_v57 = (_576_v58).dtor_cf25;
        }
        r1 = !(true);
        r1 = _495_v0;
        let _577_v59;
        let _nw90 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
        _577_v59 = _nw90;
        let _578_v60;
        _578_v60 = _module.D1.create_DC3(_577_v59);
        let _source7 = _578_v60;
        if (_source7.is_DC4) {
          let _579___mcc_h7 = (_source7).cf8;
          let _580___mcc_h8 = (_source7).cf9;
          let _581___mcc_h9 = (_source7).cf10;
          let _582___mcc_h10 = (_source7).cf11;
          let _583___mcc_h11 = (_source7).cf12;
          let _584_cf12 = _583___mcc_h11;
          let _585_cf11 = _582___mcc_h10;
          let _586_cf10 = _581___mcc_h9;
          let _587_cf9 = _580___mcc_h8;
          let _588_cf8 = _579___mcc_h7;
          let _589_v61;
          let _nw91 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _589_v61 = _nw91;
          let _590_v62;
          _590_v62 = _dafny.Seq.UnicodeFromString("q");
          let _591_v63;
          _591_v63 = _dafny.Map.Empty.slice().updateUnsafe(_590_v62,_584_cf12);
          let _592_v64;
          _592_v64 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_590_v62).length),_495_v0);
          let _593_v65;
          _593_v65 = _dafny.Map.Empty.slice().updateUnsafe(_495_v0,new BigNumber((_592_v64).length));
          let _594_v66;
          _594_v66 = _dafny.MultiSet.fromElements(_556_v42);
          let _595_v67;
          _595_v67 = _dafny.Seq.of(_585_cf11, _587_cf9, _496_v1, (_dafny.ZERO).minus(_496_v1));
          let _596_v68;
          let _nw92 = Array((new BigNumber(22)).toNumber());
          _nw92[(_dafny.ZERO).toNumber()] = _587_cf9;
          _nw92[(_dafny.ONE).toNumber()] = _587_cf9;
          _nw92[(new BigNumber(2)).toNumber()] = _496_v1;
          _nw92[(new BigNumber(3)).toNumber()] = _496_v1;
          _nw92[(new BigNumber(4)).toNumber()] = _585_cf11;
          _nw92[(new BigNumber(5)).toNumber()] = _585_cf11;
          _nw92[(new BigNumber(6)).toNumber()] = _585_cf11;
          _nw92[(new BigNumber(7)).toNumber()] = _585_cf11;
          _nw92[(new BigNumber(8)).toNumber()] = _585_cf11;
          _nw92[(new BigNumber(9)).toNumber()] = _587_cf9;
          _nw92[(new BigNumber(10)).toNumber()] = new BigNumber(((_591_v63).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(562)), ((_597_cf10) => function (_598_i3) {
            return _597_cf10;
          })(_586_cf10)), _584_cf12)).length);
          _nw92[(new BigNumber(11)).toNumber()] = _587_cf9;
          _nw92[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_590_v62).length));
          _nw92[(new BigNumber(13)).toNumber()] = _496_v1;
          _nw92[(new BigNumber(14)).toNumber()] = _585_cf11;
          _nw92[(new BigNumber(15)).toNumber()] = new BigNumber((_590_v62).length);
          _nw92[(new BigNumber(16)).toNumber()] = _496_v1;
          _nw92[(new BigNumber(17)).toNumber()] = _496_v1;
          _nw92[(new BigNumber(18)).toNumber()] = new BigNumber((_593_v65).length);
          _nw92[(new BigNumber(19)).toNumber()] = new BigNumber((_594_v66).cardinality());
          _nw92[(new BigNumber(20)).toNumber()] = _585_cf11;
          _nw92[(new BigNumber(21)).toNumber()] = new BigNumber((_595_v67).length);
          _596_v68 = _nw92;
          let _599_v69;
          _599_v69 = _dafny.Map.Empty.slice().updateUnsafe(_589_v61,_596_v68);
          _599_v69 = (_599_v69).update(_589_v61, _596_v68);
          let _index64 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_596_v68).length));
          (_596_v68)[_index64] = _496_v1;
          let _index65 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_596_v68).length));
          (_596_v68)[_index65] = _496_v1;
          _586_cf10 = _586_cf10;
          let _600_v70;
          _600_v70 = _dafny.Map.Empty.slice().updateUnsafe(_496_v1,((_495_v0) ? (_496_v1) : (new BigNumber(659))));
          _600_v70 = (_600_v70).update(_585_cf11, (_496_v1).multipliedBy(_587_cf9));
        } else if (_source7.is_DC5) {
          let _601___mcc_h12 = (_source7).cf13;
          let _602_cf13 = _601___mcc_h12;
          let _603_v71;
          let _nw93 = Array((new BigNumber(28)).toNumber()).fill(_dafny.MultiSet.Empty);
          _603_v71 = _nw93;
          _603_v71 = _603_v71;
          let _604_v72;
          _604_v72 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_602_cf13, _602_cf13),(_this).f14);
          let _605_v73;
          _605_v73 = _dafny.Seq.of(!(_495_v0), !(_495_v0));
          let _606_v74;
          _606_v74 = _dafny.Seq.of(new BigNumber(954), new BigNumber((_605_v73).length));
          let _607_v75;
          _607_v75 = _dafny.Map.Empty.slice().updateUnsafe(_554_v40,_495_v0);
          let _608_v76;
          _608_v76 = _module.D3.create_DC11(_495_v0, new BigNumber((_606_v74).length), _dafny.Set.fromElements(_495_v0, (((_607_v75).contains(_554_v40)) ? ((_607_v75).get(_554_v40)) : (_495_v0))));
          _604_v72 = (_604_v72).update((_608_v76).dtor_cf21, (_this).f14);
          let _609_v77;
          let _init11 = ((_610_v40) => function (_611_i4) {
            return _module.__default.safeModuloInt(_611_i4, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(169)), ((_612_v40) => function (_613_i5) {
              return _612_v40;
            })(_610_v40))).length)));
          })(_554_v40);
          let _nw94 = Array((new BigNumber(4)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw94.length); _i0_11++) {
            _nw94[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _609_v77 = _nw94;
          let _614_v78;
          _614_v78 = _dafny.Map.Empty.slice().updateUnsafe(_609_v77,_495_v0);
          _614_v78 = (_614_v78).update(_609_v77, _495_v0);
          let _615_v79;
          _615_v79 = _dafny.Seq.UnicodeFromString("l");
          _615_v79 = _615_v79;
        } else {
          let _616___mcc_h13 = (_source7).cf7;
          let _617_cf7 = _616___mcc_h13;
          let _618_v80;
          _618_v80 = _dafny.Seq.UnicodeFromString("egghuib");
          _618_v80 = _618_v80;
          let _619_v81;
          _619_v81 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_618_v80).length),_495_v0);
          let _620_v82;
          _620_v82 = _dafny.Seq.of(_496_v1, new BigNumber((_619_v81).length));
          _620_v82 = _620_v82;
          let _621_v83;
          let _nw95 = new _module.C0();
          _nw95.__ctor(_495_v0, false, _554_v40, _618_v80);
          _621_v83 = _nw95;
          _621_v83 = _621_v83;
          _495_v0 = (_496_v1).isEqualTo(_496_v1);
        }
      } else {
        let _622_v84;
        _622_v84 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(_496_v1)).length));
        let _623_v85;
        _623_v85 = _dafny.Seq.of(_495_v0, _495_v0);
        let _624_v86;
        _624_v86 = _dafny.Map.Empty.slice().updateUnsafe(_495_v0,_495_v0);
        let _625_v87;
        _625_v87 = _dafny.Seq.UnicodeFromString("jtb");
        let _626_v88;
        _626_v88 = _dafny.MultiSet.fromElements(_623_v85, _module.__default.fm18(_496_v1, (((_624_v86).contains(_495_v0)) ? ((_624_v86).get(_495_v0)) : (true)), globalState), _dafny.Seq.update(_623_v85, _module.__default.safeIndex(new BigNumber((_module.__default.fm18(new BigNumber((_625_v87).length), _495_v0, globalState)).length), new BigNumber((_623_v85).length)), false), _623_v85, _623_v85);
        let _627_v89;
        _627_v89 = _module.D5.create_DC15(((_495_v0) ? ((_this).f14) : ((_this).f14)), _495_v0, _622_v84, new BigNumber((_626_v88).cardinality()));
        _627_v89 = _627_v89;
        let _628_v90;
        let _nw96 = new _module.C1();
        _nw96.__ctor(_496_v1, true, _495_v0, true, _554_v40, _625_v87);
        _628_v90 = _nw96;
        let _629_v91;
        _629_v91 = _dafny.Map.Empty.slice().updateUnsafe(_628_v90,(_628_v90).f7);
        let _630_v92;
        _630_v92 = _dafny.MultiSet.fromElements((_628_v90).f7);
        let _631_v93;
        let _nw97 = Array((new BigNumber(20)).toNumber());
        _nw97[(_dafny.ZERO).toNumber()] = (((_629_v91).contains(_628_v90)) ? ((_629_v91).get(_628_v90)) : (true));
        _nw97[(_dafny.ONE).toNumber()] = _495_v0;
        _nw97[(new BigNumber(2)).toNumber()] = (_628_v90).f6;
        _nw97[(new BigNumber(3)).toNumber()] = !(_module.__default.fm7(_554_v40, _496_v1, !((_628_v90).f9), globalState));
        _nw97[(new BigNumber(4)).toNumber()] = true;
        _nw97[(new BigNumber(5)).toNumber()] = (_628_v90).f7;
        _nw97[(new BigNumber(6)).toNumber()] = (_628_v90).f9;
        _nw97[(new BigNumber(7)).toNumber()] = (_628_v90).f9;
        _nw97[(new BigNumber(8)).toNumber()] = true;
        _nw97[(new BigNumber(9)).toNumber()] = (_628_v90).f6;
        _nw97[(new BigNumber(10)).toNumber()] = !_dafny.areEqual((_628_v90).f5, _625_v87);
        _nw97[(new BigNumber(11)).toNumber()] = !((_628_v90).f7);
        _nw97[(new BigNumber(12)).toNumber()] = (_dafny.MultiSet.fromElements((_628_v90).f9)).IsProperSubsetOf(_630_v92);
        _nw97[(new BigNumber(13)).toNumber()] = ((_628_v90).f9) === ((_628_v90).f9);
        _nw97[(new BigNumber(14)).toNumber()] = _495_v0;
        _nw97[(new BigNumber(15)).toNumber()] = (_628_v90).f6;
        _nw97[(new BigNumber(16)).toNumber()] = (_628_v90).f7;
        _nw97[(new BigNumber(17)).toNumber()] = _495_v0;
        _nw97[(new BigNumber(18)).toNumber()] = false;
        _nw97[(new BigNumber(19)).toNumber()] = !((_628_v90).f9) || ((_628_v90).f7);
        _631_v93 = _nw97;
        _631_v93 = (_this).f14;
        _625_v87 = _625_v87;
        let _632_v94;
        _632_v94 = _dafny.Map.Empty.slice().updateUnsafe((_628_v90).f8,new BigNumber((_624_v86).length));
        let _633_v95;
        _633_v95 = _dafny.Set.fromElements((_628_v90).f9, (_628_v90).f7, (_628_v90).f7);
        let _634_v96;
        _634_v96 = _dafny.Seq.of((((_632_v94).contains(_496_v1)) ? ((_632_v94).get(_496_v1)) : ((_628_v90).f8)), _module.__default.safeModuloInt((_628_v90).f8, (_628_v90).f8), _module.__default.safeModuloInt((_628_v90).f8, new BigNumber((_633_v95).length)), (_dafny.ZERO).minus(_496_v1), (_628_v90).f8);
        _634_v96 = _dafny.Seq.of(_module.__default.safeModuloInt(_496_v1, new BigNumber(977)));
        if (((_628_v90).f8).isEqualTo((_628_v90).fm1(_495_v0, _dafny.Map.Empty.slice().updateUnsafe((_628_v90).f7,(_628_v90).f6), _495_v0, globalState))) {
          r1 = (_628_v90).f9;
          let _635_v97;
          _635_v97 = _dafny.Seq.of((_628_v90).f5, _dafny.Seq.UnicodeFromString("nravm"), _dafny.Seq.UnicodeFromString("uqwbaqax"));
          let _636_v98;
          let _nw98 = new _module.C0();
          _nw98.__ctor(_module.__default.fm7(_628_v90.f4, (_628_v90).f8, false, globalState), true, _628_v90.f4, (_635_v97)[_module.__default.safeIndex(new BigNumber(276), new BigNumber((_635_v97).length))]);
          _636_v98 = _nw98;
          let _637_v99;
          _637_v99 = _module.D1.create_DC4(new _dafny.CodePoint('n'.codePointAt(0)), (_628_v90).f8, _554_v40, (_dafny.ZERO).minus((_628_v90).f8), !((_628_v90).f9) || (true));
          let _638_v100;
          let _nw99 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _638_v100 = _nw99;
          let _index66 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_638_v100).length));
          (_638_v100)[_index66] = (_628_v90).f5;
          let _639_v101;
          _639_v101 = _dafny.Map.Empty.slice().updateUnsafe((_628_v90).f8,(_628_v90).f6);
          let _640_v102;
          _640_v102 = _dafny.Set.fromElements(_639_v101);
          let _pat_let_tv11 = _496_v1;
          let _pat_let_tv12 = _628_v90;
          let _pat_let_tv13 = _628_v90;
          let _pat_let_tv14 = _628_v90;
          let _index67 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_638_v100).length));
          let _rhs67 = !(_module.__default.fm7(_554_v40, new BigNumber((_dafny.Set.fromElements((_628_v90).f8, (_628_v90).f8, new BigNumber((_640_v102).length))).length), _495_v0, globalState)) || ((_628_v90).f7);
          let _rhs68 = _496_v1;
          let _rhs69 = function (_pat_let17_0) {
            return function (_641_dt__update__tmp_h2) {
              return function (_pat_let18_0) {
                return function (_642_dt__update_hcf9_h0) {
                  return function (_pat_let19_0) {
                    return function (_643_dt__update_hcf8_h0) {
                      return function (_pat_let20_0) {
                        return function (_644_dt__update_hcf10_h0) {
                          return function (_pat_let21_0) {
                            return function (_645_dt__update_hcf11_h0) {
                              return _module.D1.create_DC4(_643_dt__update_hcf8_h0, _642_dt__update_hcf9_h0, _644_dt__update_hcf10_h0, _645_dt__update_hcf11_h0, (_641_dt__update__tmp_h2).dtor_cf12);
                            }(_pat_let21_0);
                          }((_pat_let_tv14).f8);
                        }(_pat_let20_0);
                      }(_pat_let_tv13.f4);
                    }(_pat_let19_0);
                  }(_pat_let_tv12.f4);
                }(_pat_let18_0);
              }(_pat_let_tv11);
            }(_pat_let17_0);
          }(p0);
          let _rhs70 = _dafny.Seq.Concat(_625_v87, _dafny.Seq.Concat(_dafny.Seq.update(_625_v87, _module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus((_628_v90).f8)), new BigNumber((_625_v87).length)), _554_v40), _dafny.Seq.UnicodeFromString("agc")));
          let _rhs71 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_623_v85, _dafny.Seq.of(true, (_628_v90).f9))).length));
          let _lhs36 = _638_v100;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_638_v100).length));
          let _lhs38 = globalState;
          _495_v0 = _rhs67;
          _496_v1 = _rhs68;
          _637_v99 = _rhs69;
          _lhs36[_lhs37] = _rhs70;
          _lhs38.f1 = _rhs71;
          let _646_v103;
          let _nw100 = new _module.C0();
          _nw100.__ctor((_628_v90).f9, (_628_v90).f9, _628_v90.f4, (_628_v90).f5);
          _646_v103 = _nw100;
          (globalState).f1 = (_628_v90).f8;
        } else {
          let _index68 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_631_v93).length));
          (_631_v93)[_index68] = !(_dafny.Map.Empty.slice().updateUnsafe((_628_v90).f7,(_628_v90).f8)).contains(true);
          let _647_v104;
          let _nw101 = Array((new BigNumber(11)).toNumber()).fill([]);
          _647_v104 = _nw101;
          let _648_v105;
          let _nw102 = new _module.C1();
          _nw102.__ctor((_628_v90).f8, (_628_v90).f6, (_628_v90).f7, (_628_v90).f9, _628_v90.f4, (_628_v90).f5);
          _648_v105 = _nw102;
          let _649_v106;
          let _nw103 = Array((new BigNumber(7)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = _648_v105;
          _nw103[(_dafny.ONE).toNumber()] = _648_v105;
          _nw103[(new BigNumber(2)).toNumber()] = _648_v105;
          _nw103[(new BigNumber(3)).toNumber()] = _648_v105;
          _nw103[(new BigNumber(4)).toNumber()] = _648_v105;
          _nw103[(new BigNumber(5)).toNumber()] = _648_v105;
          _nw103[(new BigNumber(6)).toNumber()] = _648_v105;
          _649_v106 = _nw103;
          let _index69 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_647_v104).length));
          (_647_v104)[_index69] = _649_v106;
          let _650_v107;
          _650_v107 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sd"), (_628_v90).f5);
          let _index70 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_631_v93).length));
          let _index71 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_647_v104).length));
          let _rhs72 = !(!((_650_v107).IsDisjointFrom(_650_v107)));
          let _rhs73 = (_628_v90).f9;
          let _rhs74 = _649_v106;
          let _lhs39 = _631_v93;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_631_v93).length));
          let _lhs41 = _647_v104;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_647_v104).length));
          _lhs39[_lhs40] = _rhs72;
          _495_v0 = _rhs73;
          _lhs41[_lhs42] = _rhs74;
          r1 = !((_628_v90).f6);
          _495_v0 = !(!((((_628_v90).f7) ? (true) : ((_628_v90).f6))) || ((_628_v90).f7));
          let _651_v108;
          _651_v108 = _dafny.Map.Empty.slice().updateUnsafe((_633_v95).Union(_633_v95),_628_v90.f4);
          _651_v108 = _651_v108;
          r2 = (_628_v90).f8;
        }
      }
      let _652_i6;
      _652_i6 = _dafny.ZERO;
      L3: {
        while (!(_495_v0)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_652_i6)) {
              break L3;
            }
            _652_i6 = (_652_i6).plus(_dafny.ONE);
            _495_v0 = !(_495_v0);
            let _653_v109;
            _653_v109 = _dafny.Map.Empty.slice().updateUnsafe(_496_v1,_495_v0);
            let _654_v110;
            _654_v110 = _dafny.Seq.UnicodeFromString("dvhbgf");
            let _655_v111;
            let _nw104 = new _module.C0();
            _nw104.__ctor(_495_v0, !((((_653_v109).contains(_496_v1)) ? ((_653_v109).get(_496_v1)) : (!(_495_v0)))) || (_495_v0), _module.__default.fm13(globalState), _654_v110);
            _655_v111 = _nw104;
            let _index72 = _module.__default.safeIndex(new BigNumber(975), new BigNumber(((_this).f14).length));
            ((_this).f14)[_index72] = !(_495_v0);
            let _index73 = _module.__default.safeIndex(new BigNumber(975), new BigNumber(((_this).f14).length));
            ((_this).f14)[_index73] = _495_v0;
            let _656_v112;
            _656_v112 = _dafny.Map.Empty.slice().updateUnsafe(true,((_this).f14)[_module.__default.safeIndex(new BigNumber(975), new BigNumber(((_this).f14).length))]);
            let _657_v113;
            let _nw105 = new _module.C1();
            _nw105.__ctor(new BigNumber(873), ((((_this).f14)[_module.__default.safeIndex(new BigNumber(975), new BigNumber(((_this).f14).length))]) ? (!((((_656_v112).contains(_495_v0)) ? ((_656_v112).get(_495_v0)) : (_495_v0)))) : (_495_v0)), _495_v0, ((_this).f14)[_module.__default.safeIndex(new BigNumber(975), new BigNumber(((_this).f14).length))], _554_v40, _654_v110);
            _657_v113 = _nw105;
          }
        }
      }
      r0 = _496_v1;
      r1 = _495_v0;
      r2 = _496_v1;
      return [r0, r1, r2];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f4 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f8 = _dafny.ZERO;
      this._f9 = false;
      this._f6 = false;
      this._f7 = false;
      this._f5 = _dafny.Seq.UnicodeFromString("");
      this.f12 = false;
      this._f13 = false;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    set f4(value) {
      let _this = this;
      _this._f4 = value;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f12, f13, f8, f9, f6, f7, f4, f5) {
      let _this = this;
      (_this).f12 = f12;
      (_this)._f13 = f13;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f8;
    };
    fm2(globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(559)), function (_658_i0) {
        return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(329)), function (_659_i1) {
          return (_this).f8;
        })).length);
      }), (_module.D0.create_DC2((_this).f8, !(true), _dafny.Seq.Create(_module.__default.abs(new BigNumber(951)), function (_660_i2) {
  return (_this).f8;
}))).dtor_cf6);
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f8;
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = [];
      let _661_v0;
      _661_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f12);
      let _662_v1;
      let _nw106 = new _module.C0();
      _nw106.__ctor(!(_661_v0).contains(_this.f12), (_this).f7, _this.f4, (_this).f5);
      _662_v1 = _nw106;
      let _663_v2;
      let _nw107 = Array((new BigNumber(5)).toNumber()).fill(false);
      _663_v2 = _nw107;
      let _664_v3;
      _664_v3 = _dafny.Seq.of((_this).f9);
      let _665_v4;
      _665_v4 = _dafny.Seq.of(_664_v3);
      let _666_v5;
      let _nw108 = Array((new BigNumber(16)).toNumber());
      _nw108[(_dafny.ZERO).toNumber()] = _664_v3;
      _nw108[(_dafny.ONE).toNumber()] = _dafny.Seq.of(p0);
      _nw108[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(false, p3, (_this).f9);
      _nw108[(new BigNumber(3)).toNumber()] = _664_v3;
      _nw108[(new BigNumber(4)).toNumber()] = _664_v3;
      _nw108[(new BigNumber(5)).toNumber()] = (_665_v4)[_module.__default.safeIndex((_this).f8, new BigNumber((_665_v4).length))];
      _nw108[(new BigNumber(6)).toNumber()] = _dafny.Seq.of((_this).f7);
      _nw108[(new BigNumber(7)).toNumber()] = _664_v3;
      _nw108[(new BigNumber(8)).toNumber()] = _664_v3;
      _nw108[(new BigNumber(9)).toNumber()] = _664_v3;
      _nw108[(new BigNumber(10)).toNumber()] = _664_v3;
      _nw108[(new BigNumber(11)).toNumber()] = _664_v3;
      _nw108[(new BigNumber(12)).toNumber()] = _664_v3;
      _nw108[(new BigNumber(13)).toNumber()] = _664_v3;
      _nw108[(new BigNumber(14)).toNumber()] = _664_v3;
      _nw108[(new BigNumber(15)).toNumber()] = _664_v3;
      _666_v5 = _nw108;
      let _667_v6;
      _667_v6 = _module.D1.create_DC3(_666_v5);
      let _668_v7;
      let _nw109 = Array((new BigNumber(25)).toNumber());
      _nw109[(_dafny.ZERO).toNumber()] = _667_v6;
      _nw109[(_dafny.ONE).toNumber()] = _667_v6;
      _nw109[(new BigNumber(2)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(3)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(4)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(5)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(6)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(7)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(8)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(9)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(10)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(11)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(12)).toNumber()] = _module.D1.create_DC3(_666_v5);
      _nw109[(new BigNumber(13)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(14)).toNumber()] = _module.D1.create_DC3(_666_v5);
      _nw109[(new BigNumber(15)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(16)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(17)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(18)).toNumber()] = _module.D1.create_DC3(_666_v5);
      _nw109[(new BigNumber(19)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(20)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(21)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(22)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(23)).toNumber()] = _667_v6;
      _nw109[(new BigNumber(24)).toNumber()] = _module.D1.create_DC3(_666_v5);
      _668_v7 = _nw109;
      let _669_v8;
      let _nw110 = new _module.C2();
      _nw110.__ctor(_663_v2, _668_v7);
      _669_v8 = _nw110;
      (_this).f4 = _this.f4;
      let _670_v9;
      let _init12 = ((_671_v3) => function (_672_i0) {
        return (_672_i0).minus(new BigNumber((_671_v3).length));
      })(_664_v3);
      let _nw111 = Array((new BigNumber(7)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw111.length); _i0_12++) {
        _nw111[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _670_v9 = _nw111;
      let _index74 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_670_v9).length));
      (_670_v9)[_index74] = (_dafny.ZERO).minus((_this).f8);
      let _673_v10;
      _673_v10 = _dafny.MultiSet.fromElements(new BigNumber((p2).cardinality()));
      let _674_v11;
      _674_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_673_v10);
      let _675_v12;
      _675_v12 = _dafny.MultiSet.fromElements((((_674_v11).contains((_this).f6)) ? ((_674_v11).get((_this).f6)) : ((_673_v10).update((_this).f8, _module.__default.abs((_this).f8)))), _673_v10);
      let _index75 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_670_v9).length));
      (_670_v9)[_index75] = (_dafny.ZERO).minus(((true) ? (((_this).f8).plus((_this).f8)) : (new BigNumber((_675_v12).cardinality()))));
      (_this).f12 = (_this).f9;
      let _hi0 = new BigNumber((_dafny.MultiSet.fromElements(p1)).cardinality());
      for (let _676_i1 = (_670_v9)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_670_v9).length))]; _676_i1.isLessThan(_hi0); _676_i1 = _676_i1.plus(_dafny.ONE)) {
        let _677_v13;
        let _nw112 = new _module.C0();
        _nw112.__ctor(p3, (_this).f13, _this.f4, (_this).f5);
        _677_v13 = _nw112;
        _673_v10 = (_dafny.MultiSet.fromElements(new BigNumber(490), (_670_v9)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_670_v9).length))])).Intersect(_673_v10);
        r0 = _676_i1;
        (_this).f12 = p1;
      }
      let _678_v14;
      _678_v14 = _dafny.Set.fromElements((_670_v9)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_670_v9).length))], (_this).f8);
      r0 = ((new BigNumber(362)).multipliedBy((_dafny.ZERO).minus(new BigNumber((_678_v14).length)))).plus(new BigNumber(-419));
      let _679_v15;
      _679_v15 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f8);
      r1 = ((!(_dafny.Seq.IsPrefixOf((_this).f5, _dafny.Seq.UnicodeFromString("n")))) ? (new BigNumber(((_this).f5).length)) : (_module.__default.safeModuloInt((_670_v9)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_670_v9).length))], (((_679_v15).contains(p0)) ? ((_679_v15).get(p0)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber(704))).length))))));
      let _680_v16;
      _680_v16 = _dafny.Seq.of(_670_v9, _670_v9);
      let _681_v17;
      let _nw113 = Array((new BigNumber(20)).toNumber());
      _nw113[(_dafny.ZERO).toNumber()] = _670_v9;
      _nw113[(_dafny.ONE).toNumber()] = _670_v9;
      _nw113[(new BigNumber(2)).toNumber()] = (_680_v16)[_module.__default.safeIndex((_670_v9)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_670_v9).length))], new BigNumber((_680_v16).length))];
      _nw113[(new BigNumber(3)).toNumber()] = (_680_v16)[_module.__default.safeIndex((_670_v9)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_670_v9).length))], new BigNumber((_680_v16).length))];
      _nw113[(new BigNumber(4)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(5)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(6)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(7)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(8)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(9)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(10)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(11)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(12)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(13)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(14)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(15)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(16)).toNumber()] = (((_this).f9) ? (_670_v9) : (_670_v9));
      _nw113[(new BigNumber(17)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(18)).toNumber()] = _670_v9;
      _nw113[(new BigNumber(19)).toNumber()] = _670_v9;
      _681_v17 = _nw113;
      r2 = _681_v17;
      return [r0, r1, r2];
    }
    m5(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let _682_v0;
      _682_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,((_this).f8).plus(new BigNumber(678)));
      _682_v0 = (_682_v0).update((_this).f8, _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f8), (_this).f8));
      let _index76 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((p2).length));
      (p2)[_index76] = p0;
      let _683_v1;
      let _nw114 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _683_v1 = _nw114;
      let _index77 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_683_v1).length));
      (_683_v1)[_index77] = _dafny.Seq.Concat((_this).f5, (_this).f5);
      let _index78 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((p2).length));
      let _index79 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_683_v1).length));
      let _rhs75 = (_this).f8;
      let _rhs76 = p0;
      let _rhs77 = (_this).f5;
      let _lhs43 = globalState;
      let _lhs44 = p2;
      let _lhs45 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((p2).length));
      let _lhs46 = _683_v1;
      let _lhs47 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_683_v1).length));
      _lhs43.f1 = _rhs75;
      _lhs44[_lhs45] = _rhs76;
      _lhs46[_lhs47] = _rhs77;
      let _684_v2;
      _684_v2 = _dafny.Seq.of((_this).f8, (_this).f8);
      let _685_v3;
      _685_v3 = _dafny.Seq.of(!(!(p0)) || ((_this).f6), _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_684_v2, _module.__default.safeIndex(new BigNumber(101), new BigNumber((_684_v2).length)), (_this).f8), _684_v2), (_this).f13);
      let _index80 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((p2).length));
      (p2)[_index80] = (_685_v3)[_module.__default.safeIndex((_this).f8, new BigNumber((_685_v3).length))];
      let _hi1 = ((_this).f8).plus((_this).f8);
      for (let _686_i0 = (_this).f8; _686_i0.isLessThan(_hi1); _686_i0 = _686_i0.plus(_dafny.ONE)) {
        _685_v3 = _685_v3;
        let _687_v4;
        let _nw115 = Array((new BigNumber(24)).toNumber());
        _687_v4 = _nw115;
        let _688_v5;
        _688_v5 = _module.D0.create_DC2((_this).f8, _module.__default.fm7(p3, _686_i0, false, globalState), _dafny.Seq.of(_686_i0));
        let _689_v6;
        _689_v6 = _dafny.Map.Empty.slice().updateUnsafe(_688_v5,(_this).f8);
        let _690_v7;
        let _nw116 = new _module.C1();
        _nw116.__ctor(new BigNumber((_689_v6).length), false, (_this).f13, _this.f12, p3, (_this).f5);
        _690_v7 = _nw116;
        let _index81 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_687_v4).length));
        (_687_v4)[_index81] = _690_v7;
        let _index82 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_687_v4).length));
        (_687_v4)[_index82] = _690_v7;
        (globalState).f1 = new BigNumber((_dafny.Seq.update((_this).f5, _module.__default.safeIndex(((true) ? ((_this).f8) : (new BigNumber(((_683_v1)[_module.__default.safeIndex(new BigNumber(685), new BigNumber((_683_v1).length))]).length))), new BigNumber(((_this).f5).length)), p3)).length);
        let _691_v8;
        _691_v8 = _dafny.Set.fromElements(_685_v3);
        let _692_v9;
        _692_v9 = _dafny.Set.fromElements((_this).f6, _this.f12, (_this).f9);
        let _693_v10;
        _693_v10 = _module.D3.create_DC11((_this).f9, new BigNumber((_691_v8).length), _692_v9);
        if (!((_693_v10).dtor_cf20)) {
          (globalState).f1 = _686_i0;
          (globalState).f1 = new BigNumber(275);
          (globalState).f1 = (_686_i0).minus(new BigNumber((_684_v2).length));
          let _694_v11;
          let _nw117 = Array((new BigNumber(28)).toNumber()).fill([]);
          _694_v11 = _nw117;
          let _695_v12;
          let _nw118 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _695_v12 = _nw118;
          let _index83 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_694_v11).length));
          (_694_v11)[_index83] = _695_v12;
          let _index84 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_694_v11).length));
          (_694_v11)[_index84] = _695_v12;
          (globalState).f1 = (((_686_i0).isEqualTo(_686_i0)) ? ((_this).f8) : ((_this).f8));
        } else {
          let _696_v13;
          _696_v13 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f8).minus((_this).f8),(_this).f5);
          let _697_v14;
          let _nw119 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _697_v14 = _nw119;
          let _698_v15;
          _698_v15 = _dafny.MultiSet.fromElements(_697_v14, _697_v14, _697_v14);
          (globalState).f1 = new BigNumber(((((_696_v13).contains((((_698_v15).contains(_697_v14)) ? ((_698_v15).get(_697_v14)) : ((_this).f8)))) ? ((_696_v13).get((((_698_v15).contains(_697_v14)) ? ((_698_v15).get(_697_v14)) : ((_this).f8)))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(291)), function (_699_i1) {
            return _this.f4;
          })))).length);
          let _index85 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_697_v14).length));
          (_697_v14)[_index85] = _module.__default.safeDivisionInt((_693_v10).dtor_cf21, _module.__default.fm8(new BigNumber(((_this).f5).length), (_this).f9, true, globalState));
          let _700_v16;
          _700_v16 = _dafny.MultiSet.fromElements((_this).f13);
          let _701_v17;
          _701_v17 = _module.D5.create_DC15(p2, p0, _module.__default.fm24(_686_i0, true, p0, new BigNumber((_700_v16).cardinality()), globalState), new BigNumber(69));
          let _702_v18;
          _702_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_701_v17).dtor_cf29);
          let _index86 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_697_v14).length));
          let _index87 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_683_v1).length));
          let _rhs78 = (_this).f8;
          let _rhs79 = (_this).f8;
          let _rhs80 = (_702_v18).Merge((_702_v18).Merge(_702_v18));
          let _rhs81 = (_dafny.ZERO).minus(((_684_v2)[_module.__default.safeIndex(new BigNumber(((_683_v1)[_module.__default.safeIndex(new BigNumber(685), new BigNumber((_683_v1).length))]).length), new BigNumber((_684_v2).length))]).plus(_686_i0));
          let _rhs82 = _dafny.Seq.update(_dafny.Seq.Concat((_683_v1)[_module.__default.safeIndex(new BigNumber(685), new BigNumber((_683_v1).length))], (_this).f5), _module.__default.safeIndex(_686_i0, new BigNumber((_dafny.Seq.Concat((_683_v1)[_module.__default.safeIndex(new BigNumber(685), new BigNumber((_683_v1).length))], (_this).f5)).length)), p3);
          let _lhs48 = globalState;
          let _lhs49 = _697_v14;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_697_v14).length));
          let _lhs51 = globalState;
          let _lhs52 = _683_v1;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_683_v1).length));
          _lhs48.f1 = _rhs78;
          _lhs49[_lhs50] = _rhs79;
          _702_v18 = _rhs80;
          _lhs51.f1 = _rhs81;
          _lhs52[_lhs53] = _rhs82;
          let _703_v19;
          _703_v19 = _dafny.Map.Empty.slice().updateUnsafe(_686_i0,p3);
          _703_v19 = (_703_v19).update(new BigNumber(-494), ((p0) ? (_this.f4) : (p3)));
          let _index88 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((p2).length));
          (p2)[_index88] = false;
          (globalState).f1 = ((_dafny.ZERO).minus((_this).f8)).multipliedBy((_697_v14)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_697_v14).length))]);
        }
      }
      let _704_v20;
      let _nw120 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
      _704_v20 = _nw120;
      let _index89 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_704_v20).length));
      (_704_v20)[_index89] = (_this).f8;
      let _705_v21;
      let _init13 = function (_706_i2) {
        return _dafny.MultiSet.fromElements(true, _this.f12, true, false, !((_this).f9));
      };
      let _nw121 = Array((new BigNumber(24)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw121.length); _i0_13++) {
        _nw121[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _705_v21 = _nw121;
      let _707_v22;
      _707_v22 = _dafny.MultiSet.fromElements((_this).f8);
      let _708_v23;
      _708_v23 = _module.D0.create_DC1((_this).f9, new BigNumber((_707_v22).cardinality()), (_this).f7, new BigNumber(-704));
      let _index90 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_683_v1).length));
      let _index91 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((p2).length));
      let _index92 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_704_v20).length));
      let _rhs83 = ((((_708_v23).dtor_cf1).isEqualTo((_this).f8)) ? (_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("a"), _dafny.Seq.UnicodeFromString("dbviiq"))) : ((_this).f5));
      let _rhs84 = (_this).f13;
      let _rhs85 = _module.__default.safeModuloInt((_this).f8, _module.__default.safeModuloInt((_this).f8, (_this).f8));
      let _rhs86 = _705_v21;
      let _lhs54 = _683_v1;
      let _lhs55 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_683_v1).length));
      let _lhs56 = p2;
      let _lhs57 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((p2).length));
      let _lhs58 = _704_v20;
      let _lhs59 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_704_v20).length));
      _lhs54[_lhs55] = _rhs83;
      _lhs56[_lhs57] = _rhs84;
      _lhs58[_lhs59] = _rhs85;
      _705_v21 = _rhs86;
      let _709_v24;
      let _nw122 = new _module.C1();
      _nw122.__ctor((((_682_v0).contains((_this).f8)) ? ((_682_v0).get((_this).f8)) : (new BigNumber(444))), (_this).f13, true, (_this).f6, _this.f4, (_683_v1)[_module.__default.safeIndex(new BigNumber(685), new BigNumber((_683_v1).length))]);
      _709_v24 = _nw122;
      let _710_v25;
      _710_v25 = _dafny.Map.Empty.slice().updateUnsafe((p2)[_module.__default.safeIndex(new BigNumber(570), new BigNumber((p2).length))],_709_v24);
      let _711_v26;
      let _nw123 = Array((new BigNumber(14)).toNumber());
      _nw123[(_dafny.ZERO).toNumber()] = _710_v25;
      _nw123[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_709_v24);
      _nw123[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_709_v24).f6,_709_v24);
      _nw123[(new BigNumber(3)).toNumber()] = _710_v25;
      _nw123[(new BigNumber(4)).toNumber()] = _710_v25;
      _nw123[(new BigNumber(5)).toNumber()] = _710_v25;
      _nw123[(new BigNumber(6)).toNumber()] = _710_v25;
      _nw123[(new BigNumber(7)).toNumber()] = _710_v25;
      _nw123[(new BigNumber(8)).toNumber()] = _710_v25;
      _nw123[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_709_v24).f7,_709_v24);
      _nw123[(new BigNumber(10)).toNumber()] = (_710_v25).update((_this).f9, _709_v24);
      _nw123[(new BigNumber(11)).toNumber()] = _710_v25;
      _nw123[(new BigNumber(12)).toNumber()] = _710_v25;
      _nw123[(new BigNumber(13)).toNumber()] = (_710_v25).update((p2)[_module.__default.safeIndex(new BigNumber(570), new BigNumber((p2).length))], _709_v24);
      _711_v26 = _nw123;
      let _712_v27;
      _712_v27 = _dafny.Map.Empty.slice().updateUnsafe(_711_v26,(_685_v3)[_module.__default.safeIndex((_this).f8, new BigNumber((_685_v3).length))]);
      (_this).f12 = (((_712_v27).contains(_711_v26)) ? ((_712_v27).get(_711_v26)) : (!(p0)));
      r0 = _704_v20;
      return r0;
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f4 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f8 = _dafny.ZERO;
      this._f9 = false;
      this._f6 = false;
      this._f7 = false;
      this._f5 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T2, _module.T0, _module.T3, _module.T1];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    set f4(value) {
      let _this = this;
      _this._f4 = value;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f8, f9, f6, f7, f4, f5) {
      let _this = this;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f8;
    };
    fm2(globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements((_this).f7)).cardinality())), ((!(!((_this).f9))) ? (_dafny.Seq.of((_this).f8, (_this).f8, (_this).f8, (_this).f8, (_this).f8)) : (_dafny.Seq.of(new BigNumber(((_this).f5).length)))));
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = [];
      let _713_v0;
      _713_v0 = _dafny.Seq.of((_this).f8, (_this).f8);
      let _714_v1;
      _714_v1 = _dafny.MultiSet.fromElements(new BigNumber(((_this).f5).length));
      let _715_v2;
      _715_v2 = _dafny.Seq.of(p3, (_dafny.MultiSet.FromArray(_713_v0)).equals(_714_v1));
      _715_v2 = _715_v2;
      let _716_v3;
      let _nw124 = new _module.C1();
      _nw124.__ctor((_this).f8, false, p1, (_this).f6, _this.f4, (_this).f5);
      _716_v3 = _nw124;
      let _717_v4;
      _717_v4 = _dafny.Map.Empty.slice().updateUnsafe(_716_v3,(_this).f6);
      let _718_v5;
      _718_v5 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f8);
      let _719_v6;
      _719_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(((_718_v5).contains(p0)) ? ((_718_v5).get(p0)) : ((_this).f8)));
      let _720_v7;
      _720_v7 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber((_717_v4).length)), _719_v6);
      r0 = (((_720_v7).contains((function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of _dafny.IntegerRange(new BigNumber(908), new BigNumber(663))) {
          let _722_v8 = _compr_17;
          if (((new BigNumber(908)).isLessThanOrEqualTo(_722_v8)) && ((_722_v8).isLessThan(new BigNumber(663)))) {
            _coll17.push([(_722_v8).plus((_this).f8),(_this).f8]);
          }
        }
        return _coll17;
      }()).Merge(_719_v6))) ? ((_720_v7).get((function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of _dafny.IntegerRange(new BigNumber(908), new BigNumber(663))) {
          let _721_v8 = _compr_16;
          if (((new BigNumber(908)).isLessThanOrEqualTo(_721_v8)) && ((_721_v8).isLessThan(new BigNumber(663)))) {
            _coll16.push([(_721_v8).plus((_this).f8),(_this).f8]);
          }
        }
        return _coll16;
      }()).Merge(_719_v6))) : ((_this).f8));
      let _hi2 = (_this).f8;
      for (let _723_i0 = ((_this).f8).plus((_this).f8); _723_i0.isLessThan(_hi2); _723_i0 = _723_i0.plus(_dafny.ONE)) {
        r1 = (_dafny.ZERO).minus((((p0) ? ((_this).f8) : (new BigNumber(((_this).f5).length)))).minus((_723_i0).multipliedBy(_723_i0)));
        let _724_v9;
        let _nw125 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
        _724_v9 = _nw125;
        let _725_v10;
        let _nw126 = Array((_dafny.ONE).toNumber()).fill(false);
        _725_v10 = _nw126;
        let _726_v11;
        _726_v11 = _module.D5.create_DC15(_725_v10, true, _718_v5, new BigNumber(-268));
        let _727_v12;
        let _nw127 = Array((new BigNumber(4)).toNumber());
        _nw127[(_dafny.ZERO).toNumber()] = (_this).f6;
        _nw127[(_dafny.ONE).toNumber()] = p1;
        _nw127[(new BigNumber(2)).toNumber()] = (_726_v11).dtor_cf27;
        _nw127[(new BigNumber(3)).toNumber()] = true;
        _727_v12 = _nw127;
        let _728_v13;
        _728_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_725_v10);
        let _index93 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_724_v9).length));
        (_724_v9)[_index93] = _728_v13;
        let _index94 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_724_v9).length));
        (_724_v9)[_index94] = (_728_v13).update(new BigNumber(-894), _727_v12);
        let _729_v14;
        let _nw128 = Array((new BigNumber(5)).toNumber()).fill([]);
        _729_v14 = _nw128;
        let _730_v15;
        let _nw129 = Array((new BigNumber(6)).toNumber());
        _nw129[(_dafny.ZERO).toNumber()] = (_this).f8;
        _nw129[(_dafny.ONE).toNumber()] = _723_i0;
        _nw129[(new BigNumber(2)).toNumber()] = (_this).f8;
        _nw129[(new BigNumber(3)).toNumber()] = new BigNumber(716);
        _nw129[(new BigNumber(4)).toNumber()] = _723_i0;
        _nw129[(new BigNumber(5)).toNumber()] = (_this).f8;
        _730_v15 = _nw129;
        let _index95 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_729_v14).length));
        (_729_v14)[_index95] = _730_v15;
        let _index96 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_729_v14).length));
        let _rhs87 = _730_v15;
        let _rhs88 = _730_v15;
        let _lhs60 = _729_v14;
        let _lhs61 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_729_v14).length));
        _lhs60[_lhs61] = _rhs87;
        _730_v15 = _rhs88;
        (globalState).f1 = (((_this).f6) ? (_723_i0) : (_723_i0));
      }
      _714_v1 = ((_714_v1).Intersect((_714_v1).update(new BigNumber(94), _module.__default.abs((_this).f8)))).Difference(_714_v1);
      (_this).f4 = _this.f4;
      r0 = (_this).f8;
      r0 = (_this).f8;
      r1 = (_this).f8;
      let _731_v16;
      let _nw130 = Array((new BigNumber(3)).toNumber()).fill([]);
      _731_v16 = _nw130;
      r2 = _731_v16;
      return [r0, r1, r2];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _732_v0;
      _732_v0 = _dafny.MultiSet.fromElements((_this).f9, !(true));
      let _733_v1;
      _733_v1 = _dafny.Seq.of((_this).f7);
      if (((_732_v0).Difference(_dafny.MultiSet.fromElements((_this).f9))).IsDisjointFrom(_dafny.MultiSet.FromArray(_733_v1))) {
        _module.__default.m0(globalState);
        let _734_v2;
        _734_v2 = false;
        _734_v2 = (_this).f9;
        let _735_v3;
        let _nw131 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _735_v3 = _nw131;
        let _736_v4;
        _736_v4 = _dafny.Seq.of((_this).f8);
        let _737_v5;
        _737_v5 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_module.__default.fm27((_this).f8, _734_v2, _734_v2, globalState));
        let _738_v6;
        let _nw132 = new _module.C1();
        _nw132.__ctor(new BigNumber((_736_v4).length), (_this).f9, (((_737_v5).contains(_this.f4)) ? ((_737_v5).get(_this.f4)) : ((_this).f9)), !(_734_v2), new _dafny.CodePoint('i'.codePointAt(0)), (_this).f5);
        _738_v6 = _nw132;
        let _739_v7;
        _739_v7 = _dafny.Map.Empty.slice().updateUnsafe(_738_v6,_dafny.Map.Empty.slice().updateUnsafe((_this).f8,_734_v2));
        let _740_v8;
        _740_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,_734_v2);
        let _index97 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((_735_v3).length));
        (_735_v3)[_index97] = (_dafny.ZERO).minus(new BigNumber(((((_739_v7).contains(_738_v6)) ? ((_739_v7).get(_738_v6)) : (_740_v8))).length));
        let _index98 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((_735_v3).length));
        let _rhs89 = (_dafny.ZERO).minus((_this).f8);
        let _rhs90 = (_this).f8;
        let _rhs91 = !(((((_dafny.ZERO).minus(p0)).isLessThan(p0)) ? (_734_v2) : ((_this).f9)));
        let _lhs62 = _735_v3;
        let _lhs63 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((_735_v3).length));
        let _lhs64 = globalState;
        _lhs62[_lhs63] = _rhs89;
        _lhs64.f1 = _rhs90;
        _734_v2 = _rhs91;
        let _741_v9;
        let _nw133 = new _module.C1();
        _nw133.__ctor(_module.__default.safeModuloInt((_735_v3)[_module.__default.safeIndex(new BigNumber(688), new BigNumber((_735_v3).length))], (_this).f8), _734_v2, _734_v2, (_this).f7, _this.f4, (_this).f5);
        _741_v9 = _nw133;
        (_this).f4 = _this.f4;
      } else {
        let _742_v10;
        _742_v10 = false;
        _742_v10 = _module.__default.fm27(p0, (_this).f9, ((_this).f7) || ((_this).f6), globalState);
        let _743_v11;
        _743_v11 = _module.D1.create_DC5((_this).f8);
        let _744_v12;
        _744_v12 = _dafny.Seq.of(_743_v11, _743_v11);
        let _745_v13;
        _745_v13 = _744_v12;
        let _source8 = _745_v13;
        let _746___mcc_h0 = _source8;
        let _747_cf24 = _746___mcc_h0;
        let _748_v14;
        _748_v14 = _dafny.Seq.of(new BigNumber(272));
        let _749_v15;
        _749_v15 = _dafny.Set.fromElements(_748_v14);
        let _750_v16;
        let _nw134 = new _module.C1();
        _nw134.__ctor((_dafny.ZERO).minus(p0), !((_this).f6), ((_this).f8).isLessThan(new BigNumber(653)), _module.__default.fm27(new BigNumber((_749_v15).length), false, _742_v10, globalState), _this.f4, _dafny.Seq.Concat((_this).f5, (_this).f5));
        _750_v16 = _nw134;
        let _751_v17;
        _751_v17 = _module.D1.create_DC4(_this.f4, (_this).f8, _this.f4, p0, (_this).f6);
        _742_v10 = (((_this).f8).plus((_751_v17).dtor_cf11)).isLessThan(_module.__default.safeModuloInt((_this).f8, p0));
        let _752_v18;
        _752_v18 = _dafny.Seq.UnicodeFromString("euaygar");
        let _753_v19;
        let _nw135 = Array((new BigNumber(5)).toNumber()).fill(_dafny.MultiSet.Empty);
        _753_v19 = _nw135;
        let _index99 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((_753_v19).length));
        (_753_v19)[_index99] = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("pas"));
        let _754_v20;
        let _nw136 = Array((new BigNumber(4)).toNumber()).fill(false);
        _754_v20 = _nw136;
        let _755_v21;
        let _nw137 = Array((new BigNumber(18)).toNumber());
        _nw137[(_dafny.ZERO).toNumber()] = _754_v20;
        _nw137[(_dafny.ONE).toNumber()] = _754_v20;
        _nw137[(new BigNumber(2)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(3)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(4)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(5)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(6)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(7)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(8)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(9)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(10)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(11)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(12)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(13)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(14)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(15)).toNumber()] = _754_v20;
        _nw137[(new BigNumber(16)).toNumber()] = ((_module.__default.fm27(p0, false, _module.__default.fm27(p0, (_733_v1)[_module.__default.safeIndex(new BigNumber(635), new BigNumber((_733_v1).length))], (_this).f6, globalState), globalState)) ? (_754_v20) : (_754_v20));
        _nw137[(new BigNumber(17)).toNumber()] = _754_v20;
        _755_v21 = _nw137;
        let _index100 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_755_v21).length));
        (_755_v21)[_index100] = _754_v20;
        let _756_v22;
        _756_v22 = _dafny.Set.fromElements((_this).f7, (_this).f7);
        let _757_v23;
        _757_v23 = _module.D3.create_DC11((_this).f7, p0, _756_v22);
        let _index101 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((_753_v19).length));
        let _index102 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_755_v21).length));
        let _rhs92 = (false) === ((_this).f9);
        let _rhs93 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wxeov"), _module.__default.fm28(new BigNumber((_748_v14).length), _742_v10, _733_v1, _module.D3.create_DC11(false, (_this).f8, _756_v22), globalState));
        let _rhs94 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_module.__default.fm28((_this).f8, (_this).f9, _733_v1, _757_v23, globalState), _dafny.Seq.UnicodeFromString("nservh")));
        let _rhs95 = _754_v20;
        let _lhs65 = _753_v19;
        let _lhs66 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((_753_v19).length));
        let _lhs67 = _755_v21;
        let _lhs68 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_755_v21).length));
        _742_v10 = _rhs92;
        _752_v18 = _rhs93;
        _lhs65[_lhs66] = _rhs94;
        _lhs67[_lhs68] = _rhs95;
        let _758_v24;
        let _nw138 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _758_v24 = _nw138;
        let _759_v25;
        _759_v25 = _dafny.Map.Empty.slice().updateUnsafe(_754_v20,_758_v24);
        let _760_v26;
        _760_v26 = _dafny.MultiSet.fromElements((_this).f8);
        let _761_v27;
        _761_v27 = _dafny.Map.Empty.slice().updateUnsafe((((_759_v25).contains(_754_v20)) ? ((_759_v25).get(_754_v20)) : (_758_v24)),_760_v26);
        let _762_v28;
        let _nw139 = new _module.C3();
        _nw139.__ctor(false, _742_v10, (_dafny.ZERO).minus(new BigNumber((_761_v27).length)), (_this).f7, (_this).f7, (_this).f7, _this.f4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(657)), function (_763_i0) {
          return _this.f4;
        }));
        _762_v28 = _nw139;
        let _764_v29;
        let _init14 = function (_765_i1) {
          return false;
        };
        let _nw140 = Array((new BigNumber(14)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw140.length); _i0_14++) {
          _nw140[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _764_v29 = _nw140;
        let _766_v30;
        _766_v30 = _dafny.Set.fromElements(_764_v29);
        let _767_v31;
        _767_v31 = _module.D7.create_DC18(_766_v30);
        let _rhs96 = (_this).f7;
        let _rhs97 = !(_742_v10) || ((_this).f7);
        let _rhs98 = (_767_v31).dtor_cf33;
        let _rhs99 = (_dafny.ZERO).minus(p0);
        let _lhs69 = globalState;
        _742_v10 = _rhs96;
        _742_v10 = _rhs97;
        _766_v30 = _rhs98;
        _lhs69.f1 = _rhs99;
        let _768_v32;
        let _nw141 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _768_v32 = _nw141;
        let _index103 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_768_v32).length));
        (_768_v32)[_index103] = (_dafny.ZERO).minus(p0);
        let _index104 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_768_v32).length));
        (_768_v32)[_index104] = (_dafny.ZERO).minus(p0);
        let _index105 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_768_v32).length));
        (_768_v32)[_index105] = (_this).f8;
      }
      let _hi3 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), p0);
      for (let _769_i2 = (_dafny.ZERO).minus((_this).f8); _769_i2.isLessThan(_hi3); _769_i2 = _769_i2.plus(_dafny.ONE)) {
        _module.__default.m0(globalState);
        let _770_v33;
        _770_v33 = _dafny.Seq.UnicodeFromString("mm");
        _770_v33 = _dafny.Seq.update(_770_v33, _module.__default.safeIndex(_module.__default.safeModuloInt(p0, p0), new BigNumber((_770_v33).length)), _this.f4);
        let _771_v34;
        let _nw142 = Array((new BigNumber(7)).toNumber());
        _nw142[(_dafny.ZERO).toNumber()] = (_this).f7;
        _nw142[(_dafny.ONE).toNumber()] = true;
        _nw142[(new BigNumber(2)).toNumber()] = (_this).f6;
        _nw142[(new BigNumber(3)).toNumber()] = _module.__default.fm27(new BigNumber(((_this).f5).length), (_this).f9, (_this).f6, globalState);
        _nw142[(new BigNumber(4)).toNumber()] = (_this).f9;
        _nw142[(new BigNumber(5)).toNumber()] = (_this).f6;
        _nw142[(new BigNumber(6)).toNumber()] = (_this).f9;
        _771_v34 = _nw142;
        let _772_v35;
        let _nw143 = Array((new BigNumber(26)).toNumber());
        _nw143[(_dafny.ZERO).toNumber()] = _dafny.Seq.of((_this).f7, false);
        _nw143[(_dafny.ONE).toNumber()] = _dafny.Seq.of((_this).f7);
        _nw143[(new BigNumber(2)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(3)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(4)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(5)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(6)).toNumber()] = _dafny.Seq.of((_this).f9);
        _nw143[(new BigNumber(7)).toNumber()] = _dafny.Seq.of((_this).f9);
        _nw143[(new BigNumber(8)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(9)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(10)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(11)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(12)).toNumber()] = _dafny.Seq.of((_this).f7);
        _nw143[(new BigNumber(13)).toNumber()] = _dafny.Seq.of((_this).f7);
        _nw143[(new BigNumber(14)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(15)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(16)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(17)).toNumber()] = _dafny.Seq.of((_this).f6, (_this).f7);
        _nw143[(new BigNumber(18)).toNumber()] = _dafny.Seq.of((_this).f9);
        _nw143[(new BigNumber(19)).toNumber()] = _module.__default.fm29(globalState);
        _nw143[(new BigNumber(20)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(21)).toNumber()] = _dafny.Seq.of((_this).f9);
        _nw143[(new BigNumber(22)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(23)).toNumber()] = _733_v1;
        _nw143[(new BigNumber(24)).toNumber()] = _dafny.Seq.of((_this).f6);
        _nw143[(new BigNumber(25)).toNumber()] = _733_v1;
        _772_v35 = _nw143;
        let _773_v36;
        _773_v36 = _module.D1.create_DC3(_772_v35);
        let _774_v37;
        let _nw144 = Array((new BigNumber(26)).toNumber());
        _nw144[(_dafny.ZERO).toNumber()] = _773_v36;
        _nw144[(_dafny.ONE).toNumber()] = _773_v36;
        _nw144[(new BigNumber(2)).toNumber()] = _module.D1.create_DC3(_772_v35);
        _nw144[(new BigNumber(3)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(4)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(5)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(6)).toNumber()] = _module.D1.create_DC3(_772_v35);
        _nw144[(new BigNumber(7)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(8)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(9)).toNumber()] = _module.D1.create_DC3(_772_v35);
        _nw144[(new BigNumber(10)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(11)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(12)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(13)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(14)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(15)).toNumber()] = _module.D1.create_DC3(_772_v35);
        _nw144[(new BigNumber(16)).toNumber()] = _module.D1.create_DC3(_772_v35);
        _nw144[(new BigNumber(17)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(18)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(19)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(20)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(21)).toNumber()] = _module.D1.create_DC3(_772_v35);
        _nw144[(new BigNumber(22)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(23)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(24)).toNumber()] = _773_v36;
        _nw144[(new BigNumber(25)).toNumber()] = _773_v36;
        _774_v37 = _nw144;
        let _775_v38;
        let _nw145 = new _module.C2();
        _nw145.__ctor(_771_v34, _774_v37);
        _775_v38 = _nw145;
        let _776_v39;
        _776_v39 = _dafny.Seq.of(_775_v38, _775_v38);
        _775_v38 = (_776_v39)[_module.__default.safeIndex(_module.__default.safeModuloInt(p0, p0), new BigNumber((_776_v39).length))];
        if ((_this).f7) {
          let _777_v40;
          _777_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm30((_this).f9, globalState)).length),new BigNumber((_732_v0).cardinality()));
          let _778_v41;
          _778_v41 = _dafny.Map.Empty.slice().updateUnsafe(_777_v40,(_this).f8);
          let _779_v42;
          _779_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,false);
          let _780_v43;
          _780_v43 = _dafny.Seq.of(p0, (_this).f8);
          _778_v41 = (_778_v41).update((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_732_v0).cardinality()),p0)).update((_dafny.ZERO).minus(_769_i2), new BigNumber(((_779_v42).update(!((_this).f7), (_this).f7)).length)), (_769_i2).minus((_780_v43)[_module.__default.safeIndex(p0, new BigNumber((_780_v43).length))]));
          let _781_v44;
          _781_v44 = true;
          let _782_v45;
          _782_v45 = _dafny.Set.fromElements((_this).f7);
          let _783_v46;
          _783_v46 = _module.D3.create_DC11(false, p0, _782_v45);
          let _784_v47;
          _784_v47 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_780_v43).length), new BigNumber(646))));
          let _rhs100 = (function (_pat_let22_0) {
            return function (_785_dt__update__tmp_h0) {
              return function (_pat_let23_0) {
                return function (_786_dt__update_hcf22_h0) {
                  return _module.D3.create_DC11((_785_dt__update__tmp_h0).dtor_cf20, (_785_dt__update__tmp_h0).dtor_cf21, _786_dt__update_hcf22_h0);
                }(_pat_let23_0);
              }(_dafny.Set.fromElements((_this).f9));
            }(_pat_let22_0);
          }(_783_v46)).dtor_cf20;
          let _rhs101 = new BigNumber((_784_v47).cardinality());
          let _lhs70 = globalState;
          _781_v44 = _rhs100;
          _lhs70.f1 = _rhs101;
          let _787_v48;
          _787_v48 = _dafny.Map.Empty.slice().updateUnsafe((true) || ((_this).f9),_this.f4);
          _787_v48 = _787_v48;
          _781_v44 = (_this).f9;
          let _788_v49;
          _788_v49 = _dafny.Set.fromElements(p0);
          let _789_v50;
          _789_v50 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.of(new BigNumber(((_this).f5).length)));
          let _790_v51;
          _790_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_788_v49).length),(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_733_v1).length),(((_789_v50).contains((_this).f8)) ? ((_789_v50).get((_this).f8)) : (_780_v43)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(264),_780_v43)));
          _790_v51 = (_790_v51).update((_this).f8, _789_v50);
        } else {
          let _791_v52;
          _791_v52 = false;
          _791_v52 = !(((true) ? (_791_v52) : ((_this).f7)));
          let _index106 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_775_v38).f14).length));
          ((_775_v38).f14)[_index106] = (p0).isEqualTo(_769_i2);
          let _792_v53;
          _792_v53 = _module.D1.create_DC4(_this.f4, _769_i2, _this.f4, p0, (_this).f6);
          let _793_v54;
          _793_v54 = _dafny.Set.fromElements(_792_v53);
          let _index107 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_775_v38).f14).length));
          ((_775_v38).f14)[_index107] = !(_793_v54).equals(_793_v54);
          let _index108 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_775_v38).f14).length));
          ((_775_v38).f14)[_index108] = _dafny.Seq.contains(((!(_791_v52)) ? ((_this).f5) : ((_this).f5)), _this.f4);
          let _794_v55;
          let _nw146 = new _module.C0();
          _nw146.__ctor(((_775_v38).f14)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_775_v38).f14).length))], ((_775_v38).f14)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_775_v38).f14).length))], _this.f4, (_this).f5);
          _794_v55 = _nw146;
          let _795_v56;
          _795_v56 = _dafny.Map.Empty.slice().updateUnsafe(_794_v55,new _dafny.CodePoint('m'.codePointAt(0)));
          (globalState).f1 = _module.__default.safeModuloInt((_this).f8, new BigNumber(((_795_v56).Merge(_795_v56)).length));
          let _796_v57;
          _796_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_791_v52);
          let _797_v58;
          let _nw147 = new _module.C3();
          _nw147.__ctor(_dafny.Seq.contains(_dafny.Seq.of(p0), (_this).f8), (((_796_v57).contains(_769_i2)) ? ((_796_v57).get(_769_i2)) : ((_this).f6)), (_dafny.ZERO).minus((p0).minus((_this).f8)), (_this).f7, (_this).f7, (_this).f7, _this.f4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(665)), function (_798_i3) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          }));
          _797_v58 = _nw147;
        }
      }
      let _799_v59;
      let _nw148 = Array((new BigNumber(19)).toNumber()).fill(_module.D6.Default());
      _799_v59 = _nw148;
      let _800_v60;
      _800_v60 = _module.D6.create_DC17((_this).f5, (_this).f5);
      let _index109 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_799_v59).length));
      (_799_v59)[_index109] = (((_this).f7) ? (_800_v60) : (_module.D6.create_DC17((_this).f5, (_this).f5)));
      let _801_v61;
      _801_v61 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_dafny.Seq.update(_733_v1, _module.__default.safeIndex(p0, new BigNumber((_733_v1).length)), (_this).f9));
      let _index110 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_799_v59).length));
      let _rhs102 = (_this).f8;
      let _rhs103 = _800_v60;
      let _rhs104 = new BigNumber(26);
      let _rhs105 = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_733_v1, (((_801_v61).contains(p0)) ? ((_801_v61).get(p0)) : (_733_v1))));
      let _lhs71 = globalState;
      let _lhs72 = _799_v59;
      let _lhs73 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_799_v59).length));
      let _lhs74 = globalState;
      _lhs71.f1 = _rhs102;
      _lhs72[_lhs73] = _rhs103;
      _lhs74.f1 = _rhs104;
      _732_v0 = _rhs105;
      let _hi4 = (_dafny.ZERO).minus((_this).f8);
      for (let _802_i4 = (((_this).f7) ? (p0) : ((_this).f8)); _802_i4.isLessThan(_hi4); _802_i4 = _802_i4.plus(_dafny.ONE)) {
        if ((_733_v1)[_module.__default.safeIndex((_dafny.ZERO).minus(_802_i4), new BigNumber((_733_v1).length))]) {
          let _803_v62;
          let _nw149 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _803_v62 = _nw149;
          let _index111 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_803_v62).length));
          (_803_v62)[_index111] = (_dafny.ZERO).minus(_802_i4);
          let _index112 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_803_v62).length));
          (_803_v62)[_index112] = p0;
          let _804_v63;
          _804_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f4);
          (globalState).f1 = (((_this).f6) ? (_802_i4) : (((_803_v62)[_module.__default.safeIndex(new BigNumber(883), new BigNumber((_803_v62).length))]).multipliedBy(new BigNumber((_804_v63).length))));
          let _805_v64;
          let _nw150 = Array((new BigNumber(28)).toNumber()).fill(false);
          _805_v64 = _nw150;
          let _806_v65;
          _806_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f8);
          let _807_v66;
          _807_v66 = _dafny.MultiSet.fromElements((_this).f8);
          let _808_v67;
          _808_v67 = _module.D5.create_DC15(_805_v64, (_this).f9, _806_v65, new BigNumber((_807_v66).cardinality()));
          let _809_v68;
          _809_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_808_v67).dtor_cf27);
          let _810_v69;
          _810_v69 = _dafny.Seq.of(_809_v68, _809_v68);
          let _811_v70;
          _811_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,new BigNumber((_810_v69).length));
          _806_v65 = (_806_v65).update((_this).f6, (_this).fm1((_this).f6, _809_v68, (_this).f9, globalState));
          (globalState).f1 = _802_i4;
          let _812_v71;
          _812_v71 = true;
          let _index113 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_803_v62).length));
          let _rhs106 = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f4,_805_v64)).length)).isLessThan(p0);
          let _rhs107 = _module.__default.safeDivisionInt(_802_i4, (_this).f8);
          let _lhs75 = _803_v62;
          let _lhs76 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_803_v62).length));
          _812_v71 = _rhs106;
          _lhs75[_lhs76] = _rhs107;
        } else {
          (globalState).f1 = new BigNumber(382);
          let _813_v72;
          _813_v72 = false;
          _813_v72 = (_this).f9;
          let _814_v73;
          _814_v73 = _dafny.Set.fromElements(_this.f4, _this.f4);
          let _815_v74;
          let _nw151 = Array((new BigNumber(11)).toNumber());
          _nw151[(_dafny.ZERO).toNumber()] = _813_v72;
          _nw151[(_dafny.ONE).toNumber()] = (_this).f6;
          _nw151[(new BigNumber(2)).toNumber()] = (_this).f6;
          _nw151[(new BigNumber(3)).toNumber()] = (_this).f7;
          _nw151[(new BigNumber(4)).toNumber()] = (_this).f9;
          _nw151[(new BigNumber(5)).toNumber()] = (p0).isLessThanOrEqualTo((_this).f8);
          _nw151[(new BigNumber(6)).toNumber()] = (_dafny.Set.fromElements(_this.f4)).IsSubsetOf(_814_v73);
          _nw151[(new BigNumber(7)).toNumber()] = (_813_v72) && ((_this).f7);
          _nw151[(new BigNumber(8)).toNumber()] = _813_v72;
          _nw151[(new BigNumber(9)).toNumber()] = (_this).f7;
          _nw151[(new BigNumber(10)).toNumber()] = true;
          _815_v74 = _nw151;
          let _index114 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_815_v74).length));
          (_815_v74)[_index114] = _813_v72;
          let _index115 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_815_v74).length));
          (_815_v74)[_index115] = (_this).f7;
          let _816_v75;
          _816_v75 = _module.D2.create_DC8(p0, p0);
          let _817_v76;
          _817_v76 = _module.D2.create_DC9(_816_v75);
          _817_v76 = ((true) ? (_817_v76) : (_817_v76));
          let _init15 = ((_818_v74) => function (_819_i5) {
            return (_818_v74)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_818_v74).length))];
          })(_815_v74);
          let _nw152 = Array((new BigNumber(20)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw152.length); _i0_15++) {
            _nw152[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _815_v74 = _nw152;
        }
        let _820_v77;
        let _nw153 = Array((new BigNumber(27)).toNumber()).fill(false);
        _820_v77 = _nw153;
        let _821_v78;
        _821_v78 = _dafny.Set.fromElements(_820_v77);
        let _822_v79;
        _822_v79 = _module.D7.create_DC18(_821_v78);
        let _pat_let_tv15 = _820_v77;
        _822_v79 = function (_pat_let24_0) {
          return function (_823_dt__update__tmp_h1) {
            return function (_pat_let25_0) {
              return function (_824_dt__update_hcf33_h0) {
                return _module.D7.create_DC18(_824_dt__update_hcf33_h0);
              }(_pat_let25_0);
            }(_dafny.Set.fromElements(_pat_let_tv15));
          }(_pat_let24_0);
        }(_module.D7.create_DC18(_821_v78));
        let _825_v80;
        _825_v80 = false;
        _825_v80 = false;
        let _826_v81;
        _826_v81 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f8);
        let _827_v82;
        _827_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
        let _828_v83;
        _828_v83 = _module.D3.create_DC11((_this).f6, p0, _dafny.Set.fromElements((_this).f6));
        let _829_v84;
        let _init16 = ((_830_v82) => function (_831_i6) {
          return _module.__default.safeDivisionInt(_831_i6, new BigNumber((_830_v82).length));
        })(_827_v82);
        let _nw154 = Array((new BigNumber(28)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw154.length); _i0_16++) {
          _nw154[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _829_v84 = _nw154;
        let _832_v85;
        _832_v85 = _dafny.MultiSet.fromElements(_829_v84, _829_v84, _829_v84, _829_v84, _829_v84);
        let _833_v86;
        let _nw155 = Array((new BigNumber(24)).toNumber());
        _nw155[(_dafny.ZERO).toNumber()] = _802_i4;
        _nw155[(_dafny.ONE).toNumber()] = ((((_826_v81).contains((_this).f6)) ? ((_826_v81).get((_this).f6)) : (new BigNumber(((_this).f5).length)))).minus(new BigNumber(456));
        _nw155[(new BigNumber(2)).toNumber()] = new BigNumber(280);
        _nw155[(new BigNumber(3)).toNumber()] = ((_this).fm1((_this).f7, _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f9), (_this).f7, globalState)).multipliedBy(p0);
        _nw155[(new BigNumber(4)).toNumber()] = (_this).f8;
        _nw155[(new BigNumber(5)).toNumber()] = (_this).fm1(_825_v80, (_827_v82).update((_this).f6, (_this).f7), (_this).f7, globalState);
        _nw155[(new BigNumber(6)).toNumber()] = (_this).f8;
        _nw155[(new BigNumber(7)).toNumber()] = _802_i4;
        _nw155[(new BigNumber(8)).toNumber()] = (_828_v83).dtor_cf21;
        _nw155[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_802_i4);
        _nw155[(new BigNumber(10)).toNumber()] = (_this).f8;
        _nw155[(new BigNumber(11)).toNumber()] = (((_this).f6) ? (new BigNumber(((_this).f5).length)) : (_802_i4));
        _nw155[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw155[(new BigNumber(13)).toNumber()] = (new BigNumber((_832_v85).cardinality())).multipliedBy((_this).f8);
        _nw155[(new BigNumber(14)).toNumber()] = _802_i4;
        _nw155[(new BigNumber(15)).toNumber()] = (_802_i4).plus(p0);
        _nw155[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw155[(new BigNumber(17)).toNumber()] = _802_i4;
        _nw155[(new BigNumber(18)).toNumber()] = new BigNumber(((_this).f5).length);
        _nw155[(new BigNumber(19)).toNumber()] = (_this).f8;
        _nw155[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_802_i4, new BigNumber(((_this).f5).length)));
        _nw155[(new BigNumber(21)).toNumber()] = _802_i4;
        _nw155[(new BigNumber(22)).toNumber()] = (_this).f8;
        _nw155[(new BigNumber(23)).toNumber()] = new BigNumber(-869);
        _833_v86 = _nw155;
        let _rhs108 = (new BigNumber(777)).minus((_this).f8);
        let _rhs109 = _829_v84;
        let _rhs110 = _module.__default.safeDivisionInt(new BigNumber(578), _802_i4);
        let _rhs111 = p0;
        let _lhs77 = globalState;
        let _lhs78 = globalState;
        let _lhs79 = globalState;
        _lhs77.f1 = _rhs108;
        _833_v86 = _rhs109;
        _lhs78.f1 = _rhs110;
        _lhs79.f1 = _rhs111;
      }
      (globalState).f1 = (_this).f8;
      let _834_i7;
      _834_i7 = _dafny.ZERO;
      L4: {
        while ((p0).isLessThan(new BigNumber(((_this).f5).length))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_834_i7)) {
              break L4;
            }
            _834_i7 = (_834_i7).plus(_dafny.ONE);
            let _835_v87;
            _835_v87 = false;
            _835_v87 = _dafny.Seq.IsPrefixOf((_this).f5, (_this).f5);
            let _836_v88;
            _836_v88 = _dafny.Map.Empty.slice().updateUnsafe(_835_v87,(_this).f9);
            let _837_v89;
            _837_v89 = _dafny.Map.Empty.slice().updateUnsafe(_835_v87,_dafny.Seq.Create(_module.__default.abs(new BigNumber(58)), ((_838_p0) => function (_839_i8) {
              return _838_p0;
            })(p0)));
            (globalState).f1 = (_this).fm1(true, (_836_v88).Merge(_dafny.Map.Empty.slice().updateUnsafe(_835_v87,true)), !(!(_837_v89).equals(_837_v89)), globalState);
            let _840_v90;
            let _nw156 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
            _840_v90 = _nw156;
            let _index116 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_840_v90).length));
            (_840_v90)[_index116] = (_this).f8;
            let _index117 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_840_v90).length));
            (_840_v90)[_index117] = p0;
            let _841_v91;
            _841_v91 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(97),false);
            let _842_v92;
            _842_v92 = _dafny.Set.fromElements(true);
            let _843_v93;
            _843_v93 = _module.D3.create_DC11((((_841_v91).contains((_dafny.ZERO).minus(p0))) ? ((_841_v91).get((_dafny.ZERO).minus(p0))) : ((_this).f7)), new BigNumber(26), _842_v92);
            let _844_v94;
            let _nw157 = new _module.C0();
            _nw157.__ctor(_835_v87, (_this).f7, _this.f4, _module.__default.fm28((_840_v90)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_840_v90).length))], (_this).f6, _733_v1, _843_v93, globalState));
            _844_v94 = _nw157;
            let _845_v95;
            let _nw158 = Array((new BigNumber(7)).toNumber()).fill([]);
            _845_v95 = _nw158;
            let _846_v96;
            let _nw159 = Array((new BigNumber(12)).toNumber()).fill(false);
            _846_v96 = _nw159;
            let _847_v97;
            let _nw160 = Array((new BigNumber(3)).toNumber());
            _nw160[(_dafny.ZERO).toNumber()] = _846_v96;
            _nw160[(_dafny.ONE).toNumber()] = _846_v96;
            _nw160[(new BigNumber(2)).toNumber()] = _846_v96;
            _847_v97 = _nw160;
            let _index118 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_845_v95).length));
            (_845_v95)[_index118] = _847_v97;
            let _848_v98;
            let _init17 = ((_849_v94, _850_p0, _851_v87) => function (_852_i9) {
              return _module.D0.create_DC1((_849_v94).f6, _850_p0, _851_v87, (_this).f8);
            })(_844_v94, p0, _835_v87);
            let _nw161 = Array((new BigNumber(4)).toNumber());
            for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw161.length); _i0_17++) {
              _nw161[_i0_17] = _init17(new BigNumber(_i0_17));
            }
            _848_v98 = _nw161;
            let _853_v99;
            _853_v99 = _module.D0.create_DC1((_this).f9, (_dafny.ZERO).minus((_this).f8), (_844_v94).f7, new BigNumber(((_844_v94).f5).length));
            let _index119 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_848_v98).length));
            (_848_v98)[_index119] = _853_v99;
            let _index120 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_845_v95).length));
            let _index121 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_848_v98).length));
            let _rhs112 = (((_this).f6) ? (_844_v94) : (_844_v94));
            let _rhs113 = _847_v97;
            let _rhs114 = _853_v99;
            let _lhs80 = _845_v95;
            let _lhs81 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_845_v95).length));
            let _lhs82 = _848_v98;
            let _lhs83 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_848_v98).length));
            _844_v94 = _rhs112;
            _lhs80[_lhs81] = _rhs113;
            _lhs82[_lhs83] = _rhs114;
          }
        }
      }
      let _854_v100;
      _854_v100 = _dafny.Map.Empty.slice().updateUnsafe(p0,((_this).f8).plus(p0));
      r0 = _854_v100;
      return r0;
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f4 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f8 = _dafny.ZERO;
      this._f9 = false;
      this._f6 = false;
      this._f7 = false;
      this._f5 = _dafny.Seq.UnicodeFromString("");
      this._f10 = [];
      this._f11 = [];
    }
    _parentTraits() {
      return [_module.T1, _module.T2, _module.T3, _module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    set f4(value) {
      let _this = this;
      _this._f4 = value;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f10, f11, f6, f7, f4, f5, f8, f9) {
      let _this = this;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      return;
    }
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_this).f8);
    };
    fm2(globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f6)).length), (_this).f8), _dafny.Seq.Concat(_dafny.Seq.of((_this).f8), _dafny.Seq.Create(_module.__default.abs(new BigNumber(342)), function (_855_i0) {
        return (_this).f8;
      })));
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = [];
      let _856_v0;
      _856_v0 = _module.D0.create_DC0();
      let _source9 = function (_source10) {
        if (_source10.is_DC0) {
          return _module.D0.create_DC1((_this).f6, (_this).f8, (_this).f9, (_this).f8);
        } else if (_source10.is_DC1) {
          let _857___mcc_h7 = (_source10).cf0;
          let _858___mcc_h8 = (_source10).cf1;
          let _859___mcc_h9 = (_source10).cf2;
          let _860___mcc_h10 = (_source10).cf3;
          let _861_cf3 = _860___mcc_h10;
          let _862_cf2 = _859___mcc_h9;
          let _863_cf1 = _858___mcc_h8;
          let _864_cf0 = _857___mcc_h7;
          return _module.D0.create_DC1((_this).f9, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(528)), function (_865_i0) {
  return (_this).f8;
})).length), new BigNumber(((_this).f5).length))).length), _864_cf0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f9)).length));
        } else {
          let _866___mcc_h11 = (_source10).cf4;
          let _867___mcc_h12 = (_source10).cf5;
          let _868___mcc_h13 = (_source10).cf6;
          let _869_cf6 = _868___mcc_h13;
          let _870_cf5 = _867___mcc_h12;
          let _871_cf4 = _866___mcc_h11;
          return _module.D0.create_DC1(true, new BigNumber(556), (_this).f7, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_871_cf4,_870_cf5)).length));
        }
      }(_856_v0);
      if (_source9.is_DC0) {
        let _872_v1;
        _872_v1 = false;
        _872_v1 = (_this).f9;
        _872_v1 = _872_v1;
        let _873_v2;
        _873_v2 = _dafny.Seq.of((_this).f6);
        let _874_v3;
        _874_v3 = _dafny.Seq.of(_873_v2, _873_v2);
        let _875_v4;
        _875_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(539),(_this).f8);
        let _876_v5;
        _876_v5 = _dafny.Seq.of(_875_v4, _875_v4);
        let _877_v6;
        _877_v6 = _dafny.Set.fromElements((_this).f8);
        let _878_v7;
        _878_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(727),(_this).f7);
        let _879_v8;
        let _nw162 = Array((new BigNumber(22)).toNumber());
        _nw162[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((_874_v3)[_module.__default.safeIndex(new BigNumber((_876_v5).length), new BigNumber((_874_v3).length))], _873_v2);
        _nw162[(_dafny.ONE).toNumber()] = _873_v2;
        _nw162[(new BigNumber(2)).toNumber()] = _873_v2;
        _nw162[(new BigNumber(3)).toNumber()] = _873_v2;
        _nw162[(new BigNumber(4)).toNumber()] = _873_v2;
        _nw162[(new BigNumber(5)).toNumber()] = _873_v2;
        _nw162[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat((_874_v3)[_module.__default.safeIndex(new BigNumber((_877_v6).length), new BigNumber((_874_v3).length))], _873_v2);
        _nw162[(new BigNumber(7)).toNumber()] = _873_v2;
        _nw162[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_this).f7), _dafny.Seq.of(true));
        _nw162[(new BigNumber(9)).toNumber()] = _873_v2;
        _nw162[(new BigNumber(10)).toNumber()] = _873_v2;
        _nw162[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_873_v2, _module.__default.safeIndex(new BigNumber(102), new BigNumber((_873_v2).length)), _872_v1), _dafny.Seq.of(p0, p0));
        _nw162[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_873_v2, _873_v2);
        _nw162[(new BigNumber(13)).toNumber()] = _873_v2;
        _nw162[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_this).f9, !(false)), _873_v2);
        _nw162[(new BigNumber(15)).toNumber()] = _873_v2;
        _nw162[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(p1, (((_878_v7).contains((_this).f8)) ? ((_878_v7).get((_this).f8)) : (p1)), _module.__default.fm3(globalState));
        _nw162[(new BigNumber(17)).toNumber()] = _873_v2;
        _nw162[(new BigNumber(18)).toNumber()] = _dafny.Seq.of((_this).f6);
        _nw162[(new BigNumber(19)).toNumber()] = _873_v2;
        _nw162[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_873_v2, _873_v2);
        _nw162[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_this).f6, p0), _873_v2);
        _879_v8 = _nw162;
        let _880_v9;
        _880_v9 = _dafny.Seq.of((_this).f11);
        let _881_v10;
        let _nw163 = Array((new BigNumber(26)).toNumber());
        _nw163[(_dafny.ZERO).toNumber()] = (_this).f11;
        _nw163[(_dafny.ONE).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(2)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(3)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(4)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(5)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(6)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(7)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(8)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(9)).toNumber()] = (_880_v9)[_module.__default.safeIndex((_this).f8, new BigNumber((_880_v9).length))];
        _nw163[(new BigNumber(10)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(11)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(12)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(13)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(14)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(15)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(16)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(17)).toNumber()] = (_880_v9)[_module.__default.safeIndex((_this).f8, new BigNumber((_880_v9).length))];
        _nw163[(new BigNumber(18)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(19)).toNumber()] = (_880_v9)[_module.__default.safeIndex((_this).f8, new BigNumber((_880_v9).length))];
        _nw163[(new BigNumber(20)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(21)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(22)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(23)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(24)).toNumber()] = (_this).f11;
        _nw163[(new BigNumber(25)).toNumber()] = (_this).f11;
        _881_v10 = _nw163;
        let _882_v11;
        _882_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f11);
        let _index122 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_881_v10).length));
        (_881_v10)[_index122] = (((_882_v11).contains(new BigNumber((_module.__default.fm4(p0, (_this).f8, (_this).f7, new BigNumber((_877_v6).length), globalState)).cardinality()))) ? ((_882_v11).get(new BigNumber((_module.__default.fm4(p0, (_this).f8, (_this).f7, new BigNumber((_877_v6).length), globalState)).cardinality()))) : ((_this).f11));
        let _index123 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_881_v10).length));
        let _rhs115 = (_module.D1.create_DC3(_879_v8)).dtor_cf7;
        let _rhs116 = (_this).f11;
        let _lhs84 = _881_v10;
        let _lhs85 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_881_v10).length));
        _879_v8 = _rhs115;
        _lhs84[_lhs85] = _rhs116;
        _872_v1 = !(!(_872_v1));
      } else if (_source9.is_DC1) {
        let _883___mcc_h0 = (_source9).cf0;
        let _884___mcc_h1 = (_source9).cf1;
        let _885___mcc_h2 = (_source9).cf2;
        let _886___mcc_h3 = (_source9).cf3;
        let _887_cf3 = _886___mcc_h3;
        let _888_cf2 = _885___mcc_h2;
        let _889_cf1 = _884___mcc_h1;
        let _890_cf0 = _883___mcc_h0;
        _888_cf2 = !(_888_cf2);
        let _891_v12;
        _891_v12 = _dafny.Seq.of(p1);
        let _index124 = _module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index124] = (_891_v12)[_module.__default.safeIndex((_this).f8, new BigNumber((_891_v12).length))];
        let _index125 = _module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index125] = p0;
        let _892_v13;
        _892_v13 = _dafny.Set.fromElements(!(_890_cf0));
        let _893_v14;
        let _nw164 = Array((new BigNumber(19)).toNumber());
        _nw164[(_dafny.ZERO).toNumber()] = p1;
        _nw164[(_dafny.ONE).toNumber()] = _module.__default.fm7(_this.f4, new BigNumber((_892_v13).length), p0, globalState);
        _nw164[(new BigNumber(2)).toNumber()] = p3;
        _nw164[(new BigNumber(3)).toNumber()] = ((_this).f10)[_module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length))];
        _nw164[(new BigNumber(4)).toNumber()] = !(_module.__default.fm3(globalState));
        _nw164[(new BigNumber(5)).toNumber()] = (_this).f9;
        _nw164[(new BigNumber(6)).toNumber()] = ((_this).f10)[_module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length))];
        _nw164[(new BigNumber(7)).toNumber()] = p1;
        _nw164[(new BigNumber(8)).toNumber()] = _890_cf0;
        _nw164[(new BigNumber(9)).toNumber()] = (_this).f9;
        _nw164[(new BigNumber(10)).toNumber()] = (_this).f9;
        _nw164[(new BigNumber(11)).toNumber()] = !(_888_cf2);
        _nw164[(new BigNumber(12)).toNumber()] = (_891_v12)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(561)), function (_894_i1) {
          return _this.f4;
        })).length)), new BigNumber((_891_v12).length))];
        _nw164[(new BigNumber(13)).toNumber()] = _888_cf2;
        _nw164[(new BigNumber(14)).toNumber()] = _890_cf0;
        _nw164[(new BigNumber(15)).toNumber()] = ((_this).f10)[_module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length))];
        _nw164[(new BigNumber(16)).toNumber()] = ((_this).f10)[_module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length))];
        _nw164[(new BigNumber(17)).toNumber()] = _module.__default.fm3(globalState);
        _nw164[(new BigNumber(18)).toNumber()] = (_this).f7;
        _893_v14 = _nw164;
        let _895_v15;
        let _nw165 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
        _895_v15 = _nw165;
        let _896_v16;
        _896_v16 = _module.D1.create_DC3(_895_v15);
        let _897_v17;
        let _nw166 = Array((new BigNumber(5)).toNumber());
        _nw166[(_dafny.ZERO).toNumber()] = _896_v16;
        _nw166[(_dafny.ONE).toNumber()] = _896_v16;
        _nw166[(new BigNumber(2)).toNumber()] = _module.D1.create_DC3(_895_v15);
        _nw166[(new BigNumber(3)).toNumber()] = _896_v16;
        _nw166[(new BigNumber(4)).toNumber()] = _module.D1.create_DC3(_895_v15);
        _897_v17 = _nw166;
        let _898_v18;
        let _nw167 = new _module.C2();
        _nw167.__ctor(_893_v14, _897_v17);
        _898_v18 = _nw167;
        if (((_dafny.ZERO).minus((_this).f8)).isLessThan((new BigNumber(((_this).f5).length)).minus(new BigNumber(-466)))) {
          let _899_v19;
          let _nw168 = new _module.C0();
          _nw168.__ctor(!(_module.__default.fm7(_this.f4, _887_cf3, !(true), globalState)), p3, _this.f4, (_this).f5);
          _899_v19 = _nw168;
          let _900_v20;
          let _nw169 = Array((new BigNumber(18)).toNumber());
          _nw169[(_dafny.ZERO).toNumber()] = (((_this).f6) ? (_899_v19) : (_899_v19));
          _nw169[(_dafny.ONE).toNumber()] = _899_v19;
          _nw169[(new BigNumber(2)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(3)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(4)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(5)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(6)).toNumber()] = ((!(_890_cf0)) ? (_899_v19) : (_899_v19));
          _nw169[(new BigNumber(7)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(8)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(9)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(10)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(11)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(12)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(13)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(14)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(15)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(16)).toNumber()] = _899_v19;
          _nw169[(new BigNumber(17)).toNumber()] = _899_v19;
          _900_v20 = _nw169;
          let _index126 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_900_v20).length));
          (_900_v20)[_index126] = _899_v19;
          let _901_v21;
          _901_v21 = _dafny.Seq.of(_899_v19);
          let _902_v22;
          _902_v22 = _dafny.Seq.of(_dafny.Seq.of(true, false), _891_v12);
          let _index127 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_900_v20).length));
          let _index128 = _module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length));
          let _rhs117 = ((_this).f8).minus(_887_cf3);
          let _rhs118 = (_901_v21)[_module.__default.safeIndex(((_this).f8).multipliedBy((_899_v19).fm9(_889_cf1, (_902_v22)[_module.__default.safeIndex((_this).f8, new BigNumber((_902_v22).length))], globalState)), new BigNumber((_901_v21).length))];
          let _rhs119 = !_dafny.Seq.contains(_dafny.Seq.Concat(_891_v12, _891_v12), (_this).f7);
          let _lhs86 = _900_v20;
          let _lhs87 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_900_v20).length));
          let _lhs88 = (_this).f10;
          let _lhs89 = _module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length));
          _887_cf3 = _rhs117;
          _lhs86[_lhs87] = _rhs118;
          _lhs88[_lhs89] = _rhs119;
          let _903_v23;
          _903_v23 = _dafny.Seq.UnicodeFromString("muet");
          let _904_v24;
          _904_v24 = _dafny.Map.Empty.slice().updateUnsafe(!(_890_cf0),false);
          _903_v23 = _dafny.Seq.Concat(_dafny.Seq.Concat((_this).f5, _903_v23), _dafny.Seq.Concat((_this).f5, _module.__default.fm14((_this).fm1(((_this).f10)[_module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length))], _904_v24, !(_890_cf0), globalState), (_this).f8, false, p1, globalState)));
          let _index129 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_900_v20).length));
          (_900_v20)[_index129] = (_900_v20)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_900_v20).length))];
          let _index130 = _module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index130] = p1;
          (globalState).f1 = _889_cf1;
        } else {
          let _905_v25;
          let _nw170 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _905_v25 = _nw170;
          let _906_v26;
          _906_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f9);
          let _index131 = _module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length));
          let _rhs120 = _890_cf0;
          let _rhs121 = _905_v25;
          let _rhs122 = ((_dafny.Map.Empty.slice().updateUnsafe(((_this).f10)[_module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length))],p0)).Merge(_906_v26)).Merge((_906_v26).Merge((_906_v26).update((_this).f7, (_this).f7)));
          let _rhs123 = _module.__default.safeDivisionInt(_889_cf1, new BigNumber((_dafny.Seq.of(_887_cf3)).length));
          let _rhs124 = (_887_cf3).minus(_889_cf1);
          let _lhs90 = (_this).f10;
          let _lhs91 = _module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length));
          let _lhs92 = globalState;
          let _lhs93 = globalState;
          _lhs90[_lhs91] = _rhs120;
          _905_v25 = _rhs121;
          _906_v26 = _rhs122;
          _lhs92.f1 = _rhs123;
          _lhs93.f1 = _rhs124;
          let _index132 = _module.__default.safeIndex(new BigNumber(385), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index132] = _888_cf2;
          let _907_v27;
          _907_v27 = _dafny.Seq.UnicodeFromString("vsc");
          _907_v27 = _dafny.Seq.update(_907_v27, _module.__default.safeIndex(_889_cf1, new BigNumber((_907_v27).length)), _this.f4);
          _887_cf3 = _887_cf3;
          let _908_v28;
          _908_v28 = _module.D1.create_DC5(_889_cf1);
          let _909_v29;
          _909_v29 = _dafny.Seq.of(_908_v28, _908_v28);
          let _910_v30;
          _910_v30 = _dafny.Seq.Concat(_909_v29, _909_v29);
          _910_v30 = _dafny.Seq.update(_dafny.Seq.update(_909_v29, _module.__default.safeIndex(new BigNumber(635), new BigNumber((_909_v29).length)), _908_v28), _module.__default.safeIndex((_this).f8, new BigNumber((_dafny.Seq.update(_909_v29, _module.__default.safeIndex(new BigNumber(635), new BigNumber((_909_v29).length)), _908_v28)).length)), _module.__default.fm25(globalState));
        }
      } else {
        let _911___mcc_h4 = (_source9).cf4;
        let _912___mcc_h5 = (_source9).cf5;
        let _913___mcc_h6 = (_source9).cf6;
        let _914_cf6 = _913___mcc_h6;
        let _915_cf5 = _912___mcc_h5;
        let _916_cf4 = _911___mcc_h4;
        let _917_v31;
        _917_v31 = _dafny.Seq.UnicodeFromString("eyth");
        let _918_v32;
        _918_v32 = _dafny.MultiSet.fromElements(new BigNumber(555), (_this).f8);
        _917_v31 = ((p0) ? (_917_v31) : (_module.__default.fm14((_this).f8, (((_918_v32).contains((_this).f8)) ? ((_918_v32).get((_this).f8)) : ((_this).f8)), _915_cf5, true, globalState)));
        let _919_v33;
        let _nw171 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _919_v33 = _nw171;
        let _index133 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_919_v33).length));
        (_919_v33)[_index133] = _module.__default.fm14(_916_cf4, (_this).f8, (_this).f7, p0, globalState);
        let _index134 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_919_v33).length));
        (_919_v33)[_index134] = _dafny.Seq.update((_this).f5, _module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_917_v31).length), new BigNumber((p2).cardinality())), new BigNumber(((_this).f5).length)), _this.f4);
        if (!((new BigNumber((function () {
          let _coll18 = new _dafny.Map();
          for (const _compr_18 of (_918_v32).Elements) {
            let _920_v34 = _compr_18;
            if ((_918_v32).contains(_920_v34)) {
              _coll18.push([_module.__default.safeDivisionInt(_920_v34, new BigNumber((_dafny.Seq.of((_this).f9, true)).length)),_this.f4]);
            }
          }
          return _coll18;
        }()).length)).multipliedBy((_this).f8)).isEqualTo((_this).f8)) {
          _915_cf5 = p0;
          r1 = (_this).f8;
          let _921_v35;
          _921_v35 = _module.D1.create_DC5(_916_cf4);
          let _922_v36;
          _922_v36 = _dafny.Seq.of(_921_v35);
          _922_v36 = _dafny.Seq.of(_module.D1.create_DC5(_916_cf4), _921_v35, _921_v35);
          _915_cf5 = ((((_this).f6) ? (!(p0)) : ((_this).f7))) || (!(_915_cf5));
          let _923_v37;
          let _nw172 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
          _923_v37 = _nw172;
          let _index135 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_923_v37).length));
          (_923_v37)[_index135] = _918_v32;
          let _index136 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_923_v37).length));
          (_923_v37)[_index136] = _918_v32;
        } else {
          let _924_v38;
          _924_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_915_cf5)).length),new _dafny.CodePoint('k'.codePointAt(0)));
          (_this).f4 = (((_924_v38).contains(_module.__default.safeModuloInt(_916_cf4, new BigNumber(((_this).f5).length)))) ? ((_924_v38).get(_module.__default.safeModuloInt(_916_cf4, new BigNumber(((_this).f5).length)))) : (_this.f4));
          (globalState).f1 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_916_cf4).minus((_this).f8)))));
          let _925_v39;
          let _nw173 = new _module.C1();
          _nw173.__ctor(new BigNumber((_dafny.Seq.of(p1)).length), !(p1), _module.__default.fm3(globalState), p1, _module.__default.fm13(globalState), (_919_v33)[_module.__default.safeIndex(new BigNumber(891), new BigNumber((_919_v33).length))]);
          _925_v39 = _nw173;
          let _index137 = _module.__default.safeIndex(new BigNumber(54), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index137] = p1;
          let _index138 = _module.__default.safeIndex(new BigNumber(54), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index138] = p3;
          let _926_v40;
          let _nw174 = Array((new BigNumber(8)).toNumber()).fill(false);
          _926_v40 = _nw174;
          let _927_v41;
          _927_v41 = _dafny.Set.fromElements(_926_v40);
          _915_cf5 = !((_927_v41).IsDisjointFrom(_927_v41));
        }
        let _928_v42;
        _928_v42 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f7) ? (_916_cf4) : (_916_cf4)),(_this).f8);
        let _929_v43;
        let _nw175 = new _module.C3();
        _nw175.__ctor(p3, p3, new BigNumber(33), p0, p3, p0, _this.f4, _dafny.Seq.UnicodeFromString("qkhm"));
        _929_v43 = _nw175;
        let _930_v44;
        _930_v44 = _dafny.MultiSet.fromElements(_929_v43);
        _916_cf4 = (_dafny.ZERO).minus((((_928_v42).contains((new BigNumber(837)).plus(new BigNumber((_930_v44).cardinality())))) ? ((_928_v42).get((new BigNumber(837)).plus(new BigNumber((_930_v44).cardinality())))) : (_916_cf4)));
      }
      if (false) {
        let _931_v45;
        _931_v45 = true;
        _931_v45 = !((_this).f7);
        let _932_v46;
        _932_v46 = _dafny.Map.Empty.slice().updateUnsafe(_931_v45,true);
        let _933_v47;
        _933_v47 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).fm1(p3, _932_v46, (_this).f6, globalState));
        _931_v45 = ((_this).f8).isEqualTo((((_933_v47).contains(_931_v45)) ? ((_933_v47).get(_931_v45)) : (new BigNumber(((_this).f5).length))));
        (globalState).f1 = (((p0) ? ((_this).f8) : ((_this).f8))).multipliedBy((_this).f8);
        _931_v45 = !((_this).f6);
        let _934_v48;
        let _nw176 = new _module.C3();
        _nw176.__ctor((_this).f7, p0, (_this).f8, p0, _931_v45, p0, _this.f4, _dafny.Seq.Concat((_this).f5, (_this).f5));
        _934_v48 = _nw176;
      } else {
        if (_module.__default.fm3(globalState)) {
          let _935_v49;
          _935_v49 = _dafny.Set.fromElements((_this).f8, (((_this).f7) ? ((_this).f8) : ((_this).f8)));
          _935_v49 = function () {
            let _coll19 = new _dafny.Set();
            for (const _compr_19 of _dafny.IntegerRange(new BigNumber(190), new BigNumber(516))) {
              let _936_v50 = _compr_19;
              if (((new BigNumber(190)).isLessThanOrEqualTo(_936_v50)) && ((_936_v50).isLessThan(new BigNumber(516)))) {
                _coll19.add((_936_v50).plus(new BigNumber((p2).cardinality())));
              }
            }
            return _coll19;
          }();
          let _937_v51;
          _937_v51 = true;
          let _rhs125 = _this.f4;
          let _rhs126 = (_this).f7;
          let _lhs94 = _this;
          _lhs94.f4 = _rhs125;
          _937_v51 = _rhs126;
          r1 = (_this).f8;
          let _index139 = _module.__default.safeIndex(new BigNumber(107), new BigNumber(((_this).f11).length));
          ((_this).f11)[_index139] = (_this).f8;
          let _index140 = _module.__default.safeIndex(new BigNumber(107), new BigNumber(((_this).f11).length));
          ((_this).f11)[_index140] = (_this).f8;
          let _938_v52;
          _938_v52 = _module.D3.create_DC10((_this).f5);
          let _939_v53;
          _939_v53 = _dafny.MultiSet.fromElements((_this).f5, _dafny.Seq.Concat((_938_v52).dtor_cf19, (_this).f5), (_this).f5);
          _939_v53 = _939_v53;
        } else {
          r1 = (_module.__default.safeModuloInt(new BigNumber(((_this).f5).length), new BigNumber(134))).plus(_module.__default.safeModuloInt((_this).f8, (_this).f8));
          let _index141 = _module.__default.safeIndex(new BigNumber(694), new BigNumber(((_this).f11).length));
          ((_this).f11)[_index141] = (_this).f8;
          let _index142 = _module.__default.safeIndex(new BigNumber(694), new BigNumber(((_this).f11).length));
          ((_this).f11)[_index142] = ((_this).f8).plus(((_this).f8).plus((_this).f8));
          let _940_v54;
          _940_v54 = false;
          let _941_v55;
          _941_v55 = _dafny.Set.fromElements(((_this).f11)[_module.__default.safeIndex(new BigNumber(694), new BigNumber(((_this).f11).length))]);
          let _942_v56;
          _942_v56 = _dafny.Seq.of(((_this).f11)[_module.__default.safeIndex(new BigNumber(694), new BigNumber(((_this).f11).length))], ((_this).f11)[_module.__default.safeIndex(new BigNumber(694), new BigNumber(((_this).f11).length))], (_this).f8, new BigNumber((_941_v55).length));
          let _943_v57;
          _943_v57 = _dafny.Seq.of(_module.__default.fm3(globalState));
          let _944_v58;
          _944_v58 = _dafny.Set.fromElements(_942_v56);
          let _index143 = _module.__default.safeIndex(new BigNumber(694), new BigNumber(((_this).f11).length));
          let _rhs127 = !(((!(_module.__default.fm7(new _dafny.CodePoint('u'.codePointAt(0)), ((_this).f11)[_module.__default.safeIndex(new BigNumber(694), new BigNumber(((_this).f11).length))], p0, globalState))) ? ((_this).f9) : (((_942_v56)[_module.__default.safeIndex((_this).f8, new BigNumber((_942_v56).length))]).isLessThan(new BigNumber(-550)))));
          let _rhs128 = _module.__default.fm7(_this.f4, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_943_v57, _dafny.Seq.update(_943_v57, _module.__default.safeIndex(new BigNumber((_943_v57).length), new BigNumber((_943_v57).length)), !((_this).f6)))).length))), !((_this).f6), globalState);
          let _rhs129 = (_this).f8;
          let _rhs130 = ((_this).f11)[_module.__default.safeIndex(new BigNumber(694), new BigNumber(((_this).f11).length))];
          let _rhs131 = (_944_v58).equals(_944_v58);
          let _lhs95 = globalState;
          let _lhs96 = (_this).f11;
          let _lhs97 = _module.__default.safeIndex(new BigNumber(694), new BigNumber(((_this).f11).length));
          _940_v54 = _rhs127;
          _940_v54 = _rhs128;
          _lhs95.f1 = _rhs129;
          _lhs96[_lhs97] = _rhs130;
          _940_v54 = _rhs131;
          _940_v54 = (_941_v55).IsSubsetOf(_941_v55);
          _940_v54 = p0;
        }
        let _945_v59;
        _945_v59 = _dafny.Seq.of((_this).f9, (_this).f9, (_this).f6, (_this).f6);
        if ((_945_v59)[_module.__default.safeIndex((_this).f8, new BigNumber((_945_v59).length))]) {
          _module.__default.m0(globalState);
          let _946_v60;
          _946_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f7);
          let _947_v61;
          _947_v61 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,p1);
          let _948_v62;
          let _nw177 = new _module.C3();
          _nw177.__ctor((_this).f7, (_this).f7, (_this).f8, p1, false, p3, _this.f4, (_this).f5);
          _948_v62 = _nw177;
          let _949_v63;
          _949_v63 = _dafny.MultiSet.fromElements((_this).fm1((_this).f7, _946_v60, (_this).f7, globalState), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_948_v62)).length));
          let _950_v64;
          _950_v64 = _dafny.Seq.of((_this).f8, new BigNumber((_949_v63).cardinality()));
          let _951_v65;
          _951_v65 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_950_v64).length));
          let _952_v66;
          _952_v66 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber((p2).cardinality()));
          let _953_v67;
          _953_v67 = _dafny.Seq.of(_946_v60);
          let _954_v68;
          _954_v68 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(237),(_953_v67)[_module.__default.safeIndex((_948_v62).f8, new BigNumber((_953_v67).length))]);
          let _955_v69;
          let _nw178 = Array((new BigNumber(29)).toNumber());
          _nw178[(_dafny.ZERO).toNumber()] = (_this).f8;
          _nw178[(_dafny.ONE).toNumber()] = ((_this).fm1(p0, _946_v60, p3, globalState)).plus((_this).f8);
          _nw178[(new BigNumber(2)).toNumber()] = (_this).f8;
          _nw178[(new BigNumber(3)).toNumber()] = (_this).f8;
          _nw178[(new BigNumber(4)).toNumber()] = (_this).f8;
          _nw178[(new BigNumber(5)).toNumber()] = (_this).f8;
          _nw178[(new BigNumber(6)).toNumber()] = (_this).f8;
          _nw178[(new BigNumber(7)).toNumber()] = (_this).f8;
          _nw178[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt((_this).f8, new BigNumber((_947_v61).length));
          _nw178[(new BigNumber(9)).toNumber()] = (_this).f8;
          _nw178[(new BigNumber(10)).toNumber()] = ((_this).f8).plus(new BigNumber((_951_v65).length));
          _nw178[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update((_this).f5, _module.__default.safeIndex((_this).f8, new BigNumber(((_this).f5).length)), _948_v62.f4), (_this).f5)).length);
          _nw178[(new BigNumber(12)).toNumber()] = (_this).f8;
          _nw178[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus((_948_v62).f8);
          _nw178[(new BigNumber(14)).toNumber()] = (_948_v62).f8;
          _nw178[(new BigNumber(15)).toNumber()] = (((_952_v66).contains((_948_v62).f8)) ? ((_952_v66).get((_948_v62).f8)) : (new BigNumber(337)));
          _nw178[(new BigNumber(16)).toNumber()] = (_this).fm1((_this).f6, _946_v60, (_this).f7, globalState);
          _nw178[(new BigNumber(17)).toNumber()] = (_this).f8;
          _nw178[(new BigNumber(18)).toNumber()] = ((_this).f8).multipliedBy((_this).f8);
          _nw178[(new BigNumber(19)).toNumber()] = new BigNumber((((((_954_v68).contains((_this).f8)) ? ((_954_v68).get((_this).f8)) : (_946_v60))).update(!(!((_this).f6)), (_948_v62).f7)).length);
          _nw178[(new BigNumber(20)).toNumber()] = (_948_v62).f8;
          _nw178[(new BigNumber(21)).toNumber()] = new BigNumber(((p2).Difference(_dafny.MultiSet.FromArray(_945_v59))).cardinality());
          _nw178[(new BigNumber(22)).toNumber()] = (_948_v62).f8;
          _nw178[(new BigNumber(23)).toNumber()] = (_this).f8;
          _nw178[(new BigNumber(24)).toNumber()] = ((_this).f8).multipliedBy((_948_v62).f8);
          _nw178[(new BigNumber(25)).toNumber()] = new BigNumber(((_951_v65).Merge(_951_v65)).length);
          _nw178[(new BigNumber(26)).toNumber()] = (_dafny.ZERO).minus((_this).f8);
          _nw178[(new BigNumber(27)).toNumber()] = (_this).f8;
          _nw178[(new BigNumber(28)).toNumber()] = (_948_v62).f8;
          _955_v69 = _nw178;
          _955_v69 = _955_v69;
          let _956_v70;
          _956_v70 = _dafny.Map.Empty.slice().updateUnsafe((_948_v62).f8,_module.__default.fm26(p3, (_948_v62).f8, _948_v62.f4, globalState));
          let _957_v71;
          _957_v71 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_956_v70);
          let _958_v72;
          let _nw179 = Array((new BigNumber(7)).toNumber());
          _nw179[(_dafny.ZERO).toNumber()] = (_957_v71).update(_948_v62.f4, _956_v70);
          _nw179[(_dafny.ONE).toNumber()] = _957_v71;
          _nw179[(new BigNumber(2)).toNumber()] = (_957_v71).update(_this.f4, _956_v70);
          _nw179[(new BigNumber(3)).toNumber()] = _957_v71;
          _nw179[(new BigNumber(4)).toNumber()] = _957_v71;
          _nw179[(new BigNumber(5)).toNumber()] = _957_v71;
          _nw179[(new BigNumber(6)).toNumber()] = _957_v71;
          _958_v72 = _nw179;
          let _index144 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_958_v72).length));
          (_958_v72)[_index144] = _957_v71;
          let _index145 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_958_v72).length));
          (_958_v72)[_index145] = (_957_v71).Merge(_957_v71);
          let _959_v73;
          let _nw180 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
          _959_v73 = _nw180;
          _959_v73 = _959_v73;
          let _960_v74;
          _960_v74 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f9);
          let _961_v75;
          _961_v75 = _dafny.Map.Empty.slice().updateUnsafe(_960_v74,p1);
          let _index146 = _module.__default.safeIndex(new BigNumber(223), new BigNumber(((_this).f11).length));
          ((_this).f11)[_index146] = (_948_v62).f8;
          let _index147 = _module.__default.safeIndex(new BigNumber(223), new BigNumber(((_this).f11).length));
          let _rhs132 = _module.__default.fm8(_module.__default.safeDivisionInt((_this).f8, (_this).f8), false, ((_948_v62).f8).isLessThanOrEqualTo((_948_v62).f8), globalState);
          let _rhs133 = (_948_v62).f8;
          let _rhs134 = _961_v75;
          let _rhs135 = _module.__default.safeModuloInt((_this).fm1(p3, _dafny.Map.Empty.slice().updateUnsafe((_948_v62).f9,(_948_v62).f6), (_this).f9, globalState), (_948_v62).f8);
          let _rhs136 = _module.__default.safeModuloInt(((_this).f8).multipliedBy((_948_v62).f8), (_this).f8);
          let _lhs98 = globalState;
          let _lhs99 = (_this).f11;
          let _lhs100 = _module.__default.safeIndex(new BigNumber(223), new BigNumber(((_this).f11).length));
          r0 = _rhs132;
          _lhs98.f1 = _rhs133;
          _961_v75 = _rhs134;
          _lhs99[_lhs100] = _rhs135;
          r0 = _rhs136;
        } else {
          let _962_v76;
          _962_v76 = true;
          _962_v76 = true;
          _962_v76 = !(_dafny.Seq.IsProperPrefixOf((_this).f5, (_this).f5));
          let _index148 = _module.__default.safeIndex(new BigNumber(463), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index148] = p1;
          let _index149 = _module.__default.safeIndex(new BigNumber(463), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index149] = false;
          let _963_v77;
          let _nw181 = new _module.C0();
          _nw181.__ctor(_962_v76, ((_this).f10)[_module.__default.safeIndex(new BigNumber(463), new BigNumber(((_this).f10).length))], _this.f4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(38)), function (_964_i2) {
            return _this.f4;
          }));
          _963_v77 = _nw181;
          let _965_v78;
          _965_v78 = _dafny.Map.Empty.slice().updateUnsafe(_963_v77,_962_v76);
          let _966_v79;
          _966_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(((_965_v78).contains(_963_v77)) ? ((_965_v78).get(_963_v77)) : (!((_this).f7))));
          (_this).m4((_this).f8, _966_v79, globalState);
          let _index150 = _module.__default.safeIndex(new BigNumber(463), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index150] = !((_945_v59)[_module.__default.safeIndex((_this).f8, new BigNumber((_945_v59).length))]);
        }
        (_this).f4 = _this.f4;
        let _967_v80;
        _967_v80 = false;
        _967_v80 = p0;
        let _968_v81;
        _968_v81 = _dafny.Seq.of((_this).f8, (_dafny.ZERO).minus((_this).f8));
        let _969_v82;
        _969_v82 = _dafny.Set.fromElements((_this).f6, (_this).f7);
        let _970_v83;
        _970_v83 = _module.D3.create_DC11(false, new BigNumber((_968_v81).length), _969_v82);
        let _971_v84;
        _971_v84 = _dafny.Seq.of((_970_v83).dtor_cf22);
        let _972_v85;
        _972_v85 = _module.D0.create_DC1(_967_v80, new BigNumber(-224), (_this).f6, (_this).f8);
        (globalState).f1 = (new BigNumber((_971_v84).length)).multipliedBy(_module.__default.fm8((_972_v85).dtor_cf3, !(p3), p1, globalState));
      }
      let _973_v86;
      _973_v86 = _dafny.MultiSet.fromElements((_this).f8, (_this).f8);
      let _974_v87;
      _974_v87 = _module.D6.create_DC16(_973_v86);
      r0 = (_dafny.ZERO).minus(new BigNumber(((_974_v87).dtor_cf30).cardinality()));
      r1 = (_dafny.ZERO).minus((_this).f8);
      if ((p2).IsDisjointFrom(p2)) {
        let _975_v88;
        let _nw182 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Set.Empty);
        _975_v88 = _nw182;
        let _976_v89;
        _976_v89 = _dafny.Set.fromElements(p0, true);
        let _index151 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_975_v88).length));
        (_975_v88)[_index151] = (((_this).f6) ? (_module.__default.fm0((_this).f9, new BigNumber(924), new BigNumber((p2).cardinality()), p3, globalState)) : (_976_v89));
        let _977_v90;
        _977_v90 = false;
        let _978_v91;
        let _nw183 = new _module.C4();
        _nw183.__ctor(new BigNumber(644), p3, !((_this).f6), false, _this.f4, (_this).f5);
        _978_v91 = _nw183;
        let _979_v92;
        _979_v92 = _dafny.Map.Empty.slice().updateUnsafe(_978_v91,(_this).f11);
        let _980_v93;
        let _nw184 = Array((new BigNumber(27)).toNumber());
        _nw184[(_dafny.ZERO).toNumber()] = (_this).f11;
        _nw184[(_dafny.ONE).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(2)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(3)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(4)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(5)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(6)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(7)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(8)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(9)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(10)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(11)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(12)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(13)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(14)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(15)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(16)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(17)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(18)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(19)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(20)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(21)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(22)).toNumber()] = ((p0) ? ((_this).f11) : ((_this).f11));
        _nw184[(new BigNumber(23)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(24)).toNumber()] = (_this).f11;
        _nw184[(new BigNumber(25)).toNumber()] = (((_979_v92).contains(_978_v91)) ? ((_979_v92).get(_978_v91)) : ((_this).f11));
        _nw184[(new BigNumber(26)).toNumber()] = (_this).f11;
        _980_v93 = _nw184;
        let _index152 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length));
        (_980_v93)[_index152] = (_this).f11;
        let _981_v94;
        let _nw185 = new _module.C0();
        _nw185.__ctor(p1, p1, _978_v91.f4, _dafny.Seq.Concat((_978_v91).f5, (_this).f5));
        _981_v94 = _nw185;
        let _index153 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_975_v88).length));
        let _index154 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length));
        let _rhs137 = _976_v89;
        let _rhs138 = _978_v91.f4;
        let _rhs139 = !((_this).f7);
        let _rhs140 = ((((_this).f8).isLessThanOrEqualTo((_this).f8)) ? ((_this).f11) : ((_this).f11));
        let _rhs141 = _981_v94;
        let _lhs101 = _975_v88;
        let _lhs102 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_975_v88).length));
        let _lhs103 = _this;
        let _lhs104 = _980_v93;
        let _lhs105 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length));
        _lhs101[_lhs102] = _rhs137;
        _lhs103.f4 = _rhs138;
        _977_v90 = _rhs139;
        _lhs104[_lhs105] = _rhs140;
        _981_v94 = _rhs141;
        let _arr0 = (_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))];
        let _index155 = _module.__default.safeIndex(new BigNumber(341), new BigNumber(((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))]).length));
        _arr0[_index155] = (_this).f8;
        let _982_v95;
        _982_v95 = _dafny.Map.Empty.slice().updateUnsafe(_973_v86,(_978_v91).f9);
        let _arr1 = (_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))];
        let _index156 = _module.__default.safeIndex(new BigNumber(341), new BigNumber(((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))]).length));
        _arr1[_index156] = new BigNumber(((_982_v95).Merge(_982_v95)).length);
        (globalState).f1 = ((_this).f8).multipliedBy((_978_v91).f8);
        if ((_978_v91).f9) {
          (_978_v91).f4 = _this.f4;
          let _983_v96;
          _983_v96 = _module.D1.create_DC4(_978_v91.f4, (_this).f8, _978_v91.f4, (_this).f8, (_978_v91).f6);
          let _984_v97;
          _984_v97 = _dafny.Seq.of((_this).f8);
          let _985_v98;
          let _nw186 = new _module.C4();
          _nw186.__ctor((_dafny.ZERO).minus((_978_v91).f8), ((_978_v91).f9) && (false), (_983_v96).dtor_cf12, (_973_v86).IsSubsetOf(_dafny.MultiSet.fromElements((_978_v91).f8, (_this).f8)), ((_978_v91).f5)[_module.__default.safeIndex(new BigNumber((_984_v97).length), new BigNumber(((_978_v91).f5).length))], (_this).f5);
          _985_v98 = _nw186;
          let _986_v99;
          _986_v99 = _dafny.Seq.of((_this).f10);
          let _index157 = _module.__default.safeIndex(new BigNumber(492), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index157] = (((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))])[_module.__default.safeIndex(new BigNumber(341), new BigNumber(((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))]).length))]).isLessThan(((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))])[_module.__default.safeIndex(new BigNumber(341), new BigNumber(((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))]).length))]);
          let _987_v100;
          _987_v100 = _dafny.Seq.of((_this).f9);
          let _988_v101;
          let _nw187 = new _module.C1();
          _nw187.__ctor((_this).f8, (_987_v100)[_module.__default.safeIndex((_this).f8, new BigNumber((_987_v100).length))], (_this).f6, p3, _978_v91.f4, (_this).f5);
          _988_v101 = _nw187;
          let _989_v102;
          _989_v102 = _module.D0.create_DC1(!(p1), (_978_v91).f8, p0, (_978_v91).f8);
          let _990_v103;
          _990_v103 = _dafny.Set.fromElements(_989_v102, _989_v102);
          let _991_v104;
          _991_v104 = _dafny.Set.fromElements((_this).f8);
          let _992_v105;
          _992_v105 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? (_988_v101) : (_988_v101)),new BigNumber((_dafny.Seq.Concat(_module.__default.fm14(new BigNumber((_990_v103).length), new BigNumber(((_988_v101).f5).length), (_978_v91).f6, (_988_v101).f9, globalState), _dafny.Seq.update((_988_v101).f5, _module.__default.safeIndex(new BigNumber(718), new BigNumber(((_988_v101).f5).length)), ((_this).f5)[_module.__default.safeIndex(new BigNumber((_991_v104).length), new BigNumber(((_this).f5).length))]))).length));
          let _index158 = _module.__default.safeIndex(new BigNumber(492), new BigNumber(((_this).f10).length));
          let _index159 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length));
          let _rhs142 = _986_v99;
          let _rhs143 = (_988_v101).f6;
          let _rhs144 = (_992_v105).Merge(_992_v105);
          let _rhs145 = (_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))];
          let _rhs146 = (new BigNumber((_dafny.Seq.Concat(_987_v100, _987_v100)).length)).multipliedBy((_978_v91).f8);
          let _lhs106 = (_this).f10;
          let _lhs107 = _module.__default.safeIndex(new BigNumber(492), new BigNumber(((_this).f10).length));
          let _lhs108 = _980_v93;
          let _lhs109 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length));
          let _lhs110 = globalState;
          _986_v99 = _rhs142;
          _lhs106[_lhs107] = _rhs143;
          _992_v105 = _rhs144;
          _lhs108[_lhs109] = _rhs145;
          _lhs110.f1 = _rhs146;
          (_this).f4 = _988_v101.f4;
          let _993_v106;
          let _nw188 = Array((new BigNumber(21)).toNumber());
          _nw188[(_dafny.ZERO).toNumber()] = p3;
          _nw188[(_dafny.ONE).toNumber()] = _module.__default.fm3(globalState);
          _nw188[(new BigNumber(2)).toNumber()] = p1;
          _nw188[(new BigNumber(3)).toNumber()] = p1;
          _nw188[(new BigNumber(4)).toNumber()] = (_978_v91).f7;
          _nw188[(new BigNumber(5)).toNumber()] = !(!(p1));
          _nw188[(new BigNumber(6)).toNumber()] = (_this).f6;
          _nw188[(new BigNumber(7)).toNumber()] = p3;
          _nw188[(new BigNumber(8)).toNumber()] = false;
          _nw188[(new BigNumber(9)).toNumber()] = _module.__default.fm3(globalState);
          _nw188[(new BigNumber(10)).toNumber()] = (_this).f7;
          _nw188[(new BigNumber(11)).toNumber()] = (_this).f9;
          _nw188[(new BigNumber(12)).toNumber()] = false;
          _nw188[(new BigNumber(13)).toNumber()] = p3;
          _nw188[(new BigNumber(14)).toNumber()] = false;
          _nw188[(new BigNumber(15)).toNumber()] = (_978_v91).f6;
          _nw188[(new BigNumber(16)).toNumber()] = (_978_v91).f9;
          _nw188[(new BigNumber(17)).toNumber()] = _module.__default.fm7(_978_v91.f4, new BigNumber(-833), p1, globalState);
          _nw188[(new BigNumber(18)).toNumber()] = (_978_v91).f9;
          _nw188[(new BigNumber(19)).toNumber()] = (_978_v91).f7;
          _nw188[(new BigNumber(20)).toNumber()] = (_978_v91).f6;
          _993_v106 = _nw188;
          let _994_v107;
          let _init18 = ((_995_v100) => function (_996_i3) {
            return (_module.D8.create_DC22(_995_v100)).dtor_cf41;
          })(_987_v100);
          let _nw189 = Array((new BigNumber(6)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw189.length); _i0_18++) {
            _nw189[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _994_v107 = _nw189;
          let _997_v108;
          _997_v108 = _module.D1.create_DC3(_994_v107);
          let _pat_let_tv16 = _994_v107;
          let _998_v109;
          let _nw190 = Array((new BigNumber(25)).toNumber());
          _nw190[(_dafny.ZERO).toNumber()] = _997_v108;
          _nw190[(_dafny.ONE).toNumber()] = _997_v108;
          _nw190[(new BigNumber(2)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(3)).toNumber()] = _module.D1.create_DC3(_994_v107);
          _nw190[(new BigNumber(4)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(5)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(6)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(7)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(8)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(9)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(10)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(11)).toNumber()] = _module.D1.create_DC3(_994_v107);
          _nw190[(new BigNumber(12)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(13)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(14)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(15)).toNumber()] = function (_pat_let26_0) {
            return function (_999_dt__update__tmp_h2) {
              return function (_pat_let27_0) {
                return function (_1000_dt__update_hcf7_h0) {
                  return _module.D1.create_DC3(_1000_dt__update_hcf7_h0);
                }(_pat_let27_0);
              }(_pat_let_tv16);
            }(_pat_let26_0);
          }(_module.D1.create_DC3(_994_v107));
          _nw190[(new BigNumber(16)).toNumber()] = _module.D1.create_DC3(_994_v107);
          _nw190[(new BigNumber(17)).toNumber()] = _module.D1.create_DC3(_994_v107);
          _nw190[(new BigNumber(18)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(19)).toNumber()] = _module.D1.create_DC3(_994_v107);
          _nw190[(new BigNumber(20)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(21)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(22)).toNumber()] = _997_v108;
          _nw190[(new BigNumber(23)).toNumber()] = _module.D1.create_DC3(_994_v107);
          _nw190[(new BigNumber(24)).toNumber()] = _997_v108;
          _998_v109 = _nw190;
          let _1001_v110;
          let _nw191 = new _module.C2();
          _nw191.__ctor(_993_v106, _998_v109);
          _1001_v110 = _nw191;
          let _1002_v111;
          _1002_v111 = _dafny.Set.fromElements((_978_v91).f5);
          let _index160 = _module.__default.safeIndex(new BigNumber(492), new BigNumber(((_this).f10).length));
          let _rhs147 = (_1002_v111).IsSubsetOf(((true) ? (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("lhuqmt"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(737)), ((_1003_v91) => function (_1004_i4) {
            return _1003_v91.f4;
          })(_978_v91)))) : (_1002_v111)));
          let _rhs148 = _1001_v110;
          let _lhs111 = (_this).f10;
          let _lhs112 = _module.__default.safeIndex(new BigNumber(492), new BigNumber(((_this).f10).length));
          _lhs111[_lhs112] = _rhs147;
          _1001_v110 = _rhs148;
        } else {
          let _1005_v112;
          _1005_v112 = _dafny.Seq.of(new BigNumber(910));
          let _1006_v113;
          _1006_v113 = _dafny.Seq.of(new BigNumber((_1005_v112).length));
          let _1007_v114;
          _1007_v114 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_973_v86).IsDisjointFrom(_dafny.MultiSet.FromArray(_1006_v113)));
          let _1008_v115;
          _1008_v115 = _dafny.Map.Empty.slice().updateUnsafe((_978_v91).f6,(_this).f6);
          _1007_v114 = (_1007_v114).update((_this).f10, (((_1008_v115).contains(false)) ? ((_1008_v115).get(false)) : ((_978_v91).f9)));
          (globalState).f1 = (_978_v91).f8;
          let _1009_v116;
          let _init19 = ((_1010_v91, _1011_v93) => function (_1012_i5) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(_1010_v91.f4, (_1010_v91).f8, _1010_v91.f4, (_this).f8, (_1010_v91).f7),(_1010_v91).f8)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(_1010_v91.f4, ((_1011_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_1011_v93).length))])[_module.__default.safeIndex(new BigNumber(341), new BigNumber(((_1011_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_1011_v93).length))]).length))], _this.f4, (_1010_v91).f8, (_1010_v91).f6),((_1011_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_1011_v93).length))])[_module.__default.safeIndex(new BigNumber(341), new BigNumber(((_1011_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_1011_v93).length))]).length))]));
          })(_978_v91, _980_v93);
          let _nw192 = Array((_dafny.ONE).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw192.length); _i0_19++) {
            _nw192[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _1009_v116 = _nw192;
          let _1013_v117;
          _1013_v117 = _module.D1.create_DC4(_978_v91.f4, (_this).f8, _978_v91.f4, (_this).f8, false);
          let _1014_v118;
          _1014_v118 = _dafny.Map.Empty.slice().updateUnsafe(_1013_v117,(_978_v91).f8);
          let _index161 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_1009_v116).length));
          (_1009_v116)[_index161] = _1014_v118;
          let _1015_v119;
          _1015_v119 = _dafny.Seq.UnicodeFromString("ciyeo");
          let _1016_v121;
          let _nw193 = new _module.C0();
          _nw193.__ctor(_module.__default.fm3(globalState), _977_v90, new _dafny.CodePoint('f'.codePointAt(0)), (_978_v91).f5);
          _1016_v121 = _nw193;
          let _1017_v122;
          _1017_v122 = _dafny.MultiSet.fromElements(_1016_v121);
          let _1018_v123;
          _1018_v123 = _dafny.Map.Empty.slice().updateUnsafe(((_978_v91).f8).multipliedBy(((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))])[_module.__default.safeIndex(new BigNumber(341), new BigNumber(((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))]).length))]),(_this).f5);
          let _index162 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_1009_v116).length));
          let _rhs149 = ((p0) ? (function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of (_dafny.MultiSet.fromElements(_module.D1.create_DC4(_this.f4, new BigNumber((_1017_v122).cardinality()), _1016_v121.f4, new BigNumber(((_978_v91).f5).length), p0))).Elements) {
              let _1019_v120 = _compr_20;
              if ((_dafny.MultiSet.fromElements(_module.D1.create_DC4(_this.f4, new BigNumber((_1017_v122).cardinality()), _1016_v121.f4, new BigNumber(((_978_v91).f5).length), p0))).contains(_1019_v120)) {
                _coll20.push([_1019_v120,((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))])[_module.__default.safeIndex(new BigNumber(341), new BigNumber(((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))]).length))]]);
              }
            }
            return _coll20;
          }()) : (_dafny.Map.Empty.slice().updateUnsafe(_1013_v117,(_this).f8)));
          let _rhs150 = ((new BigNumber(((_978_v91).f5).length)).multipliedBy((_978_v91).f8)).isLessThanOrEqualTo(((_this).f8).plus((_978_v91).f8));
          let _rhs151 = (((_1018_v123).contains((_978_v91).f8)) ? ((_1018_v123).get((_978_v91).f8)) : ((_1016_v121).f5));
          let _lhs113 = _1009_v116;
          let _lhs114 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_1009_v116).length));
          _lhs113[_lhs114] = _rhs149;
          _977_v90 = _rhs150;
          _1015_v119 = _rhs151;
          let _index163 = _module.__default.safeIndex(new BigNumber(653), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index163] = !((((_this).f9) ? ((_978_v91).f7) : ((_978_v91).f9)));
          let _index164 = _module.__default.safeIndex(new BigNumber(653), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index164] = (_978_v91).f6;
          let _1020_v124;
          _1020_v124 = _dafny.Map.Empty.slice().updateUnsafe((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))],(_dafny.ZERO).minus((_978_v91).f8));
          let _index165 = _module.__default.safeIndex(new BigNumber(653), new BigNumber(((_this).f10).length));
          let _arr2 = (_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))];
          let _index166 = _module.__default.safeIndex(new BigNumber(341), new BigNumber(((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))]).length));
          let _rhs152 = (_1016_v121).f7;
          let _rhs153 = (new BigNumber((_1020_v124).length)).isLessThan(new BigNumber(((_978_v91).f5).length));
          let _rhs154 = _1015_v119;
          let _rhs155 = new BigNumber(339);
          let _lhs115 = (_this).f10;
          let _lhs116 = _module.__default.safeIndex(new BigNumber(653), new BigNumber(((_this).f10).length));
          let _lhs117 = (_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))];
          let _lhs118 = _module.__default.safeIndex(new BigNumber(341), new BigNumber(((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))]).length));
          _lhs115[_lhs116] = _rhs152;
          _977_v90 = _rhs153;
          _1015_v119 = _rhs154;
          _lhs117[_lhs118] = _rhs155;
        }
        let _arr3 = (_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))];
        let _index167 = _module.__default.safeIndex(new BigNumber(341), new BigNumber(((_980_v93)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_980_v93).length))]).length));
        _arr3[_index167] = (_978_v91).f8;
      } else {
        let _1021_v125;
        _1021_v125 = _dafny.MultiSet.fromElements((_this).f9);
        _1021_v125 = _dafny.MultiSet.fromElements((function (_pat_let28_0) {
          return function (_1022_dt__update__tmp_h3) {
            return function (_pat_let29_0) {
              return function (_1023_dt__update_hcf43_h0) {
                return _module.D8.create_DC23((_1022_dt__update__tmp_h3).dtor_cf42, _1023_dt__update_hcf43_h0);
              }(_pat_let29_0);
            }(true);
          }(_pat_let28_0);
        }(_module.D8.create_DC23(new BigNumber(937), p3))).dtor_cf43);
        let _1024_v126;
        _1024_v126 = true;
        _1024_v126 = p1;
        let _index168 = _module.__default.safeIndex(new BigNumber(757), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index168] = _dafny.areEqual((_this).f5, (_this).f5);
        let _1025_v127;
        let _nw194 = Array((new BigNumber(9)).toNumber());
        _nw194[(_dafny.ZERO).toNumber()] = (_this).f10;
        _nw194[(_dafny.ONE).toNumber()] = (_this).f10;
        _nw194[(new BigNumber(2)).toNumber()] = (_this).f10;
        _nw194[(new BigNumber(3)).toNumber()] = (_this).f10;
        _nw194[(new BigNumber(4)).toNumber()] = (_this).f10;
        _nw194[(new BigNumber(5)).toNumber()] = (_this).f10;
        _nw194[(new BigNumber(6)).toNumber()] = (_this).f10;
        _nw194[(new BigNumber(7)).toNumber()] = (_this).f10;
        _nw194[(new BigNumber(8)).toNumber()] = (_this).f10;
        _1025_v127 = _nw194;
        let _1026_v128;
        let _nw195 = Array((_dafny.ONE).toNumber()).fill(_module.D6.Default());
        _1026_v128 = _nw195;
        let _index169 = _module.__default.safeIndex(new BigNumber(432), new BigNumber((_1026_v128).length));
        (_1026_v128)[_index169] = _974_v87;
        let _1027_v129;
        _1027_v129 = _dafny.Seq.of(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f8), (_this).f8));
        let _index170 = _module.__default.safeIndex(new BigNumber(757), new BigNumber(((_this).f10).length));
        let _index171 = _module.__default.safeIndex(new BigNumber(432), new BigNumber((_1026_v128).length));
        let _rhs156 = !((new BigNumber(718)).isEqualTo((_this).f8));
        let _rhs157 = _1025_v127;
        let _rhs158 = _974_v87;
        let _rhs159 = _dafny.Seq.Concat(_1027_v129, _1027_v129);
        let _lhs119 = (_this).f10;
        let _lhs120 = _module.__default.safeIndex(new BigNumber(757), new BigNumber(((_this).f10).length));
        let _lhs121 = _1026_v128;
        let _lhs122 = _module.__default.safeIndex(new BigNumber(432), new BigNumber((_1026_v128).length));
        _lhs119[_lhs120] = _rhs156;
        _1025_v127 = _rhs157;
        _lhs121[_lhs122] = _rhs158;
        _1027_v129 = _rhs159;
        let _1028_v130;
        _1028_v130 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,((true) ? (!(!(_1024_v126))) : (true)));
        let _1029_v131;
        let _init20 = ((_1030_v126) => function (_1031_i6) {
          return _1030_v126;
        })(_1024_v126);
        let _nw196 = Array((new BigNumber(27)).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw196.length); _i0_20++) {
          _nw196[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _1029_v131 = _nw196;
        let _1032_v132;
        _1032_v132 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f8);
        let _1033_v133;
        _1033_v133 = _module.D5.create_DC15(_1029_v131, p1, _1032_v132, (_this).f8);
        let _index172 = _module.__default.safeIndex(new BigNumber(757), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index172] = (((_1028_v130).contains(((_this).f8).minus((_this).f8))) ? ((_1028_v130).get(((_this).f8).minus((_this).f8))) : ((_1033_v133).dtor_cf27));
        let _1034_v134;
        _1034_v134 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber(-229));
        let _1035_v135;
        let _nw197 = new _module.C3();
        _nw197.__ctor(((_this).f10)[_module.__default.safeIndex(new BigNumber(757), new BigNumber(((_this).f10).length))], p1, (((_1034_v134).contains((_this).f8)) ? ((_1034_v134).get((_this).f8)) : (new BigNumber((_973_v86).cardinality()))), (_this).f6, _1024_v126, (_this).f7, _this.f4, (_this).f5);
        _1035_v135 = _nw197;
        let _1036_v136;
        _1036_v136 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-343),_1035_v135);
        let _1037_v137;
        _1037_v137 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(((_1028_v130).contains((_this).f8)) ? ((_1028_v130).get((_this).f8)) : ((_1035_v135).f13)));
        _1036_v136 = (_1036_v136).update((_this).fm1(p0, _1037_v137, p0, globalState), (((_this).f6) ? (_1035_v135) : (_1035_v135)));
      }
      r1 = _module.__default.safeDivisionInt((_this).f8, (_this).f8);
      let _1038_v138;
      _1038_v138 = _dafny.MultiSet.fromElements((_this).f11, (_this).f11, (_this).f11, (_this).f11);
      let _1039_v139;
      _1039_v139 = _dafny.Seq.of(_1038_v138);
      r0 = (_dafny.ZERO).minus(new BigNumber(((_1038_v138).Intersect(((_1039_v139)[_module.__default.safeIndex((_this).f8, new BigNumber((_1039_v139).length))]).Union(_1038_v138))).cardinality()));
      r1 = (_this).f8;
      let _1040_v140;
      let _nw198 = Array((new BigNumber(16)).toNumber()).fill([]);
      _1040_v140 = _nw198;
      r2 = _1040_v140;
      return [r0, r1, r2];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _1041_v0;
      _1041_v0 = _dafny.Seq.of((_this).f7);
      let _1042_v1;
      _1042_v1 = _dafny.Seq.of(false, (_this).f6, (_1041_v0)[_module.__default.safeIndex(p0, new BigNumber((_1041_v0).length))], true);
      let _1043_v2;
      _1043_v2 = _dafny.MultiSet.fromElements((_this).f7, (_this).f9, (_this).f6, (_this).f9, true);
      if (_module.__default.fm27(_module.__default.safeDivisionInt(new BigNumber((_1042_v1).length), new BigNumber((_1043_v2).cardinality())), (_this).f6, (_this).f7, globalState)) {
        let _1044_v3;
        _1044_v3 = _dafny.Seq.UnicodeFromString("bgo");
        let _1045_v4;
        _1045_v4 = _dafny.Set.fromElements((_this).f9);
        let _1046_v5;
        _1046_v5 = _module.D3.create_DC11((_this).f6, (_this).f8, _1045_v4);
        let _1047_v6;
        _1047_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Seq.Create(_module.__default.abs(new BigNumber(281)), function (_1048_i0) {
          return _this.f4;
        }));
        _1044_v3 = _dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm28((_this).f8, (_this).f7, _module.__default.fm29(globalState), _1046_v5, globalState), (((_1047_v6).contains((_this).f9)) ? ((_1047_v6).get((_this).f9)) : (_1044_v3))), (_this).f5);
        _1044_v3 = _1044_v3;
        _1044_v3 = _dafny.Seq.UnicodeFromString("xolxenqc");
        let _1049_v7;
        let _nw199 = Array((new BigNumber(22)).toNumber()).fill(_module.D1.Default());
        _1049_v7 = _nw199;
        let _1050_v8;
        let _nw200 = new _module.C2();
        _nw200.__ctor((_this).f10, _1049_v7);
        _1050_v8 = _nw200;
        let _index173 = _module.__default.safeIndex(new BigNumber(111), new BigNumber(((_1050_v8).f14).length));
        ((_1050_v8).f14)[_index173] = (_this).f7;
        let _1051_v9;
        let _init21 = function (_1052_i1) {
          return _module.D6.create_DC17(_dafny.Seq.UnicodeFromString("xy"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-640)), function (_1053_i2) {
  return _this.f4;
}));
        };
        let _nw201 = Array((new BigNumber(6)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw201.length); _i0_21++) {
          _nw201[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _1051_v9 = _nw201;
        let _1054_v10;
        _1054_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1051_v9,_1050_v8);
        let _1055_v11;
        _1055_v11 = _dafny.Seq.of(_1051_v9);
        let _1056_v12;
        _1056_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber(((_this).f5).length));
        let _index174 = _module.__default.safeIndex(new BigNumber(111), new BigNumber(((_1050_v8).f14).length));
        let _rhs160 = (((_1054_v10).contains((_1055_v11)[_module.__default.safeIndex(new BigNumber((_1056_v12).length), new BigNumber((_1055_v11).length))])) ? ((_1054_v10).get((_1055_v11)[_module.__default.safeIndex(new BigNumber((_1056_v12).length), new BigNumber((_1055_v11).length))])) : (_1050_v8));
        let _rhs161 = (_this).f9;
        let _lhs123 = (_1050_v8).f14;
        let _lhs124 = _module.__default.safeIndex(new BigNumber(111), new BigNumber(((_1050_v8).f14).length));
        _1050_v8 = _rhs160;
        _lhs123[_lhs124] = _rhs161;
        let _1057_v13;
        _1057_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1044_v3,(_this).f6);
        let _index175 = _module.__default.safeIndex(new BigNumber(111), new BigNumber(((_1050_v8).f14).length));
        ((_1050_v8).f14)[_index175] = (((_1057_v13).contains(_dafny.Seq.Concat(_1044_v3, _1044_v3))) ? ((_1057_v13).get(_dafny.Seq.Concat(_1044_v3, _1044_v3))) : ((_this).f9));
      } else {
        let _1058_v14;
        _1058_v14 = false;
        _1058_v14 = _1058_v14;
        let _1059_v15;
        _1059_v15 = _dafny.MultiSet.fromElements((_this).f8);
        let _1060_v16;
        _1060_v16 = _dafny.Seq.of(p0);
        let _1061_v17;
        _1061_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1060_v16,p0);
        let _1062_v18;
        _1062_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1058_v14,(((_1059_v15).contains(new BigNumber((_1061_v17).length))) ? ((_1059_v15).get(new BigNumber((_1061_v17).length))) : (p0)));
        _1062_v18 = _1062_v18;
        let _1063_v19;
        _1063_v19 = _module.D3.create_DC10((_this).f5);
        let _1064_v20;
        _1064_v20 = _module.D3.create_DC12(_1063_v19);
        let _source11 = _1064_v20;
        if (_source11.is_DC11) {
          let _1065___mcc_h0 = (_source11).cf20;
          let _1066___mcc_h1 = (_source11).cf21;
          let _1067___mcc_h2 = (_source11).cf22;
          let _1068_cf22 = _1067___mcc_h2;
          let _1069_cf21 = _1066___mcc_h1;
          let _1070_cf20 = _1065___mcc_h0;
          let _1071_v21;
          let _nw202 = new _module.C0();
          _nw202.__ctor((_this).f6, _1070_cf20, ((false) ? (_this.f4) : (new _dafny.CodePoint('v'.codePointAt(0)))), _dafny.Seq.Concat((_this).f5, (_this).f5));
          _1071_v21 = _nw202;
          let _1072_v22;
          let _nw203 = Array((new BigNumber(16)).toNumber()).fill(_module.D1.Default());
          _1072_v22 = _nw203;
          let _1073_v23;
          let _nw204 = new _module.C2();
          _nw204.__ctor((_this).f10, _1072_v22);
          _1073_v23 = _nw204;
          _1069_cf21 = (_this).f8;
          (globalState).f1 = (_this).f8;
        } else if (_source11.is_DC10) {
          let _1074___mcc_h3 = (_source11).cf19;
          let _1075_cf19 = _1074___mcc_h3;
          let _1076_v24;
          _1076_v24 = _dafny.Set.fromElements((_this).f10);
          let _1077_v25;
          _1077_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,false);
          let _1078_v26;
          _1078_v26 = _module.D0.create_DC1((_this).f7, (new BigNumber((_1059_v15).cardinality())).minus((_this).f8), _module.__default.fm7(_this.f4, new BigNumber((_1076_v24).length), (_this).f9, globalState), new BigNumber((_1077_v25).length));
          let _rhs162 = _1078_v26;
          let _rhs163 = _1059_v15;
          let _rhs164 = false;
          _1078_v26 = _rhs162;
          _1059_v15 = _rhs163;
          _1058_v14 = _rhs164;
          let _index176 = _module.__default.safeIndex(new BigNumber(273), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index176] = (_this).f9;
          let _index177 = _module.__default.safeIndex(new BigNumber(273), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index177] = ((_1059_v15).update((_this).f8, _module.__default.abs((_this).f8))).IsSubsetOf(_1059_v15);
          let _1079_v27;
          let _nw205 = new _module.C4();
          _nw205.__ctor(new BigNumber(-331), (_this).f6, ((_this).f10)[_module.__default.safeIndex(new BigNumber(273), new BigNumber(((_this).f10).length))], ((_this).f8).isEqualTo(new BigNumber(629)), _this.f4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(112)), function (_1080_i3) {
            return _this.f4;
          }));
          _1079_v27 = _nw205;
          let _1081_v28;
          _1081_v28 = _module.D0.create_DC2(new BigNumber((_1060_v16).length), (_this).f6, _dafny.Seq.of((_this).f8, p0));
          let _1082_v29;
          _1082_v29 = _dafny.Map.Empty.slice().updateUnsafe((_1081_v28).dtor_cf6,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1059_v15).cardinality()),p0));
          let _index178 = _module.__default.safeIndex(new BigNumber(662), new BigNumber(((_this).f11).length));
          ((_this).f11)[_index178] = (new BigNumber((_1082_v29).length)).multipliedBy((_this).f8);
          let _1083_v30;
          _1083_v30 = _dafny.Set.fromElements(_dafny.MultiSet.FromArray(_1041_v0), _1043_v2, _1043_v2);
          let _index179 = _module.__default.safeIndex(new BigNumber(662), new BigNumber(((_this).f11).length));
          ((_this).f11)[_index179] = (new BigNumber((_1083_v30).length)).minus((_this).f8);
        } else {
          let _1084___mcc_h4 = (_source11).cf23;
          let _1085_cf23 = _1084___mcc_h4;
          let _1086_v31;
          let _nw206 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _1086_v31 = _nw206;
          _1086_v31 = (_this).f11;
          let _1087_v32;
          _1087_v32 = _module.D8.create_DC22(_dafny.Seq.of((_this).f6, (_this).f7));
          let _1088_v33;
          let _nw207 = Array((new BigNumber(29)).toNumber());
          _nw207[(_dafny.ZERO).toNumber()] = _1042_v1;
          _nw207[(_dafny.ONE).toNumber()] = _1041_v0;
          _nw207[(new BigNumber(2)).toNumber()] = _1042_v1;
          _nw207[(new BigNumber(3)).toNumber()] = (_1087_v32).dtor_cf41;
          _nw207[(new BigNumber(4)).toNumber()] = _module.__default.fm18((_this).f8, (_this).f6, globalState);
          _nw207[(new BigNumber(5)).toNumber()] = _1042_v1;
          _nw207[(new BigNumber(6)).toNumber()] = _1041_v0;
          _nw207[(new BigNumber(7)).toNumber()] = _1041_v0;
          _nw207[(new BigNumber(8)).toNumber()] = _1042_v1;
          _nw207[(new BigNumber(9)).toNumber()] = _1041_v0;
          _nw207[(new BigNumber(10)).toNumber()] = _1041_v0;
          _nw207[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_1058_v14);
          _nw207[(new BigNumber(12)).toNumber()] = _1041_v0;
          _nw207[(new BigNumber(13)).toNumber()] = _1041_v0;
          _nw207[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(true);
          _nw207[(new BigNumber(15)).toNumber()] = _1042_v1;
          _nw207[(new BigNumber(16)).toNumber()] = _1042_v1;
          _nw207[(new BigNumber(17)).toNumber()] = _1041_v0;
          _nw207[(new BigNumber(18)).toNumber()] = _1042_v1;
          _nw207[(new BigNumber(19)).toNumber()] = _1041_v0;
          _nw207[(new BigNumber(20)).toNumber()] = _dafny.Seq.of(_1058_v14);
          _nw207[(new BigNumber(21)).toNumber()] = _1041_v0;
          _nw207[(new BigNumber(22)).toNumber()] = _1042_v1;
          _nw207[(new BigNumber(23)).toNumber()] = _1041_v0;
          _nw207[(new BigNumber(24)).toNumber()] = _1042_v1;
          _nw207[(new BigNumber(25)).toNumber()] = _1041_v0;
          _nw207[(new BigNumber(26)).toNumber()] = _1042_v1;
          _nw207[(new BigNumber(27)).toNumber()] = _1041_v0;
          _nw207[(new BigNumber(28)).toNumber()] = _1041_v0;
          _1088_v33 = _nw207;
          let _1089_v34;
          _1089_v34 = _module.D1.create_DC3(_1088_v33);
          let _1090_v35;
          _1090_v35 = _dafny.Seq.of(_1089_v34);
          let _1091_v36;
          _1091_v36 = _module.D9.create_DC24(_1090_v35);
          let _1092_v37;
          let _nw208 = Array((new BigNumber(23)).toNumber());
          _nw208[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_1090_v35, _1090_v35), _module.__default.safeIndex((_this).f8, new BigNumber((_dafny.Seq.Concat(_1090_v35, _1090_v35)).length)), _1089_v34);
          _nw208[(_dafny.ONE).toNumber()] = _1090_v35;
          _nw208[(new BigNumber(2)).toNumber()] = _1090_v35;
          _nw208[(new BigNumber(3)).toNumber()] = _1090_v35;
          _nw208[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_1090_v35, _module.__default.safeIndex((_this).f8, new BigNumber((_1090_v35).length)), _1089_v34);
          _nw208[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1090_v35, _1090_v35);
          _nw208[(new BigNumber(6)).toNumber()] = _1090_v35;
          _nw208[(new BigNumber(7)).toNumber()] = _1090_v35;
          _nw208[(new BigNumber(8)).toNumber()] = _1090_v35;
          _nw208[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_1089_v34, _module.D1.create_DC3(_1088_v33), _1089_v34, _1089_v34), _module.__default.safeIndex((_this).f8, new BigNumber((_dafny.Seq.of(_1089_v34, _module.D1.create_DC3(_1088_v33), _1089_v34, _1089_v34)).length)), _module.D1.create_DC3(_1088_v33)), _1090_v35);
          _nw208[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1090_v35, _1090_v35);
          _nw208[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_1089_v34);
          _nw208[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(_1089_v34);
          _nw208[(new BigNumber(13)).toNumber()] = _1090_v35;
          _nw208[(new BigNumber(14)).toNumber()] = _1090_v35;
          _nw208[(new BigNumber(15)).toNumber()] = _1090_v35;
          _nw208[(new BigNumber(16)).toNumber()] = _1090_v35;
          _nw208[(new BigNumber(17)).toNumber()] = _1090_v35;
          _nw208[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_1090_v35, _1090_v35);
          _nw208[(new BigNumber(19)).toNumber()] = _1090_v35;
          _nw208[(new BigNumber(20)).toNumber()] = _1090_v35;
          _nw208[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_1090_v35, _1090_v35);
          _nw208[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat((_1091_v36).dtor_cf44, _1090_v35);
          _1092_v37 = _nw208;
          _1092_v37 = _1092_v37;
          _1058_v14 = !(false);
          let _index180 = _module.__default.safeIndex(new BigNumber(771), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index180] = (_this).f6;
          let _index181 = _module.__default.safeIndex(new BigNumber(771), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index181] = (_this).f9;
        }
        (globalState).f1 = (_dafny.ZERO).minus(p0);
        let _1093_v38;
        _1093_v38 = _module.D6.create_DC17((_this).f5, (_this).f5);
        _1093_v38 = _module.__default.fm31(globalState);
      }
      if ((_this).f7) {
        let _1094_v39;
        let _nw209 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
        _1094_v39 = _nw209;
        let _1095_v40;
        _1095_v40 = _module.D1.create_DC3(_1094_v39);
        let _pat_let_tv17 = _1094_v39;
        let _1096_v41;
        let _nw210 = Array((new BigNumber(21)).toNumber());
        _nw210[(_dafny.ZERO).toNumber()] = _1095_v40;
        _nw210[(_dafny.ONE).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(2)).toNumber()] = _module.D1.create_DC3(_1094_v39);
        _nw210[(new BigNumber(3)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(4)).toNumber()] = _module.D1.create_DC3(_1094_v39);
        _nw210[(new BigNumber(5)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(6)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(7)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(8)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(9)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(10)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(11)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(12)).toNumber()] = _module.D1.create_DC3(_1094_v39);
        _nw210[(new BigNumber(13)).toNumber()] = function (_pat_let30_0) {
          return function (_1097_dt__update__tmp_h0) {
            return function (_pat_let31_0) {
              return function (_1098_dt__update_hcf7_h0) {
                return _module.D1.create_DC3(_1098_dt__update_hcf7_h0);
              }(_pat_let31_0);
            }(_pat_let_tv17);
          }(_pat_let30_0);
        }(_module.D1.create_DC3(_1094_v39));
        _nw210[(new BigNumber(14)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(15)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(16)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(17)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(18)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(19)).toNumber()] = _1095_v40;
        _nw210[(new BigNumber(20)).toNumber()] = _1095_v40;
        _1096_v41 = _nw210;
        let _1099_v42;
        let _nw211 = new _module.C2();
        _nw211.__ctor((_this).f10, _1096_v41);
        _1099_v42 = _nw211;
        let _1100_v43;
        _1100_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber(612));
        r0 = _1100_v43;
        let _source12 = _1095_v40;
        if (_source12.is_DC4) {
          let _1101___mcc_h5 = (_source12).cf8;
          let _1102___mcc_h6 = (_source12).cf9;
          let _1103___mcc_h7 = (_source12).cf10;
          let _1104___mcc_h8 = (_source12).cf11;
          let _1105___mcc_h9 = (_source12).cf12;
          let _1106_cf12 = _1105___mcc_h9;
          let _1107_cf11 = _1104___mcc_h8;
          let _1108_cf10 = _1103___mcc_h7;
          let _1109_cf9 = _1102___mcc_h6;
          let _1110_cf8 = _1101___mcc_h5;
          _1109_cf9 = (_1109_cf9).plus(new BigNumber((_dafny.Seq.Concat(_1041_v0, _1042_v1)).length));
          let _1111_v44;
          _1111_v44 = _dafny.MultiSet.fromElements((_this).f11, (_this).f11);
          let _index182 = _module.__default.safeIndex(new BigNumber(769), new BigNumber(((_1099_v42).f14).length));
          ((_1099_v42).f14)[_index182] = (_1111_v44).IsDisjointFrom(_1111_v44);
          let _index183 = _module.__default.safeIndex(new BigNumber(769), new BigNumber(((_1099_v42).f14).length));
          ((_1099_v42).f14)[_index183] = (_this).f7;
          let _1112_v45;
          let _nw212 = new _module.C4();
          _nw212.__ctor((_this).f8, (_this).f6, (_this).f9, !(((_1099_v42).f14)[_module.__default.safeIndex(new BigNumber(769), new BigNumber(((_1099_v42).f14).length))]), _1110_cf8, (_this).f5);
          _1112_v45 = _nw212;
          let _1113_v46;
          _1113_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1112_v45,(_this).f6);
          let _1114_v47;
          _1114_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1108_cf10,(_this).f9);
          let _1115_v48;
          _1115_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1113_v46,_1114_v47);
          let _1116_v49;
          _1116_v49 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f7);
          let _rhs165 = _1115_v48;
          let _rhs166 = (_this).fm1(_module.__default.fm3(globalState), _1116_v49, !((_this).f7), globalState);
          let _lhs125 = globalState;
          _1115_v48 = _rhs165;
          _lhs125.f1 = _rhs166;
          let _1117_v50;
          _1117_v50 = _dafny.Seq.of((_this).f8);
          (globalState).f1 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(_1109_cf9, _1109_cf9), new BigNumber((_1117_v50).length));
        } else if (_source12.is_DC5) {
          let _1118___mcc_h10 = (_source12).cf13;
          let _1119_cf13 = _1118___mcc_h10;
          let _1120_v51;
          _1120_v51 = _dafny.Seq.of((_this).f8);
          let _1121_v52;
          _1121_v52 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1041_v0).length)), _1100_v43);
          let _1122_v53;
          _1122_v53 = _module.D6.create_DC17(_dafny.Seq.update((_this).f5, _module.__default.safeIndex(new BigNumber(555), new BigNumber(((_this).f5).length)), _this.f4), (_this).f5);
          let _1123_v54;
          _1123_v54 = _dafny.Seq.of((_1120_v51)[_module.__default.safeIndex(p0, new BigNumber((_1120_v51).length))], new BigNumber(((_1121_v52)[_module.__default.safeIndex(new BigNumber(((_1122_v53).dtor_cf31).length), new BigNumber((_1121_v52).length))]).length));
          let _1124_v55;
          let _nw213 = Array((new BigNumber(5)).toNumber());
          _nw213[(_dafny.ZERO).toNumber()] = _dafny.Seq.contains(_1120_v51, new BigNumber(((_1100_v43).update(p0, _1119_cf13)).length));
          _nw213[(_dafny.ONE).toNumber()] = (_this).f7;
          _nw213[(new BigNumber(2)).toNumber()] = (_this).f9;
          _nw213[(new BigNumber(3)).toNumber()] = (p0).isLessThan((_this).f8);
          _nw213[(new BigNumber(4)).toNumber()] = (((_this).f9) ? ((_this).f7) : ((_this).f9));
          _1124_v55 = _nw213;
          _1124_v55 = (_1099_v42).f14;
          let _1125_v56;
          _1125_v56 = false;
          _1125_v56 = (_this).f7;
          let _1126_v57;
          let _nw214 = Array((new BigNumber(10)).toNumber()).fill([]);
          _1126_v57 = _nw214;
          let _1127_v58;
          _1127_v58 = _module.D2.create_DC8((_this).f8, new BigNumber(-507));
          let _1128_v59;
          _1128_v59 = _module.D0.create_DC0();
          let _1129_v60;
          _1129_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1128_v59,_1125_v56);
          let _1130_v61;
          _1130_v61 = _module.D0.create_DC1((_this).f9, p0, (_this).f7, new BigNumber((_1129_v60).length));
          let _1131_v62;
          let _nw215 = Array((new BigNumber(26)).toNumber());
          _nw215[(_dafny.ZERO).toNumber()] = _1127_v58;
          _nw215[(_dafny.ONE).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(2)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(3)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(4)).toNumber()] = _module.D2.create_DC8((_this).f8, (_1130_v61).dtor_cf3);
          _nw215[(new BigNumber(5)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(6)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(7)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(8)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(9)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(10)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(11)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(12)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(13)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(14)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(15)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(16)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(17)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(18)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(19)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(20)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(21)).toNumber()] = _module.D2.create_DC8(p0, (_this).f8);
          _nw215[(new BigNumber(22)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(23)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(24)).toNumber()] = _1127_v58;
          _nw215[(new BigNumber(25)).toNumber()] = _1127_v58;
          _1131_v62 = _nw215;
          let _1132_v63;
          _1132_v63 = _module.D10.create_DC26(_1131_v62);
          let _1133_v64;
          let _nw216 = Array((new BigNumber(8)).toNumber());
          _nw216[(_dafny.ZERO).toNumber()] = _1131_v62;
          _nw216[(_dafny.ONE).toNumber()] = (_1132_v63).dtor_cf48;
          _nw216[(new BigNumber(2)).toNumber()] = _1131_v62;
          _nw216[(new BigNumber(3)).toNumber()] = _1131_v62;
          _nw216[(new BigNumber(4)).toNumber()] = _1131_v62;
          _nw216[(new BigNumber(5)).toNumber()] = _1131_v62;
          _nw216[(new BigNumber(6)).toNumber()] = _1131_v62;
          _nw216[(new BigNumber(7)).toNumber()] = _1131_v62;
          _1133_v64 = _nw216;
          let _index184 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1126_v57).length));
          (_1126_v57)[_index184] = _1133_v64;
          let _index185 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1126_v57).length));
          (_1126_v57)[_index185] = _1133_v64;
          let _index186 = _module.__default.safeIndex(new BigNumber(54), new BigNumber(((_this).f11).length));
          ((_this).f11)[_index186] = new BigNumber(((_this).f5).length);
          let _1134_v65;
          _1134_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,!(_1125_v56));
          let _index187 = _module.__default.safeIndex(new BigNumber(54), new BigNumber(((_this).f11).length));
          let _rhs167 = _this.f4;
          let _rhs168 = new BigNumber(223);
          let _rhs169 = _module.__default.safeDivisionInt((_this).f8, _module.__default.safeModuloInt((_this).f8, (_this).f8));
          let _rhs170 = _this.f4;
          let _rhs171 = _module.__default.fm7(new _dafny.CodePoint('v'.codePointAt(0)), new BigNumber((_1134_v65).length), (_dafny.Set.fromElements(new BigNumber((_module.__default.fm24(_1119_cf13, (_this).f9, (_this).f7, p0, globalState)).length))).IsDisjointFrom(_module.__default.fm30((_this).f6, globalState)), globalState);
          let _lhs126 = _this;
          let _lhs127 = (_this).f11;
          let _lhs128 = _module.__default.safeIndex(new BigNumber(54), new BigNumber(((_this).f11).length));
          let _lhs129 = _this;
          _lhs126.f4 = _rhs167;
          _lhs127[_lhs128] = _rhs168;
          _1119_cf13 = _rhs169;
          _lhs129.f4 = _rhs170;
          _1125_v56 = _rhs171;
        } else {
          let _1135___mcc_h11 = (_source12).cf7;
          let _1136_cf7 = _1135___mcc_h11;
          (globalState).f1 = p0;
          (globalState).f1 = (_dafny.ZERO).minus((_this).f8);
          let _1137_v66;
          let _init22 = ((_1138_p0) => function (_1139_i4) {
            return _module.D2.create_DC8(_1138_p0, new BigNumber(730));
          })(p0);
          let _nw217 = Array((new BigNumber(17)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw217.length); _i0_22++) {
            _nw217[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _1137_v66 = _nw217;
          let _1140_v67;
          _1140_v67 = _module.D10.create_DC26(_1137_v66);
          let _pat_let_tv18 = _1137_v66;
          _1140_v67 = function (_pat_let32_0) {
            return function (_1141_dt__update__tmp_h1) {
              return function (_pat_let33_0) {
                return function (_1142_dt__update_hcf48_h0) {
                  return _module.D10.create_DC26(_1142_dt__update_hcf48_h0);
                }(_pat_let33_0);
              }(_pat_let_tv18);
            }(_pat_let32_0);
          }(_1140_v67);
          let _1143_v68;
          _1143_v68 = true;
          _1143_v68 = true;
        }
        (globalState).f1 = p0;
        let _1144_v69;
        _1144_v69 = true;
        _1144_v69 = false;
      } else {
        let _1145_v70;
        _1145_v70 = true;
        let _1146_v71;
        _1146_v71 = _dafny.Set.fromElements(new BigNumber(((_this).f5).length), (_this).f8);
        _1145_v70 = ((((_this).f6) ? (_1146_v71) : (_1146_v71))).IsSubsetOf(_1146_v71);
        _1145_v70 = !(false);
        let _1147_v72;
        let _nw218 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _1147_v72 = _nw218;
        _1147_v72 = _1147_v72;
        let _index188 = _module.__default.safeIndex(new BigNumber(136), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index188] = (_this).f7;
        let _index189 = _module.__default.safeIndex(new BigNumber(136), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index189] = !(_1145_v70);
        let _1148_v73;
        _1148_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f8);
        r0 = ((_1148_v73).Merge(_1148_v73)).Merge(_1148_v73);
      }
      (globalState).f1 = (_dafny.ZERO).minus((_this).f8);
      if (function (_source13) {
        if (_source13.is_DC7) {
          let _1149___mcc_h12 = (_source13).cf15;
          let _1150_cf15 = _1149___mcc_h12;
          return (_this).f6;
        } else if (_source13.is_DC8) {
          let _1151___mcc_h13 = (_source13).cf16;
          let _1152___mcc_h14 = (_source13).cf17;
          let _1153_cf17 = _1152___mcc_h14;
          let _1154_cf16 = _1151___mcc_h13;
          return (_this).f6;
        } else if (_source13.is_DC6) {
          let _1155___mcc_h15 = (_source13).cf14;
          let _1156_cf14 = _1155___mcc_h15;
          return (_this).f9;
        } else {
          let _1157___mcc_h16 = (_source13).cf18;
          let _1158_cf18 = _1157___mcc_h16;
          return _dafny.areEqual(_dafny.Seq.of(_dafny.Set.fromElements((_this).f7)), _dafny.Seq.of(_dafny.Set.fromElements((_this).f6, (_this).f9, (_this).f9, (_this).f7, (_this).f6), _dafny.Set.fromElements((_this).f6, (_this).f7, (_this).f7, (_this).f9)));
        }
      }(function (_pat_let34_0) {
        return function (_1159_dt__update__tmp_h2) {
          return function (_pat_let35_0) {
            return function (_1160_dt__update_hcf15_h0) {
              return _module.D2.create_DC7(_1160_dt__update_hcf15_h0);
            }(_pat_let35_0);
          }((_this).f9);
        }(_pat_let34_0);
      }(_module.D2.create_DC7((_this).f6)))) {
        (globalState).f1 = p0;
        let _1161_v74;
        _1161_v74 = _dafny.Set.fromElements((_this).f9, (_this).f6);
        if (((_this).f8).isEqualTo(new BigNumber((_1161_v74).length))) {
          let _1162_v75;
          _1162_v75 = false;
          let _1163_v76;
          _1163_v76 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f9);
          let _rhs172 = _module.__default.safeModuloInt(new BigNumber(-556), p0);
          let _rhs173 = (_1162_v75) || ((_this).f7);
          let _rhs174 = _1162_v75;
          let _rhs175 = _1163_v76;
          let _lhs130 = globalState;
          _lhs130.f1 = _rhs172;
          _1162_v75 = _rhs173;
          _1162_v75 = _rhs174;
          _1163_v76 = _rhs175;
          (globalState).f1 = (_this).f8;
          let _1164_v77;
          _1164_v77 = _dafny.Seq.UnicodeFromString("rc");
          _1164_v77 = _dafny.Seq.UnicodeFromString("jdhcdkvjy");
          let _1165_v78;
          let _nw219 = new _module.C3();
          _nw219.__ctor(false, true, new BigNumber(317), true, !(((_dafny.ZERO).minus(p0)).isLessThanOrEqualTo((_this).f8)), (_this).f9, _this.f4, _module.__default.fm14((_dafny.ZERO).minus((_this).f8), new BigNumber((_1164_v77).length), (_this).f7, (_this).f6, globalState));
          _1165_v78 = _nw219;
          let _1166_v79;
          _1166_v79 = _dafny.Seq.of((_this).f8, p0);
          let _1167_v80;
          _1167_v80 = _dafny.MultiSet.fromElements(_1166_v79, _1166_v79);
          let _1168_v81;
          _1168_v81 = _module.D11.create_DC29(_1167_v80);
          let _1169_v82;
          let _nw220 = new _module.C4();
          _nw220.__ctor(new BigNumber(738), ((_1168_v81).dtor_cf55).contains(_1166_v79), (_this).f6, (_this).f9, _this.f4, _1164_v77);
          _1169_v82 = _nw220;
        } else {
          let _1170_v83;
          _1170_v83 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          _1170_v83 = (_1170_v83).update(p0, p0);
          let _1171_v84;
          let _nw221 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Set.Empty);
          _1171_v84 = _nw221;
          let _index190 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_1171_v84).length));
          (_1171_v84)[_index190] = _1161_v74;
          let _1172_v85;
          _1172_v85 = true;
          let _index191 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_1171_v84).length));
          let _rhs176 = _1161_v74;
          let _rhs177 = ((p0).multipliedBy((_this).f8)).plus((_this).f8);
          let _rhs178 = !((((_this).f6) ? (_1172_v85) : (!((_this).f7))));
          let _rhs179 = p0;
          let _lhs131 = _1171_v84;
          let _lhs132 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_1171_v84).length));
          let _lhs133 = globalState;
          let _lhs134 = globalState;
          _lhs131[_lhs132] = _rhs176;
          _lhs133.f1 = _rhs177;
          _1172_v85 = _rhs178;
          _lhs134.f1 = _rhs179;
          let _1173_v86;
          _1173_v86 = _dafny.Seq.of(_1041_v0);
          let _1174_v87;
          let _nw222 = Array((new BigNumber(11)).toNumber());
          _nw222[(_dafny.ZERO).toNumber()] = _1042_v1;
          _nw222[(_dafny.ONE).toNumber()] = _1042_v1;
          _nw222[(new BigNumber(2)).toNumber()] = (_1173_v86)[_module.__default.safeIndex(p0, new BigNumber((_1173_v86).length))];
          _nw222[(new BigNumber(3)).toNumber()] = _dafny.Seq.of((_this).f9);
          _nw222[(new BigNumber(4)).toNumber()] = _1041_v0;
          _nw222[(new BigNumber(5)).toNumber()] = _1042_v1;
          _nw222[(new BigNumber(6)).toNumber()] = _1041_v0;
          _nw222[(new BigNumber(7)).toNumber()] = _1041_v0;
          _nw222[(new BigNumber(8)).toNumber()] = _1042_v1;
          _nw222[(new BigNumber(9)).toNumber()] = _1042_v1;
          _nw222[(new BigNumber(10)).toNumber()] = _1042_v1;
          _1174_v87 = _nw222;
          let _1175_v88;
          _1175_v88 = _module.D1.create_DC3(_1174_v87);
          let _1176_v89;
          let _nw223 = Array((new BigNumber(19)).toNumber());
          _nw223[(_dafny.ZERO).toNumber()] = _module.D1.create_DC3(_1174_v87);
          _nw223[(_dafny.ONE).toNumber()] = _1175_v88;
          _nw223[(new BigNumber(2)).toNumber()] = _1175_v88;
          _nw223[(new BigNumber(3)).toNumber()] = _1175_v88;
          _nw223[(new BigNumber(4)).toNumber()] = _1175_v88;
          _nw223[(new BigNumber(5)).toNumber()] = _module.D1.create_DC3(_1174_v87);
          _nw223[(new BigNumber(6)).toNumber()] = _1175_v88;
          _nw223[(new BigNumber(7)).toNumber()] = _1175_v88;
          _nw223[(new BigNumber(8)).toNumber()] = _module.D1.create_DC3(_1174_v87);
          _nw223[(new BigNumber(9)).toNumber()] = _1175_v88;
          _nw223[(new BigNumber(10)).toNumber()] = _1175_v88;
          _nw223[(new BigNumber(11)).toNumber()] = _1175_v88;
          _nw223[(new BigNumber(12)).toNumber()] = _1175_v88;
          _nw223[(new BigNumber(13)).toNumber()] = _module.D1.create_DC3(_1174_v87);
          _nw223[(new BigNumber(14)).toNumber()] = _1175_v88;
          _nw223[(new BigNumber(15)).toNumber()] = _1175_v88;
          _nw223[(new BigNumber(16)).toNumber()] = _1175_v88;
          _nw223[(new BigNumber(17)).toNumber()] = _module.D1.create_DC3(_1174_v87);
          _nw223[(new BigNumber(18)).toNumber()] = _1175_v88;
          _1176_v89 = _nw223;
          let _1177_v90;
          let _nw224 = new _module.C2();
          _nw224.__ctor((_this).f10, _1176_v89);
          _1177_v90 = _nw224;
          let _1178_v91;
          let _nw225 = new _module.C0();
          _nw225.__ctor(((_1172_v85) ? ((_this).f7) : ((_this).f7)), (_this).f6, _this.f4, _dafny.Seq.Concat((_this).f5, _dafny.Seq.UnicodeFromString("mum")));
          _1178_v91 = _nw225;
          let _1179_v92;
          _1179_v92 = _dafny.Seq.of(p0, (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(-577))).length)), new BigNumber(258));
          let _1180_v93;
          _1180_v93 = _dafny.Seq.of(_1179_v92, _dafny.Seq.of(p0, new BigNumber(((_this).f5).length)), _1179_v92, _1179_v92);
          _1180_v93 = _1180_v93;
        }
        let _1181_v94;
        _1181_v94 = false;
        _1181_v94 = (p0).isEqualTo(p0);
        let _index192 = _module.__default.safeIndex(new BigNumber(782), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index192] = true;
        let _index193 = _module.__default.safeIndex(new BigNumber(782), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index193] = _1181_v94;
        let _1182_v95;
        let _init23 = function (_1183_i5) {
          return (_this).f6;
        };
        let _nw226 = Array((new BigNumber(16)).toNumber());
        for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw226.length); _i0_23++) {
          _nw226[_i0_23] = _init23(new BigNumber(_i0_23));
        }
        _1182_v95 = _nw226;
        let _1184_v96;
        let _nw227 = Array((new BigNumber(20)).toNumber()).fill(_module.D1.Default());
        _1184_v96 = _nw227;
        let _1185_v97;
        let _nw228 = new _module.C2();
        _nw228.__ctor(_1182_v95, _1184_v96);
        _1185_v97 = _nw228;
      } else {
        let _1186_v98;
        _1186_v98 = _dafny.Seq.of(_1042_v1);
        let _1187_v99;
        let _nw229 = Array((new BigNumber(23)).toNumber());
        _nw229[(_dafny.ZERO).toNumber()] = _1041_v0;
        _nw229[(_dafny.ONE).toNumber()] = _1042_v1;
        _nw229[(new BigNumber(2)).toNumber()] = _1042_v1;
        _nw229[(new BigNumber(3)).toNumber()] = _1042_v1;
        _nw229[(new BigNumber(4)).toNumber()] = _1042_v1;
        _nw229[(new BigNumber(5)).toNumber()] = _1042_v1;
        _nw229[(new BigNumber(6)).toNumber()] = _1041_v0;
        _nw229[(new BigNumber(7)).toNumber()] = _1042_v1;
        _nw229[(new BigNumber(8)).toNumber()] = _dafny.Seq.of((_this).f7, (_this).f9);
        _nw229[(new BigNumber(9)).toNumber()] = _1041_v0;
        _nw229[(new BigNumber(10)).toNumber()] = _1042_v1;
        _nw229[(new BigNumber(11)).toNumber()] = _1041_v0;
        _nw229[(new BigNumber(12)).toNumber()] = _1041_v0;
        _nw229[(new BigNumber(13)).toNumber()] = _1042_v1;
        _nw229[(new BigNumber(14)).toNumber()] = _1042_v1;
        _nw229[(new BigNumber(15)).toNumber()] = _1041_v0;
        _nw229[(new BigNumber(16)).toNumber()] = _1042_v1;
        _nw229[(new BigNumber(17)).toNumber()] = _1041_v0;
        _nw229[(new BigNumber(18)).toNumber()] = _dafny.Seq.of((_this).f9);
        _nw229[(new BigNumber(19)).toNumber()] = _dafny.Seq.of((_this).f6, (_this).f7);
        _nw229[(new BigNumber(20)).toNumber()] = _1041_v0;
        _nw229[(new BigNumber(21)).toNumber()] = _1041_v0;
        _nw229[(new BigNumber(22)).toNumber()] = (_1186_v98)[_module.__default.safeIndex((_this).f8, new BigNumber((_1186_v98).length))];
        _1187_v99 = _nw229;
        let _1188_v100;
        _1188_v100 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(((false) ? (_1042_v1) : (_1041_v0)))).cardinality()),_module.D1.create_DC3(_1187_v99));
        let _1189_v101;
        _1189_v101 = _module.D1.create_DC3(_1187_v99);
        _1188_v100 = (_1188_v100).update((((_this).f7) ? (p0) : (p0)), _1189_v101);
        let _1190_v102;
        _1190_v102 = false;
        _1190_v102 = (_this).f6;
        let _1191_v103;
        _1191_v103 = _dafny.Set.fromElements((_this).f9, false, _1190_v102);
        let _1192_v104;
        _1192_v104 = _module.D3.create_DC11(!((_this).f6), (_this).f8, _1191_v103);
        let _1193_v105;
        _1193_v105 = _dafny.Seq.of(new BigNumber((_module.__default.fm28(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(66)), function (_1194_i6) {
          return _this.f4;
        })).length), (_this).f6, _1042_v1, _1192_v104, globalState)).length));
        let _1195_v106;
        _1195_v106 = _dafny.Set.fromElements((_this).f8, (_1193_v105)[_module.__default.safeIndex(p0, new BigNumber((_1193_v105).length))]);
        let _1196_v107;
        _1196_v107 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f6),_dafny.Seq.update(_1193_v105, _module.__default.safeIndex((_1193_v105)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1195_v106).length)), new BigNumber((_1193_v105).length))], new BigNumber((_1193_v105).length)), (_this).f8));
        _1196_v107 = (_1196_v107).Merge(_1196_v107);
        _1190_v102 = (_1191_v103).contains((_this).f6);
        let _1197_v108;
        _1197_v108 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f6);
        let _1198_v109;
        _1198_v109 = _module.D1.create_DC5(p0);
        let _1199_v110;
        _1199_v110 = _dafny.MultiSet.fromElements(p0, p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).fm1((_this).f7, _1197_v108, _module.__default.fm27((_this).f8, _1190_v102, (_this).f6, globalState), globalState)),(_this).f5)).length), (new BigNumber((_dafny.Seq.of(_1198_v109, _module.D1.create_DC5(new BigNumber(((_this).f5).length)), _1198_v109, _module.D1.create_DC5(new BigNumber(528)), _1198_v109)).length)).multipliedBy(new BigNumber(896)));
        (globalState).f1 = (((_1199_v110).contains(p0)) ? ((_1199_v110).get(p0)) : (p0));
      }
      let _1200_v111;
      _1200_v111 = _module.D1.create_DC5(new BigNumber(-740));
      let _1201_v112;
      _1201_v112 = _dafny.Seq.of(_1200_v111, _1200_v111);
      let _1202_v113;
      _1202_v113 = _1201_v112;
      if (function (_source14) {
        let _1203___mcc_h17 = _source14;
        let _1204_cf24 = _1203___mcc_h17;
        return (_this).f7;
      }(_1202_v113)) {
        (globalState).f1 = p0;
        let _1205_v114;
        _1205_v114 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,(_this).f6);
        let _1206_v115;
        _1206_v115 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f7);
        (_this).m4(p0, (_dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f9)).update((_this).f8, (((_1205_v114).contains(_this.f4)) ? ((_1205_v114).get(_this.f4)) : ((((_1206_v115).contains((_this).f7)) ? ((_1206_v115).get((_this).f7)) : ((_this).f7))))), globalState);
        (globalState).f1 = new BigNumber((_1041_v0).length);
        let _1207_v116;
        let _nw230 = new _module.C0();
        _nw230.__ctor((_1043_v2).IsDisjointFrom(_1043_v2), (_this).f9, new _dafny.CodePoint('q'.codePointAt(0)), _dafny.Seq.Concat((_this).f5, (_this).f5));
        _1207_v116 = _nw230;
        let _1208_v118;
        _1208_v118 = _dafny.Seq.of(_module.__default.fm8(p0, (_this).f9, (_this).f6, globalState));
        let _1209_v119;
        _1209_v119 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,p0);
        let _1210_v120;
        _1210_v120 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1208_v118).length),_1209_v119);
        let _1211_v121;
        _1211_v121 = _dafny.Seq.of((((_1210_v120).contains((_this).f8)) ? ((_1210_v120).get((_this).f8)) : (_1209_v119)));
        let _1212_v122;
        _1212_v122 = _dafny.Seq.of((_this).f8, new BigNumber((function () {
          let _coll21 = new _dafny.Map();
          for (const _compr_21 of (_1211_v121).Elements) {
            let _1213_v117 = _compr_21;
            if (_dafny.Seq.contains(_1211_v121, _1213_v117)) {
              _coll21.push([_1213_v117,(_this).f6]);
            }
          }
          return _coll21;
        }()).length));
        (globalState).f1 = ((_1212_v122)[_module.__default.safeIndex((_this).f8, new BigNumber((_1212_v122).length))]).minus((p0).minus(p0));
      } else {
        let _1214_v123;
        _1214_v123 = _dafny.MultiSet.fromElements((_this).f5, _dafny.Seq.UnicodeFromString("lbxa"));
        let _1215_v124;
        _1215_v124 = _dafny.Set.fromElements((_this).f7, _module.__default.fm27(p0, (_this).f9, (_this).f6, globalState));
        let _1216_v125;
        _1216_v125 = _module.D3.create_DC11((_this).f9, new BigNumber((_1214_v123).cardinality()), _1215_v124);
        let _1217_v126;
        _1217_v126 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_module.__default.fm28(p0, false, _dafny.Seq.of((_this).f7), _1216_v125, globalState));
        _1217_v126 = (((_this).f9) ? (_1217_v126) : (_1217_v126));
        let _1218_v127;
        let _out21;
        _out21 = (_this).m3(p0, (_this).f8, p0, globalState);
        _1218_v127 = _out21;
        let _1219_v128;
        _1219_v128 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(424)), function (_1220_i7) {
          return new BigNumber(462);
        })).length));
        let _1221_v129;
        _1221_v129 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
        (globalState).f1 = _module.__default.safeDivisionInt((((_1219_v128).contains(p0)) ? ((_1219_v128).get(p0)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-830)), function (_1222_i8) {
          return _dafny.Seq.UnicodeFromString("tfpkek");
        })).length))), (_this).fm1(_module.__default.fm3(globalState), _1221_v129, (_this).f7, globalState));
        let _1223_v130;
        let _nw231 = new _module.C0();
        _nw231.__ctor(_1218_v127, (_this).f9, _this.f4, (_this).f5);
        _1223_v130 = _nw231;
        let _index194 = _module.__default.safeIndex(new BigNumber(21), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index194] = (_this).f9;
        let _index195 = _module.__default.safeIndex(new BigNumber(739), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index195] = (_1043_v2).IsProperSubsetOf(_1043_v2);
        let _index196 = _module.__default.safeIndex(new BigNumber(21), new BigNumber(((_this).f10).length));
        let _index197 = _module.__default.safeIndex(new BigNumber(739), new BigNumber(((_this).f10).length));
        let _rhs180 = !((!((_this).f9) || ((_this).f7)) || ((_this).f6));
        let _rhs181 = ((_this).f8).isLessThan((_this).f8);
        let _rhs182 = (_this).f9;
        let _rhs183 = (_this).f7;
        let _rhs184 = _dafny.Seq.Concat(_1041_v0, _dafny.Seq.Concat(_1041_v0, _1042_v1));
        let _lhs135 = (_this).f10;
        let _lhs136 = _module.__default.safeIndex(new BigNumber(21), new BigNumber(((_this).f10).length));
        let _lhs137 = (_this).f10;
        let _lhs138 = _module.__default.safeIndex(new BigNumber(739), new BigNumber(((_this).f10).length));
        _lhs135[_lhs136] = _rhs180;
        _lhs137[_lhs138] = _rhs181;
        _1218_v127 = _rhs182;
        _1218_v127 = _rhs183;
        _1042_v1 = _rhs184;
      }
      if ((_this).f9) {
        let _1224_v131;
        _1224_v131 = _dafny.Set.fromElements(_dafny.Seq.IsPrefixOf(_1041_v0, _dafny.Seq.of((_this).f9, (_this).f7)), !((_this).f9) || ((_this).f9), true);
        _1224_v131 = _1224_v131;
        let _1225_v132;
        let _nw232 = Array((new BigNumber(28)).toNumber()).fill([]);
        _1225_v132 = _nw232;
        let _index198 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_1225_v132).length));
        (_1225_v132)[_index198] = (_this).f10;
        let _index199 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_1225_v132).length));
        (_1225_v132)[_index199] = (_this).f10;
        let _1226_v133;
        _1226_v133 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
        let _1227_v134;
        _1227_v134 = _module.D3.create_DC11((_this).f7, new BigNumber(((_this).f5).length), _1224_v131);
        let _pat_let_tv19 = _1224_v131;
        let _1228_v135;
        let _nw233 = new _module.C3();
        _nw233.__ctor(!_dafny.Seq.contains(_dafny.Seq.of(p0), new BigNumber(428)), !(_1224_v131).contains((_this).f6), (_this).fm1((_this).f6, _1226_v133, (_this).f6, globalState), (false) && (_module.__default.fm7(_module.__default.fm13(globalState), _module.__default.fm8((_this).f8, (_this).f6, (_1042_v1)[_module.__default.safeIndex((_this).f8, new BigNumber((_1042_v1).length))], globalState), true, globalState)), (_this).f6, (_this).f9, _this.f4, _module.__default.fm28(p0, (_this).f6, _1041_v0, function (_pat_let36_0) {
          return function (_1229_dt__update__tmp_h3) {
            return function (_pat_let37_0) {
              return function (_1230_dt__update_hcf22_h0) {
                return function (_pat_let38_0) {
                  return function (_1231_dt__update_hcf20_h0) {
                    return _module.D3.create_DC11(_1231_dt__update_hcf20_h0, (_1229_dt__update__tmp_h3).dtor_cf21, _1230_dt__update_hcf22_h0);
                  }(_pat_let38_0);
                }((_this).f7);
              }(_pat_let37_0);
            }(_pat_let_tv19);
          }(_pat_let36_0);
        }(_1227_v134), globalState));
        _1228_v135 = _nw233;
        _1228_v135 = _1228_v135;
        let _1232_v136;
        _1232_v136 = _dafny.Seq.of(p0, (_dafny.ZERO).minus(_module.__default.fm8(p0, (_1228_v135).f13, !((_this).f9), globalState)), (_this).f8);
        _1232_v136 = _1232_v136;
        let _1233_v137;
        _1233_v137 = _module.D1.create_DC4(_this.f4, (_dafny.ZERO).minus(p0), _this.f4, p0, _1228_v135.f12);
        let _1234_v138;
        _1234_v138 = _dafny.MultiSet.fromElements(_1233_v137);
        (_1228_v135).f12 = (!((_1224_v131).IsDisjointFrom(_1224_v131))) === ((_1234_v138).IsProperSubsetOf(_1234_v138));
      } else {
        let _1235_v139;
        _1235_v139 = true;
        let _1236_v140;
        let _nw234 = Array((new BigNumber(15)).toNumber()).fill([]);
        _1236_v140 = _nw234;
        let _index200 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_1236_v140).length));
        (_1236_v140)[_index200] = (_this).f10;
        let _index201 = _module.__default.safeIndex(new BigNumber(461), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index201] = _dafny.Seq.contains((_this).f5, _this.f4);
        let _index202 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_1236_v140).length));
        let _index203 = _module.__default.safeIndex(new BigNumber(461), new BigNumber(((_this).f10).length));
        let _rhs185 = !(_module.__default.fm3(globalState)) || ((_this).f6);
        let _rhs186 = _this.f4;
        let _rhs187 = (_this).f10;
        let _rhs188 = (_this).f8;
        let _rhs189 = (_this).f9;
        let _lhs139 = _this;
        let _lhs140 = _1236_v140;
        let _lhs141 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_1236_v140).length));
        let _lhs142 = globalState;
        let _lhs143 = (_this).f10;
        let _lhs144 = _module.__default.safeIndex(new BigNumber(461), new BigNumber(((_this).f10).length));
        _1235_v139 = _rhs185;
        _lhs139.f4 = _rhs186;
        _lhs140[_lhs141] = _rhs187;
        _lhs142.f1 = _rhs188;
        _lhs143[_lhs144] = _rhs189;
        (globalState).f1 = (_this).f8;
        let _1237_v141;
        let _init24 = function (_1238_i9) {
          return (_1238_i9).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f9)).length));
        };
        let _nw235 = Array((new BigNumber(29)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw235.length); _i0_24++) {
          _nw235[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _1237_v141 = _nw235;
        let _index204 = _module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length));
        ((_this).f11)[_index204] = (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f6)).length));
        let _index205 = _module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length));
        let _rhs190 = _1237_v141;
        let _rhs191 = _1235_v139;
        let _rhs192 = (_this).f6;
        let _rhs193 = (_this).f6;
        let _rhs194 = new BigNumber((_1041_v0).length);
        let _lhs145 = (_this).f11;
        let _lhs146 = _module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length));
        _1237_v141 = _rhs190;
        _1235_v139 = _rhs191;
        _1235_v139 = _rhs192;
        _1235_v139 = _rhs193;
        _lhs145[_lhs146] = _rhs194;
        let _1239_v142;
        _1239_v142 = _dafny.Seq.of((_this).f8);
        if (_dafny.Seq.IsPrefixOf(_1239_v142, (((_this).f6) ? (_dafny.Seq.of(((_this).f11)[_module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length))], ((_this).f11)[_module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length))], p0)) : (_1239_v142)))) {
          let _1240_v143;
          _1240_v143 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f9);
          (_this).m4((_dafny.ZERO).minus(new BigNumber((_module.__default.fm18(p0, !(((_this).f10)[_module.__default.safeIndex(new BigNumber(461), new BigNumber(((_this).f10).length))]), globalState)).length)), _1240_v143, globalState);
          let _1241_v144;
          let _nw236 = Array((new BigNumber(17)).toNumber()).fill([]);
          _1241_v144 = _nw236;
          let _1242_v145;
          let _nw237 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1242_v145 = _nw237;
          let _index206 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1241_v144).length));
          (_1241_v144)[_index206] = _1242_v145;
          let _1243_v146;
          let _nw238 = Array((new BigNumber(22)).toNumber()).fill(_module.D1.Default());
          _1243_v146 = _nw238;
          let _1244_v147;
          let _nw239 = new _module.C2();
          _nw239.__ctor((_1236_v140)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_1236_v140).length))], _1243_v146);
          _1244_v147 = _nw239;
          let _1245_v148;
          _1245_v148 = _dafny.MultiSet.fromElements(((_this).f11)[_module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length))]);
          let _index207 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1241_v144).length));
          let _index208 = _module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length));
          let _rhs195 = _1242_v145;
          let _rhs196 = (_dafny.ZERO).minus((_this).f8);
          let _rhs197 = p0;
          let _rhs198 = (((p0).isLessThanOrEqualTo(new BigNumber((_1245_v148).cardinality()))) ? (_1244_v147) : (_1244_v147));
          let _lhs147 = _1241_v144;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1241_v144).length));
          let _lhs149 = globalState;
          let _lhs150 = (_this).f11;
          let _lhs151 = _module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length));
          _lhs147[_lhs148] = _rhs195;
          _lhs149.f1 = _rhs196;
          _lhs150[_lhs151] = _rhs197;
          _1244_v147 = _rhs198;
          let _index209 = _module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length));
          ((_this).f11)[_index209] = p0;
          _1042_v1 = _1042_v1;
          let _1246_v149;
          _1246_v149 = _module.D1.create_DC4(_module.__default.fm13(globalState), p0, _this.f4, (_this).f8, (_this).f9);
          let _1247_v150;
          let _nw240 = Array((new BigNumber(21)).toNumber());
          _nw240[(_dafny.ZERO).toNumber()] = _1041_v0;
          _nw240[(_dafny.ONE).toNumber()] = _1041_v0;
          _nw240[(new BigNumber(2)).toNumber()] = _1041_v0;
          _nw240[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_1042_v1, _module.__default.safeIndex(p0, new BigNumber((_1042_v1).length)), (_this).f7);
          _nw240[(new BigNumber(4)).toNumber()] = _1041_v0;
          _nw240[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(((_this).f10)[_module.__default.safeIndex(new BigNumber(461), new BigNumber(((_this).f10).length))], (_this).f6, (_this).f7, _1235_v139);
          _nw240[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(_1235_v139, (_this).f7);
          _nw240[(new BigNumber(7)).toNumber()] = _1042_v1;
          _nw240[(new BigNumber(8)).toNumber()] = _1042_v1;
          _nw240[(new BigNumber(9)).toNumber()] = _1041_v0;
          _nw240[(new BigNumber(10)).toNumber()] = _dafny.Seq.of((_this).f7);
          _nw240[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(((_this).f10)[_module.__default.safeIndex(new BigNumber(461), new BigNumber(((_this).f10).length))]);
          _nw240[(new BigNumber(12)).toNumber()] = _1042_v1;
          _nw240[(new BigNumber(13)).toNumber()] = _1042_v1;
          _nw240[(new BigNumber(14)).toNumber()] = _dafny.Seq.of((_this).f9);
          _nw240[(new BigNumber(15)).toNumber()] = _1041_v0;
          _nw240[(new BigNumber(16)).toNumber()] = _1042_v1;
          _nw240[(new BigNumber(17)).toNumber()] = _1041_v0;
          _nw240[(new BigNumber(18)).toNumber()] = _1042_v1;
          _nw240[(new BigNumber(19)).toNumber()] = _1041_v0;
          _nw240[(new BigNumber(20)).toNumber()] = _1041_v0;
          _1247_v150 = _nw240;
          let _1248_v151;
          _1248_v151 = _dafny.Map.Empty.slice().updateUnsafe(_1235_v139,_module.D1.create_DC3(_1247_v150));
          let _1249_v152;
          let _1250_v153;
          let _1251_v154;
          let _out22;
          let _out23;
          let _out24;
          let _outcollector6 = (_1244_v147).m7(_1246_v149, _module.__default.fm32(_dafny.Seq.update(_1041_v0, _module.__default.safeIndex(new BigNumber(767), new BigNumber((_1041_v0).length)), _1235_v139), true, globalState), _1248_v151, globalState);
          _out22 = _outcollector6[0];
          _out23 = _outcollector6[1];
          _out24 = _outcollector6[2];
          _1249_v152 = _out22;
          _1250_v153 = _out23;
          _1251_v154 = _out24;
        } else {
          (_this).f4 = _this.f4;
          _1235_v139 = (_this).f9;
          let _1252_v155;
          _1252_v155 = _dafny.Seq.UnicodeFromString("cklypmb");
          _1252_v155 = (_this).f5;
          let _1253_v156;
          _1253_v156 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_1235_v139);
          let _1254_v157;
          _1254_v157 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(316),_1253_v156);
          let _1255_v158;
          _1255_v158 = _dafny.Seq.of((_this).f5);
          let _1256_v159;
          _1256_v159 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_1254_v157).length)),(_1255_v158)[_module.__default.safeIndex((_this).f8, new BigNumber((_1255_v158).length))]);
          _1256_v159 = (_1256_v159).update(new BigNumber((_1041_v0).length), _1252_v155);
          let _1257_v160;
          _1257_v160 = _module.D9.create_DC25(_module.__default.fm27((_this).f8, (_this).f6, !(((_this).f10)[_module.__default.safeIndex(new BigNumber(461), new BigNumber(((_this).f10).length))]), globalState), (_this).f7, new BigNumber((_module.__default.fm33(globalState)).cardinality()));
          let _index210 = _module.__default.safeIndex(new BigNumber(461), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index210] = ((_dafny.ZERO).minus(((_this).f11)[_module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length))])).isLessThanOrEqualTo(((_1257_v160).dtor_cf47).multipliedBy(p0));
        }
        let _1258_v161;
        _1258_v161 = _dafny.Map.Empty.slice().updateUnsafe(_1235_v139,((_this).f11)[_module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length))]);
        let _1259_v162;
        _1259_v162 = _module.D5.create_DC15((_1236_v140)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_1236_v140).length))], false, _1258_v161, ((_this).f11)[_module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length))]);
        let _1260_v163;
        _1260_v163 = _dafny.Map.Empty.slice().updateUnsafe(_1235_v139,(_1042_v1)[_module.__default.safeIndex(((_this).f11)[_module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length))], new BigNumber((_1042_v1).length))]);
        let _1261_v164;
        let _nw241 = Array((new BigNumber(22)).toNumber());
        _nw241[(_dafny.ZERO).toNumber()] = _1259_v162;
        _nw241[(_dafny.ONE).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(2)).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(3)).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(4)).toNumber()] = _module.D5.create_DC15((_1236_v140)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_1236_v140).length))], (_this).f7, _1258_v161, ((_this).f11)[_module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length))]);
        _nw241[(new BigNumber(5)).toNumber()] = _module.D5.create_DC15((_1236_v140)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_1236_v140).length))], (_this).f7, _1258_v161, (_this).f8);
        _nw241[(new BigNumber(6)).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(7)).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(8)).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(9)).toNumber()] = _module.D5.create_DC15((_1236_v140)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_1236_v140).length))], _1235_v139, _dafny.Map.Empty.slice().updateUnsafe((_1042_v1)[_module.__default.safeIndex((_this).fm1((_this).f9, _1260_v163, ((_this).f10)[_module.__default.safeIndex(new BigNumber(461), new BigNumber(((_this).f10).length))], globalState), new BigNumber((_1042_v1).length))],p0), ((_this).f11)[_module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length))]);
        _nw241[(new BigNumber(10)).toNumber()] = _module.D5.create_DC15((_1236_v140)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_1236_v140).length))], (_this).f7, _1258_v161, new BigNumber(831));
        _nw241[(new BigNumber(11)).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(12)).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(13)).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(14)).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(15)).toNumber()] = _module.D5.create_DC15((_1236_v140)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_1236_v140).length))], ((_this).f10)[_module.__default.safeIndex(new BigNumber(461), new BigNumber(((_this).f10).length))], _1258_v161, new BigNumber(((_this).f5).length));
        _nw241[(new BigNumber(16)).toNumber()] = _module.D5.create_DC15((_1236_v140)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_1236_v140).length))], _module.__default.fm3(globalState), _1258_v161, new BigNumber(576));
        _nw241[(new BigNumber(17)).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(18)).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(19)).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(20)).toNumber()] = _1259_v162;
        _nw241[(new BigNumber(21)).toNumber()] = _1259_v162;
        _1261_v164 = _nw241;
        let _index211 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_1261_v164).length));
        (_1261_v164)[_index211] = _1259_v162;
        let _1262_v165;
        _1262_v165 = _module.D3.create_DC11(((_this).f10)[_module.__default.safeIndex(new BigNumber(461), new BigNumber(((_this).f10).length))], ((_this).f11)[_module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length))], _dafny.Set.fromElements((_this).f7));
        let _pat_let_tv20 = _1258_v161;
        let _pat_let_tv21 = _1262_v165;
        let _pat_let_tv22 = p0;
        let _index212 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_1261_v164).length));
        (_1261_v164)[_index212] = function (_pat_let39_0) {
          return function (_1265_dt__update__tmp_h5) {
            return function (_pat_let42_0) {
              return function (_1266_dt__update_hcf27_h0) {
                return function (_pat_let43_0) {
                  return function (_1267_dt__update_hcf28_h1) {
                    return _module.D5.create_DC15((_1265_dt__update__tmp_h5).dtor_cf26, _1266_dt__update_hcf27_h0, _1267_dt__update_hcf28_h1, (_1265_dt__update__tmp_h5).dtor_cf29);
                  }(_pat_let43_0);
                }((_dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv21).dtor_cf20,((_this).f11)[_module.__default.safeIndex(new BigNumber(162), new BigNumber(((_this).f11).length))])).update(!(((_this).f10)[_module.__default.safeIndex(new BigNumber(461), new BigNumber(((_this).f10).length))]), _pat_let_tv22));
              }(_pat_let42_0);
            }(!((_this).f6));
          }(_pat_let39_0);
        }(function (_pat_let40_0) {
          return function (_1263_dt__update__tmp_h4) {
            return function (_pat_let41_0) {
              return function (_1264_dt__update_hcf28_h0) {
                return _module.D5.create_DC15((_1263_dt__update__tmp_h4).dtor_cf26, (_1263_dt__update__tmp_h4).dtor_cf27, _1264_dt__update_hcf28_h0, (_1263_dt__update__tmp_h4).dtor_cf29);
              }(_pat_let41_0);
            }(_pat_let_tv20);
          }(_pat_let40_0);
        }(_1259_v162));
      }
      let _1268_v166;
      _1268_v166 = _dafny.MultiSet.fromElements(new BigNumber(537), p0);
      let _1269_v167;
      _1269_v167 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber((_1268_v166).cardinality()));
      r0 = (_1269_v167).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber(66))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(!((_this).f7))).cardinality()),p0)));
      return r0;
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _1270_v0;
      _1270_v0 = _dafny.MultiSet.fromElements((_this).f5, (_this).f5, (_this).f5);
      let _1271_v1;
      _1271_v1 = _dafny.Seq.of(_1270_v0, _1270_v0);
      let _1272_v2;
      _1272_v2 = _dafny.Map.Empty.slice().updateUnsafe((_1270_v0).IsProperSubsetOf((_1271_v1)[_module.__default.safeIndex(p1, new BigNumber((_1271_v1).length))]),_this.f4);
      _1272_v2 = (_1272_v2).update(!((_this).f9) || ((_this).f7), _this.f4);
      let _hi5 = p0;
      for (let _1273_i0 = p2; _1273_i0.isLessThan(_hi5); _1273_i0 = _1273_i0.plus(_dafny.ONE)) {
        let _1274_v3;
        _1274_v3 = _dafny.Set.fromElements((_this).f9);
        let _1275_v4;
        _1275_v4 = _dafny.Seq.of(_1274_v3, _1274_v3);
        let _1276_v5;
        _1276_v5 = _dafny.Seq.of((_this).f7);
        let _1277_v6;
        _1277_v6 = _dafny.Seq.of(!((_1274_v3).IsProperSubsetOf((_1275_v4)[_module.__default.safeIndex(new BigNumber((_1276_v5).length), new BigNumber((_1275_v4).length))])));
        _1276_v5 = _dafny.Seq.update(_dafny.Seq.Concat(_1277_v6, (((_this).f7) ? (_1277_v6) : (_dafny.Seq.of((_this).f6)))), _module.__default.safeIndex(new BigNumber(637), new BigNumber((_dafny.Seq.Concat(_1277_v6, (((_this).f7) ? (_1277_v6) : (_dafny.Seq.of((_this).f6))))).length)), (_this).f6);
        r0 = _module.__default.fm7(_this.f4, _1273_i0, (_this).f9, globalState);
        let _1278_v7;
        _1278_v7 = _module.D1.create_DC4(_this.f4, new BigNumber(501), _this.f4, p2, (_this).f6);
        let _1279_v8;
        let _nw242 = new _module.C0();
        _nw242.__ctor((_1278_v7).dtor_cf12, (_1277_v6)[_module.__default.safeIndex(p0, new BigNumber((_1277_v6).length))], _this.f4, (_this).f5);
        _1279_v8 = _nw242;
        let _1280_v9;
        _1280_v9 = _dafny.Seq.UnicodeFromString("jne");
        let _index213 = _module.__default.safeIndex(new BigNumber(596), new BigNumber(((_this).f11).length));
        ((_this).f11)[_index213] = p1;
        let _1281_v10;
        _1281_v10 = _dafny.MultiSet.fromElements((_this).f9);
        let _index214 = _module.__default.safeIndex(new BigNumber(596), new BigNumber(((_this).f11).length));
        let _rhs199 = _dafny.Seq.Concat((_this).f5, (_this).f5);
        let _rhs200 = true;
        let _rhs201 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm4((_this).f9, new BigNumber((_1280_v9).length), (_this).f9, p0, globalState)).IsProperSubsetOf(_1281_v10),_dafny.Seq.Concat(_1276_v5, _dafny.Seq.of((_this).f7, (_this).f9)))).length);
        let _rhs202 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qxofl"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(951)), function (_1282_i1) {
          return _this.f4;
        }));
        let _lhs152 = (_this).f11;
        let _lhs153 = _module.__default.safeIndex(new BigNumber(596), new BigNumber(((_this).f11).length));
        _1280_v9 = _rhs199;
        r0 = _rhs200;
        _lhs152[_lhs153] = _rhs201;
        _1280_v9 = _rhs202;
      }
      r0 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(414)), function (_1283_i2) {
        return _this.f4;
      }), (_this).f5);
      let _hi6 = new BigNumber((_dafny.Seq.of((_this).f8)).length);
      for (let _1284_i3 = p1; _1284_i3.isLessThan(_hi6); _1284_i3 = _1284_i3.plus(_dafny.ONE)) {
        let _1285_v11;
        let _nw243 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _1285_v11 = _nw243;
        _1285_v11 = _1285_v11;
        let _1286_v12;
        _1286_v12 = _dafny.Set.fromElements(new BigNumber(77));
        let _1287_v13;
        _1287_v13 = _dafny.Map.Empty.slice().updateUnsafe(!((_1286_v12).IsProperSubsetOf(_1286_v12)),(_this).f5);
        _1287_v13 = (_1287_v13).update((_this).f7, (_this).f5);
        let _1288_v14;
        _1288_v14 = _dafny.Set.fromElements(_this.f4);
        let _1289_v15;
        _1289_v15 = _dafny.Seq.of((_1288_v14).Difference(_1288_v14), _1288_v14, _1288_v14, (_1288_v14).Union(_1288_v14), _1288_v14);
        (globalState).f1 = new BigNumber(((_1289_v15)[_module.__default.safeIndex(_1284_i3, new BigNumber((_1289_v15).length))]).length);
        let _1290_v16;
        _1290_v16 = _dafny.Seq.of((_this).f9, _dafny.Seq.IsProperPrefixOf((_this).f5, (_this).f5));
        _1290_v16 = _dafny.Seq.update(_1290_v16, _module.__default.safeIndex((_1284_i3).minus((_this).f8), new BigNumber((_1290_v16).length)), _module.__default.fm27(p2, !((_this).f7), (_this).f9, globalState));
      }
      let _1291_v17;
      let _nw244 = new _module.C1();
      _nw244.__ctor(p2, false, _module.__default.fm27((_this).f8, (_this).f6, (_this).f9, globalState), (_this).f9, _this.f4, (_this).f5);
      _1291_v17 = _nw244;
      let _1292_v18;
      _1292_v18 = _module.D10.create_DC28((_this).f7, p2, (_this).f6, (_this).f9, _1291_v17);
      let _1293_v19;
      _1293_v19 = _dafny.MultiSet.fromElements((_1292_v18).dtor_cf51);
      (globalState).f1 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_1293_v19).cardinality()), p1));
      let _1294_i4;
      _1294_i4 = _dafny.ZERO;
      L5: {
        while ((_this).f7) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1294_i4)) {
              break L5;
            }
            _1294_i4 = (_1294_i4).plus(_dafny.ONE);
            let _1295_v20;
            let _nw245 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
            _1295_v20 = _nw245;
            let _1296_v21;
            _1296_v21 = _dafny.Seq.of(false);
            let _1297_v22;
            _1297_v22 = _dafny.Seq.of(_1296_v21);
            let _index215 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_1295_v20).length));
            (_1295_v20)[_index215] = _1297_v22;
            let _1298_v23;
            _1298_v23 = _dafny.Set.fromElements((_this).f9, (_this).f6, (_this).f7, (_this).f9, (_this).f9);
            let _1299_v24;
            _1299_v24 = _module.D8.create_DC22(_1296_v21);
            let _index216 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_1295_v20).length));
            let _rhs203 = _module.__default.fm34(!((_this).f9), _1299_v24, _module.__default.safeDivisionInt((_this).f8, _module.__default.fm8(p2, (_this).f7, (_this).f6, globalState)), globalState);
            let _rhs204 = (new BigNumber((_1296_v21).length)).plus(p2);
            let _rhs205 = _1298_v23;
            let _rhs206 = (_this).f6;
            let _lhs154 = _1295_v20;
            let _lhs155 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_1295_v20).length));
            let _lhs156 = globalState;
            _lhs154[_lhs155] = _rhs203;
            _lhs156.f1 = _rhs204;
            _1298_v23 = _rhs205;
            r0 = _rhs206;
            let _1300_v25;
            _1300_v25 = _dafny.Seq.UnicodeFromString("etoajigf");
            _1300_v25 = _dafny.Seq.UnicodeFromString("a");
            (_this).f4 = _this.f4;
            let _1301_v26;
            _1301_v26 = _dafny.Seq.of((_dafny.MultiSet.fromElements((_this).f9)).update((_this).f9, _module.__default.abs(p0)), _dafny.MultiSet.fromElements((_this).f6, (_this).f7, (_this).f6));
            let _1302_v27;
            _1302_v27 = _module.D3.create_DC11((_this).f9, p0, _1298_v23);
            let _1303_v28;
            let _init25 = function (_1304_i5) {
              return _this.f4;
            };
            let _nw246 = Array((new BigNumber(28)).toNumber());
            for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw246.length); _i0_25++) {
              _nw246[_i0_25] = _init25(new BigNumber(_i0_25));
            }
            _1303_v28 = _nw246;
            let _index217 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_1303_v28).length));
            (_1303_v28)[_index217] = ((_this).f5)[_module.__default.safeIndex(p0, new BigNumber(((_this).f5).length))];
            let _1305_v29;
            let _nw247 = Array((new BigNumber(6)).toNumber()).fill(_module.D1.Default());
            _1305_v29 = _nw247;
            let _1306_v30;
            let _nw248 = new _module.C2();
            _nw248.__ctor((_this).f10, _1305_v29);
            _1306_v30 = _nw248;
            let _1307_v31;
            _1307_v31 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f9);
            let _1308_v32;
            _1308_v32 = _dafny.MultiSet.fromElements((((_1307_v31).contains((_this).f7)) ? ((_1307_v31).get((_this).f7)) : (true)), (_this).f6, (_this).f6, !((_this).f9), (_this).f7);
            let _1309_v33;
            _1309_v33 = _dafny.Map.Empty.slice().updateUnsafe(true,p2);
            let _pat_let_tv23 = _1298_v23;
            let _pat_let_tv24 = globalState;
            let _index218 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_1303_v28).length));
            let _rhs207 = !((_this).f6);
            let _rhs208 = _dafny.Seq.update(_dafny.Seq.Concat(_1301_v26, _dafny.Seq.of(_dafny.MultiSet.FromArray(_1296_v21), _1308_v32)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_1301_v26, _dafny.Seq.of(_dafny.MultiSet.FromArray(_1296_v21), _1308_v32))).length)), _1308_v32);
            let _rhs209 = function (_pat_let44_0) {
              return function (_1312_dt__update__tmp_h1) {
                return function (_pat_let47_0) {
                  return function (_1313_dt__update_hcf20_h0) {
                    return _module.D3.create_DC11(_1313_dt__update_hcf20_h0, (_1312_dt__update__tmp_h1).dtor_cf21, (_1312_dt__update__tmp_h1).dtor_cf22);
                  }(_pat_let47_0);
                }(_module.__default.fm3(_pat_let_tv24));
              }(_pat_let44_0);
            }(function (_pat_let45_0) {
              return function (_1310_dt__update__tmp_h0) {
                return function (_pat_let46_0) {
                  return function (_1311_dt__update_hcf22_h0) {
                    return _module.D3.create_DC11((_1310_dt__update__tmp_h0).dtor_cf20, (_1310_dt__update__tmp_h0).dtor_cf21, _1311_dt__update_hcf22_h0);
                  }(_pat_let46_0);
                }(_pat_let_tv23);
              }(_pat_let45_0);
            }(_module.D3.create_DC11((_this).f9, new BigNumber((_1309_v33).length), _1298_v23)));
            let _rhs210 = _this.f4;
            let _rhs211 = _1306_v30;
            let _lhs157 = _1303_v28;
            let _lhs158 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_1303_v28).length));
            r0 = _rhs207;
            _1301_v26 = _rhs208;
            _1302_v27 = _rhs209;
            _lhs157[_lhs158] = _rhs210;
            _1306_v30 = _rhs211;
          }
        }
      }
      r0 = !(((false) ? ((_this).f7) : ((_this).f9)));
      return r0;
    }
    m4(p0, p1, globalState) {
      let _this = this;
      let _1314_v0;
      let _nw249 = new _module.C4();
      _nw249.__ctor((_this).f8, (_this).f6, (_this).f9, (_this).f7, _this.f4, _dafny.Seq.UnicodeFromString("vknfksjs"));
      _1314_v0 = _nw249;
      let _1315_v1;
      _1315_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1314_v0,(_this).f8);
      let _1316_v2;
      _1316_v2 = _module.D2.create_DC8(p0, new BigNumber((_1315_v1).length));
      let _pat_let_tv25 = _1314_v0;
      let _pat_let_tv26 = _1314_v0;
      let _pat_let_tv27 = _1314_v0;
      (globalState).f1 = function (_source15) {
        if (_source15.is_DC7) {
          let _1317___mcc_h0 = (_source15).cf15;
          let _1318_cf15 = _1317___mcc_h0;
          return (_this).f8;
        } else if (_source15.is_DC8) {
          let _1319___mcc_h1 = (_source15).cf16;
          let _1320___mcc_h2 = (_source15).cf17;
          let _1321_cf17 = _1320___mcc_h2;
          let _1322_cf16 = _1319___mcc_h1;
          return _1321_cf17;
        } else if (_source15.is_DC6) {
          let _1323___mcc_h3 = (_source15).cf14;
          let _1324_cf14 = _1323___mcc_h3;
          return (_dafny.ZERO).minus(new BigNumber(((_pat_let_tv25).f5).length));
        } else {
          let _1325___mcc_h4 = (_source15).cf18;
          let _1326_cf18 = _1325___mcc_h4;
          return new BigNumber(((_dafny.MultiSet.fromElements((_pat_let_tv26).f7)).Difference(_dafny.MultiSet.fromElements((_pat_let_tv27).f6))).cardinality());
        }
      }(_1316_v2);
      let _1327_v3;
      _1327_v3 = _dafny.MultiSet.fromElements(false, (_1314_v0).f7);
      let _1328_i0;
      _1328_i0 = _dafny.ZERO;
      L6: {
        while ((_1327_v3).IsDisjointFrom(_module.__default.fm4((_1314_v0).f6, (_this).f8, (_1314_v0).f6, p0, globalState))) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1328_i0)) {
              break L6;
            }
            _1328_i0 = (_1328_i0).plus(_dafny.ONE);
            let _1329_v4;
            _1329_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_1314_v0).f9);
            let _1330_v5;
            _1330_v5 = _module.D12.create_DC31((_1329_v4).update((_this).f6, (_1314_v0).f7));
            (globalState).f1 = _module.__default.safeModuloInt((_this).fm1((_this).f6, (_1330_v5).dtor_cf60, (_1314_v0).f6, globalState), _module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements(true)).length), (_this).f8));
            let _1331_v6;
            _1331_v6 = _dafny.Seq.of((_this).f8, (_1314_v0).fm1((_1314_v0).f9, _1329_v4, (_this).f7, globalState), _module.__default.safeModuloInt(new BigNumber((_1329_v4).length), (_1314_v0).f8), ((_1314_v0).f8).multipliedBy(p0), (_this).f8);
            _1331_v6 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1331_v6, _1331_v6), _1331_v6);
            let _index219 = _module.__default.safeIndex(new BigNumber(374), new BigNumber(((_this).f10).length));
            ((_this).f10)[_index219] = (_1314_v0).f7;
            let _index220 = _module.__default.safeIndex(new BigNumber(374), new BigNumber(((_this).f10).length));
            ((_this).f10)[_index220] = (_1314_v0).f9;
            let _1332_v7;
            _1332_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1331_v6,true);
            let _index221 = _module.__default.safeIndex(new BigNumber(374), new BigNumber(((_this).f10).length));
            ((_this).f10)[_index221] = !(_1332_v7).equals(((_module.__default.fm27((_1314_v0).f8, (_this).f9, false, globalState)) ? (_dafny.Map.Empty.slice().updateUnsafe(_1331_v6,(_1314_v0).f7)) : (_1332_v7)));
          }
        }
      }
      let _1333_v8;
      let _nw250 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1333_v8 = _nw250;
      let _nw251 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1333_v8 = _nw251;
      let _1334_v9;
      let _nw252 = new _module.C1();
      _nw252.__ctor((_1314_v0).f8, (_this).f9, (_this).f9, (_1314_v0).f7, _1314_v0.f4, _dafny.Seq.UnicodeFromString("gedhynvk"));
      _1334_v9 = _nw252;
      let _1335_v10;
      _1335_v10 = _dafny.MultiSet.fromElements((_module.D10.create_DC28((_this).f9, (_this).f8, (_1314_v0).f7, !((_this).f7), _1334_v9)).dtor_cf51);
      let _hi7 = new BigNumber(104);
      for (let _1336_i1 = new BigNumber((_1335_v10).cardinality()); _1336_i1.isLessThan(_hi7); _1336_i1 = _1336_i1.plus(_dafny.ONE)) {
        let _1337_v11;
        _1337_v11 = _dafny.Map.Empty.slice().updateUnsafe((_1314_v0).f8,(_1314_v0).f8);
        _1337_v11 = (_1337_v11).update(p0, (_dafny.ZERO).minus((_1314_v0).f8));
        if ((_1314_v0).f6) {
          let _1338_v12;
          _1338_v12 = true;
          let _1339_v13;
          _1339_v13 = _module.D12.create_DC33(_1338_v12, _dafny.Seq.UnicodeFromString("o"));
          let _1340_v14;
          _1340_v14 = _dafny.Map.Empty.slice().updateUnsafe(false,!((_1314_v0).f7));
          let _pat_let_tv28 = _1314_v0;
          let _pat_let_tv29 = globalState;
          let _pat_let_tv30 = _1340_v14;
          let _pat_let_tv31 = _1314_v0;
          let _pat_let_tv32 = globalState;
          let _pat_let_tv33 = _1314_v0;
          let _pat_let_tv34 = globalState;
          _1338_v12 = (function (_pat_let48_0) {
            return function (_1341_dt__update__tmp_h0) {
              return function (_pat_let49_0) {
                return function (_1342_dt__update_hcf63_h0) {
                  return _module.D12.create_DC33(_1342_dt__update_hcf63_h0, (_1341_dt__update__tmp_h0).dtor_cf64);
                }(_pat_let49_0);
              }(_module.__default.fm7(_pat_let_tv28.f4, (_this).fm1(_module.__default.fm3(_pat_let_tv29), _pat_let_tv30, (_pat_let_tv31).f6, _pat_let_tv32), (_pat_let_tv33).f7, _pat_let_tv34));
            }(_pat_let48_0);
          }(_1339_v13)).dtor_cf63;
          let _1343_v15;
          let _1344_v16;
          let _1345_v17;
          let _1346_v18;
          let _out25;
          let _out26;
          let _out27;
          let _out28;
          let _outcollector7 = (_1334_v9).m8((((_this).f7) ? ((_1314_v0).f6) : ((_this).f6)), (_1314_v0).f7, globalState);
          _out25 = _outcollector7[0];
          _out26 = _outcollector7[1];
          _out27 = _outcollector7[2];
          _out28 = _outcollector7[3];
          _1343_v15 = _out25;
          _1344_v16 = _out26;
          _1345_v17 = _out27;
          _1346_v18 = _out28;
          let _1347_v19;
          let _nw253 = Array((new BigNumber(17)).toNumber()).fill([]);
          _1347_v19 = _nw253;
          let _nw254 = Array((new BigNumber(16)).toNumber()).fill([]);
          _1347_v19 = _nw254;
          _1338_v12 = _1338_v12;
          let _index222 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_1346_v18).length));
          (_1346_v18)[_index222] = (_this).f7;
          let _index223 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_1346_v18).length));
          (_1346_v18)[_index223] = (_1314_v0).f7;
        } else {
          let _1348_v20;
          _1348_v20 = true;
          let _rhs212 = !((((p1).contains((_this).f8)) ? ((p1).get((_this).f8)) : ((_1314_v0).f6))) || ((_1314_v0).f7);
          let _rhs213 = !((_1314_v0).f7);
          _1348_v20 = _rhs212;
          _1348_v20 = _rhs213;
          let _rhs214 = (_this).f8;
          let _rhs215 = _1314_v0;
          let _lhs159 = globalState;
          _lhs159.f1 = _rhs214;
          _1314_v0 = _rhs215;
          _1348_v20 = true;
          let _1349_v21;
          _1349_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1348_v20,(_this).f7);
          (globalState).f1 = (_1334_v9).fm1(true, _1349_v21, (_this).f6, globalState);
          let _1350_v22;
          _1350_v22 = _dafny.Map.Empty.slice().updateUnsafe((_1314_v0).f8,true);
          let _1351_v23;
          _1351_v23 = _dafny.Seq.of((_1314_v0).f9);
          _1350_v22 = (_1350_v22).update(new BigNumber(((_1337_v11).Merge(_1337_v11)).length), ((_1351_v23)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1337_v11).length)), new BigNumber((_1351_v23).length))]) || (!((_1314_v0).f7)));
        }
        (globalState).f1 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(((_1335_v10).Difference(_1335_v10)).cardinality()), new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(603)), ((_1352_v0) => function (_1353_i2) {
          return _1352_v0.f4;
        })(_1314_v0)), _module.__default.safeIndex(_module.__default.fm8((_this).f8, (_this).f9, (_1314_v0).f6, globalState), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(603)), ((_1354_v0) => function (_1355_i2) {
          return _1354_v0.f4;
        })(_1314_v0))).length)), _1314_v0.f4)).length)));
        let _index224 = _module.__default.safeIndex(new BigNumber(446), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index224] = !(new BigNumber((_dafny.Seq.UnicodeFromString("bhpixis")).length)).isEqualTo(_1336_i1);
        let _1356_v24;
        _1356_v24 = true;
        let _1357_v25;
        let _nw255 = Array((new BigNumber(6)).toNumber()).fill([]);
        _1357_v25 = _nw255;
        let _1358_v26;
        _1358_v26 = _dafny.Seq.of((_1314_v0).f7, (_this).f7);
        let _1359_v27;
        _1359_v27 = _dafny.Map.Empty.slice().updateUnsafe((_1314_v0).f6,new BigNumber((_1358_v26).length));
        let _1360_v28;
        _1360_v28 = _module.D5.create_DC15((_this).f10, (_1314_v0).f9, _1359_v27, new BigNumber(368));
        let _1361_v29;
        _1361_v29 = _dafny.Seq.of((_this).f10, (_this).f10, (_this).f10, (_this).f10);
        let _1362_v30;
        let _nw256 = Array((new BigNumber(29)).toNumber());
        _nw256[(_dafny.ZERO).toNumber()] = (_this).f10;
        _nw256[(_dafny.ONE).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(2)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(3)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(4)).toNumber()] = (_1360_v28).dtor_cf26;
        _nw256[(new BigNumber(5)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(6)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(7)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(8)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(9)).toNumber()] = (_1361_v29)[_module.__default.safeIndex((_this).f8, new BigNumber((_1361_v29).length))];
        _nw256[(new BigNumber(10)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(11)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(12)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(13)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(14)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(15)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(16)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(17)).toNumber()] = (_1361_v29)[_module.__default.safeIndex((_this).f8, new BigNumber((_1361_v29).length))];
        _nw256[(new BigNumber(18)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(19)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(20)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(21)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(22)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(23)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(24)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(25)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(26)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(27)).toNumber()] = (_this).f10;
        _nw256[(new BigNumber(28)).toNumber()] = (_this).f10;
        _1362_v30 = _nw256;
        let _index225 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_1357_v25).length));
        (_1357_v25)[_index225] = _1362_v30;
        let _1363_v31;
        _1363_v31 = _dafny.Seq.of((_this).f8);
        let _1364_v32;
        _1364_v32 = _dafny.Seq.of(_1363_v31);
        let _index226 = _module.__default.safeIndex(new BigNumber(779), new BigNumber(((_this).f11).length));
        ((_this).f11)[_index226] = new BigNumber((_dafny.Seq.update(_dafny.Seq.of(p0, (_this).f8), _module.__default.safeIndex((_this).f8, new BigNumber((_dafny.Seq.of(p0, (_this).f8)).length)), new BigNumber((_1364_v32).length))).length);
        let _1365_v33;
        _1365_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,false);
        let _1366_v34;
        _1366_v34 = _module.D8.create_DC23((_1334_v9).fm1((_1314_v0).f9, _1365_v33, (_1314_v0).f6, globalState), (_this).f9);
        let _1367_v35;
        _1367_v35 = _dafny.Set.fromElements((_this).f8);
        let _index227 = _module.__default.safeIndex(new BigNumber(446), new BigNumber(((_this).f10).length));
        let _index228 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_1357_v25).length));
        let _index229 = _module.__default.safeIndex(new BigNumber(779), new BigNumber(((_this).f11).length));
        let _rhs216 = ((false) ? ((_1314_v0).f9) : ((_1366_v34).dtor_cf43));
        let _rhs217 = _module.__default.fm27((_dafny.ZERO).minus(new BigNumber((_1335_v10).cardinality())), (_1367_v35).IsSubsetOf(_1367_v35), (_this).f7, globalState);
        let _rhs218 = _1362_v30;
        let _rhs219 = (_dafny.ZERO).minus(new BigNumber(-822));
        let _lhs160 = (_this).f10;
        let _lhs161 = _module.__default.safeIndex(new BigNumber(446), new BigNumber(((_this).f10).length));
        let _lhs162 = _1357_v25;
        let _lhs163 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_1357_v25).length));
        let _lhs164 = (_this).f11;
        let _lhs165 = _module.__default.safeIndex(new BigNumber(779), new BigNumber(((_this).f11).length));
        _lhs160[_lhs161] = _rhs216;
        _1356_v24 = _rhs217;
        _lhs162[_lhs163] = _rhs218;
        _lhs164[_lhs165] = _rhs219;
      }
      let _index230 = _module.__default.safeIndex(new BigNumber(340), new BigNumber(((_this).f11).length));
      ((_this).f11)[_index230] = ((_dafny.ZERO).minus(p0)).plus((_this).f8);
      let _1368_v36;
      _1368_v36 = _dafny.Set.fromElements((_1314_v0).f8, (_1314_v0).f8, (_1314_v0).f8, p0, new BigNumber(852));
      let _1369_v37;
      _1369_v37 = _dafny.Map.Empty.slice().updateUnsafe((_1314_v0).f8,_1335_v10);
      let _1370_v38;
      _1370_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1368_v36,(_module.D9.create_DC25((_this).f9, (_this).f6, new BigNumber((_1369_v37).length))).dtor_cf47);
      let _index231 = _module.__default.safeIndex(new BigNumber(340), new BigNumber(((_this).f11).length));
      ((_this).f11)[_index231] = (((_1370_v38).contains(_dafny.Set.fromElements((_this).f8, (_1314_v0).f8, (_this).f8, new BigNumber(((_this).f5).length), p0))) ? ((_1370_v38).get(_dafny.Set.fromElements((_this).f8, (_1314_v0).f8, (_this).f8, new BigNumber(((_this).f5).length), p0))) : (_module.__default.safeModuloInt(p0, (_1314_v0).f8)));
      let _1371_v39;
      _1371_v39 = false;
      _1371_v39 = !((_this).f6);
      return;
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
