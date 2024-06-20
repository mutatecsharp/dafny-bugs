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
    static fm0(p0, globalState) {
      return !(new BigNumber(834)).isEqualTo((new BigNumber(-650)).multipliedBy(new BigNumber((_dafny.Seq.of(!(true))).length)));
    };
    static fm1(globalState) {
      return (_dafny.ZERO).minus(new BigNumber((function (_source0) {
        let _0___mcc_h0 = _source0;
        let _1_cf0 = _0___mcc_h0;
        return (_dafny.Set.fromElements(_1_cf0)).Union(_dafny.Set.fromElements(_1_cf0, _1_cf0, _1_cf0));
      }(_dafny.Seq.Create(_module.__default.abs(new BigNumber(597)), function (_2_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      }))).length));
    };
    static fm8(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vpftgh"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(639)), function (_3_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("jimxsttfl")));
    };
    static fm9(p0, globalState) {
      let _source1 = _dafny.Set.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true));
      let _4___mcc_h0 = _source1;
      let _5_cf62 = _4___mcc_h0;
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(796)), function (_6_i0) {
        return new BigNumber(875);
      });
    };
    static fm10(globalState) {
      return (((false) ? (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ouixhpwjj")).length), new BigNumber(((_module.D9.create_DC20(new BigNumber((_dafny.Seq.of(false, true, false, false, false)).length), true, false, false, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(315))).length),new BigNumber(106)))).dtor_cf38).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-680),_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0))))).length))) : ((_module.D9.create_DC17(function () {
  let _coll0 = new _dafny.Set();
  for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(786),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-631),new BigNumber(-19))).length))).Keys.Elements) {
    let _7_v0 = _compr_0;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(786),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-631),new BigNumber(-19))).length))).contains(_7_v0)) {
      _coll0.add(_module.__default.safeDivisionInt(_7_v0, new BigNumber(389)));
    }
  }
  return _coll0;
}())).dtor_cf31))).Union((_dafny.Set.fromElements(new BigNumber(531), new BigNumber((_dafny.Seq.of(false, true, false)).length), new BigNumber((_dafny.Set.fromElements(!(false))).length))).Intersect(function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-219)), function (_8_i0) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("suq")).length),!(false))).length);
        })).Elements) {
          let _9_v1 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-219)), function (_8_i0) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("suq")).length),!(false))).length);
          }), _9_v1)) {
            _coll1.add((_9_v1).plus(new BigNumber(199)));
          }
        }
        return _coll1;
      }()));
    };
    static fm12(p0, p1, globalState) {
      if (true) {
        if (true) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }
      } else {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }
    };
    static fm13(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false, false, true, false, false), (_module.D7.create_DC13(_dafny.Seq.of(true, true))).dtor_cf23), ((false) ? (_dafny.Seq.of(false, true)) : ((_module.D7.create_DC13(_dafny.Seq.of(false))).dtor_cf23)));
    };
    static fm14(p0, p1, p2, p3, globalState) {
      let _source2 = _module.D18.create_DC46(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,false)));
      if (_source2.is_DC47) {
        let _10___mcc_h0 = (_source2).cf86;
        let _11___mcc_h1 = (_source2).cf87;
        let _12_cf87 = _11___mcc_h1;
        let _13_cf86 = _10___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("a"), _dafny.Seq.UnicodeFromString("obdympjc"));
      } else {
        let _14___mcc_h2 = (_source2).cf85;
        let _15_cf85 = _14___mcc_h2;
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(886)), function (_16_i0) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("rw"));
      }
    };
    static fm15(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(-721)).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(771), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).cardinality()))).length)),new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(true, !((_module.D16.create_DC40(true, new BigNumber(499), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false)).length)))).dtor_cf74)), _dafny.Seq.of(true, true, false))).length));
    };
    static fm16(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(767)), _dafny.Seq.of(new BigNumber(676))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false, true)).length)), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),false)).length),new BigNumber((_dafny.Seq.UnicodeFromString("otnqi")).length))).length))));
    };
    static fm17(p0, p1, p2, globalState) {
      let _source3 = _module.D15.create_DC36(new _dafny.CodePoint('c'.codePointAt(0)), new BigNumber(604));
      if (_source3.is_DC36) {
        let _17___mcc_h0 = (_source3).cf64;
        let _18___mcc_h1 = (_source3).cf65;
        let _19_cf65 = _18___mcc_h1;
        let _20_cf64 = _17___mcc_h0;
        return _dafny.Set.fromElements(false, !(true), !(true), !(true));
      } else if (_source3.is_DC37) {
        let _21___mcc_h2 = (_source3).cf66;
        let _22___mcc_h3 = (_source3).cf67;
        let _23___mcc_h4 = (_source3).cf68;
        let _24___mcc_h5 = (_source3).cf69;
        let _25_cf69 = _24___mcc_h5;
        let _26_cf68 = _23___mcc_h4;
        let _27_cf67 = _22___mcc_h3;
        let _28_cf66 = _21___mcc_h2;
        return (_dafny.Set.fromElements(_28_cf66)).Intersect(_dafny.Set.fromElements(_25_cf69));
      } else if (_source3.is_DC38) {
        let _29___mcc_h6 = (_source3).cf70;
        let _30___mcc_h7 = (_source3).cf71;
        let _31___mcc_h8 = (_source3).cf72;
        let _32_cf72 = _31___mcc_h8;
        let _33_cf71 = _30___mcc_h7;
        let _34_cf70 = _29___mcc_h6;
        return (_module.D24.create_DC58(_dafny.Set.fromElements(_33_cf71, _33_cf71))).dtor_cf109;
      } else {
        let _35___mcc_h9 = (_source3).cf63;
        let _36_cf63 = _35___mcc_h9;
        return _dafny.Set.fromElements(true);
      }
    };
    static fm18(globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(936)), function (_37_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length))), _dafny.Seq.of(new BigNumber(-501)))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(878))));
    };
    static fm19(globalState) {
      if ((new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),false)).length))).length)).isLessThanOrEqualTo(new BigNumber(896))) {
        return _dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)));
      } else {
        return _dafny.Set.fromElements(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)));
      }
    };
    static fm20(p0, globalState) {
      return _dafny.Seq.of(_dafny.Seq.of(true, false), ((true) ? (_dafny.Seq.of(false)) : (_dafny.Seq.of(!(true)))));
    };
    static fm23(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(((false) ? (_dafny.Map.Empty.slice().updateUnsafe(false,!(true))) : (_dafny.Map.Empty.slice().updateUnsafe(true,false))));
    };
    static fm24(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new BigNumber(295))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(856)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("skaroir")).length)));
    };
    static fm25(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false, false), _dafny.Seq.of(false, !(true))), _dafny.Seq.of(!(true), false));
    };
    static fm26(p0, p1, p2, globalState) {
      let _source4 = _dafny.MultiSet.fromElements(new BigNumber(520), new BigNumber(-680));
      let _38___mcc_h0 = _source4;
      let _39_cf18 = _38___mcc_h0;
      return new _dafny.CodePoint('h'.codePointAt(0));
    };
    static fm27(p0, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber(264))).Union(_dafny.MultiSet.fromElements(new BigNumber(((_module.D9.create_DC17(_dafny.Set.fromElements(new BigNumber(743)))).dtor_cf31).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(307)), function (_40_i0) {
        return _dafny.Set.fromElements(true);
      })).length)));
    };
    static fm28(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC1(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, true)).length),new BigNumber((_dafny.Seq.UnicodeFromString("g")).length))).length),true));
    };
    static fm29(p0, p1, p2, globalState) {
      return ((true) ? (true) : (false));
    };
    static fm30(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(229),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(978),true),false)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(138),new BigNumber(704)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-387),new BigNumber(884)));
    };
    static fm31(p0, p1, globalState) {
      let _source5 = _module.D9.create_DC18(!(true));
      if (_source5.is_DC18) {
        let _41___mcc_h0 = (_source5).cf32;
        let _42_cf32 = _41___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rpb"), _dafny.Seq.UnicodeFromString("ypsbjt"));
      } else if (_source5.is_DC19) {
        let _43___mcc_h1 = (_source5).cf33;
        let _44_cf33 = _43___mcc_h1;
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xa"), _dafny.Seq.UnicodeFromString("vxt"));
      } else if (_source5.is_DC20) {
        let _45___mcc_h2 = (_source5).cf34;
        let _46___mcc_h3 = (_source5).cf35;
        let _47___mcc_h4 = (_source5).cf36;
        let _48___mcc_h5 = (_source5).cf37;
        let _49___mcc_h6 = (_source5).cf38;
        let _50_cf38 = _49___mcc_h6;
        let _51_cf37 = _48___mcc_h5;
        let _52_cf36 = _47___mcc_h4;
        let _53_cf35 = _46___mcc_h3;
        let _54_cf34 = _45___mcc_h2;
        if (_51_cf37) {
          return _dafny.Seq.UnicodeFromString("eg");
        } else {
          return _dafny.Seq.UnicodeFromString("uj");
        }
      } else {
        let _55___mcc_h7 = (_source5).cf31;
        let _56_cf31 = _55___mcc_h7;
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(885)), function (_57_i0) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(865)), function (_58_i1) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        }));
      }
    };
    static fm32(globalState) {
      return _dafny.MultiSet.fromElements(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("y"), _dafny.Seq.UnicodeFromString("qaed")), (_dafny.MultiSet.fromElements(new BigNumber(672))).IsDisjointFrom(_dafny.MultiSet.fromElements(new BigNumber(573), new BigNumber(-43), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false, true)).length))).cardinality()))).length), new BigNumber((_dafny.Seq.UnicodeFromString("m")).length), new BigNumber(950))).length), new BigNumber(615), new BigNumber(-460))));
    };
    static fm33(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true, true), _dafny.Seq.of(false, true)), _dafny.Seq.of(true, !(false)));
    };
    static fm36(p0, globalState) {
      let _source6 = _module.D20.create_DC51();
      if (_source6.is_DC51) {
        if (!(true)) {
          return _module.D2.create_DC4(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()), new BigNumber(608), new BigNumber(-857), new _dafny.CodePoint('f'.codePointAt(0)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(199)), function (_59_i0) {
  return new _dafny.CodePoint('r'.codePointAt(0));
}));
        } else {
          return _module.D2.create_DC4(new BigNumber(478), new BigNumber((_dafny.Seq.UnicodeFromString("crgitqdfk")).length), new BigNumber((_dafny.Seq.UnicodeFromString("hlk")).length), new _dafny.CodePoint('k'.codePointAt(0)), _dafny.Seq.UnicodeFromString("ovgnevf"));
        }
      } else {
        let _60___mcc_h0 = (_source6).cf92;
        let _61_cf92 = _60___mcc_h0;
        return _module.D2.create_DC4((_dafny.ZERO).minus(new BigNumber(-718)), new BigNumber(896), new BigNumber(-841), new _dafny.CodePoint('u'.codePointAt(0)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-641)), function (_62_i1) {
  return new _dafny.CodePoint('b'.codePointAt(0));
}));
      }
    };
    static fm37(p0, p1, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(659)), function (_63_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      });
    };
    static fm38(p0, globalState) {
      return new _dafny.CodePoint('g'.codePointAt(0));
    };
    static fm39(globalState) {
      let _source7 = function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),false)).Keys.Elements) {
            let _64_v0 = _compr_3;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),false)).contains(_64_v0)) {
              _coll3.push([_64_v0,!(false)]);
            }
          }
          return _coll3;
        }()).Keys.Elements) {
          let _65_v1 = _compr_2;
          if ((function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),false)).Keys.Elements) {
              let _64_v0 = _compr_4;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),false)).contains(_64_v0)) {
                _coll4.push([_64_v0,!(false)]);
              }
            }
            return _coll4;
          }()).contains(_65_v1)) {
            _coll2.add(_65_v1);
          }
        }
        return _coll2;
      }();
      let _66___mcc_h0 = _source7;
      let _67_cf62 = _66___mcc_h0;
      return _module.D11.create_DC24(new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality()), _module.D4.create_DC8(new _dafny.CodePoint('n'.codePointAt(0)), new BigNumber(190), false));
    };
    static fm40(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ofqxguj"),_dafny.Seq.of(true, true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vh"),_dafny.Seq.of(!(!(false)))));
    };
    static fm41(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(!(true)), _dafny.Seq.Concat(_dafny.Seq.of(false, false), _dafny.Seq.of(false)));
    };
    static fm42(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(((false) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),false)).length)) : (new BigNumber(213))),_module.__default.safeDivisionInt(new BigNumber(661), new BigNumber(501)));
    };
    static fm43(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(545)), function (_68_i0) {
        return new BigNumber((function () {
          let _coll5 = new _dafny.Set();
          for (const _compr_5 of (_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))).Elements) {
            let _69_v0 = _compr_5;
            if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))), _69_v0)) {
              _coll5.add(_69_v0);
            }
          }
          return _coll5;
        }()).length);
      }), _dafny.Seq.of(new BigNumber(676), new BigNumber(-823))), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(313)), function (_70_i1) {
        return new BigNumber(-438);
      }), _dafny.Seq.of(new BigNumber(631))));
    };
    static fm44(globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-592)), function (_71_i0) {
        return _dafny.Seq.of(false, true, false);
      });
    };
    static fm47(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(false)).length)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_module.D6.create_DC12(false, _dafny.Seq.of(_dafny.Seq.of(true, true), _dafny.Seq.of(true)), _dafny.Seq.UnicodeFromString("kvltui")), _module.D6.create_DC12(false, _dafny.Seq.of(_dafny.Seq.of(false)), (_module.D6.create_DC12(false, _dafny.Seq.of(_dafny.Seq.of(!(true), true)), _dafny.Seq.UnicodeFromString("dyruvhv"))).dtor_cf22))).length))).length)),!(true)));
    };
    static fm48(p0, p1, p2, globalState) {
      return (_module.D1.create_DC1(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(876), new BigNumber(396))).length),true))).dtor_cf1;
    };
    static fm49(p0, globalState) {
      if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(587)))).IsProperSubsetOf((_dafny.MultiSet.fromElements(new BigNumber(-927))))) {
        return _module.D4.create_DC9(new BigNumber((_dafny.Seq.of(false, true, true)).length), !(false));
      } else {
        return _module.D4.create_DC9(new BigNumber((_dafny.Seq.UnicodeFromString("okgwkrflp")).length), true);
      }
    };
    static fm50(globalState) {
      return (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(-469), new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)))).cardinality())), _dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-827)))))).Difference(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(256)), function (_72_i0) {
        return new BigNumber(((_module.D24.create_DC58(_dafny.Set.fromElements(false, true))).dtor_cf109).length);
      }), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(880))).length)))).length)), new BigNumber(942), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-878), new BigNumber(888))).cardinality())), _dafny.Seq.of(new BigNumber((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of _dafny.IntegerRange(new BigNumber(996), new BigNumber(410))) {
          let _73_v0 = _compr_6;
          if (((new BigNumber(996)).isLessThanOrEqualTo(_73_v0)) && ((_73_v0).isLessThan(new BigNumber(410)))) {
            _coll6.push([_module.__default.safeDivisionInt(_73_v0, new BigNumber(72)),false]);
          }
        }
        return _coll6;
      }()).length))));
    };
    static fm51(p0, p1, p2, globalState) {
      return _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),!(false))), (function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)))).Elements) {
          let _74_v0 = _compr_7;
          if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)))).contains(_74_v0)) {
            _coll7.push([_74_v0,true]);
          }
        }
        return _coll7;
      }()).Merge(function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),false)).Keys.Elements) {
          let _75_v1 = _compr_8;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),false)).contains(_75_v1)) {
            _coll8.push([_75_v1,true]);
          }
        }
        return _coll8;
      }()));
    };
    static fm52(p0, p1, p2, p3, globalState) {
      return _module.D9.create_DC18(false);
    };
    static fm53(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of(false, true, true), _dafny.Seq.of(false)),(_dafny.Set.fromElements(false)).IsProperSubsetOf(_dafny.Set.fromElements(true, false, true, true, false)));
    };
    static fm54(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-113)), function (_76_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })).length),new BigNumber(80))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('n'.codePointAt(0)))).length),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(524), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-541)), function (_77_i1) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length)))).length)))),new BigNumber((_dafny.Set.fromElements(new BigNumber(868), new BigNumber(-924))).length));
    };
    static fm55(p0, p1, p2, p3, globalState) {
      return _module.D16.create_DC40(!(!(true)), new BigNumber(((function () {
  let _coll9 = new _dafny.Map();
  for (const _compr_9 of (_dafny.Set.fromElements(new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)))).Elements) {
    let _78_v0 = _compr_9;
    if ((_dafny.Set.fromElements(new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)))).contains(_78_v0)) {
      _coll9.push([_78_v0,false]);
    }
  }
  return _coll9;
}()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),true))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(415),new BigNumber(651))).length));
    };
    static fm56(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(82)), function (_79_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("vgydj"), _dafny.Seq.UnicodeFromString("mxhrte"))).Union(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("fqymdm")))).Intersect(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("hydj")));
    };
    static fm57(p0, p1, globalState) {
      return _module.D15.create_DC35(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((function () {
  let _coll10 = new _dafny.Map();
  for (const _compr_10 of _dafny.IntegerRange(new BigNumber(364), new BigNumber(789))) {
    let _80_v0 = _compr_10;
    if (((new BigNumber(364)).isLessThanOrEqualTo(_80_v0)) && ((_80_v0).isLessThan(new BigNumber(789)))) {
      _coll10.push([(_80_v0).multipliedBy(new BigNumber(171)),_dafny.MultiSet.fromElements(false)]);
    }
  }
  return _coll10;
}()).length)),new _dafny.CodePoint('q'.codePointAt(0)))));
    };
    static m0(p0, p1, p2, p3, globalState) {
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _pat_let_tv0 = p1;
      let _source8 = function (_pat_let0_0) {
        return function (_81_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_82_dt__update_hcf56_h0) {
              return _module.D13.create_DC31(_82_dt__update_hcf56_h0);
            }(_pat_let1_0);
          }((_dafny.ZERO).minus((_dafny.ZERO).minus(_pat_let_tv0)));
        }(_pat_let0_0);
      }(_module.D13.create_DC31(p2));
      if (_source8.is_DC31) {
        let _83___mcc_h0 = (_source8).cf56;
        let _84_cf56 = _83___mcc_h0;
        let _85_v0;
        let _init0 = function (_86_i0) {
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wqcmjr"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(180)), function (_87_i1) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }));
        };
        let _nw0 = Array((new BigNumber(14)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
          _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _85_v0 = _nw0;
        _85_v0 = _85_v0;
        let _88_v1;
        _88_v1 = _module.D4.create_DC9(new BigNumber(-389), p0);
        _88_v1 = ((p0) ? (_88_v1) : (_88_v1));
        if (_module.__default.fm0((p3).minus(_84_cf56), globalState)) {
          let _89_v2;
          _89_v2 = _dafny.Set.fromElements(p2, p2);
          let _90_v3;
          _90_v3 = _dafny.Seq.UnicodeFromString("hqrd");
          let _91_v4;
          _91_v4 = _dafny.Map.Empty.slice().updateUnsafe(true,_90_v3);
          let _92_v5;
          _92_v5 = _dafny.MultiSet.fromElements(p0, p0, p0, p0, p0);
          let _93_v6;
          _93_v6 = _module.D10.create_DC22(new BigNumber((_91_v4).length), _92_v5);
          let _94_v7;
          _94_v7 = _dafny.Seq.of(p0);
          let _95_v8;
          _95_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,_85_v0);
          let _96_v9;
          let _nw1 = new _module.C1();
          _nw1.__ctor((_89_v2).IsDisjointFrom(_module.__default.fm10(globalState)), (_93_v6).dtor_cf41, _module.__default.fm16(p3, globalState), new BigNumber((_94_v7).length), (((_95_v8).contains(new BigNumber(471))) ? ((_95_v8).get(new BigNumber(471))) : (_85_v0)));
          _96_v9 = _nw1;
          let _97_v10;
          let _nw2 = new _module.C0();
          _nw2.__ctor(_84_cf56, _dafny.Seq.Concat(_90_v3, _90_v3));
          _97_v10 = _nw2;
          _97_v10 = _97_v10;
          let _98_v12;
          let _nw3 = new _module.C3();
          _nw3.__ctor(p1, _85_v0);
          _98_v12 = _nw3;
          let _99_v13;
          _99_v13 = _dafny.MultiSet.fromElements(_98_v12, _98_v12);
          let _100_v14;
          _100_v14 = _dafny.MultiSet.fromElements(new BigNumber((_90_v3).length), new BigNumber((_99_v13).cardinality()));
          let _101_v15;
          _101_v15 = _dafny.Seq.of((_97_v10).f36, p2);
          let _102_v17;
          _102_v17 = _dafny.Map.Empty.slice().updateUnsafe(_96_v9.f35,p0);
          (globalState).f3 = (_module.__default.safeModuloInt(new BigNumber(((function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of (_100_v14).Elements) {
              let _103_v11 = _compr_11;
              if ((_100_v14).contains(_103_v11)) {
                _coll11.push([_module.__default.safeDivisionInt(_103_v11, (_97_v10).f36),new BigNumber(((_97_v10).f37).length)]);
              }
            }
            return _coll11;
          }()).update((_101_v15)[_module.__default.safeIndex((_dafny.ZERO).minus(p3), new BigNumber((_101_v15).length))], p2)).length), new BigNumber((function () {
            let _coll12 = new _dafny.Set();
            for (const _compr_12 of _dafny.IntegerRange(new BigNumber(719), new BigNumber(727))) {
              let _104_v16 = _compr_12;
              if (((new BigNumber(719)).isLessThanOrEqualTo(_104_v16)) && ((_104_v16).isLessThan(new BigNumber(727)))) {
                _coll12.add((_104_v16).multipliedBy(new BigNumber(105)));
              }
            }
            return _coll12;
          }()).length))).multipliedBy((new BigNumber((_102_v17).length)).plus(_84_cf56));
          let _105_v18;
          let _nw4 = new _module.C6();
          _nw4.__ctor((_96_v9).fm3(_92_v5, globalState), new BigNumber(-792), (_92_v5).Difference(_module.__default.fm32(globalState)), _101_v15, p2, _85_v0, (_93_v6).dtor_cf41, new BigNumber((_module.__default.fm8(globalState)).length));
          _105_v18 = _nw4;
          let _106_v19;
          let _nw5 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
          _106_v19 = _nw5;
          let _index0 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_106_v19).length));
          (_106_v19)[_index0] = new BigNumber((_100_v14).cardinality());
          let _index1 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_106_v19).length));
          (_106_v19)[_index1] = ((true) ? ((_97_v10).f36) : ((((_92_v5).contains((_105_v18).f33)) ? ((_92_v5).get((_105_v18).f33)) : (_105_v18.f34))));
        } else {
          _84_cf56 = _module.__default.safeDivisionInt(p2, p1);
          let _107_v20;
          _107_v20 = new _dafny.CodePoint('s'.codePointAt(0));
          let _108_v21;
          _108_v21 = _dafny.Set.fromElements(_107_v20);
          let _109_v24;
          let _nw6 = Array((new BigNumber(19)).toNumber());
          _nw6[(_dafny.ZERO).toNumber()] = _108_v21;
          _nw6[(_dafny.ONE).toNumber()] = _108_v21;
          _nw6[(new BigNumber(2)).toNumber()] = _108_v21;
          _nw6[(new BigNumber(3)).toNumber()] = _108_v21;
          _nw6[(new BigNumber(4)).toNumber()] = _module.__default.fm19(globalState);
          _nw6[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(_107_v20);
          _nw6[(new BigNumber(6)).toNumber()] = _108_v21;
          _nw6[(new BigNumber(7)).toNumber()] = _108_v21;
          _nw6[(new BigNumber(8)).toNumber()] = _108_v21;
          _nw6[(new BigNumber(9)).toNumber()] = _108_v21;
          _nw6[(new BigNumber(10)).toNumber()] = _108_v21;
          _nw6[(new BigNumber(11)).toNumber()] = (function () {
            let _coll13 = new _dafny.Set();
            for (const _compr_13 of (function () {
              let _coll14 = new _dafny.Map();
              for (const _compr_14 of (_dafny.Map.Empty.slice().updateUnsafe(_107_v20,_107_v20)).Keys.Elements) {
                let _110_v22 = _compr_14;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_107_v20,_107_v20)).contains(_110_v22)) {
                  _coll14.push([_110_v22,false]);
                }
              }
              return _coll14;
            }()).Keys.Elements) {
              let _111_v23 = _compr_13;
              if ((function () {
                let _coll15 = new _dafny.Map();
                for (const _compr_15 of (_dafny.Map.Empty.slice().updateUnsafe(_107_v20,_107_v20)).Keys.Elements) {
                  let _110_v22 = _compr_15;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(_107_v20,_107_v20)).contains(_110_v22)) {
                    _coll15.push([_110_v22,false]);
                  }
                }
                return _coll15;
              }()).contains(_111_v23)) {
                _coll13.add(_111_v23);
              }
            }
            return _coll13;
          }()).Union(_108_v21);
          _nw6[(new BigNumber(12)).toNumber()] = (_108_v21).Difference(_108_v21);
          _nw6[(new BigNumber(13)).toNumber()] = _108_v21;
          _nw6[(new BigNumber(14)).toNumber()] = _dafny.Set.fromElements(_107_v20);
          _nw6[(new BigNumber(15)).toNumber()] = _108_v21;
          _nw6[(new BigNumber(16)).toNumber()] = _108_v21;
          _nw6[(new BigNumber(17)).toNumber()] = (_108_v21).Union(_108_v21);
          _nw6[(new BigNumber(18)).toNumber()] = _108_v21;
          _109_v24 = _nw6;
          let _112_v25;
          _112_v25 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("khlhpemq")).length));
          let _113_v26;
          let _nw7 = Array((new BigNumber(5)).toNumber());
          _113_v26 = _nw7;
          let _114_v27;
          _114_v27 = _dafny.Map.Empty.slice().updateUnsafe(_113_v26,p0);
          let _115_v28;
          let _nw8 = Array((new BigNumber(7)).toNumber());
          _nw8[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_112_v25, _dafny.Seq.of(p3, new BigNumber((_114_v27).length)));
          _nw8[(_dafny.ONE).toNumber()] = _112_v25;
          _nw8[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(p1, p3, new BigNumber(7));
          _nw8[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_112_v25, _dafny.Seq.update(_112_v25, _module.__default.safeIndex(p1, new BigNumber((_112_v25).length)), new BigNumber(-307)));
          _nw8[(new BigNumber(4)).toNumber()] = _112_v25;
          _nw8[(new BigNumber(5)).toNumber()] = _112_v25;
          _nw8[(new BigNumber(6)).toNumber()] = _112_v25;
          _115_v28 = _nw8;
          let _index2 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_115_v28).length));
          (_115_v28)[_index2] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(754)), ((_116_p2) => function (_117_i2) {
            return _116_p2;
          })(p2));
          let _index3 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_115_v28).length));
          let _rhs0 = _109_v24;
          let _rhs1 = ((p0) ? (_dafny.Seq.Concat(_112_v25, _112_v25)) : (_112_v25));
          let _lhs0 = _115_v28;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_115_v28).length));
          _109_v24 = _rhs0;
          _lhs0[_lhs1] = _rhs1;
          let _118_v29;
          _118_v29 = _dafny.Map.Empty.slice().updateUnsafe(_84_cf56,new BigNumber(690));
          let _119_v31;
          _119_v31 = _dafny.Seq.UnicodeFromString("adhmxyo");
          let _120_v32;
          let _nw9 = Array((new BigNumber(27)).toNumber());
          _nw9[(_dafny.ZERO).toNumber()] = _118_v29;
          _nw9[(_dafny.ONE).toNumber()] = _118_v29;
          _nw9[(new BigNumber(2)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(3)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(4)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(5)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(6)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(7)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(8)).toNumber()] = function () {
            let _coll16 = new _dafny.Map();
            for (const _compr_16 of _dafny.IntegerRange(new BigNumber(-415), new BigNumber(552))) {
              let _121_v30 = _compr_16;
              if (((new BigNumber(-415)).isLessThanOrEqualTo(_121_v30)) && ((_121_v30).isLessThan(new BigNumber(552)))) {
                _coll16.push([_module.__default.safeDivisionInt(_121_v30, _84_cf56),p3]);
              }
            }
            return _coll16;
          }();
          _nw9[(new BigNumber(9)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(10)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(11)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(12)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(13)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p3,_84_cf56);
          _nw9[(new BigNumber(15)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(16)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p3,_84_cf56);
          _nw9[(new BigNumber(18)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(19)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(20)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(21)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(22)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(23)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(24)).toNumber()] = _118_v29;
          _nw9[(new BigNumber(25)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_119_v31).length));
          _nw9[(new BigNumber(26)).toNumber()] = _118_v29;
          _120_v32 = _nw9;
          let _122_v33;
          _122_v33 = _dafny.Map.Empty.slice().updateUnsafe(p3,_module.D16.create_DC39(_120_v32));
          let _123_v34;
          _123_v34 = _dafny.MultiSet.fromElements(_122_v33);
          _123_v34 = _123_v34;
          let _124_v35;
          _124_v35 = _dafny.MultiSet.fromElements(_107_v20);
          let _125_v36;
          let _nw10 = Array((new BigNumber(10)).toNumber());
          _nw10[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_84_cf56);
          _nw10[(_dafny.ONE).toNumber()] = _84_cf56;
          _nw10[(new BigNumber(2)).toNumber()] = _84_cf56;
          _nw10[(new BigNumber(3)).toNumber()] = _84_cf56;
          _nw10[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_124_v35).cardinality()), _84_cf56);
          _nw10[(new BigNumber(5)).toNumber()] = p3;
          _nw10[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("eiavcl")).length);
          _nw10[(new BigNumber(7)).toNumber()] = _84_cf56;
          _nw10[(new BigNumber(8)).toNumber()] = p2;
          _nw10[(new BigNumber(9)).toNumber()] = p3;
          _125_v36 = _nw10;
          _125_v36 = _125_v36;
          let _126_v37;
          _126_v37 = _dafny.MultiSet.fromElements(p0);
          let _127_v38;
          let _nw11 = new _module.C1();
          _nw11.__ctor(p0, _126_v37, _112_v25, _84_cf56, _85_v0);
          _127_v38 = _nw11;
        }
        let _128_v39;
        _128_v39 = _dafny.MultiSet.fromElements(p0);
        let _129_v40;
        let _nw12 = new _module.C2();
        _nw12.__ctor(_128_v39, ((p0) ? (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(p1, globalState),p0)).update(_module.__default.fm0(p3, globalState), p0)).length)) : (new BigNumber(207))));
        _129_v40 = _nw12;
        _129_v40 = _129_v40;
      } else if (_source8.is_DC32) {
        let _130___mcc_h1 = (_source8).cf57;
        let _131___mcc_h2 = (_source8).cf58;
        let _132___mcc_h3 = (_source8).cf59;
        let _133___mcc_h4 = (_source8).cf60;
        let _134_cf60 = _133___mcc_h4;
        let _135_cf59 = _132___mcc_h3;
        let _136_cf58 = _131___mcc_h2;
        let _137_cf57 = _130___mcc_h1;
        let _138_v41;
        _138_v41 = _dafny.Seq.UnicodeFromString("jc");
        let _139_v42;
        let _init1 = ((_140_cf57, _141_p0) => function (_142_i3) {
          return (_140_cf57).isEqualTo((_module.D4.create_DC9(_140_cf57, _141_p0)).dtor_cf16);
        })(_137_cf57, p0);
        let _nw13 = Array((new BigNumber(6)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw13.length); _i0_1++) {
          _nw13[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _139_v42 = _nw13;
        let _rhs2 = _136_cf58;
        let _rhs3 = _dafny.Seq.UnicodeFromString("fkc");
        let _rhs4 = false;
        let _rhs5 = _139_v42;
        let _rhs6 = ((false) ? (p3) : (_136_cf58));
        _136_cf58 = _rhs2;
        _138_v41 = _rhs3;
        _134_cf60 = _rhs4;
        _139_v42 = _rhs5;
        _136_cf58 = _rhs6;
        (globalState).f15 = _136_cf58;
        if (!(_134_cf60) || (p0)) {
          let _143_v43;
          _143_v43 = _dafny.MultiSet.fromElements(p0);
          let _144_v44;
          let _nw14 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _144_v44 = _nw14;
          let _145_v45;
          _145_v45 = _dafny.Map.Empty.slice().updateUnsafe(_134_cf60,_144_v44);
          let _146_v46;
          let _nw15 = new _module.C6();
          _nw15.__ctor(true, p2, _143_v43, _135_cf59, p3, (((_145_v45).contains(!(!(_134_cf60)))) ? ((_145_v45).get(!(!(_134_cf60)))) : (_144_v44)), _143_v43, p2);
          _146_v46 = _nw15;
          let _147_v47;
          let _nw16 = Array((new BigNumber(13)).toNumber());
          _nw16[(_dafny.ZERO).toNumber()] = _146_v46;
          _nw16[(_dafny.ONE).toNumber()] = _146_v46;
          _nw16[(new BigNumber(2)).toNumber()] = _146_v46;
          _nw16[(new BigNumber(3)).toNumber()] = _146_v46;
          _nw16[(new BigNumber(4)).toNumber()] = _146_v46;
          _nw16[(new BigNumber(5)).toNumber()] = _146_v46;
          _nw16[(new BigNumber(6)).toNumber()] = _146_v46;
          _nw16[(new BigNumber(7)).toNumber()] = _146_v46;
          _nw16[(new BigNumber(8)).toNumber()] = _146_v46;
          _nw16[(new BigNumber(9)).toNumber()] = _146_v46;
          _nw16[(new BigNumber(10)).toNumber()] = _146_v46;
          _nw16[(new BigNumber(11)).toNumber()] = _146_v46;
          _nw16[(new BigNumber(12)).toNumber()] = _146_v46;
          _147_v47 = _nw16;
          _147_v47 = _147_v47;
          let _148_v48;
          let _nw17 = new _module.C0();
          _nw17.__ctor(_137_cf57, _138_v41);
          _148_v48 = _nw17;
          let _149_v49;
          _149_v49 = _dafny.Seq.of(_148_v48, _148_v48);
          let _150_v50;
          _150_v50 = _dafny.Seq.of(_149_v49);
          let _151_v51;
          _151_v51 = new _dafny.CodePoint('w'.codePointAt(0));
          let _152_v52;
          let _nw18 = Array((new BigNumber(10)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = (_146_v46).f32;
          _nw18[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_150_v50, _150_v50)).length);
          _nw18[(new BigNumber(2)).toNumber()] = (_148_v48).f36;
          _nw18[(new BigNumber(3)).toNumber()] = (_148_v48).f36;
          _nw18[(new BigNumber(4)).toNumber()] = _137_cf57;
          _nw18[(new BigNumber(5)).toNumber()] = new BigNumber(913);
          _nw18[(new BigNumber(6)).toNumber()] = new BigNumber((_module.__default.fm56(_module.D4.create_DC8(_151_v51, new BigNumber(470), p0), p0, _module.__default.fm57(p0, p0, globalState), globalState)).length);
          _nw18[(new BigNumber(7)).toNumber()] = (_146_v46).f32;
          _nw18[(new BigNumber(8)).toNumber()] = _136_cf58;
          _nw18[(new BigNumber(9)).toNumber()] = new BigNumber(543);
          _152_v52 = _nw18;
          let _index4 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_152_v52).length));
          (_152_v52)[_index4] = _137_cf57;
          let _153_v53;
          _153_v53 = _dafny.Map.Empty.slice().updateUnsafe(p3,_137_cf57);
          let _index5 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_152_v52).length));
          let _rhs7 = _module.__default.fm8(globalState);
          let _rhs8 = (_module.D9.create_DC20(_136_cf58, !(_134_cf60), p0, p0, _153_v53)).dtor_cf34;
          let _rhs9 = false;
          let _rhs10 = !(p0);
          let _lhs2 = _152_v52;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_152_v52).length));
          let _lhs4 = globalState;
          _138_v41 = _rhs7;
          _lhs2[_lhs3] = _rhs8;
          _lhs4.f26 = _rhs9;
          _134_cf60 = _rhs10;
          let _index6 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_139_v42).length));
          (_139_v42)[_index6] = p0;
          let _index7 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_139_v42).length));
          (_139_v42)[_index7] = p0;
          let _154_v54;
          _154_v54 = _dafny.MultiSet.fromElements(new BigNumber((_135_cf59).length), (p2).minus(p1), _module.__default.safeDivisionInt((_152_v52)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_152_v52).length))], (_148_v48).f36), new BigNumber(-492));
          let _index8 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_152_v52).length));
          let _rhs11 = (_154_v54).Union(_154_v54);
          let _rhs12 = new BigNumber(((_148_v48).f37).length);
          let _lhs5 = _152_v52;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_152_v52).length));
          _154_v54 = _rhs11;
          _lhs5[_lhs6] = _rhs12;
          (globalState).f26 = (_139_v42)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_139_v42).length))];
        } else {
          let _155_v55;
          _155_v55 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_138_v41)[_module.__default.safeIndex(new BigNumber(843), new BigNumber((_138_v41).length))]);
          let _156_v56;
          _156_v56 = _dafny.Map.Empty.slice().updateUnsafe(p2,_155_v55);
          let _157_v57;
          _157_v57 = _dafny.Seq.of((((_156_v56).contains(_136_cf58)) ? ((_156_v56).get(_136_cf58)) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(678)), function (_158_i4) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          })).length),new _dafny.CodePoint('t'.codePointAt(0))))));
          let _159_v58;
          _159_v58 = _module.D15.create_DC35(_157_v57);
          _159_v58 = ((p0) ? (_159_v58) : (_module.D15.create_DC35(_157_v57)));
          (globalState).f18 = _module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements(_134_cf60)).length), new BigNumber(-822));
          let _160_v59;
          _160_v59 = _module.D13.create_DC32(new BigNumber(14), p3, _135_cf59, _134_cf60);
          let _161_v60;
          _161_v60 = _dafny.Seq.of(_160_v59);
          _161_v60 = _dafny.Seq.Concat(_161_v60, _dafny.Seq.of(_160_v59, _160_v59));
          let _162_v61;
          _162_v61 = new _dafny.CodePoint('o'.codePointAt(0));
          let _163_v62;
          let _nw19 = Array((new BigNumber(12)).toNumber());
          _nw19[(_dafny.ZERO).toNumber()] = _162_v61;
          _nw19[(_dafny.ONE).toNumber()] = _162_v61;
          _nw19[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('d'.codePointAt(0));
          _nw19[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
          _nw19[(new BigNumber(4)).toNumber()] = _162_v61;
          _nw19[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
          _nw19[(new BigNumber(6)).toNumber()] = _162_v61;
          _nw19[(new BigNumber(7)).toNumber()] = _162_v61;
          _nw19[(new BigNumber(8)).toNumber()] = _162_v61;
          _nw19[(new BigNumber(9)).toNumber()] = _162_v61;
          _nw19[(new BigNumber(10)).toNumber()] = _162_v61;
          _nw19[(new BigNumber(11)).toNumber()] = _162_v61;
          _163_v62 = _nw19;
          let _index9 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_163_v62).length));
          (_163_v62)[_index9] = _162_v61;
          let _index10 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_163_v62).length));
          (_163_v62)[_index10] = _162_v61;
          (globalState).f21 = (_dafny.ZERO).minus(p3);
        }
        r0 = p0;
      } else if (_source8.is_DC30) {
        let _164___mcc_h5 = (_source8).cf55;
        let _165_cf55 = _164___mcc_h5;
        let _166_v63;
        _166_v63 = _dafny.Seq.of(p3);
        if (!(!(_dafny.Seq.IsProperPrefixOf(_166_v63, _dafny.Seq.Concat(_dafny.Seq.of(p3, p3, (_dafny.ZERO).minus(p3)), _166_v63))))) {
          let _167_v64;
          _167_v64 = _dafny.Seq.of(p0, _165_cf55.f35);
          let _168_v65;
          _168_v65 = _dafny.Seq.UnicodeFromString("qmttqthg");
          let _169_v66;
          let _nw20 = Array((new BigNumber(27)).toNumber());
          _nw20[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-212)), function (_170_i5) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          });
          _nw20[(_dafny.ONE).toNumber()] = _168_v65;
          _nw20[(new BigNumber(2)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(3)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(4)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(5)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(6)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(7)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("cfdjji");
          _nw20[(new BigNumber(9)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(10)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(11)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("eet");
          _nw20[(new BigNumber(13)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(14)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("tqv");
          _nw20[(new BigNumber(16)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("fwkfera");
          _nw20[(new BigNumber(18)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(19)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(20)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(21)).toNumber()] = _dafny.Seq.UnicodeFromString("y");
          _nw20[(new BigNumber(22)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(23)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(24)).toNumber()] = _168_v65;
          _nw20[(new BigNumber(25)).toNumber()] = _dafny.Seq.UnicodeFromString("ldwxsm");
          _nw20[(new BigNumber(26)).toNumber()] = _dafny.Seq.UnicodeFromString("fytgb");
          _169_v66 = _nw20;
          let _171_v67;
          _171_v67 = _dafny.MultiSet.fromElements(_165_cf55.f35, p0, p0, p0);
          let _172_v68;
          let _nw21 = new _module.C6();
          _nw21.__ctor((_165_cf55).fm3(_dafny.MultiSet.fromElements(p0), globalState), p2, _dafny.MultiSet.FromArray(_167_v64), _dafny.Seq.update(_module.__default.fm16(new BigNumber(511), globalState), _module.__default.safeIndex(p1, new BigNumber((_module.__default.fm16(new BigNumber(511), globalState)).length)), p2), (p1).minus(p2), _169_v66, (_171_v67).update(true, _module.__default.abs(new BigNumber((_171_v67).cardinality()))), new BigNumber((_168_v65).length));
          _172_v68 = _nw21;
          let _173_v69;
          let _nw22 = new _module.C4();
          _nw22.__ctor(new BigNumber((_dafny.Seq.Concat(_166_v63, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-21)), ((_174_v65) => function (_175_i6) {
            return (_dafny.ZERO).minus(new BigNumber((_174_v65).length));
          })(_168_v65)))).length), _169_v66);
          _173_v69 = _nw22;
          _173_v69 = _173_v69;
          let _176_v70;
          _176_v70 = _168_v65;
          let _177_v71;
          _177_v71 = _dafny.Set.fromElements((_165_cf55).fm11(false, _172_v68.f34, _176_v70, p3, globalState));
          let _178_v72;
          _178_v72 = _dafny.Map.Empty.slice().updateUnsafe(_165_cf55.f35,(_177_v71).IsSubsetOf(_177_v71));
          _178_v72 = (_178_v72).update((_172_v68).f33, !(_165_cf55.f35));
          let _179_v73;
          _179_v73 = new _dafny.CodePoint('n'.codePointAt(0));
          _179_v73 = _179_v73;
          let _180_v74;
          _180_v74 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_168_v65).length),(new BigNumber(926)).multipliedBy(p1));
          _180_v74 = (_180_v74).update(new BigNumber(((_180_v74).update((_dafny.ZERO).minus(p3), new BigNumber(93))).length), _172_v68.f34);
        } else {
          let _181_v75;
          _181_v75 = _dafny.MultiSet.fromElements(_165_cf55.f35);
          let _182_v76;
          let _init2 = function (_183_i7) {
            return _dafny.Seq.UnicodeFromString("mte");
          };
          let _nw23 = Array((new BigNumber(26)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw23.length); _i0_2++) {
            _nw23[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _182_v76 = _nw23;
          let _184_v77;
          let _nw24 = new _module.C6();
          _nw24.__ctor(_165_cf55.f35, p3, _181_v75, _dafny.Seq.of(p1), new BigNumber((_166_v63).length), _182_v76, _181_v75, p3);
          _184_v77 = _nw24;
          let _185_v78;
          _185_v78 = _dafny.Seq.of(_165_cf55.f35);
          let _186_v79;
          _186_v79 = _dafny.Map.Empty.slice().updateUnsafe(_184_v77,new BigNumber((_185_v78).length));
          let _187_v80;
          _187_v80 = _dafny.Map.Empty.slice().updateUnsafe(_165_cf55.f35,_186_v79);
          let _188_v81;
          _188_v81 = _dafny.MultiSet.fromElements(_186_v79, (((_187_v80).contains(p0)) ? ((_187_v80).get(p0)) : (_186_v79)), _dafny.Map.Empty.slice().updateUnsafe(_184_v77,p3));
          let _189_v82;
          let _nw25 = new _module.C1();
          _nw25.__ctor(_165_cf55.f35, _dafny.MultiSet.fromElements(p0), _166_v63, new BigNumber((_188_v81).cardinality()), _182_v76);
          _189_v82 = _nw25;
          let _190_v83;
          _190_v83 = new _dafny.CodePoint('u'.codePointAt(0));
          let _rhs13 = true;
          let _rhs14 = _module.__default.fm0((_189_v82).fm4(_dafny.Set.fromElements(_dafny.Seq.of(_190_v83)), p0, _185_v78, globalState), globalState);
          let _rhs15 = _189_v82;
          let _lhs7 = _165_cf55;
          let _lhs8 = globalState;
          _lhs7.f35 = _rhs13;
          _lhs8.f26 = _rhs14;
          _189_v82 = _rhs15;
          let _191_v84;
          let _nw26 = Array((new BigNumber(8)).toNumber()).fill(false);
          _191_v84 = _nw26;
          let _index11 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_191_v84).length));
          (_191_v84)[_index11] = (new BigNumber((function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of _dafny.IntegerRange(new BigNumber(446), new BigNumber(851))) {
              let _192_v85 = _compr_17;
              if (((new BigNumber(446)).isLessThanOrEqualTo(_192_v85)) && ((_192_v85).isLessThan(new BigNumber(851)))) {
                _coll17.push([(_192_v85).minus(new BigNumber(531)),p0]);
              }
            }
            return _coll17;
          }()).length)).isLessThanOrEqualTo(p2);
          let _193_v86;
          _193_v86 = _dafny.Map.Empty.slice().updateUnsafe(_165_cf55.f35,new BigNumber(-355));
          let _194_v87;
          _194_v87 = _dafny.Seq.UnicodeFromString("ostrgckj");
          let _195_v88;
          _195_v88 = _194_v87;
          let _index12 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_191_v84).length));
          (_191_v84)[_index12] = !((_165_cf55).fm11(true, (new BigNumber((_193_v86).length)).minus((_189_v82).f27), _195_v88, (_184_v77).f32, globalState));
          r0 = p0;
          let _196_v89;
          let _init3 = ((_197_v77, _198_v82, _199_v75) => function (_200_i8) {
            return _dafny.MultiSet.fromElements((_197_v77).f31, (_198_v82).f29, _199_v75, (_197_v77).f31);
          })(_184_v77, _189_v82, _181_v75);
          let _nw27 = Array((new BigNumber(20)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw27.length); _i0_3++) {
            _nw27[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _196_v89 = _nw27;
          let _201_v90;
          _201_v90 = _dafny.MultiSet.fromElements(_181_v75);
          let _index13 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_196_v89).length));
          (_196_v89)[_index13] = _201_v90;
          let _202_v91;
          _202_v91 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_194_v87).length),_201_v90);
          let _index14 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_196_v89).length));
          (_196_v89)[_index14] = ((((_202_v91).contains(p3)) ? ((_202_v91).get(p3)) : (_201_v90))).Intersect(_201_v90);
          (globalState).f3 = (new BigNumber((_193_v86).length)).minus(new BigNumber(904));
        }
        let _203_v92;
        let _nw28 = Array((new BigNumber(4)).toNumber());
        _203_v92 = _nw28;
        let _204_v93;
        _204_v93 = _dafny.MultiSet.fromElements(true, !(_165_cf55.f35));
        let _205_v94;
        _205_v94 = _dafny.Seq.of(_204_v93, _204_v93);
        let _206_v95;
        let _nw29 = new _module.C2();
        _nw29.__ctor((_205_v94)[_module.__default.safeIndex(p3, new BigNumber((_205_v94).length))], p3);
        _206_v95 = _nw29;
        let _index15 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_203_v92).length));
        (_203_v92)[_index15] = _206_v95;
        let _207_v96;
        _207_v96 = _dafny.Seq.UnicodeFromString("tytaticn");
        let _208_v97;
        _208_v97 = _dafny.Map.Empty.slice().updateUnsafe(_165_cf55.f35,p1);
        let _209_v98;
        _209_v98 = new _dafny.CodePoint('l'.codePointAt(0));
        let _210_v99;
        _210_v99 = _dafny.Set.fromElements(_dafny.Seq.of(_209_v98));
        let _211_v100;
        _211_v100 = _dafny.Seq.of(_module.__default.fm0(p1, globalState));
        let _212_v101;
        let _nw30 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _212_v101 = _nw30;
        let _213_v102;
        let _nw31 = new _module.C3();
        _nw31.__ctor(p1, _212_v101);
        _213_v102 = _nw31;
        let _214_v103;
        _214_v103 = _dafny.Map.Empty.slice().updateUnsafe((((_208_v97).contains(false)) ? ((_208_v97).get(false)) : (new BigNumber((_204_v93).cardinality()))),_213_v102);
        let _215_v104;
        _215_v104 = _dafny.MultiSet.fromElements((new BigNumber((_207_v96).length)).multipliedBy(new BigNumber((_208_v97).length)), p2, p1, (_165_cf55).fm4(_210_v99, false, _211_v100, globalState), new BigNumber((_214_v103).length));
        let _216_v105;
        _216_v105 = _dafny.Set.fromElements(p0, _165_cf55.f35);
        let _217_v106;
        _217_v106 = _dafny.Set.fromElements(p1);
        let _218_v107;
        _218_v107 = _dafny.Map.Empty.slice().updateUnsafe(_165_cf55.f35,p0);
        let _219_v108;
        _219_v108 = _module.D10.create_DC22(new BigNumber((_207_v96).length), _204_v93);
        let _220_v109;
        _220_v109 = _dafny.MultiSet.fromElements(_216_v105, _216_v105);
        let _221_v110;
        _221_v110 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
        let _222_v111;
        let _nw32 = Array((new BigNumber(21)).toNumber());
        _nw32[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(p1, p1);
        _nw32[(_dafny.ONE).toNumber()] = p1;
        _nw32[(new BigNumber(2)).toNumber()] = p3;
        _nw32[(new BigNumber(3)).toNumber()] = p1;
        _nw32[(new BigNumber(4)).toNumber()] = p2;
        _nw32[(new BigNumber(5)).toNumber()] = (_module.__default.fm1(globalState)).multipliedBy(new BigNumber((_216_v105).length));
        _nw32[(new BigNumber(6)).toNumber()] = p3;
        _nw32[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(p1, new BigNumber((_217_v106).length));
        _nw32[(new BigNumber(8)).toNumber()] = p2;
        _nw32[(new BigNumber(9)).toNumber()] = new BigNumber((_218_v107).length);
        _nw32[(new BigNumber(10)).toNumber()] = p1;
        _nw32[(new BigNumber(11)).toNumber()] = new BigNumber(307);
        _nw32[(new BigNumber(12)).toNumber()] = p1;
        _nw32[(new BigNumber(13)).toNumber()] = p2;
        _nw32[(new BigNumber(14)).toNumber()] = (_219_v108).dtor_cf40;
        _nw32[(new BigNumber(15)).toNumber()] = p1;
        _nw32[(new BigNumber(16)).toNumber()] = new BigNumber(((_220_v109).update(_216_v105, _module.__default.abs((((_204_v93).contains(p0)) ? ((_204_v93).get(p0)) : (p1))))).cardinality());
        _nw32[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm1(globalState));
        _nw32[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus((p2).plus(p1));
        _nw32[(new BigNumber(19)).toNumber()] = p2;
        _nw32[(new BigNumber(20)).toNumber()] = (p2).minus(new BigNumber((_221_v110).length));
        _222_v111 = _nw32;
        let _index16 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_222_v111).length));
        (_222_v111)[_index16] = _module.__default.safeModuloInt(new BigNumber((_208_v97).length), p2);
        let _223_v112;
        let _nw33 = Array((new BigNumber(18)).toNumber());
        _223_v112 = _nw33;
        let _224_v113;
        _224_v113 = _dafny.Seq.of(_223_v112);
        let _225_v114;
        _225_v114 = _dafny.Map.Empty.slice().updateUnsafe(p0,_207_v96);
        let _226_v115;
        _226_v115 = _module.D22.create_DC53(_224_v113);
        let _index17 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_203_v92).length));
        let _index18 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_222_v111).length));
        let _rhs16 = _206_v95;
        let _rhs17 = ((_215_v104).Intersect(_215_v104)).Difference(_215_v104);
        let _rhs18 = new BigNumber((((!(_165_cf55.f35) || (p0)) ? (_166_v63) : (_dafny.Seq.Concat(_166_v63, _166_v63)))).length);
        let _rhs19 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p0,_207_v96)).Merge((_225_v114).Merge(_225_v114))).length);
        let _rhs20 = (_226_v115).dtor_cf94;
        let _lhs9 = _203_v92;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_203_v92).length));
        let _lhs11 = globalState;
        let _lhs12 = _222_v111;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_222_v111).length));
        _lhs9[_lhs10] = _rhs16;
        _215_v104 = _rhs17;
        _lhs11.f3 = _rhs18;
        _lhs12[_lhs13] = _rhs19;
        _224_v113 = _rhs20;
        let _227_v116;
        let _nw34 = new _module.C6();
        _nw34.__ctor(!(_165_cf55.f35), new BigNumber((_207_v96).length), _204_v93, _module.__default.fm16(p1, globalState), (_222_v111)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_222_v111).length))], _212_v101, _204_v93, p3);
        _227_v116 = _nw34;
        let _228_v117;
        _228_v117 = _dafny.Map.Empty.slice().updateUnsafe(_227_v116,(_222_v111)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_222_v111).length))]);
        let _229_v118;
        _229_v118 = _dafny.Map.Empty.slice().updateUnsafe(_222_v111,_228_v117);
        _229_v118 = (_229_v118).update(_222_v111, (_228_v117).update(_227_v116, (_222_v111)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_222_v111).length))]));
        let _230_v119;
        _230_v119 = _module.D4.create_DC9(new BigNumber(-28), p0);
        let _231_v120;
        let _nw35 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
        _231_v120 = _nw35;
        let _232_v121;
        _232_v121 = _231_v120;
        let _source9 = ((!(((_222_v111)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_222_v111).length))]).isLessThanOrEqualTo((_230_v119).dtor_cf16))) ? (_232_v121) : (_232_v121));
        let _233___mcc_h7 = _source9;
        let _234_cf93 = _233___mcc_h7;
        let _235_v122;
        let _init4 = ((_236_p0, _237_v98, _238_v116) => function (_239_i9) {
          return _dafny.Map.Empty.slice().updateUnsafe(_236_p0,_module.D4.create_DC8(_237_v98, (_dafny.ZERO).minus(_238_v116.f34), _236_p0));
        })(p0, _209_v98, _227_v116);
        let _nw36 = Array((_dafny.ONE).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw36.length); _i0_4++) {
          _nw36[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _235_v122 = _nw36;
        let _240_v123;
        _240_v123 = _module.D4.create_DC8(_209_v98, new BigNumber((_218_v107).length), (_227_v116).f33);
        let _index19 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_235_v122).length));
        (_235_v122)[_index19] = _dafny.Map.Empty.slice().updateUnsafe(true,_240_v123);
        let _241_v124;
        _241_v124 = _dafny.Map.Empty.slice().updateUnsafe(_165_cf55.f35,_240_v123);
        let _index20 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_235_v122).length));
        (_235_v122)[_index20] = ((_241_v124).update(_165_cf55.f35, _240_v123)).update(!(p0) || (false), _240_v123);
        let _242_v125;
        let _nw37 = new _module.C0();
        _nw37.__ctor(p1, (((_225_v114).contains(!(_165_cf55.f35))) ? ((_225_v114).get(!(_165_cf55.f35))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(751)), ((_243_v98) => function (_244_i10) {
          return _243_v98;
        })(_209_v98)))));
        _242_v125 = _nw37;
        let _245_v126;
        _245_v126 = _dafny.Seq.of((_203_v92)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_203_v92).length))], _206_v95);
        let _246_v127;
        _246_v127 = _dafny.Map.Empty.slice().updateUnsafe((_222_v111)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_222_v111).length))],new BigNumber((_dafny.Seq.UnicodeFromString("kqkukel")).length));
        let _247_v128;
        let _nw38 = new _module.C5();
        _nw38.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_246_v127,p1), _204_v93, _166_v63, new BigNumber((_207_v96).length), _212_v101);
        _247_v128 = _nw38;
        let _248_v129;
        _248_v129 = _dafny.MultiSet.fromElements(_247_v128);
        let _nw39 = new _module.C1();
        _nw39.__ctor(_dafny.areEqual(_245_v126, _245_v126), _204_v93, _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-554), new BigNumber((_225_v114).length), new BigNumber((_248_v129).cardinality())), _dafny.Seq.of(p2)), new BigNumber((_166_v63).length), _212_v101);
        _165_cf55 = _nw39;
        let _249_v130;
        let _init5 = ((_250_v98, _251_v63) => function (_252_i11) {
          return _dafny.Map.Empty.slice().updateUnsafe(_250_v98,_251_v63);
        })(_209_v98, _166_v63);
        let _nw40 = Array((new BigNumber(11)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw40.length); _i0_5++) {
          _nw40[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _249_v130 = _nw40;
        let _253_v131;
        _253_v131 = _module.D2.create_DC4(p2, (_222_v111)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_222_v111).length))], (_242_v125).f36, _209_v98, _207_v96);
        let _254_v132;
        _254_v132 = _dafny.Map.Empty.slice().updateUnsafe((_253_v131).dtor_cf8,_module.__default.fm43((_227_v116).f33, p0, !(p0), globalState));
        let _index21 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_249_v130).length));
        (_249_v130)[_index21] = _254_v132;
        let _255_v133;
        _255_v133 = _dafny.Map.Empty.slice().updateUnsafe(_165_cf55,_254_v132);
        let _index22 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_249_v130).length));
        (_249_v130)[_index22] = (((_255_v133).contains(_165_cf55)) ? ((_255_v133).get(_165_cf55)) : (_254_v132));
      } else {
        let _256___mcc_h6 = (_source8).cf61;
        let _257_cf61 = _256___mcc_h6;
        (globalState).f26 = p0;
        let _258_v134;
        let _nw41 = Array((new BigNumber(11)).toNumber());
        _258_v134 = _nw41;
        let _259_v135;
        _259_v135 = _dafny.Set.fromElements(_258_v134, _258_v134, _258_v134);
        if (!(p0) || ((_259_v135).IsProperSubsetOf(_259_v135))) {
          (globalState).f26 = !(false);
          let _260_v136;
          let _nw42 = new _module.C0();
          _nw42.__ctor(p1, _module.__default.fm8(globalState));
          _260_v136 = _nw42;
          let _261_v137;
          _261_v137 = _module.D6.create_DC11(_260_v136);
          let _262_v138;
          _262_v138 = _module.D19.create_DC49(p0, p0, new BigNumber(807));
          let _263_v139;
          _263_v139 = _dafny.Map.Empty.slice().updateUnsafe(_261_v137,_262_v138);
          _263_v139 = (_263_v139).update(_261_v137, _262_v138);
          let _264_v140;
          let _nw43 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _264_v140 = _nw43;
          let _index23 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_264_v140).length));
          (_264_v140)[_index23] = new BigNumber((_dafny.Seq.Concat((_260_v136).f37, (_260_v136).f37)).length);
          let _index24 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_264_v140).length));
          (_264_v140)[_index24] = p1;
          let _265_v141;
          _265_v141 = _dafny.Map.Empty.slice().updateUnsafe(((_260_v136).f36).multipliedBy(p3),p0);
          _265_v141 = (_265_v141).update(p2, p0);
          let _266_v142;
          _266_v142 = _dafny.Seq.UnicodeFromString("dsshj");
          let _267_v143;
          _267_v143 = _dafny.Set.fromElements((_260_v136).f36, p2);
          let _268_v144;
          _268_v144 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_267_v143).Intersect(_267_v143));
          let _269_v145;
          _269_v145 = new _dafny.CodePoint('e'.codePointAt(0));
          let _rhs21 = _dafny.Seq.Concat(_dafny.Seq.update((_260_v136).f37, _module.__default.safeIndex(p2, new BigNumber(((_260_v136).f37).length)), _269_v145), (_260_v136).f37);
          let _rhs22 = _268_v144;
          _266_v142 = _rhs21;
          _268_v144 = _rhs22;
        } else {
          let _270_v146;
          _270_v146 = _dafny.Seq.of((p2).isLessThan(p1), p0, p0, p0, p0);
          let _271_v147;
          _271_v147 = _module.D4.create_DC9(new BigNumber(534), p0);
          let _rhs23 = _dafny.Seq.of(p0, p0, p0);
          let _rhs24 = (((p0) ? (_271_v147) : (_module.__default.fm49(p1, globalState)))).dtor_cf17;
          let _rhs25 = p3;
          let _lhs14 = globalState;
          let _lhs15 = globalState;
          _270_v146 = _rhs23;
          _lhs14.f26 = _rhs24;
          _lhs15.f3 = _rhs25;
          let _272_v148;
          _272_v148 = _module.D10.create_DC22(p3, _dafny.MultiSet.FromArray(_270_v146));
          let _273_v149;
          _273_v149 = _dafny.Seq.of(p3);
          let _274_v150;
          let _nw44 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _274_v150 = _nw44;
          let _275_v151;
          let _nw45 = new _module.C1();
          _nw45.__ctor(true, (_272_v148).dtor_cf41, _273_v149, p2, _274_v150);
          _275_v151 = _nw45;
          let _276_v152;
          _276_v152 = _dafny.Map.Empty.slice().updateUnsafe(p1,_275_v151);
          let _277_v153;
          _277_v153 = _dafny.Seq.UnicodeFromString("f");
          _276_v152 = (_276_v152).update(new BigNumber((_277_v153).length), _275_v151);
          let _278_v154;
          _278_v154 = _dafny.MultiSet.fromElements(_275_v151.f35, p0, _275_v151.f35, true);
          let _279_v155;
          _279_v155 = _dafny.Map.Empty.slice().updateUnsafe(_278_v154,p2);
          let _280_v156;
          _280_v156 = _dafny.Map.Empty.slice().updateUnsafe(_279_v155,!(p0));
          _280_v156 = (_280_v156).update((_dafny.Map.Empty.slice().updateUnsafe(_278_v154,new BigNumber(231))).Merge(_279_v155), p0);
          let _281_v157;
          let _nw46 = new _module.C3();
          _nw46.__ctor(new BigNumber(-596), _274_v150);
          _281_v157 = _nw46;
          let _282_v158;
          let _init6 = ((_283_p0, _284_p3) => function (_285_i12) {
            return _dafny.Map.Empty.slice().updateUnsafe(_283_p0,_284_p3);
          })(p0, p3);
          let _nw47 = Array((new BigNumber(25)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw47.length); _i0_6++) {
            _nw47[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _282_v158 = _nw47;
          let _rhs26 = _281_v157;
          let _rhs27 = _282_v158;
          _281_v157 = _rhs26;
          _282_v158 = _rhs27;
          let _286_v159;
          _286_v159 = _dafny.Set.fromElements(new BigNumber((_270_v146).length), p3, new BigNumber((_dafny.Set.fromElements((_278_v154).update(p0, _module.__default.abs(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(163)), function (_287_i13) {
            return new _dafny.CodePoint('q'.codePointAt(0));
          })).length))))).length));
          let _288_v160;
          _288_v160 = new _dafny.CodePoint('r'.codePointAt(0));
          let _289_v161;
          _289_v161 = _dafny.Map.Empty.slice().updateUnsafe(_288_v160,_286_v159);
          _286_v159 = (((_289_v161).contains(_288_v160)) ? ((_289_v161).get(_288_v160)) : (_286_v159));
        }
        let _290_v162;
        _290_v162 = _dafny.MultiSet.fromElements(p0, p0, p0, false);
        let _291_v163;
        _291_v163 = _dafny.Map.Empty.slice().updateUnsafe(_290_v162,p0);
        let _292_v164;
        _292_v164 = _dafny.Seq.of(p0, p0);
        let _293_v165;
        _293_v165 = _dafny.Set.fromElements(p3);
        let _294_v166;
        _294_v166 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
        let _295_v167;
        let _nw48 = Array((new BigNumber(29)).toNumber());
        _nw48[(_dafny.ZERO).toNumber()] = p1;
        _nw48[(_dafny.ONE).toNumber()] = p2;
        _nw48[(new BigNumber(2)).toNumber()] = p3;
        _nw48[(new BigNumber(3)).toNumber()] = p2;
        _nw48[(new BigNumber(4)).toNumber()] = p2;
        _nw48[(new BigNumber(5)).toNumber()] = p3;
        _nw48[(new BigNumber(6)).toNumber()] = p3;
        _nw48[(new BigNumber(7)).toNumber()] = new BigNumber((_291_v163).length);
        _nw48[(new BigNumber(8)).toNumber()] = p1;
        _nw48[(new BigNumber(9)).toNumber()] = new BigNumber((_292_v164).length);
        _nw48[(new BigNumber(10)).toNumber()] = p1;
        _nw48[(new BigNumber(11)).toNumber()] = new BigNumber(37);
        _nw48[(new BigNumber(12)).toNumber()] = p1;
        _nw48[(new BigNumber(13)).toNumber()] = p3;
        _nw48[(new BigNumber(14)).toNumber()] = p3;
        _nw48[(new BigNumber(15)).toNumber()] = p1;
        _nw48[(new BigNumber(16)).toNumber()] = p1;
        _nw48[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-917));
        _nw48[(new BigNumber(18)).toNumber()] = new BigNumber((_293_v165).length);
        _nw48[(new BigNumber(19)).toNumber()] = p2;
        _nw48[(new BigNumber(20)).toNumber()] = p2;
        _nw48[(new BigNumber(21)).toNumber()] = new BigNumber((_293_v165).length);
        _nw48[(new BigNumber(22)).toNumber()] = new BigNumber(-608);
        _nw48[(new BigNumber(23)).toNumber()] = new BigNumber(877);
        _nw48[(new BigNumber(24)).toNumber()] = p3;
        _nw48[(new BigNumber(25)).toNumber()] = new BigNumber((_294_v166).length);
        _nw48[(new BigNumber(26)).toNumber()] = p3;
        _nw48[(new BigNumber(27)).toNumber()] = p3;
        _nw48[(new BigNumber(28)).toNumber()] = p2;
        _295_v167 = _nw48;
        let _296_v168;
        _296_v168 = _dafny.Set.fromElements(p0, p0);
        let _297_v169;
        _297_v169 = _dafny.Map.Empty.slice().updateUnsafe(_296_v168,new BigNumber((_dafny.Set.fromElements(true)).length));
        let _298_v170;
        _298_v170 = _dafny.Map.Empty.slice().updateUnsafe(_295_v167,new BigNumber((_297_v169).length));
        let _299_v171;
        _299_v171 = _module.D23.create_DC56(_298_v170);
        (globalState).f18 = new BigNumber(((((_299_v171).dtor_cf104).Merge(_298_v170)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_295_v167,p3))).length);
        let _300_v172;
        _300_v172 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_292_v164).length),p1);
        let _301_v173;
        _301_v173 = _dafny.Map.Empty.slice().updateUnsafe(_300_v172,p3);
        let _302_v174;
        _302_v174 = _dafny.Seq.UnicodeFromString("b");
        let _303_v175;
        _303_v175 = _dafny.Seq.of(new BigNumber(590), new BigNumber((_302_v174).length), p1);
        let _304_v176;
        let _nw49 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _304_v176 = _nw49;
        let _305_v177;
        _305_v177 = _dafny.Seq.of(_304_v176);
        let _306_v178;
        let _nw50 = new _module.C3();
        _nw50.__ctor(p1, _304_v176);
        _306_v178 = _nw50;
        let _307_v179;
        _307_v179 = _dafny.Seq.of(_306_v178);
        let _308_v180;
        let _nw51 = new _module.C5();
        _nw51.__ctor((_301_v173).Merge(_301_v173), (_dafny.MultiSet.fromElements(p0, p0)).Union(_290_v162), _303_v175, p2, (_305_v177)[_module.__default.safeIndex(new BigNumber((_307_v179).length), new BigNumber((_305_v177).length))]);
        _308_v180 = _nw51;
      }
      let _309_v181;
      _309_v181 = _dafny.Seq.UnicodeFromString("ym");
      let _310_v182;
      _310_v182 = new _dafny.CodePoint('l'.codePointAt(0));
      let _311_v183;
      _311_v183 = _dafny.MultiSet.fromElements(_309_v181, _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("w"), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.UnicodeFromString("w")).length)), _310_v182), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("w"), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.UnicodeFromString("w")).length)), _310_v182)).length)), _310_v182));
      let _312_v184;
      _312_v184 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("fuf"));
      let _313_i14;
      _313_i14 = _dafny.ZERO;
      L0: {
        while ((_dafny.MultiSet.FromArray(_312_v184)).IsProperSubsetOf(_311_v183)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_313_i14)) {
              break L0;
            }
            _313_i14 = (_313_i14).plus(_dafny.ONE);
            let _314_v186;
            _314_v186 = _dafny.Set.fromElements(p3, p1);
            let _315_v187;
            _315_v187 = _dafny.Seq.of(function () {
              let _coll18 = new _dafny.Set();
              for (const _compr_18 of _dafny.IntegerRange(new BigNumber(-577), new BigNumber(565))) {
                let _316_v185 = _compr_18;
                if (((new BigNumber(-577)).isLessThanOrEqualTo(_316_v185)) && ((_316_v185).isLessThan(new BigNumber(565)))) {
                  _coll18.add((_316_v185).plus(p3));
                }
              }
              return _coll18;
            }(), _314_v186);
            let _317_v188;
            _317_v188 = _module.D11.create_DC25(_module.D11.create_DC23(_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.update(_315_v187, _module.__default.safeIndex(p1, new BigNumber((_315_v187).length)), _314_v186), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.update(_315_v187, _module.__default.safeIndex(p1, new BigNumber((_315_v187).length)), _314_v186)).length)), _314_v186))));
            let _318_v189;
            _318_v189 = _module.D4.create_DC8(_310_v182, new BigNumber(-919), p0);
            let _319_v190;
            _319_v190 = _module.D11.create_DC24(new BigNumber(731), _318_v189);
            let _pat_let_tv1 = _319_v190;
            let _pat_let_tv2 = _319_v190;
            let _320_v191;
            let _nw52 = Array((new BigNumber(18)).toNumber());
            _nw52[(_dafny.ZERO).toNumber()] = _317_v188;
            _nw52[(_dafny.ONE).toNumber()] = _317_v188;
            _nw52[(new BigNumber(2)).toNumber()] = _module.D11.create_DC25(_319_v190);
            _nw52[(new BigNumber(3)).toNumber()] = _module.__default.fm39(globalState);
            _nw52[(new BigNumber(4)).toNumber()] = _317_v188;
            _nw52[(new BigNumber(5)).toNumber()] = _317_v188;
            _nw52[(new BigNumber(6)).toNumber()] = _317_v188;
            _nw52[(new BigNumber(7)).toNumber()] = _module.D11.create_DC25(_319_v190);
            _nw52[(new BigNumber(8)).toNumber()] = _317_v188;
            _nw52[(new BigNumber(9)).toNumber()] = _317_v188;
            _nw52[(new BigNumber(10)).toNumber()] = function (_pat_let2_0) {
              return function (_321_dt__update__tmp_h1) {
                return function (_pat_let3_0) {
                  return function (_322_dt__update_hcf45_h0) {
                    return _module.D11.create_DC25(_322_dt__update_hcf45_h0);
                  }(_pat_let3_0);
                }(_pat_let_tv1);
              }(_pat_let2_0);
            }(_317_v188);
            _nw52[(new BigNumber(11)).toNumber()] = _317_v188;
            _nw52[(new BigNumber(12)).toNumber()] = _317_v188;
            _nw52[(new BigNumber(13)).toNumber()] = function (_pat_let4_0) {
              return function (_323_dt__update__tmp_h2) {
                return function (_pat_let5_0) {
                  return function (_324_dt__update_hcf45_h1) {
                    return _module.D11.create_DC25(_324_dt__update_hcf45_h1);
                  }(_pat_let5_0);
                }(_pat_let_tv2);
              }(_pat_let4_0);
            }(_317_v188);
            _nw52[(new BigNumber(14)).toNumber()] = _317_v188;
            _nw52[(new BigNumber(15)).toNumber()] = _317_v188;
            _nw52[(new BigNumber(16)).toNumber()] = _317_v188;
            _nw52[(new BigNumber(17)).toNumber()] = _module.__default.fm39(globalState);
            _320_v191 = _nw52;
            let _index25 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_320_v191).length));
            (_320_v191)[_index25] = _module.__default.fm39(globalState);
            let _pat_let_tv3 = _319_v190;
            let _index26 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_320_v191).length));
            (_320_v191)[_index26] = function (_pat_let6_0) {
              return function (_325_dt__update__tmp_h3) {
                return function (_pat_let7_0) {
                  return function (_326_dt__update_hcf45_h2) {
                    return _module.D11.create_DC25(_326_dt__update_hcf45_h2);
                  }(_pat_let7_0);
                }(_pat_let_tv3);
              }(_pat_let6_0);
            }(_module.D11.create_DC25(_319_v190));
            r0 = !(!(p0)) || (((_dafny.ZERO).minus(p2)).isLessThan(p1));
            let _327_v192;
            _327_v192 = _dafny.MultiSet.fromElements(!(false), p0);
            let _328_v193;
            _328_v193 = p0;
            let _329_v194;
            let _init7 = ((_330_v181) => function (_331_i15) {
              return _330_v181;
            })(_309_v181);
            let _nw53 = Array((new BigNumber(12)).toNumber());
            for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw53.length); _i0_7++) {
              _nw53[_i0_7] = _init7(new BigNumber(_i0_7));
            }
            _329_v194 = _nw53;
            let _332_v195;
            let _nw54 = new _module.C1();
            _nw54.__ctor((p3).isLessThanOrEqualTo(p1), _327_v192, _module.__default.fm43(p0, _328_v193, p0, globalState), _module.__default.safeModuloInt(p2, p3), _329_v194);
            _332_v195 = _nw54;
            (globalState).f18 = (_dafny.ZERO).minus((new BigNumber((_327_v192).cardinality())).minus((new BigNumber(450)).minus(p1)));
          }
        }
      }
      let _333_v196;
      let _nw55 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
      _333_v196 = _nw55;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_333_v196).length))) {
        let _334_i16 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_334_i16)) && ((_334_i16).isLessThan(new BigNumber((_333_v196).length))))) {
          (_333_v196)[(_334_i16)] = ((_dafny.Map.Empty.slice().updateUnsafe(p0,p3)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,p1))).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,p3));
        }
      }
      let _335_i17;
      _335_i17 = _dafny.ZERO;
      L1: {
        while (p0) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_335_i17)) {
              break L1;
            }
            _335_i17 = (_335_i17).plus(_dafny.ONE);
            (globalState).f24 = p2;
            r0 = ((new BigNumber((_309_v181).length)).plus(p1)).isLessThanOrEqualTo(((_dafny.ZERO).minus(p3)).multipliedBy(new BigNumber((_309_v181).length)));
            let _336_v197;
            _336_v197 = _dafny.Seq.of(p0, p0, p0, p0, p0);
            let _337_v198;
            _337_v198 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_336_v197).length),p0);
            r0 = (p0) || (!((((_337_v198).contains(p1)) ? ((_337_v198).get(p1)) : (p0))));
            let _338_v200;
            _338_v200 = _module.D9.create_DC20(new BigNumber(89), p0, p0, p0, function () {
  let _coll19 = new _dafny.Map();
  for (const _compr_19 of (_337_v198).Keys.Elements) {
    let _339_v199 = _compr_19;
    if ((_337_v198).contains(_339_v199)) {
      _coll19.push([(_339_v199).multipliedBy(new BigNumber(-886)),p3]);
    }
  }
  return _coll19;
}());
            let _340_v201;
            _340_v201 = _dafny.Map.Empty.slice().updateUnsafe((_338_v200).dtor_cf34,p3);
            let _341_v202;
            _341_v202 = _dafny.Map.Empty.slice().updateUnsafe(_340_v201,new BigNumber((_309_v181).length));
            let _342_v203;
            _342_v203 = _dafny.MultiSet.fromElements(false);
            let _343_v204;
            _343_v204 = _dafny.Seq.of(p1);
            let _344_v205;
            _344_v205 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
            let _345_v206;
            let _init8 = ((_346_v181) => function (_347_i18) {
              return _346_v181;
            })(_309_v181);
            let _nw56 = Array((new BigNumber(29)).toNumber());
            for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw56.length); _i0_8++) {
              _nw56[_i0_8] = _init8(new BigNumber(_i0_8));
            }
            _345_v206 = _nw56;
            let _348_v207;
            _348_v207 = _dafny.Map.Empty.slice().updateUnsafe((((_344_v205).contains(p0)) ? ((_344_v205).get(p0)) : (p0)),_345_v206);
            let _349_v208;
            let _nw57 = new _module.C5();
            _nw57.__ctor(_341_v202, _342_v203, _343_v204, p3, (((_348_v207).contains(p0)) ? ((_348_v207).get(p0)) : (_345_v206)));
            _349_v208 = _nw57;
          }
        }
      }
      let _350_v209;
      _350_v209 = _dafny.Set.fromElements(_310_v182);
      let _351_v210;
      _351_v210 = _dafny.MultiSet.fromElements(new BigNumber(365));
      let _352_v211;
      _352_v211 = _dafny.MultiSet.fromElements(new BigNumber((_350_v209).length), p1, p3, new BigNumber((_351_v210).cardinality()));
      _351_v210 = _351_v210;
      let _353_v212;
      let _nw58 = Array((new BigNumber(5)).toNumber()).fill(false);
      _353_v212 = _nw58;
      let _354_v213;
      _354_v213 = _dafny.MultiSet.fromElements(_353_v212, _353_v212);
      let _hi0 = new BigNumber((_354_v213).cardinality());
      for (let _355_i19 = new BigNumber((_module.__default.fm14(p0, p0, p0, new BigNumber((_dafny.Seq.UnicodeFromString("owyumxl")).length), globalState)).length); _355_i19.isLessThan(_hi0); _355_i19 = _355_i19.plus(_dafny.ONE)) {
        let _356_v214;
        let _init9 = ((_357_v181, _358_p0, _359_i19) => function (_360_i20) {
          return _module.D13.create_DC32(new BigNumber((_357_v181).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_358_p0)).length)), _dafny.Seq.of(_359_i19), _358_p0);
        })(_309_v181, p0, _355_i19);
        let _nw59 = Array((new BigNumber(2)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw59.length); _i0_9++) {
          _nw59[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _356_v214 = _nw59;
        let _361_v215;
        _361_v215 = _dafny.Seq.of(p1, p3, _355_i19, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-99)), ((_362_v182) => function (_363_i21) {
          return _362_v182;
        })(_310_v182))).length), p3);
        let _364_v216;
        _364_v216 = false;
        let _365_v217;
        _365_v217 = _module.D13.create_DC32(p3, p3, _361_v215, (_364_v216));
        let _index27 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_356_v214).length));
        (_356_v214)[_index27] = _365_v217;
        let _index28 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_356_v214).length));
        (_356_v214)[_index28] = _365_v217;
        (globalState).f26 = p0;
        _309_v181 = _dafny.Seq.Concat(_309_v181, _dafny.Seq.Create(_module.__default.abs(new BigNumber(809)), ((_366_v182) => function (_367_i22) {
          return _366_v182;
        })(_310_v182)));
        let _368_v218;
        let _nw60 = Array((new BigNumber(29)).toNumber()).fill([]);
        _368_v218 = _nw60;
        let _369_v219;
        let _nw61 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _369_v219 = _nw61;
        let _index29 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_368_v218).length));
        (_368_v218)[_index29] = _369_v219;
        let _index30 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_368_v218).length));
        (_368_v218)[_index30] = _369_v219;
      }
      r0 = p0;
      let _370_v220;
      _370_v220 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
      let _371_v221;
      _371_v221 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _372_v222;
      _372_v222 = _dafny.Map.Empty.slice().updateUnsafe(_370_v220,new BigNumber((_371_v221).length));
      let _373_v223;
      _373_v223 = _dafny.Seq.of(p1, p3);
      let _374_v224;
      let _nw62 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _374_v224 = _nw62;
      let _375_v225;
      let _nw63 = new _module.C5();
      _nw63.__ctor(_372_v222, _dafny.MultiSet.fromElements(p0, p0), _373_v223, p2, _374_v224);
      _375_v225 = _nw63;
      let _376_v226;
      _376_v226 = _module.D22.create_DC54(_310_v182, _375_v225, p0, new BigNumber((_309_v181).length));
      r1 = (_376_v226).dtor_cf98;
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _377_v0;
      let _nw64 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
      _377_v0 = _nw64;
      let _378_v1;
      _378_v1 = _dafny.Seq.UnicodeFromString("yn");
      let _379_v2;
      _379_v2 = _dafny.MultiSet.fromElements(_378_v1);
      let _380_v3;
      _380_v3 = new _dafny.CodePoint('g'.codePointAt(0));
      let _381_v4;
      _381_v4 = false;
      let _382_v5;
      _382_v5 = _dafny.Set.fromElements(_381_v4);
      let _383_v6;
      _383_v6 = new BigNumber(491);
      let _384_v7;
      _384_v7 = _dafny.Set.fromElements(new BigNumber((_382_v5).length), _383_v6, _383_v6);
      let _385_v8;
      _385_v8 = _dafny.Map.Empty.slice().updateUnsafe(_380_v3,new BigNumber((_384_v7).length));
      let _386_v9;
      _386_v9 = _dafny.Seq.of(_383_v6);
      let _387_v10;
      let _nw65 = Array((new BigNumber(3)).toNumber());
      _nw65[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(new BigNumber((_385_v8).length));
      _nw65[(_dafny.ONE).toNumber()] = _386_v9;
      _nw65[(new BigNumber(2)).toNumber()] = _386_v9;
      _387_v10 = _nw65;
      let _388_v11;
      let _nw66 = Array((new BigNumber(26)).toNumber()).fill(false);
      _388_v11 = _nw66;
      let _389_v12;
      _389_v12 = _dafny.Map.Empty.slice().updateUnsafe(_378_v1,_388_v11);
      let _390_v13;
      _390_v13 = _dafny.Set.fromElements(_378_v1);
      let _391_globalState;
      let _nw67 = new _module.GlobalState();
      _nw67.__ctor(new BigNumber(367), _dafny.Seq.UnicodeFromString("thcvtep"), _377_v0, new BigNumber(880), new BigNumber(597), new BigNumber(980), false, true, (_379_v2).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("kseko"))), false, _387_v10, new BigNumber(683), new BigNumber(646), _389_v12, false, new BigNumber(368), false, (_390_v13).Intersect(_390_v13), new BigNumber(-783), new BigNumber(756), true, new BigNumber(853), function () {
        let _coll20 = new _dafny.Map();
        for (const _compr_20 of _dafny.IntegerRange(new BigNumber(979), new BigNumber(540))) {
          let _392_v14 = _compr_20;
          if (((new BigNumber(979)).isLessThanOrEqualTo(_392_v14)) && ((_392_v14).isLessThan(new BigNumber(540)))) {
            _coll20.push([(_392_v14).minus(new BigNumber((_384_v7).length)),_383_v6]);
          }
        }
        return _coll20;
      }(), _388_v11, new BigNumber(539), new BigNumber(-892), true);
      _391_globalState = _nw67;
      let _rhs28 = _381_v4;
      let _rhs29 = _383_v6;
      let _lhs16 = _391_globalState;
      _381_v4 = _rhs28;
      _lhs16.f0 = _rhs29;
      if (_module.__default.fm0(_module.__default.safeModuloInt(_383_v6, _383_v6), _391_globalState)) {
        if (!(_381_v4) || (_381_v4)) {
          _381_v4 = _381_v4;
          _378_v1 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_378_v1, _378_v1), _module.__default.safeIndex(_383_v6, new BigNumber((_dafny.Seq.Concat(_378_v1, _378_v1)).length)), _380_v3), _dafny.Seq.Concat(_378_v1, _378_v1));
          let _393_v15;
          _393_v15 = _dafny.Map.Empty.slice().updateUnsafe(_381_v4,_383_v6);
          _393_v15 = (_393_v15).update(_381_v4, new BigNumber((_378_v1).length));
          (_391_globalState).f26 = _381_v4;
          let _394_v16;
          _394_v16 = _dafny.MultiSet.fromElements(_381_v4);
          let _395_v17;
          _395_v17 = _dafny.Seq.of((_dafny.MultiSet.fromElements(_381_v4, _381_v4)).Intersect(_394_v16), _394_v16, (_dafny.MultiSet.fromElements(_381_v4, _381_v4, _381_v4)).Union(_394_v16));
          _395_v17 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-117)), ((_396_v4) => function (_397_i0) {
            return _dafny.MultiSet.fromElements(_396_v4);
          })(_381_v4));
        } else {
          (_391_globalState).f26 = _381_v4;
          (_391_globalState).f18 = _module.__default.fm1(_391_globalState);
          let _398_v18;
          let _nw68 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
          _398_v18 = _nw68;
          let _rhs30 = _383_v6;
          let _rhs31 = _398_v18;
          let _lhs17 = _391_globalState;
          _lhs17.f24 = _rhs30;
          _398_v18 = _rhs31;
          let _399_v19;
          let _400_v20;
          let _out0;
          let _out1;
          let _outcollector0 = _module.__default.m0((_381_v4) === (_381_v4), _383_v6, _383_v6, _383_v6, _391_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _399_v19 = _out0;
          _400_v20 = _out1;
          _399_v19 = false;
        }
        let _401_v21;
        _401_v21 = _dafny.Seq.of(_377_v0, _377_v0, _377_v0);
        _401_v21 = _dafny.Seq.Concat(_401_v21, _dafny.Seq.Concat(_401_v21, _401_v21));
        if (!(_381_v4)) {
          (_391_globalState).f0 = _383_v6;
          let _402_v22;
          _402_v22 = _dafny.Map.Empty.slice().updateUnsafe(_383_v6,_381_v4);
          _402_v22 = (_402_v22).update(_383_v6, true);
          let _403_v23;
          let _nw69 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _403_v23 = _nw69;
          let _404_v24;
          let _nw70 = new _module.C4();
          _nw70.__ctor(_383_v6, _403_v23);
          _404_v24 = _nw70;
          let _405_v25;
          _405_v25 = _dafny.MultiSet.fromElements(_383_v6);
          let _406_v26;
          let _nw71 = new _module.C0();
          _nw71.__ctor((_dafny.ZERO).minus((((_405_v25).contains(new BigNumber((_386_v9).length))) ? ((_405_v25).get(new BigNumber((_386_v9).length))) : (_383_v6))), _378_v1);
          _406_v26 = _nw71;
          _383_v6 = (_406_v26).f36;
        } else {
          _388_v11 = _388_v11;
          let _407_v27;
          _407_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(205),_381_v4);
          (_391_globalState).f26 = _dafny.areEqual(_dafny.Seq.update(_dafny.Seq.Concat(_386_v9, _386_v9), _module.__default.safeIndex(_383_v6, new BigNumber((_dafny.Seq.Concat(_386_v9, _386_v9)).length)), new BigNumber((_407_v27).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-955)), ((_408_v6) => function (_409_i1) {
            return _408_v6;
          })(_383_v6)));
          (_391_globalState).f15 = (_383_v6).minus(new BigNumber(-306));
          let _410_v28;
          _410_v28 = _module.D16.create_DC40(true, (_dafny.ZERO).minus(_383_v6), new BigNumber(855));
          let _411_v29;
          _411_v29 = _dafny.Map.Empty.slice().updateUnsafe(_383_v6,_410_v28);
          (_391_globalState).f21 = (_dafny.ZERO).minus((_383_v6).minus(new BigNumber((_411_v29).length)));
          let _412_v30;
          _412_v30 = _dafny.Map.Empty.slice().updateUnsafe(_383_v6,new BigNumber(189));
          _412_v30 = _412_v30;
        }
        let _index31 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_388_v11).length));
        (_388_v11)[_index31] = _381_v4;
        let _index32 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_377_v0).length));
        (_377_v0)[_index32] = _383_v6;
        let _413_v31;
        _413_v31 = _dafny.Seq.of(_386_v9, _386_v9);
        let _414_v32;
        _414_v32 = _module.D9.create_DC17(_384_v7);
        let _415_v33;
        _415_v33 = _dafny.Seq.of(_414_v32);
        let _416_v34;
        _416_v34 = _dafny.Map.Empty.slice().updateUnsafe(_415_v33,_377_v0);
        let _index33 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_388_v11).length));
        let _index34 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_377_v0).length));
        let _rhs32 = !_dafny.areEqual(_413_v31, _dafny.Seq.of(_386_v9));
        let _rhs33 = !(_381_v4);
        let _rhs34 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_383_v6, _383_v6));
        let _rhs35 = !(_381_v4) || ((_381_v4) || (false));
        let _rhs36 = !((_416_v34).Merge(_416_v34)).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-536)), ((_417_v6) => function (_418_i2) {
          return _module.D9.create_DC17(_dafny.Set.fromElements(_417_v6));
        })(_383_v6)));
        let _lhs18 = _391_globalState;
        let _lhs19 = _388_v11;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_388_v11).length));
        let _lhs21 = _377_v0;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_377_v0).length));
        _lhs18.f26 = _rhs32;
        _lhs19[_lhs20] = _rhs33;
        _lhs21[_lhs22] = _rhs34;
        _381_v4 = _rhs35;
        _381_v4 = _rhs36;
        let _419_v35;
        _419_v35 = _dafny.MultiSet.fromElements((_388_v11)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_388_v11).length))], _381_v4, ((_388_v11)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_388_v11).length))]) === (_381_v4));
        _419_v35 = _419_v35;
      } else {
        let _420_v36;
        _420_v36 = _dafny.Seq.of((_383_v6).isLessThan(_383_v6));
        _420_v36 = _dafny.Seq.Concat(_420_v36, _420_v36);
        let _421_v37;
        let _422_v38;
        let _out2;
        let _out3;
        let _outcollector1 = _module.__default.m0(_381_v4, new BigNumber((_386_v9).length), _383_v6, _383_v6, _391_globalState);
        _out2 = _outcollector1[0];
        _out3 = _outcollector1[1];
        _421_v37 = _out2;
        _422_v38 = _out3;
        if (_381_v4) {
          let _423_v39;
          _423_v39 = _dafny.Map.Empty.slice().updateUnsafe(_383_v6,_381_v4);
          let _424_v40;
          let _425_v41;
          let _out4;
          let _out5;
          let _outcollector2 = _module.__default.m0(_421_v37, (_422_v38).minus(_383_v6), _422_v38, new BigNumber((_dafny.MultiSet.fromElements(_421_v37, _421_v37, (((_423_v39).contains(_422_v38)) ? ((_423_v39).get(_422_v38)) : (_381_v4)))).cardinality()), _391_globalState);
          _out4 = _outcollector2[0];
          _out5 = _outcollector2[1];
          _424_v40 = _out4;
          _425_v41 = _out5;
          _380_v3 = _module.__default.fm12(_module.__default.safeModuloInt(_383_v6, _383_v6), _424_v40, _391_globalState);
          let _426_v42;
          _426_v42 = _dafny.Map.Empty.slice().updateUnsafe(_380_v3,_424_v40);
          let _427_v43;
          _427_v43 = _module.D2.create_DC3(_426_v42);
          let _428_v44;
          _428_v44 = _dafny.Map.Empty.slice().updateUnsafe(_427_v43,_383_v6);
          _428_v44 = (_428_v44).update(_427_v43, (_425_v41).minus(_383_v6));
          (_391_globalState).f26 = _381_v4;
          let _429_v45;
          _429_v45 = _dafny.MultiSet.fromElements(false, _381_v4, !(_381_v4));
          let _430_v46;
          let _nw72 = Array((new BigNumber(10)).toNumber());
          _nw72[(_dafny.ZERO).toNumber()] = _378_v1;
          _nw72[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("w");
          _nw72[(new BigNumber(2)).toNumber()] = _378_v1;
          _nw72[(new BigNumber(3)).toNumber()] = _378_v1;
          _nw72[(new BigNumber(4)).toNumber()] = _378_v1;
          _nw72[(new BigNumber(5)).toNumber()] = _module.__default.fm8(_391_globalState);
          _nw72[(new BigNumber(6)).toNumber()] = _378_v1;
          _nw72[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_378_v1, _module.__default.safeIndex(_422_v38, new BigNumber((_378_v1).length)), _380_v3);
          _nw72[(new BigNumber(8)).toNumber()] = _378_v1;
          _nw72[(new BigNumber(9)).toNumber()] = _378_v1;
          _430_v46 = _nw72;
          let _431_v47;
          let _nw73 = new _module.C1();
          _nw73.__ctor(_424_v40, _429_v45, _386_v9, _383_v6, _430_v46);
          _431_v47 = _nw73;
          _431_v47 = _431_v47;
        } else {
          _381_v4 = ((_421_v37) ? (_381_v4) : (_381_v4));
          let _432_v48;
          _432_v48 = _module.D13.create_DC32((_dafny.ZERO).minus(_383_v6), _422_v38, _386_v9, _421_v37);
          let _433_v49;
          let _434_v50;
          let _out6;
          let _out7;
          let _outcollector3 = _module.__default.m0(true, (new BigNumber(((_432_v48).dtor_cf59).length)).multipliedBy(_422_v38), ((_dafny.ZERO).minus(new BigNumber(-481))).multipliedBy(_422_v38), _module.__default.safeModuloInt(_module.__default.fm1(_391_globalState), _383_v6), _391_globalState);
          _out6 = _outcollector3[0];
          _out7 = _outcollector3[1];
          _433_v49 = _out6;
          _434_v50 = _out7;
          (_391_globalState).f26 = true;
          let _435_v51;
          _435_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-243),_422_v38);
          let _436_v52;
          _436_v52 = _dafny.Map.Empty.slice().updateUnsafe(_435_v51,_434_v50);
          let _437_v53;
          _437_v53 = _dafny.MultiSet.fromElements(_433_v49);
          let _438_v54;
          let _init10 = ((_439_v1) => function (_440_i3) {
            return _439_v1;
          })(_378_v1);
          let _nw74 = Array((new BigNumber(16)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw74.length); _i0_10++) {
            _nw74[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _438_v54 = _nw74;
          let _441_v55;
          let _nw75 = new _module.C5();
          _nw75.__ctor((_436_v52).Merge(_module.__default.fm54(_381_v4, _433_v49, _391_globalState)), _437_v53, _386_v9, _434_v50, _438_v54);
          _441_v55 = _nw75;
          let _442_v56;
          _442_v56 = _dafny.Map.Empty.slice().updateUnsafe(_434_v50,_dafny.Seq.UnicodeFromString("ijuo"));
          let _443_v57;
          let _nw76 = new _module.C0();
          _nw76.__ctor(new BigNumber(448), (((_442_v56).contains(new BigNumber(-214))) ? ((_442_v56).get(new BigNumber(-214))) : (_378_v1)));
          _443_v57 = _nw76;
          _443_v57 = _443_v57;
        }
        let _444_v58;
        _444_v58 = _dafny.MultiSet.fromElements(_420_v36, _420_v36, _420_v36, _420_v36, _420_v36);
        let _445_v60;
        _445_v60 = (_dafny.Set.fromElements(_420_v36, _dafny.Seq.of(_421_v37, _381_v4))).Difference(function () {
          let _coll21 = new _dafny.Set();
          for (const _compr_21 of ((_444_v58).update(_dafny.Seq.of(true), _module.__default.abs(_422_v38))).Elements) {
            let _446_v59 = _compr_21;
            if (((_444_v58).update(_dafny.Seq.of(true), _module.__default.abs(_422_v38))).contains(_446_v59)) {
              _coll21.add(_446_v59);
            }
          }
          return _coll21;
        }());
        let _source10 = _445_v60;
        let _447___mcc_h0 = _source10;
        let _448_cf62 = _447___mcc_h0;
        (_391_globalState).f26 = !(!(!(_381_v4))) || (_381_v4);
        let _init11 = ((_449_v38) => function (_450_i4) {
          return _module.__default.safeDivisionInt(_450_i4, _449_v38);
        })(_422_v38);
        let _nw77 = Array((new BigNumber(13)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw77.length); _i0_11++) {
          _nw77[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _377_v0 = _nw77;
        _378_v1 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(675)), ((_451_v3) => function (_452_i5) {
          return _451_v3;
        })(_380_v3));
        _380_v3 = _380_v3;
        (_391_globalState).f18 = _383_v6;
      }
      let _453_v61;
      _453_v61 = _module.D11.create_DC25(_module.D11.create_DC23(_dafny.MultiSet.fromElements(_384_v7)));
      let _454_v62;
      _454_v62 = _dafny.MultiSet.fromElements(_384_v7);
      let _455_v63;
      _455_v63 = _module.D11.create_DC23(_454_v62);
      let _pat_let_tv4 = _455_v63;
      let _source11 = function (_pat_let8_0) {
        return function (_456_dt__update__tmp_h0) {
          return function (_pat_let9_0) {
            return function (_457_dt__update_hcf45_h0) {
              return _module.D11.create_DC25(_457_dt__update_hcf45_h0);
            }(_pat_let9_0);
          }(_pat_let_tv4);
        }(_pat_let8_0);
      }(_453_v61);
      if (_source11.is_DC24) {
        let _458___mcc_h1 = (_source11).cf43;
        let _459___mcc_h2 = (_source11).cf44;
        let _460_cf44 = _459___mcc_h2;
        let _461_cf43 = _458___mcc_h1;
        let _462_v64;
        _462_v64 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_378_v1).length),_461_cf43);
        let _463_v65;
        _463_v65 = _dafny.MultiSet.fromElements(_381_v4);
        let _464_v66;
        let _nw78 = new _module.C2();
        _nw78.__ctor(_463_v65, _383_v6);
        _464_v66 = _nw78;
        let _465_v67;
        _465_v67 = _dafny.Set.fromElements(_464_v66, _464_v66, _464_v66);
        _462_v64 = (_462_v64).update((_386_v9)[_module.__default.safeIndex(_383_v6, new BigNumber((_386_v9).length))], new BigNumber((_465_v67).length));
        let _466_v68;
        let _467_v69;
        let _out8;
        let _out9;
        let _outcollector4 = (_464_v66).m2(_381_v4, _381_v4, new BigNumber((_463_v65).cardinality()), (_dafny.ZERO).minus(_461_cf43), _391_globalState);
        _out8 = _outcollector4[0];
        _out9 = _outcollector4[1];
        _466_v68 = _out8;
        _467_v69 = _out9;
        let _468_v70;
        let _nw79 = Array((new BigNumber(22)).toNumber());
        _nw79[(_dafny.ZERO).toNumber()] = _378_v1;
        _nw79[(_dafny.ONE).toNumber()] = _378_v1;
        _nw79[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-447)), ((_469_v3) => function (_470_i6) {
          return _469_v3;
        })(_380_v3));
        _nw79[(new BigNumber(3)).toNumber()] = _378_v1;
        _nw79[(new BigNumber(4)).toNumber()] = _378_v1;
        _nw79[(new BigNumber(5)).toNumber()] = _378_v1;
        _nw79[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("ukgaklleg");
        _nw79[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("hl");
        _nw79[(new BigNumber(8)).toNumber()] = _378_v1;
        _nw79[(new BigNumber(9)).toNumber()] = _378_v1;
        _nw79[(new BigNumber(10)).toNumber()] = _378_v1;
        _nw79[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("wkbli");
        _nw79[(new BigNumber(12)).toNumber()] = _378_v1;
        _nw79[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("usfy");
        _nw79[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("yakla");
        _nw79[(new BigNumber(15)).toNumber()] = _378_v1;
        _nw79[(new BigNumber(16)).toNumber()] = _378_v1;
        _nw79[(new BigNumber(17)).toNumber()] = _378_v1;
        _nw79[(new BigNumber(18)).toNumber()] = _378_v1;
        _nw79[(new BigNumber(19)).toNumber()] = _378_v1;
        _nw79[(new BigNumber(20)).toNumber()] = _378_v1;
        _nw79[(new BigNumber(21)).toNumber()] = _378_v1;
        _468_v70 = _nw79;
        let _471_v71;
        _471_v71 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_391_globalState),_468_v70);
        let _472_v72;
        let _nw80 = new _module.C4();
        _nw80.__ctor(_module.__default.fm1(_391_globalState), (((_471_v71).contains(_467_v69)) ? ((_471_v71).get(_467_v69)) : (_468_v70)));
        _472_v72 = _nw80;
        let _rhs37 = _381_v4;
        let _rhs38 = _381_v4;
        let _rhs39 = _467_v69;
        let _rhs40 = _381_v4;
        let _lhs23 = _391_globalState;
        let _lhs24 = _391_globalState;
        let _lhs25 = _391_globalState;
        _381_v4 = _rhs37;
        _lhs23.f26 = _rhs38;
        _lhs24.f15 = _rhs39;
        _lhs25.f26 = _rhs40;
      } else if (_source11.is_DC23) {
        let _473___mcc_h3 = (_source11).cf42;
        let _474_cf42 = _473___mcc_h3;
        let _475_v73;
        let _476_v74;
        let _out10;
        let _out11;
        let _outcollector5 = _module.__default.m0((_383_v6).isLessThanOrEqualTo(_383_v6), (new BigNumber(84)).minus(_383_v6), _383_v6, ((_381_v4) ? (_383_v6) : (new BigNumber((_378_v1).length))), _391_globalState);
        _out10 = _outcollector5[0];
        _out11 = _outcollector5[1];
        _475_v73 = _out10;
        _476_v74 = _out11;
        if (_381_v4) {
          let _477_v75;
          let _478_v76;
          let _out12;
          let _out13;
          let _outcollector6 = _module.__default.m0(_381_v4, _383_v6, _476_v74, _383_v6, _391_globalState);
          _out12 = _outcollector6[0];
          _out13 = _outcollector6[1];
          _477_v75 = _out12;
          _478_v76 = _out13;
          let _479_v77;
          _479_v77 = _dafny.MultiSet.fromElements(_381_v4, _477_v75);
          let _480_v78;
          _480_v78 = _module.D13.create_DC32(_383_v6, (_dafny.ZERO).minus((_386_v9)[_module.__default.safeIndex((_dafny.ZERO).minus(_476_v74), new BigNumber((_386_v9).length))]), _386_v9, _381_v4);
          let _481_v79;
          _481_v79 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_378_v1).length),_477_v75);
          let _482_v80;
          _482_v80 = _dafny.Map.Empty.slice().updateUnsafe(_477_v75,true);
          let _483_v81;
          _483_v81 = _dafny.Map.Empty.slice().updateUnsafe((((_481_v79).contains(new BigNumber((_482_v80).length))) ? ((_481_v79).get(new BigNumber((_482_v80).length))) : (_381_v4)),_478_v76);
          let _484_v82;
          let _nw81 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _484_v82 = _nw81;
          let _485_v83;
          _485_v83 = _module.D17.create_DC43(_484_v82);
          let _486_v84;
          let _nw82 = new _module.C6();
          _nw82.__ctor(_475_v73, _module.__default.fm1(_391_globalState), _479_v77, (_480_v78).dtor_cf59, (_476_v74).minus(new BigNumber((_483_v81).length)), (_485_v83).dtor_cf82, _479_v77, _478_v76);
          _486_v84 = _nw82;
          _486_v84 = ((_477_v75) ? (_486_v84) : (_486_v84));
          let _487_v85;
          let _nw83 = Array((_dafny.ONE).toNumber());
          _nw83[(_dafny.ZERO).toNumber()] = _481_v79;
          _487_v85 = _nw83;
          let _488_v86;
          _488_v86 = _487_v85;
          let _489_v87;
          _489_v87 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_383_v6),(_487_v85));
          _489_v87 = (_489_v87).update(_478_v76, _487_v85);
          let _490_v88;
          let _nw84 = new _module.C1();
          _nw84.__ctor(true, (_479_v77).update(_475_v73, _module.__default.abs(new BigNumber((_378_v1).length))), _386_v9, new BigNumber(805), _484_v82);
          _490_v88 = _nw84;
          let _491_v89;
          _491_v89 = _dafny.MultiSet.fromElements(_490_v88);
          let _492_v90;
          let _nw85 = new _module.C6();
          _nw85.__ctor(((_486_v84).f32).isEqualTo(new BigNumber((_491_v89).cardinality())), _module.__default.safeModuloInt((_486_v84).f32, new BigNumber((_378_v1).length)), (_486_v84).f31, _386_v9, _module.__default.safeDivisionInt((_dafny.ZERO).minus(_383_v6), new BigNumber(-584)), _484_v82, (_486_v84).f31, _383_v6);
          _492_v90 = _nw85;
          let _493_v91;
          _493_v91 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(4)), ((_494_v3) => function (_495_i7) {
            return _494_v3;
          })(_380_v3)),(((_492_v90).f33) ? (_492_v90) : (_492_v90)));
          _492_v90 = (((_493_v91).contains(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bcvfglbs"), _378_v1))) ? ((_493_v91).get(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bcvfglbs"), _378_v1))) : (_492_v90));
          let _496_v92;
          _496_v92 = _378_v1;
          let _497_v93;
          let _498_v94;
          let _499_v95;
          let _out14;
          let _out15;
          let _out16;
          let _outcollector7 = (_490_v88).m5((_490_v88).fm11(!((_492_v90).f33), new BigNumber((_378_v1).length), _496_v92, new BigNumber((_378_v1).length), _391_globalState), _module.__default.fm0(_383_v6, _391_globalState), _391_globalState);
          _out14 = _outcollector7[0];
          _out15 = _outcollector7[1];
          _out16 = _outcollector7[2];
          _497_v93 = _out14;
          _498_v94 = _out15;
          _499_v95 = _out16;
        } else {
          (_391_globalState).f26 = (_384_v7).IsDisjointFrom(_384_v7);
          let _500_v96;
          _500_v96 = _module.D10.create_DC21(_386_v9);
          let _501_v97;
          _501_v97 = _dafny.Map.Empty.slice().updateUnsafe(((_381_v4) ? (_381_v4) : (_381_v4)),_500_v96);
          _501_v97 = (_501_v97).update(true, _500_v96);
          let _502_v98;
          _502_v98 = _378_v1;
          let _503_v99;
          _503_v99 = _dafny.Seq.of(_module.__default.fm14(_381_v4, false, _475_v73, _383_v6, _391_globalState));
          let _504_v100;
          _504_v100 = _dafny.Map.Empty.slice().updateUnsafe(_476_v74,true);
          let _505_v101;
          let _nw86 = Array((new BigNumber(27)).toNumber());
          _nw86[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("fgvfq");
          _nw86[(_dafny.ONE).toNumber()] = _378_v1;
          _nw86[(new BigNumber(2)).toNumber()] = _378_v1;
          _nw86[(new BigNumber(3)).toNumber()] = _378_v1;
          _nw86[(new BigNumber(4)).toNumber()] = _378_v1;
          _nw86[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(433)), ((_506_v3) => function (_507_i8) {
            return _506_v3;
          })(_380_v3)), _module.__default.safeIndex(_476_v74, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(433)), ((_508_v3) => function (_509_i8) {
            return _508_v3;
          })(_380_v3))).length)), _380_v3);
          _nw86[(new BigNumber(6)).toNumber()] = _dafny.Seq.update((_502_v98), _module.__default.safeIndex(_module.__default.fm1(_391_globalState), new BigNumber(((_502_v98)).length)), _380_v3);
          _nw86[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("a");
          _nw86[(new BigNumber(8)).toNumber()] = _378_v1;
          _nw86[(new BigNumber(9)).toNumber()] = _378_v1;
          _nw86[(new BigNumber(10)).toNumber()] = _module.__default.fm31(_475_v73, _381_v4, _391_globalState);
          _nw86[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(104)), ((_510_v3) => function (_511_i9) {
            return _510_v3;
          })(_380_v3));
          _nw86[(new BigNumber(12)).toNumber()] = (_503_v99)[_module.__default.safeIndex(new BigNumber((_504_v100).length), new BigNumber((_503_v99).length))];
          _nw86[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("qgwx");
          _nw86[(new BigNumber(14)).toNumber()] = _378_v1;
          _nw86[(new BigNumber(15)).toNumber()] = _378_v1;
          _nw86[(new BigNumber(16)).toNumber()] = _378_v1;
          _nw86[(new BigNumber(17)).toNumber()] = _378_v1;
          _nw86[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("kx");
          _nw86[(new BigNumber(19)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(486)), ((_512_v3) => function (_513_i10) {
            return _512_v3;
          })(_380_v3));
          _nw86[(new BigNumber(20)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("olpvqovqh"), _module.__default.safeIndex(_476_v74, new BigNumber((_dafny.Seq.UnicodeFromString("olpvqovqh")).length)), _380_v3);
          _nw86[(new BigNumber(21)).toNumber()] = _378_v1;
          _nw86[(new BigNumber(22)).toNumber()] = _378_v1;
          _nw86[(new BigNumber(23)).toNumber()] = _378_v1;
          _nw86[(new BigNumber(24)).toNumber()] = _378_v1;
          _nw86[(new BigNumber(25)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(590)), ((_514_v3) => function (_515_i11) {
            return _514_v3;
          })(_380_v3));
          _nw86[(new BigNumber(26)).toNumber()] = _dafny.Seq.UnicodeFromString("sfnpbmx");
          _505_v101 = _nw86;
          let _516_v102;
          let _nw87 = new _module.C3();
          _nw87.__ctor(_476_v74, _505_v101);
          _516_v102 = _nw87;
          let _517_v103;
          _517_v103 = _module.D13.create_DC32(_383_v6, _476_v74, _386_v9, _475_v73);
          let _518_v104;
          _518_v104 = _dafny.Map.Empty.slice().updateUnsafe((_517_v103).dtor_cf59,_516_v102);
          _516_v102 = (((_518_v104).contains(_386_v9)) ? ((_518_v104).get(_386_v9)) : (_516_v102));
          let _519_v105;
          _519_v105 = _dafny.MultiSet.fromElements(new BigNumber(201), _476_v74);
          let _520_v106;
          let _nw88 = new _module.C3();
          _nw88.__ctor(_383_v6, _505_v101);
          _520_v106 = _nw88;
          let _521_v107;
          _521_v107 = _dafny.Seq.of(_520_v106, _520_v106);
          let _522_v108;
          _522_v108 = _module.D9.create_DC18(true);
          let _523_v109;
          _523_v109 = _dafny.Seq.of((_522_v108).dtor_cf32);
          let _524_v110;
          _524_v110 = _module.D16.create_DC40(_475_v73, (_dafny.ZERO).minus((((_519_v105).contains(new BigNumber((_dafny.Seq.update(_378_v1, _module.__default.safeIndex(new BigNumber((_521_v107).length), new BigNumber((_378_v1).length)), _380_v3)).length))) ? ((_519_v105).get(new BigNumber((_dafny.Seq.update(_378_v1, _module.__default.safeIndex(new BigNumber((_521_v107).length), new BigNumber((_378_v1).length)), _380_v3)).length))) : (_383_v6))), new BigNumber((_523_v109).length));
          let _525_v111;
          _525_v111 = _dafny.MultiSet.fromElements(_381_v4, (_524_v110).dtor_cf74);
          let _526_v112;
          let _nw89 = new _module.C6();
          _nw89.__ctor(_475_v73, new BigNumber(549), _525_v111, _386_v9, _476_v74, _505_v101, _525_v111, (_520_v106).f27);
          _526_v112 = _nw89;
          let _527_v113;
          _527_v113 = _dafny.Map.Empty.slice().updateUnsafe(_380_v3,_526_v112);
          _527_v113 = (_527_v113).update(_380_v3, _526_v112);
          _385_v8 = (_385_v8).update(_380_v3, _526_v112.f34);
        }
        let _528_v114;
        let _529_v115;
        let _out17;
        let _out18;
        let _outcollector8 = _module.__default.m0(_475_v73, _476_v74, _476_v74, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(391)), ((_530_v1) => function (_531_i12) {
          return new BigNumber((_530_v1).length);
        })(_378_v1))).length), _391_globalState);
        _out17 = _outcollector8[0];
        _out18 = _outcollector8[1];
        _528_v114 = _out17;
        _529_v115 = _out18;
        let _index35 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_388_v11).length));
        (_388_v11)[_index35] = _475_v73;
        let _532_v116;
        _532_v116 = _dafny.MultiSet.fromElements(_529_v115, _529_v115);
        let _index36 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_388_v11).length));
        let _rhs41 = _528_v114;
        let _rhs42 = _475_v73;
        let _rhs43 = !(!(!(_528_v114) || (_528_v114)) || (((_module.__default.fm0(new BigNumber((_532_v116).cardinality()), _391_globalState)) ? (_528_v114) : (!(_528_v114)))));
        let _lhs26 = _388_v11;
        let _lhs27 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_388_v11).length));
        let _lhs28 = _391_globalState;
        _lhs26[_lhs27] = _rhs41;
        _lhs28.f26 = _rhs42;
        _528_v114 = _rhs43;
      } else {
        let _533___mcc_h4 = (_source11).cf45;
        let _534_cf45 = _533___mcc_h4;
        if (_381_v4) {
          let _535_v117;
          _535_v117 = _dafny.Map.Empty.slice().updateUnsafe(_383_v6,_dafny.Seq.UnicodeFromString("jrjwq"));
          let _536_v118;
          _536_v118 = _dafny.Map.Empty.slice().updateUnsafe(_381_v4,_381_v4);
          let _537_v119;
          let _nw90 = Array((new BigNumber(27)).toNumber());
          _nw90[(_dafny.ZERO).toNumber()] = _378_v1;
          _nw90[(_dafny.ONE).toNumber()] = _378_v1;
          _nw90[(new BigNumber(2)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(3)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(4)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(5)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(6)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(7)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(8)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(9)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(10)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(11)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(12)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(13)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(14)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(15)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(16)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(43)), ((_538_v3) => function (_539_i13) {
            return _538_v3;
          })(_380_v3));
          _nw90[(new BigNumber(18)).toNumber()] = _module.__default.fm37(_383_v6, _383_v6, _391_globalState);
          _nw90[(new BigNumber(19)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(20)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(21)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(22)).toNumber()] = _dafny.Seq.UnicodeFromString("by");
          _nw90[(new BigNumber(23)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(31)), ((_540_v3) => function (_541_i14) {
            return _540_v3;
          })(_380_v3));
          _nw90[(new BigNumber(24)).toNumber()] = _378_v1;
          _nw90[(new BigNumber(25)).toNumber()] = (((_535_v117).contains(new BigNumber((_536_v118).length))) ? ((_535_v117).get(new BigNumber((_536_v118).length))) : (_378_v1));
          _nw90[(new BigNumber(26)).toNumber()] = _378_v1;
          _537_v119 = _nw90;
          let _542_v120;
          let _nw91 = new _module.C4();
          _nw91.__ctor(new BigNumber((_386_v9).length), _537_v119);
          _542_v120 = _nw91;
          let _543_v121;
          _543_v121 = _dafny.MultiSet.fromElements(_542_v120);
          let _544_v122;
          _544_v122 = _dafny.Map.Empty.slice().updateUnsafe(_383_v6,_386_v9);
          let _545_v123;
          _545_v123 = _dafny.Seq.of(_381_v4);
          (_391_globalState).f21 = ((new BigNumber((_543_v121).cardinality())).multipliedBy(new BigNumber(((((_544_v122).contains(new BigNumber((_545_v123).length))) ? ((_544_v122).get(new BigNumber((_545_v123).length))) : (_386_v9))).length))).multipliedBy((_383_v6).minus(_383_v6));
          let _546_v124;
          _546_v124 = _dafny.MultiSet.fromElements(_388_v11);
          let _rhs44 = (_546_v124).Union(_546_v124);
          let _rhs45 = (_383_v6).isLessThanOrEqualTo((new BigNumber((_378_v1).length)).plus((_dafny.ZERO).minus(_383_v6)));
          let _lhs29 = _391_globalState;
          _546_v124 = _rhs44;
          _lhs29.f26 = _rhs45;
          _378_v1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_378_v1, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(8)), ((_547_v3) => function (_548_i15) {
            return _547_v3;
          })(_380_v3)), _module.__default.safeIndex(_383_v6, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(8)), ((_549_v3) => function (_550_i15) {
            return _549_v3;
          })(_380_v3))).length)), _380_v3)), _378_v1);
          (_391_globalState).f26 = (_381_v4) === (_381_v4);
          let _551_v125;
          let _nw92 = new _module.C0();
          _nw92.__ctor(new BigNumber(745), _378_v1);
          _551_v125 = _nw92;
          let _552_v126;
          let _out19;
          _out19 = (_542_v120).m11(_551_v125, (_381_v4) && (_381_v4), _391_globalState);
          _552_v126 = _out19;
        } else {
          (_391_globalState).f26 = _381_v4;
          (_391_globalState).f26 = _381_v4;
          _381_v4 = (_379_v2).IsDisjointFrom((_379_v2).Union(_379_v2));
          let _553_v127;
          let _nw93 = Array((new BigNumber(24)).toNumber());
          _nw93[(_dafny.ZERO).toNumber()] = _378_v1;
          _nw93[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(232)), ((_554_v3) => function (_555_i16) {
            return _554_v3;
          })(_380_v3));
          _nw93[(new BigNumber(2)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(3)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(4)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(195)), function (_556_i17) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          });
          _nw93[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("wfnbai");
          _nw93[(new BigNumber(7)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(8)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(9)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(10)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("apfximgjv");
          _nw93[(new BigNumber(12)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(13)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(14)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(983)), ((_557_v3) => function (_558_i18) {
            return _557_v3;
          })(_380_v3));
          _nw93[(new BigNumber(16)).toNumber()] = _dafny.Seq.UnicodeFromString("gadjbunu");
          _nw93[(new BigNumber(17)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(18)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(684)), ((_559_v3) => function (_560_i19) {
            return _559_v3;
          })(_380_v3));
          _nw93[(new BigNumber(19)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(20)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(21)).toNumber()] = _dafny.Seq.UnicodeFromString("gwlwggmnv");
          _nw93[(new BigNumber(22)).toNumber()] = _378_v1;
          _nw93[(new BigNumber(23)).toNumber()] = _dafny.Seq.UnicodeFromString("ukpupoan");
          _553_v127 = _nw93;
          let _561_v128;
          _561_v128 = _dafny.MultiSet.fromElements(_381_v4);
          let _562_v129;
          let _nw94 = new _module.C6();
          _nw94.__ctor(_381_v4, _383_v6, _dafny.MultiSet.fromElements(_381_v4, _381_v4), _386_v9, new BigNumber((_384_v7).length), _553_v127, _561_v128, _383_v6);
          _562_v129 = _nw94;
          let _563_v130;
          _563_v130 = _dafny.Set.fromElements(_562_v129);
          (_391_globalState).f21 = (((_563_v130).IsDisjointFrom(_563_v130)) ? ((_383_v6).plus(_383_v6)) : (new BigNumber((_378_v1).length)));
          let _564_v131;
          _564_v131 = _dafny.Map.Empty.slice().updateUnsafe(_377_v0,(_562_v129).f32);
          (_391_globalState).f15 = (((_564_v131).contains(_377_v0)) ? ((_564_v131).get(_377_v0)) : ((_dafny.ZERO).minus(_module.__default.fm1(_391_globalState))));
        }
        let _index37 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length));
        (_388_v11)[_index37] = (_381_v4) || (_module.__default.fm0(_383_v6, _391_globalState));
        let _index38 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length));
        (_388_v11)[_index38] = !(_381_v4) || (_381_v4);
        let _565_v132;
        _565_v132 = _dafny.MultiSet.fromElements(_381_v4, (_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))]);
        let _566_v133;
        let _nw95 = new _module.C2();
        _nw95.__ctor((_565_v132).update(_module.__default.fm0((_dafny.ZERO).minus(_383_v6), _391_globalState), _module.__default.abs(new BigNumber(869))), _383_v6);
        _566_v133 = _nw95;
        let _567_v134;
        _567_v134 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-78),_566_v133);
        let _568_v136;
        _568_v136 = _dafny.Seq.of(function () {
          let _coll22 = new _dafny.Map();
          for (const _compr_22 of _dafny.IntegerRange(new BigNumber(987), new BigNumber(775))) {
            let _569_v135 = _compr_22;
            if (((new BigNumber(987)).isLessThanOrEqualTo(_569_v135)) && ((_569_v135).isLessThan(new BigNumber(775)))) {
              _coll22.push([(_569_v135).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))],_383_v6)).length)),(_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))]]);
            }
          }
          return _coll22;
        }());
        let _570_v137;
        _570_v137 = _dafny.Seq.of(_566_v133, _566_v133);
        _566_v133 = (((_567_v134).contains(new BigNumber((_dafny.Seq.Concat(_568_v136, _568_v136)).length))) ? ((_567_v134).get(new BigNumber((_dafny.Seq.Concat(_568_v136, _568_v136)).length))) : ((_570_v137)[_module.__default.safeIndex(new BigNumber((_module.__default.fm48((_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))], _381_v4, _383_v6, _391_globalState)).length), new BigNumber((_570_v137).length))]));
        let _571_v138;
        let _nw96 = Array((new BigNumber(3)).toNumber());
        _nw96[(_dafny.ZERO).toNumber()] = _378_v1;
        _nw96[(_dafny.ONE).toNumber()] = _module.__default.fm37(_383_v6, _383_v6, _391_globalState);
        _nw96[(new BigNumber(2)).toNumber()] = _378_v1;
        _571_v138 = _nw96;
        let _572_v139;
        _572_v139 = _module.D17.create_DC43(_571_v138);
        let _source12 = _572_v139;
        if (_source12.is_DC44) {
          let _573___mcc_h5 = (_source12).cf83;
          let _574_cf83 = _573___mcc_h5;
          let _575_v140;
          _575_v140 = _dafny.Seq.of(_381_v4);
          let _576_v141;
          let _nw97 = Array((new BigNumber(23)).toNumber());
          _nw97[(_dafny.ZERO).toNumber()] = (_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))];
          _nw97[(_dafny.ONE).toNumber()] = (_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))];
          _nw97[(new BigNumber(2)).toNumber()] = false;
          _nw97[(new BigNumber(3)).toNumber()] = _381_v4;
          _nw97[(new BigNumber(4)).toNumber()] = _381_v4;
          _nw97[(new BigNumber(5)).toNumber()] = (_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))];
          _nw97[(new BigNumber(6)).toNumber()] = (_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))];
          _nw97[(new BigNumber(7)).toNumber()] = _381_v4;
          _nw97[(new BigNumber(8)).toNumber()] = ((false) ? ((_566_v133).fm5(_391_globalState)) : (_381_v4));
          _nw97[(new BigNumber(9)).toNumber()] = true;
          _nw97[(new BigNumber(10)).toNumber()] = (_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))];
          _nw97[(new BigNumber(11)).toNumber()] = _dafny.Seq.IsPrefixOf(_575_v140, _575_v140);
          _nw97[(new BigNumber(12)).toNumber()] = !(_381_v4) || ((_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))]);
          _nw97[(new BigNumber(13)).toNumber()] = _381_v4;
          _nw97[(new BigNumber(14)).toNumber()] = (_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))];
          _nw97[(new BigNumber(15)).toNumber()] = (_module.__default.fm1(_391_globalState)).isEqualTo(_383_v6);
          _nw97[(new BigNumber(16)).toNumber()] = (_566_v133).fm21(_391_globalState);
          _nw97[(new BigNumber(17)).toNumber()] = true;
          _nw97[(new BigNumber(18)).toNumber()] = false;
          _nw97[(new BigNumber(19)).toNumber()] = (_381_v4) && ((_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))]);
          _nw97[(new BigNumber(20)).toNumber()] = _381_v4;
          _nw97[(new BigNumber(21)).toNumber()] = _381_v4;
          _nw97[(new BigNumber(22)).toNumber()] = (_565_v132).IsProperSubsetOf(_dafny.MultiSet.FromArray(_575_v140));
          _576_v141 = _nw97;
          let _577_v142;
          _577_v142 = _dafny.Set.fromElements(_576_v141, _576_v141);
          let _index39 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length));
          let _rhs46 = _576_v141;
          let _rhs47 = _378_v1;
          let _rhs48 = ((_577_v142).Difference(_577_v142)).IsProperSubsetOf(_dafny.Set.fromElements(_576_v141));
          let _lhs30 = _388_v11;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length));
          _576_v141 = _rhs46;
          _378_v1 = _rhs47;
          _lhs30[_lhs31] = _rhs48;
          let _index40 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length));
          (_388_v11)[_index40] = !((_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))]);
          let _578_v143;
          _578_v143 = _dafny.Map.Empty.slice().updateUnsafe(_381_v4,_377_v0);
          let _579_v144;
          _579_v144 = _module.D15.create_DC37((_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))], _574_cf83, _377_v0, (_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))]);
          let _580_v145;
          let _nw98 = Array((new BigNumber(18)).toNumber());
          _nw98[(_dafny.ZERO).toNumber()] = _377_v0;
          _nw98[(_dafny.ONE).toNumber()] = _377_v0;
          _nw98[(new BigNumber(2)).toNumber()] = _377_v0;
          _nw98[(new BigNumber(3)).toNumber()] = (((_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))]) ? (_377_v0) : ((((_578_v143).contains((_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))])) ? ((_578_v143).get((_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))])) : (_377_v0))));
          _nw98[(new BigNumber(4)).toNumber()] = _377_v0;
          _nw98[(new BigNumber(5)).toNumber()] = (_579_v144).dtor_cf68;
          _nw98[(new BigNumber(6)).toNumber()] = _377_v0;
          _nw98[(new BigNumber(7)).toNumber()] = _377_v0;
          _nw98[(new BigNumber(8)).toNumber()] = _377_v0;
          _nw98[(new BigNumber(9)).toNumber()] = _377_v0;
          _nw98[(new BigNumber(10)).toNumber()] = _377_v0;
          _nw98[(new BigNumber(11)).toNumber()] = _377_v0;
          _nw98[(new BigNumber(12)).toNumber()] = _377_v0;
          _nw98[(new BigNumber(13)).toNumber()] = _377_v0;
          _nw98[(new BigNumber(14)).toNumber()] = _377_v0;
          _nw98[(new BigNumber(15)).toNumber()] = _377_v0;
          _nw98[(new BigNumber(16)).toNumber()] = _377_v0;
          _nw98[(new BigNumber(17)).toNumber()] = _377_v0;
          _580_v145 = _nw98;
          let _index41 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_580_v145).length));
          (_580_v145)[_index41] = _377_v0;
          let _581_v146;
          _581_v146 = _module.D12.create_DC28(_574_cf83, _383_v6, _381_v4);
          let _index42 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_580_v145).length));
          let _nw99 = Array((new BigNumber(19)).toNumber());
          _nw99[(_dafny.ZERO).toNumber()] = new BigNumber(-444);
          _nw99[(_dafny.ONE).toNumber()] = _574_cf83;
          _nw99[(new BigNumber(2)).toNumber()] = (_574_cf83).multipliedBy(_574_cf83);
          _nw99[(new BigNumber(3)).toNumber()] = new BigNumber((_575_v140).length);
          _nw99[(new BigNumber(4)).toNumber()] = new BigNumber(963);
          _nw99[(new BigNumber(5)).toNumber()] = _383_v6;
          _nw99[(new BigNumber(6)).toNumber()] = _383_v6;
          _nw99[(new BigNumber(7)).toNumber()] = (_383_v6).multipliedBy(_383_v6);
          _nw99[(new BigNumber(8)).toNumber()] = (_383_v6).minus(_383_v6);
          _nw99[(new BigNumber(9)).toNumber()] = new BigNumber(909);
          _nw99[(new BigNumber(10)).toNumber()] = _383_v6;
          _nw99[(new BigNumber(11)).toNumber()] = _383_v6;
          _nw99[(new BigNumber(12)).toNumber()] = ((_581_v146).dtor_cf52).multipliedBy(_383_v6);
          _nw99[(new BigNumber(13)).toNumber()] = (_574_cf83).multipliedBy(new BigNumber((_module.__default.fm9((_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))], _391_globalState)).length));
          _nw99[(new BigNumber(14)).toNumber()] = new BigNumber((_378_v1).length);
          _nw99[(new BigNumber(15)).toNumber()] = _383_v6;
          _nw99[(new BigNumber(16)).toNumber()] = _574_cf83;
          _nw99[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm1(_391_globalState));
          _nw99[(new BigNumber(18)).toNumber()] = _383_v6;
          (_580_v145)[_index42] = _nw99;
          (_391_globalState).f15 = (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_381_v4), _dafny.Seq.of(true))).length)).multipliedBy((_383_v6).plus(_574_cf83)));
        } else if (_source12.is_DC43) {
          let _582___mcc_h6 = (_source12).cf82;
          let _583_cf82 = _582___mcc_h6;
          _381_v4 = _381_v4;
          let _584_v147;
          _584_v147 = _dafny.Seq.of(_386_v9);
          let _585_v148;
          _585_v148 = _dafny.Set.fromElements((_584_v147)[_module.__default.safeIndex(_383_v6, new BigNumber((_584_v147).length))], _module.__default.fm43(_381_v4, _381_v4, _381_v4, _391_globalState));
          (_391_globalState).f15 = new BigNumber((_585_v148).length);
          let _586_v149;
          _586_v149 = _dafny.Map.Empty.slice().updateUnsafe(false,(_386_v9)[_module.__default.safeIndex(new BigNumber(814), new BigNumber((_386_v9).length))]);
          let _587_v150;
          _587_v150 = _dafny.Map.Empty.slice().updateUnsafe(_586_v149,(new BigNumber(423)).isLessThanOrEqualTo(_383_v6));
          let _588_v151;
          let _nw100 = new _module.C1();
          _nw100.__ctor(false, _565_v132, _386_v9, _383_v6, _583_cf82);
          _588_v151 = _nw100;
          let _589_v152;
          _589_v152 = _dafny.Map.Empty.slice().updateUnsafe(_588_v151,_585_v148);
          _587_v150 = (_587_v150).update(_dafny.Map.Empty.slice().updateUnsafe(_381_v4,new BigNumber(((((_589_v152).contains(_588_v151)) ? ((_589_v152).get(_588_v151)) : (_585_v148))).length)), !_dafny.areEqual(_dafny.Seq.of((_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))]), _dafny.Seq.of(_588_v151.f35, false, _588_v151.f35, _381_v4, (_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))])));
          let _590_v153;
          _590_v153 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(664),_382_v5);
          let _591_v154;
          _591_v154 = _module.D11.create_DC23(_dafny.MultiSet.fromElements(_384_v7));
          let _rhs49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(230),_dafny.Set.fromElements(_588_v151.f35, (_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))]));
          let _rhs50 = _591_v154;
          _590_v153 = _rhs49;
          _591_v154 = _rhs50;
        } else {
          let _592___mcc_h7 = (_source12).cf84;
          let _593_cf84 = _592___mcc_h7;
          let _594_v155;
          let _init12 = function (_595_i20) {
            return _dafny.Seq.of(false);
          };
          let _nw101 = Array((new BigNumber(17)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw101.length); _i0_12++) {
            _nw101[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _594_v155 = _nw101;
          let _596_v156;
          _596_v156 = _dafny.Map.Empty.slice().updateUnsafe(_594_v155,_386_v9);
          _386_v9 = (((_596_v156).contains(_594_v155)) ? ((_596_v156).get(_594_v155)) : (_386_v9));
          (_391_globalState).f26 = (_388_v11)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_388_v11).length))];
          let _597_v157;
          let _nw102 = new _module.C3();
          _nw102.__ctor(new BigNumber(-958), _571_v138);
          _597_v157 = _nw102;
          let _598_v158;
          _598_v158 = _module.D15.create_DC36(_380_v3, _383_v6);
          _598_v158 = _598_v158;
        }
      }
      let _599_v159;
      _599_v159 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_383_v6),_383_v6);
      let _600_v160;
      _600_v160 = _dafny.Map.Empty.slice().updateUnsafe(_599_v159,_module.__default.fm1(_391_globalState));
      let _601_v161;
      _601_v161 = _dafny.MultiSet.fromElements(_381_v4);
      let _602_v162;
      let _nw103 = Array((new BigNumber(5)).toNumber());
      _nw103[(_dafny.ZERO).toNumber()] = _378_v1;
      _nw103[(_dafny.ONE).toNumber()] = _378_v1;
      _nw103[(new BigNumber(2)).toNumber()] = _378_v1;
      _nw103[(new BigNumber(3)).toNumber()] = _378_v1;
      _nw103[(new BigNumber(4)).toNumber()] = _378_v1;
      _602_v162 = _nw103;
      let _603_v163;
      let _nw104 = new _module.C5();
      _nw104.__ctor(_600_v160, (_601_v161).update(_381_v4, _module.__default.abs(new BigNumber((_382_v5).length))), (_module.D10.create_DC21(_386_v9)).dtor_cf39, _383_v6, _602_v162);
      _603_v163 = _nw104;
      let _hi1 = _383_v6;
      for (let _604_i21 = _383_v6; _604_i21.isLessThan(_hi1); _604_i21 = _604_i21.plus(_dafny.ONE)) {
        let _index43 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_377_v0).length));
        (_377_v0)[_index43] = new BigNumber(-775);
        let _index44 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_377_v0).length));
        (_377_v0)[_index44] = _604_i21;
        let _605_v164;
        let _606_v165;
        let _607_v166;
        let _out20;
        let _out21;
        let _out22;
        let _outcollector9 = (_603_v163).m1(_391_globalState);
        _out20 = _outcollector9[0];
        _out21 = _outcollector9[1];
        _out22 = _outcollector9[2];
        _605_v164 = _out20;
        _606_v165 = _out21;
        _607_v166 = _out22;
        let _608_v167;
        _608_v167 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_381_v4, _381_v4));
        let _609_v168;
        _609_v168 = _module.D19.create_DC49(_381_v4, _381_v4, _604_i21);
        let _610_v169;
        _610_v169 = _dafny.Seq.of(_602_v162);
        let _pat_let_tv5 = _381_v4;
        let _611_v170;
        let _nw105 = new _module.C6();
        _nw105.__ctor(_381_v4, (new BigNumber(615)).plus(_383_v6), (_608_v167)[_module.__default.safeIndex(_604_i21, new BigNumber((_608_v167).length))], _dafny.Seq.of((_377_v0)[_module.__default.safeIndex(new BigNumber(416), new BigNumber((_377_v0).length))], (function (_pat_let10_0) {
          return function (_612_dt__update__tmp_h2) {
            return function (_pat_let11_0) {
              return function (_613_dt__update_hcf89_h0) {
                return _module.D19.create_DC49(_613_dt__update_hcf89_h0, (_612_dt__update__tmp_h2).dtor_cf90, (_612_dt__update__tmp_h2).dtor_cf91);
              }(_pat_let11_0);
            }(_pat_let_tv5);
          }(_pat_let10_0);
        }(_609_v168)).dtor_cf91), _383_v6, (_610_v169)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_610_v169).length))], _601_v161, _606_v165);
        _611_v170 = _nw105;
        _611_v170 = _611_v170;
        let _614_v171;
        _614_v171 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ugxwp")).length));
        (_391_globalState).f26 = ((new BigNumber((_386_v9).length)).plus(new BigNumber((_614_v171).cardinality()))).isLessThanOrEqualTo(((_377_v0)[_module.__default.safeIndex(new BigNumber(416), new BigNumber((_377_v0).length))]).multipliedBy(new BigNumber((_386_v9).length)));
      }
      let _615_v172;
      _615_v172 = _module.D4.create_DC9(_383_v6, false);
      let _616_v173;
      _616_v173 = _dafny.Map.Empty.slice().updateUnsafe(_381_v4,(_615_v172).dtor_cf17);
      _616_v173 = (_616_v173).update(_381_v4, _381_v4);
      let _617_v174;
      let _nw106 = new _module.C3();
      _nw106.__ctor(new BigNumber(923), _602_v162);
      _617_v174 = _nw106;
      let _618_v175;
      _618_v175 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_378_v1).length),_617_v174);
      let _619_v176;
      _619_v176 = _dafny.Map.Empty.slice().updateUnsafe(_618_v175,_383_v6);
      _619_v176 = (_619_v176).update(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_383_v6),_617_v174), _383_v6);
      let _620_v177;
      _620_v177 = _dafny.MultiSet.fromElements(_383_v6, _383_v6);
      if (((_603_v163).fm4(_dafny.Set.fromElements(_378_v1, _378_v1, _dafny.Seq.of(_module.__default.fm26(_380_v3, new BigNumber((_620_v177).cardinality()), _383_v6, _391_globalState))), _381_v4, _dafny.Seq.update(_dafny.Seq.of(!(_381_v4)), _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_381_v4)).length), new BigNumber((_dafny.Seq.of(!(_381_v4))).length)), _381_v4), _391_globalState)).isLessThan(_383_v6)) {
        (_391_globalState).f26 = true;
        _378_v1 = _dafny.Seq.Concat(_378_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(684)), function (_621_i22) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }));
        (_391_globalState).f0 = new BigNumber(522);
        let _index45 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_388_v11).length));
        (_388_v11)[_index45] = !(_381_v4);
        let _index46 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_388_v11).length));
        (_388_v11)[_index46] = _381_v4;
        _620_v177 = _620_v177;
      } else {
        let _622_v178;
        _622_v178 = _dafny.Seq.of(_378_v1);
        _378_v1 = (_622_v178)[_module.__default.safeIndex(_module.__default.safeModuloInt(_383_v6, _383_v6), new BigNumber((_622_v178).length))];
        let _623_v179;
        let _624_v180;
        let _625_v181;
        let _out23;
        let _out24;
        let _out25;
        let _outcollector10 = (_603_v163).m9(_391_globalState);
        _out23 = _outcollector10[0];
        _out24 = _outcollector10[1];
        _out25 = _outcollector10[2];
        _623_v179 = _out23;
        _624_v180 = _out24;
        _625_v181 = _out25;
        let _index47 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_377_v0).length));
        (_377_v0)[_index47] = new BigNumber((_378_v1).length);
        let _index48 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_377_v0).length));
        (_377_v0)[_index48] = _383_v6;
        _380_v3 = ((_625_v181) ? (_380_v3) : (_380_v3));
        let _626_v182;
        let _627_v183;
        let _628_v184;
        let _out26;
        let _out27;
        let _out28;
        let _outcollector11 = (_603_v163).m1(_391_globalState);
        _out26 = _outcollector11[0];
        _out27 = _outcollector11[1];
        _out28 = _outcollector11[2];
        _626_v182 = _out26;
        _627_v183 = _out27;
        _628_v184 = _out28;
      }
      _599_v159 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(_383_v6)).length), (_dafny.ZERO).minus(_383_v6)),_383_v6);
      let _index49 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length));
      (_377_v0)[_index49] = _383_v6;
      let _629_v185;
      _629_v185 = _dafny.Seq.of(_381_v4);
      let _index50 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length));
      (_377_v0)[_index50] = (_dafny.ZERO).minus(((_603_v163).fm4(_390_v13, _381_v4, _629_v185, _391_globalState)).plus(_module.__default.safeDivisionInt((_386_v9)[_module.__default.safeIndex(_383_v6, new BigNumber((_386_v9).length))], _383_v6)));
      let _rhs51 = ((_381_v4) ? (_381_v4) : (_381_v4));
      let _rhs52 = _383_v6;
      let _rhs53 = ((((_381_v4) ? (_381_v4) : (_381_v4))) ? (_381_v4) : (_381_v4));
      let _rhs54 = new BigNumber(451);
      let _lhs32 = _391_globalState;
      let _lhs33 = _391_globalState;
      _lhs32.f26 = _rhs51;
      _lhs33.f18 = _rhs52;
      _381_v4 = _rhs53;
      _383_v6 = _rhs54;
      _388_v11 = _388_v11;
      (_391_globalState).f26 = !(new BigNumber((_616_v173).length)).isEqualTo((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]);
      let _hi2 = (_383_v6).minus((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]);
      for (let _630_i23 = (_617_v174).fm45(_391_globalState); _630_i23.isLessThan(_hi2); _630_i23 = _630_i23.plus(_dafny.ONE)) {
        let _631_v186;
        _631_v186 = _module.D16.create_DC40(_381_v4, (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], new BigNumber(((_385_v8).update(_380_v3, (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))])).length));
        let _632_v187;
        _632_v187 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(419),_380_v3);
        let _633_v188;
        _633_v188 = _module.D4.create_DC8((((_632_v187).contains((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))])) ? ((_632_v187).get((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))])) : (_380_v3)), new BigNumber((_378_v1).length), _381_v4);
        let _pat_let_tv6 = _380_v3;
        let _634_v189;
        _634_v189 = _dafny.Map.Empty.slice().updateUnsafe(_381_v4,function (_pat_let12_0) {
          return function (_635_dt__update__tmp_h3) {
            return function (_pat_let13_0) {
              return function (_636_dt__update_hcf13_h0) {
                return _module.D4.create_DC8(_636_dt__update_hcf13_h0, (_635_dt__update__tmp_h3).dtor_cf14, (_635_dt__update__tmp_h3).dtor_cf15);
              }(_pat_let13_0);
            }(_pat_let_tv6);
          }(_pat_let12_0);
        }(_633_v188));
        let _637_v190;
        _637_v190 = _dafny.MultiSet.fromElements((_dafny.Map.Empty.slice().updateUnsafe(_381_v4,_633_v188)).update(_381_v4, _633_v188), _634_v189);
        let _638_v191;
        _638_v191 = _module.D2.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),_381_v4));
        let _639_v192;
        _639_v192 = _module.D17.create_DC43(_602_v162);
        let _640_v193;
        _640_v193 = _module.D17.create_DC45(_639_v192);
        let _641_v194;
        _641_v194 = _dafny.Seq.of(_640_v193, _module.D17.create_DC45(_639_v192), _640_v193);
        let _642_v195;
        _642_v195 = _dafny.Map.Empty.slice().updateUnsafe(_641_v194,_381_v4);
        let _index51 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length));
        let _rhs55 = !(!((_630_i23).isLessThan((_603_v163).fm4(_390_v13, false, _629_v185, _391_globalState))));
        let _rhs56 = (((!(true)) || (_381_v4)) ? (_module.__default.fm55(_383_v6, _638_v191, _381_v4, new BigNumber((_dafny.Set.fromElements(_381_v4, _381_v4)).length), _391_globalState)) : (_module.D16.create_DC40(_381_v4, (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], _630_i23)));
        let _rhs57 = _637_v190;
        let _rhs58 = new BigNumber((_642_v195).length);
        let _rhs59 = _381_v4;
        let _lhs34 = _391_globalState;
        let _lhs35 = _377_v0;
        let _lhs36 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length));
        _lhs34.f26 = _rhs55;
        _631_v186 = _rhs56;
        _637_v190 = _rhs57;
        _lhs35[_lhs36] = _rhs58;
        _381_v4 = _rhs59;
        (_391_globalState).f21 = _630_i23;
        let _643_v196;
        _643_v196 = _module.D2.create_DC4(_383_v6, new BigNumber((_378_v1).length), (_617_v174).fm45(_391_globalState), _380_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(440)), function (_644_i24) {
  return new _dafny.CodePoint('j'.codePointAt(0));
}));
        _380_v3 = ((((_620_v177).update(_630_i23, _module.__default.abs(_630_i23))).equals(_620_v177)) ? (_380_v3) : ((_643_v196).dtor_cf8));
        if (_381_v4) {
          let _645_v197;
          let _nw107 = Array((new BigNumber(7)).toNumber()).fill([]);
          _645_v197 = _nw107;
          _645_v197 = _645_v197;
          let _646_v198;
          _646_v198 = _module.D15.create_DC37(_381_v4, (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], _377_v0, true);
          let _647_v199;
          _647_v199 = _dafny.Map.Empty.slice().updateUnsafe(_630_i23,false);
          let _648_v200;
          _648_v200 = _dafny.Map.Empty.slice().updateUnsafe(!(_381_v4),_383_v6);
          let _649_v201;
          let _nw108 = new _module.C2();
          _nw108.__ctor(_601_v161, new BigNumber((_648_v200).length));
          _649_v201 = _nw108;
          let _650_v202;
          let _nw109 = Array((new BigNumber(13)).toNumber());
          _nw109[(_dafny.ZERO).toNumber()] = _646_v198;
          _nw109[(_dafny.ONE).toNumber()] = _646_v198;
          _nw109[(new BigNumber(2)).toNumber()] = _646_v198;
          _nw109[(new BigNumber(3)).toNumber()] = _module.D15.create_DC37(_381_v4, (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], _377_v0, _381_v4);
          _nw109[(new BigNumber(4)).toNumber()] = _646_v198;
          _nw109[(new BigNumber(5)).toNumber()] = _module.D15.create_DC37(_381_v4, new BigNumber((_647_v199).length), _377_v0, _381_v4);
          _nw109[(new BigNumber(6)).toNumber()] = _646_v198;
          _nw109[(new BigNumber(7)).toNumber()] = _646_v198;
          _nw109[(new BigNumber(8)).toNumber()] = _646_v198;
          _nw109[(new BigNumber(9)).toNumber()] = ((false) ? (_646_v198) : (_646_v198));
          _nw109[(new BigNumber(10)).toNumber()] = _646_v198;
          _nw109[(new BigNumber(11)).toNumber()] = _module.D15.create_DC37(!((_603_v163).fm3(_601_v161, _391_globalState)), new BigNumber((_dafny.Seq.of(_649_v201)).length), _377_v0, _381_v4);
          _nw109[(new BigNumber(12)).toNumber()] = _646_v198;
          _650_v202 = _nw109;
          let _index52 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_602_v162).length));
          (_602_v162)[_index52] = _378_v1;
          let _651_v203;
          _651_v203 = _dafny.Seq.of(_382_v5);
          let _index53 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_602_v162).length));
          let _rhs60 = _module.__default.safeModuloInt((_383_v6).plus(_630_i23), _630_i23);
          let _rhs61 = ((_386_v9)[_module.__default.safeIndex(new BigNumber((_651_v203).length), new BigNumber((_386_v9).length))]).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_649_v201).f32, (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]));
          let _rhs62 = _650_v202;
          let _rhs63 = _378_v1;
          let _rhs64 = _383_v6;
          let _lhs37 = _391_globalState;
          let _lhs38 = _391_globalState;
          let _lhs39 = _602_v162;
          let _lhs40 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_602_v162).length));
          let _lhs41 = _391_globalState;
          _lhs37.f18 = _rhs60;
          _lhs38.f26 = _rhs61;
          _650_v202 = _rhs62;
          _lhs39[_lhs40] = _rhs63;
          _lhs41.f3 = _rhs64;
          let _652_v204;
          _652_v204 = _dafny.Set.fromElements(_380_v3);
          let _653_v205;
          _653_v205 = _dafny.Map.Empty.slice().updateUnsafe(_381_v4,_652_v204);
          let _654_v206;
          _654_v206 = _dafny.Map.Empty.slice().updateUnsafe(_616_v173,(((_653_v205).contains(_381_v4)) ? ((_653_v205).get(_381_v4)) : (_652_v204)));
          _381_v4 = ((_dafny.areEqual(_629_v185, _629_v185)) ? ((((_647_v199).contains(new BigNumber((_654_v206).length))) ? ((_647_v199).get(new BigNumber((_654_v206).length))) : (_381_v4))) : (_381_v4));
          _383_v6 = new BigNumber(897);
          (_391_globalState).f0 = _383_v6;
        } else {
          let _655_v207;
          let _nw110 = Array((new BigNumber(16)).toNumber());
          _nw110[(_dafny.ZERO).toNumber()] = _388_v11;
          _nw110[(_dafny.ONE).toNumber()] = _388_v11;
          _nw110[(new BigNumber(2)).toNumber()] = _388_v11;
          _nw110[(new BigNumber(3)).toNumber()] = _388_v11;
          _nw110[(new BigNumber(4)).toNumber()] = _388_v11;
          _nw110[(new BigNumber(5)).toNumber()] = _388_v11;
          _nw110[(new BigNumber(6)).toNumber()] = _388_v11;
          _nw110[(new BigNumber(7)).toNumber()] = _388_v11;
          _nw110[(new BigNumber(8)).toNumber()] = _388_v11;
          _nw110[(new BigNumber(9)).toNumber()] = _388_v11;
          _nw110[(new BigNumber(10)).toNumber()] = _388_v11;
          _nw110[(new BigNumber(11)).toNumber()] = _388_v11;
          _nw110[(new BigNumber(12)).toNumber()] = _388_v11;
          _nw110[(new BigNumber(13)).toNumber()] = _388_v11;
          _nw110[(new BigNumber(14)).toNumber()] = _388_v11;
          _nw110[(new BigNumber(15)).toNumber()] = _388_v11;
          _655_v207 = _nw110;
          let _656_v209;
          _656_v209 = _dafny.Map.Empty.slice().updateUnsafe(_380_v3,_381_v4);
          let _657_v210;
          _657_v210 = _module.D2.create_DC3(function () {
  let _coll23 = new _dafny.Map();
  for (const _compr_23 of (_656_v209).Keys.Elements) {
    let _658_v208 = _compr_23;
    if ((_656_v209).contains(_658_v208)) {
      _coll23.push([_658_v208,_381_v4]);
    }
  }
  return _coll23;
}());
          let _659_v211;
          _659_v211 = _module.D2.create_DC5(_657_v210);
          let _660_v212;
          _660_v212 = _module.D2.create_DC5(_659_v211);
          (_391_globalState).f26 = (((_616_v173).contains(_module.__default.fm0(_630_i23, _391_globalState))) ? ((_616_v173).get(_module.__default.fm0(_630_i23, _391_globalState))) : (((_381_v4) ? ((_module.D16.create_DC42((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], _655_v207, (_617_v174).fm46(_383_v6, _380_v3, _391_globalState), _660_v212)).dtor_cf80) : (_381_v4))));
          let _661_v213;
          let _662_v214;
          let _out29;
          let _out30;
          let _outcollector12 = _module.__default.m0(((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]).isLessThan((_dafny.ZERO).minus((_617_v174).fm45(_391_globalState))), new BigNumber((_378_v1).length), _module.__default.safeDivisionInt(_630_i23, _383_v6), (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], _391_globalState);
          _out29 = _outcollector12[0];
          _out30 = _outcollector12[1];
          _661_v213 = _out29;
          _662_v214 = _out30;
          let _663_v215;
          let _init13 = ((_664_v214, _665_v4) => function (_666_i25) {
            return _dafny.Map.Empty.slice().updateUnsafe(_664_v214,!(_665_v4));
          })(_662_v214, _381_v4);
          let _nw111 = Array((new BigNumber(23)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw111.length); _i0_13++) {
            _nw111[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _663_v215 = _nw111;
          let _667_v216;
          _667_v216 = _dafny.Seq.of(_663_v215, _663_v215);
          let _668_v217;
          _668_v217 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(((false) ? (_dafny.Seq.update(_378_v1, _module.__default.safeIndex(_662_v214, new BigNumber((_378_v1).length)), _380_v3)) : (_378_v1)), _module.__default.safeIndex((_dafny.ZERO).minus(_630_i23), new BigNumber((((false) ? (_dafny.Seq.update(_378_v1, _module.__default.safeIndex(_662_v214, new BigNumber((_378_v1).length)), _380_v3)) : (_378_v1))).length)), _380_v3)).length),(_667_v216)[_module.__default.safeIndex(_662_v214, new BigNumber((_667_v216).length))]);
          _668_v217 = (_668_v217).update((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], _663_v215);
          let _669_v218;
          _669_v218 = _dafny.Seq.of(_602_v162, _602_v162, _602_v162);
          let _670_v219;
          _670_v219 = _dafny.Map.Empty.slice().updateUnsafe((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))],(_669_v218)[_module.__default.safeIndex(new BigNumber(283), new BigNumber((_669_v218).length))]);
          let _671_v220;
          let _nw112 = new _module.C5();
          _nw112.__ctor((_603_v163).f38, _601_v161, _386_v9, new BigNumber((_379_v2).cardinality()), ((_381_v4) ? ((((_670_v219).contains(new BigNumber((_629_v185).length))) ? ((_670_v219).get(new BigNumber((_629_v185).length))) : (_602_v162))) : (_602_v162)));
          _671_v220 = _nw112;
          let _672_v221;
          _672_v221 = _dafny.Seq.of(((_381_v4) ? (_378_v1) : (_378_v1)), _378_v1, _378_v1);
          _672_v221 = _dafny.Seq.update(_dafny.Seq.update(_672_v221, _module.__default.safeIndex(new BigNumber(938), new BigNumber((_672_v221).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(582)), ((_673_v3) => function (_674_i26) {
            return _673_v3;
          })(_380_v3))), _module.__default.safeIndex(_662_v214, new BigNumber((_dafny.Seq.update(_672_v221, _module.__default.safeIndex(new BigNumber(938), new BigNumber((_672_v221).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(582)), ((_675_v3) => function (_676_i26) {
            return _675_v3;
          })(_380_v3)))).length)), _module.__default.fm14(!(_661_v213), _381_v4, _381_v4, _383_v6, _391_globalState));
        }
      }
      if (_381_v4) {
        let _rhs65 = ((false) ? ((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]) : (_module.__default.safeDivisionInt((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], new BigNumber((_616_v173).length))));
        let _rhs66 = _dafny.Seq.Concat(_386_v9, _386_v9);
        let _lhs42 = _391_globalState;
        _lhs42.f21 = _rhs65;
        _386_v9 = _rhs66;
        let _677_v222;
        _677_v222 = _module.D9.create_DC20(_module.__default.fm1(_391_globalState), _381_v4, ((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]).isLessThan((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]), _381_v4, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_380_v3, _module.__default.fm12((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], _381_v4, _391_globalState))).length),new BigNumber(870)));
        let _source13 = _677_v222;
        if (_source13.is_DC18) {
          let _678___mcc_h8 = (_source13).cf32;
          let _679_cf32 = _678___mcc_h8;
          _383_v6 = (_617_v174).fm45(_391_globalState);
          let _680_v223;
          _680_v223 = _dafny.Seq.of(((_601_v161).update(_679_cf32, _module.__default.abs((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]))).Union(_601_v161));
          _680_v223 = _dafny.Seq.update(_680_v223, _module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_382_v5).length), (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]), new BigNumber((_680_v223).length)), _module.__default.fm32(_391_globalState));
          _386_v9 = _module.__default.fm9(_679_cf32, _391_globalState);
          let _681_v224;
          _681_v224 = _dafny.Map.Empty.slice().updateUnsafe(_679_cf32,_601_v161);
          let _682_v225;
          _682_v225 = _module.D11.create_DC23(_454_v62);
          let _683_v226;
          _683_v226 = _dafny.Map.Empty.slice().updateUnsafe((_681_v224).contains(_679_cf32),_682_v225);
          _683_v226 = (_683_v226).update((_module.__default.fm1(_391_globalState)).isEqualTo((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]), _682_v225);
        } else if (_source13.is_DC19) {
          let _684___mcc_h9 = (_source13).cf33;
          let _685_cf33 = _684___mcc_h9;
          let _686_v227;
          _686_v227 = _dafny.Map.Empty.slice().updateUnsafe(_629_v185,_602_v162);
          let _687_v228;
          _687_v228 = _module.D7.create_DC13(_629_v185);
          let _688_v229;
          _688_v229 = _dafny.MultiSet.fromElements(_599_v159);
          let _689_v230;
          let _nw113 = new _module.C6();
          _nw113.__ctor(_381_v4, (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], _601_v161, _dafny.Seq.Concat(_module.__default.fm9(_685_cf33, _391_globalState), _dafny.Seq.of((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))])), (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], (((_686_v227).contains((_687_v228).dtor_cf23)) ? ((_686_v227).get((_687_v228).dtor_cf23)) : (_602_v162)), _601_v161, (((_688_v229).contains(_599_v159)) ? ((_688_v229).get(_599_v159)) : ((_dafny.ZERO).minus(new BigNumber((_382_v5).length)))));
          _689_v230 = _nw113;
          _689_v230 = _689_v230;
          let _index54 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_602_v162).length));
          (_602_v162)[_index54] = _dafny.Seq.Concat(_378_v1, _378_v1);
          let _index55 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_602_v162).length));
          (_602_v162)[_index55] = _378_v1;
          let _690_v231;
          _690_v231 = _module.D17.create_DC45(_module.D17.create_DC44(new BigNumber((_378_v1).length)));
          let _691_v232;
          _691_v232 = _module.D17.create_DC45(_690_v231);
          let _692_v233;
          _692_v233 = _module.D12.create_DC27(true, _620_v177, _383_v6, _377_v0);
          let _693_v234;
          _693_v234 = _dafny.Map.Empty.slice().updateUnsafe(_691_v232,(((_689_v230).f33) ? ((_692_v233).dtor_cf50) : (_377_v0)));
          _693_v234 = (_693_v234).update(_691_v232, _377_v0);
          let _694_v235;
          _694_v235 = _module.D17.create_DC44(_383_v6);
          let _695_v236;
          let _nw114 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _695_v236 = _nw114;
          let _696_v237;
          let _nw115 = new _module.C4();
          _nw115.__ctor((_694_v235).dtor_cf83, _695_v236);
          _696_v237 = _nw115;
          (_391_globalState).f26 = (new BigNumber(744)).isLessThan((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_696_v237,(_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))])).length)).minus(_383_v6));
        } else if (_source13.is_DC20) {
          let _697___mcc_h10 = (_source13).cf34;
          let _698___mcc_h11 = (_source13).cf35;
          let _699___mcc_h12 = (_source13).cf36;
          let _700___mcc_h13 = (_source13).cf37;
          let _701___mcc_h14 = (_source13).cf38;
          let _702_cf38 = _701___mcc_h14;
          let _703_cf37 = _700___mcc_h13;
          let _704_cf36 = _699___mcc_h12;
          let _705_cf35 = _698___mcc_h11;
          let _706_cf34 = _697___mcc_h10;
          (_391_globalState).f26 = false;
          _381_v4 = (_384_v7).IsProperSubsetOf(_384_v7);
          let _707_v238;
          _707_v238 = _dafny.Map.Empty.slice().updateUnsafe(_388_v11,_383_v6);
          let _708_v239;
          _708_v239 = _dafny.Map.Empty.slice().updateUnsafe(!(_381_v4),_378_v1);
          let _709_v240;
          _709_v240 = _dafny.Seq.of(_707_v238, _dafny.Map.Empty.slice().updateUnsafe(_388_v11,new BigNumber(((((_708_v239).contains(_381_v4)) ? ((_708_v239).get(_381_v4)) : (_dafny.Seq.update(_378_v1, _module.__default.safeIndex(_383_v6, new BigNumber((_378_v1).length)), _380_v3)))).length)));
          let _710_v241;
          _710_v241 = _dafny.Map.Empty.slice().updateUnsafe(_383_v6,_707_v238);
          let _711_v242;
          _711_v242 = _dafny.Set.fromElements(_380_v3, _380_v3);
          (_391_globalState).f26 = ((_709_v240)[_module.__default.safeIndex(_383_v6, new BigNumber((_709_v240).length))]).equals((((_710_v241).contains(new BigNumber((_711_v242).length))) ? ((_710_v241).get(new BigNumber((_711_v242).length))) : (_dafny.Map.Empty.slice().updateUnsafe(_388_v11,(((_601_v161).contains(_703_cf37)) ? ((_601_v161).get(_703_cf37)) : ((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]))))));
          (_391_globalState).f18 = ((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]).plus((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]);
        } else {
          let _712___mcc_h15 = (_source13).cf31;
          let _713_cf31 = _712___mcc_h15;
          (_603_v163).m10(false, _dafny.Seq.Concat(_386_v9, _386_v9), _391_globalState);
          let _714_v243;
          _714_v243 = _dafny.Set.fromElements(_377_v0, _377_v0);
          let _index56 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_388_v11).length));
          (_388_v11)[_index56] = (_714_v243).IsDisjointFrom(_714_v243);
          let _index57 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_388_v11).length));
          let _rhs67 = _381_v4;
          let _rhs68 = _381_v4;
          let _lhs43 = _388_v11;
          let _lhs44 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_388_v11).length));
          let _lhs45 = _391_globalState;
          _lhs43[_lhs44] = _rhs67;
          _lhs45.f26 = _rhs68;
          _377_v0 = _377_v0;
          let _715_v244;
          let _nw116 = Array((new BigNumber(6)).toNumber());
          _nw116[(_dafny.ZERO).toNumber()] = true;
          _nw116[(_dafny.ONE).toNumber()] = !((_388_v11)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_388_v11).length))]);
          _nw116[(new BigNumber(2)).toNumber()] = false;
          _nw116[(new BigNumber(3)).toNumber()] = (_388_v11)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_388_v11).length))];
          _nw116[(new BigNumber(4)).toNumber()] = _381_v4;
          _nw116[(new BigNumber(5)).toNumber()] = (_603_v163).fm3(_601_v161, _391_globalState);
          _715_v244 = _nw116;
          let _716_v245;
          let _nw117 = new _module.C4();
          _nw117.__ctor(new BigNumber(231), _602_v162);
          _716_v245 = _nw117;
          let _717_v246;
          _717_v246 = _dafny.Map.Empty.slice().updateUnsafe(_715_v244,_716_v245);
          _717_v246 = (_717_v246).update(_715_v244, _716_v245);
        }
        let _718_v247;
        _718_v247 = _dafny.Set.fromElements(_382_v5, _382_v5);
        let _719_v249;
        _719_v249 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(_381_v4), _382_v5, _382_v5);
        let _index58 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_388_v11).length));
        (_388_v11)[_index58] = (function () {
          let _coll24 = new _dafny.Set();
          for (const _compr_24 of (_718_v247).Elements) {
            let _720_v248 = _compr_24;
            if ((_718_v247).contains(_720_v248)) {
              _coll24.add(_720_v248);
            }
          }
          return _coll24;
        }()).IsDisjointFrom(function () {
          let _coll25 = new _dafny.Set();
          for (const _compr_25 of (_719_v249).Elements) {
            let _721_v250 = _compr_25;
            if ((_719_v249).contains(_721_v250)) {
              _coll25.add(_721_v250);
            }
          }
          return _coll25;
        }());
        let _722_v251;
        _722_v251 = _dafny.Set.fromElements(_617_v174, _617_v174);
        let _723_v252;
        _723_v252 = _dafny.Seq.of(_722_v251);
        let _index59 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_388_v11).length));
        (_388_v11)[_index59] = !(!_dafny.areEqual(_723_v252, _723_v252));
        (_391_globalState).f26 = (_388_v11)[_module.__default.safeIndex(new BigNumber(144), new BigNumber((_388_v11).length))];
        let _724_v253;
        let _nw118 = Array((new BigNumber(29)).toNumber()).fill([]);
        _724_v253 = _nw118;
        let _725_v254;
        let _init14 = ((_726_v0, _727_v4) => function (_728_i27) {
          return _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_726_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_726_v0).length))]),_727_v4);
        })(_377_v0, _381_v4);
        let _nw119 = Array((new BigNumber(23)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw119.length); _i0_14++) {
          _nw119[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _725_v254 = _nw119;
        let _index60 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_724_v253).length));
        (_724_v253)[_index60] = _725_v254;
        let _index61 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_724_v253).length));
        (_724_v253)[_index61] = _725_v254;
      } else {
        _377_v0 = _377_v0;
        let _729_v255;
        _729_v255 = _dafny.Set.fromElements(_module.__default.fm26(_380_v3, new BigNumber((_386_v9).length), (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], _391_globalState), _380_v3);
        let _730_v256;
        let _nw120 = new _module.C2();
        _nw120.__ctor(_dafny.MultiSet.fromElements(false), (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]);
        _730_v256 = _nw120;
        let _731_v257;
        _731_v257 = _dafny.MultiSet.fromElements(_730_v256, _730_v256);
        let _732_v258;
        let _nw121 = new _module.C4();
        _nw121.__ctor(_383_v6, _602_v162);
        _732_v258 = _nw121;
        let _nw122 = Array((new BigNumber(19)).toNumber());
        _nw122[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_729_v255).length), _module.__default.fm1(_391_globalState));
        _nw122[(_dafny.ONE).toNumber()] = _383_v6;
        _nw122[(new BigNumber(2)).toNumber()] = _383_v6;
        _nw122[(new BigNumber(3)).toNumber()] = ((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]).minus(_383_v6);
        _nw122[(new BigNumber(4)).toNumber()] = (new BigNumber((_dafny.Seq.update(_386_v9, _module.__default.safeIndex((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], new BigNumber((_386_v9).length)), (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))])).length)).plus(_383_v6);
        _nw122[(new BigNumber(5)).toNumber()] = new BigNumber(802);
        _nw122[(new BigNumber(6)).toNumber()] = ((_381_v4) ? (new BigNumber((_731_v257).cardinality())) : (_383_v6));
        _nw122[(new BigNumber(7)).toNumber()] = ((_603_v163).fm4(_dafny.Set.fromElements(_378_v1, _378_v1), _381_v4, _629_v185, _391_globalState)).minus(new BigNumber(-311));
        _nw122[(new BigNumber(8)).toNumber()] = _383_v6;
        _nw122[(new BigNumber(9)).toNumber()] = _383_v6;
        _nw122[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], new BigNumber((_629_v185).length));
        _nw122[(new BigNumber(11)).toNumber()] = new BigNumber(956);
        _nw122[(new BigNumber(12)).toNumber()] = ((_381_v4) ? ((_730_v256).f32) : ((_386_v9)[_module.__default.safeIndex(new BigNumber((_629_v185).length), new BigNumber((_386_v9).length))]));
        _nw122[(new BigNumber(13)).toNumber()] = (new BigNumber((_601_v161).cardinality())).multipliedBy((_730_v256).f32);
        _nw122[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_383_v6, (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]));
        _nw122[(new BigNumber(15)).toNumber()] = (_730_v256).f32;
        _nw122[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_732_v258,function () {
          let _coll26 = new _dafny.Set();
          for (const _compr_26 of (_378_v1).Elements) {
            let _733_v259 = _compr_26;
            if (_dafny.Seq.contains(_378_v1, _733_v259)) {
              _coll26.add(_733_v259);
            }
          }
          return _coll26;
        }())).length);
        _nw122[(new BigNumber(17)).toNumber()] = (_603_v163).fm4(_390_v13, _381_v4, _629_v185, _391_globalState);
        _nw122[(new BigNumber(18)).toNumber()] = (_732_v258).f27;
        _377_v0 = _nw122;
        let _index62 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_388_v11).length));
        (_388_v11)[_index62] = (_381_v4) || (!(_381_v4));
        let _734_v260;
        _734_v260 = _dafny.Map.Empty.slice().updateUnsafe(_377_v0,_384_v7);
        let _735_v263;
        let _nw123 = Array((new BigNumber(29)).toNumber());
        _nw123[(_dafny.ZERO).toNumber()] = _734_v260;
        _nw123[(_dafny.ONE).toNumber()] = _734_v260;
        _nw123[(new BigNumber(2)).toNumber()] = _734_v260;
        _nw123[(new BigNumber(3)).toNumber()] = _734_v260;
        _nw123[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_377_v0,_384_v7);
        _nw123[(new BigNumber(5)).toNumber()] = ((_381_v4) ? (_734_v260) : (_734_v260));
        _nw123[(new BigNumber(6)).toNumber()] = _734_v260;
        _nw123[(new BigNumber(7)).toNumber()] = (_734_v260).Merge(_734_v260);
        _nw123[(new BigNumber(8)).toNumber()] = _734_v260;
        _nw123[(new BigNumber(9)).toNumber()] = _734_v260;
        _nw123[(new BigNumber(10)).toNumber()] = (_734_v260).Merge(_734_v260);
        _nw123[(new BigNumber(11)).toNumber()] = (_734_v260).update(_377_v0, _384_v7);
        _nw123[(new BigNumber(12)).toNumber()] = (_734_v260).Merge(_734_v260);
        _nw123[(new BigNumber(13)).toNumber()] = _734_v260;
        _nw123[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_377_v0,function () {
          let _coll27 = new _dafny.Set();
          for (const _compr_27 of (function () {
            let _coll28 = new _dafny.Map();
            for (const _compr_28 of (_dafny.Seq.of((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))])).Elements) {
              let _736_v261 = _compr_28;
              if (_dafny.Seq.contains(_dafny.Seq.of((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]), _736_v261)) {
                _coll28.push([(_736_v261).multipliedBy((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]),_381_v4]);
              }
            }
            return _coll28;
          }()).Keys.Elements) {
            let _737_v262 = _compr_27;
            if ((function () {
              let _coll29 = new _dafny.Map();
              for (const _compr_29 of (_dafny.Seq.of((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))])).Elements) {
                let _736_v261 = _compr_29;
                if (_dafny.Seq.contains(_dafny.Seq.of((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]), _736_v261)) {
                  _coll29.push([(_736_v261).multipliedBy((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]),_381_v4]);
                }
              }
              return _coll29;
            }()).contains(_737_v262)) {
              _coll27.add((_737_v262).minus(new BigNumber(458)));
            }
          }
          return _coll27;
        }());
        _nw123[(new BigNumber(15)).toNumber()] = _734_v260;
        _nw123[(new BigNumber(16)).toNumber()] = (_734_v260).Merge(_734_v260);
        _nw123[(new BigNumber(17)).toNumber()] = (_734_v260).Merge(_dafny.Map.Empty.slice().updateUnsafe(_377_v0,_384_v7));
        _nw123[(new BigNumber(18)).toNumber()] = _734_v260;
        _nw123[(new BigNumber(19)).toNumber()] = _734_v260;
        _nw123[(new BigNumber(20)).toNumber()] = _734_v260;
        _nw123[(new BigNumber(21)).toNumber()] = _734_v260;
        _nw123[(new BigNumber(22)).toNumber()] = _734_v260;
        _nw123[(new BigNumber(23)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_377_v0,_384_v7);
        _nw123[(new BigNumber(24)).toNumber()] = (_734_v260).Merge(_734_v260);
        _nw123[(new BigNumber(25)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_377_v0,_384_v7);
        _nw123[(new BigNumber(26)).toNumber()] = (_734_v260).Merge(_734_v260);
        _nw123[(new BigNumber(27)).toNumber()] = _734_v260;
        _nw123[(new BigNumber(28)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_377_v0,_384_v7);
        _735_v263 = _nw123;
        let _index63 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_735_v263).length));
        (_735_v263)[_index63] = _734_v260;
        let _738_v264;
        _738_v264 = _module.D2.create_DC4((_dafny.ZERO).minus(_383_v6), (_732_v258).f27, ((_730_v256).f32).multipliedBy((_732_v258).f27), _380_v3, _378_v1);
        let _739_v265;
        _739_v265 = _module.D1.create_DC2(_380_v3, (_730_v256).f32);
        let _index64 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_388_v11).length));
        let _index65 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_735_v263).length));
        let _rhs69 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_732_v258).f27), _module.__default.safeModuloInt((_730_v256).f32, (_730_v256).f32));
        let _rhs70 = _381_v4;
        let _rhs71 = ((_739_v265).dtor_cf3).multipliedBy((_732_v258).f27);
        let _rhs72 = (((_381_v4) || (_381_v4)) ? (_734_v260) : (_734_v260));
        let _rhs73 = _738_v264;
        let _lhs46 = _391_globalState;
        let _lhs47 = _388_v11;
        let _lhs48 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_388_v11).length));
        let _lhs49 = _391_globalState;
        let _lhs50 = _735_v263;
        let _lhs51 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_735_v263).length));
        _lhs46.f15 = _rhs69;
        _lhs47[_lhs48] = _rhs70;
        _lhs49.f15 = _rhs71;
        _lhs50[_lhs51] = _rhs72;
        _738_v264 = _rhs73;
        let _740_v266;
        let _nw124 = new _module.C2();
        _nw124.__ctor((_730_v256).f31, new BigNumber(-102));
        _740_v266 = _nw124;
        _740_v266 = _740_v266;
        (_391_globalState).f26 = ((_381_v4) ? (true) : (!_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(274)), ((_741_v256) => function (_742_i28) {
          return (_741_v256).f32;
        })(_730_v256)), (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))])));
      }
      let _743_i29;
      _743_i29 = _dafny.ZERO;
      L2: {
        while (!(_381_v4)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_743_i29)) {
              break L2;
            }
            _743_i29 = (_743_i29).plus(_dafny.ONE);
            let _744_v267;
            _744_v267 = _module.D8.create_DC16((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], (_383_v6).multipliedBy(_383_v6));
            _744_v267 = _744_v267;
            (_391_globalState).f0 = (_dafny.ZERO).minus((new BigNumber((_378_v1).length)).minus(((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]).plus(new BigNumber((_382_v5).length))));
            if (_381_v4) {
              (_391_globalState).f21 = _module.__default.safeModuloInt(((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]).plus(new BigNumber((_386_v9).length)), new BigNumber((_620_v177).cardinality()));
              let _745_v268;
              _745_v268 = _dafny.Seq.of(_386_v9, _386_v9);
              _599_v159 = (_599_v159).update((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], new BigNumber((_745_v268).length));
              let _index66 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length));
              (_377_v0)[_index66] = (_386_v9)[_module.__default.safeIndex((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], new BigNumber((_386_v9).length))];
              let _746_v269;
              let _init15 = ((_747_v4) => function (_748_i30) {
                return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_747_v4, _747_v4),_747_v4);
              })(_381_v4);
              let _nw125 = Array((new BigNumber(10)).toNumber());
              for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw125.length); _i0_15++) {
                _nw125[_i0_15] = _init15(new BigNumber(_i0_15));
              }
              _746_v269 = _nw125;
              let _749_v270;
              _749_v270 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_381_v4, _381_v4),_381_v4);
              let _index67 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_746_v269).length));
              (_746_v269)[_index67] = _749_v270;
              let _index68 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_746_v269).length));
              (_746_v269)[_index68] = _749_v270;
              (_603_v163).m10(!(_381_v4) || (_381_v4), _386_v9, _391_globalState);
            } else {
              _381_v4 = _381_v4;
              (_391_globalState).f26 = _381_v4;
              let _750_v271;
              _750_v271 = _module.D12.create_DC27(!(_381_v4), (_620_v177).update(new BigNumber(489), _module.__default.abs(_383_v6)), _383_v6, _377_v0);
              let _751_v272;
              let _nw126 = Array((new BigNumber(3)).toNumber());
              _nw126[(_dafny.ZERO).toNumber()] = _750_v271;
              _nw126[(_dafny.ONE).toNumber()] = _750_v271;
              _nw126[(new BigNumber(2)).toNumber()] = _750_v271;
              _751_v272 = _nw126;
              let _index69 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_751_v272).length));
              (_751_v272)[_index69] = _750_v271;
              let _752_v274;
              _752_v274 = _dafny.Map.Empty.slice().updateUnsafe(_380_v3,_381_v4);
              let _753_v275;
              _753_v275 = _dafny.Seq.of(function () {
                let _coll30 = new _dafny.Map();
                for (const _compr_30 of (_752_v274).Keys.Elements) {
                  let _754_v273 = _compr_30;
                  if ((_752_v274).contains(_754_v273)) {
                    _coll30.push([_754_v273,(_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]]);
                  }
                }
                return _coll30;
              }());
              let _index70 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_751_v272).length));
              let _rhs74 = _750_v271;
              let _rhs75 = (new BigNumber(166)).multipliedBy(((_386_v9)[_module.__default.safeIndex(_383_v6, new BigNumber((_386_v9).length))]).plus((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]));
              let _rhs76 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber(-560)).plus(new BigNumber((_753_v275).length)))), _383_v6);
              let _rhs77 = (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))];
              let _rhs78 = !(true) || (!(_381_v4));
              let _lhs52 = _751_v272;
              let _lhs53 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_751_v272).length));
              let _lhs54 = _391_globalState;
              let _lhs55 = _391_globalState;
              let _lhs56 = _391_globalState;
              let _lhs57 = _391_globalState;
              _lhs52[_lhs53] = _rhs74;
              _lhs54.f0 = _rhs75;
              _lhs55.f0 = _rhs76;
              _lhs56.f21 = _rhs77;
              _lhs57.f26 = _rhs78;
              let _755_v276;
              _755_v276 = _dafny.Map.Empty.slice().updateUnsafe(_383_v6,_381_v4);
              _755_v276 = _755_v276;
              let _rhs79 = _381_v4;
              let _rhs80 = _381_v4;
              let _rhs81 = (new BigNumber(139)).plus(_383_v6);
              let _rhs82 = _module.__default.fm41((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))], _383_v6, _381_v4, _391_globalState);
              let _rhs83 = _386_v9;
              let _lhs58 = _391_globalState;
              let _lhs59 = _391_globalState;
              let _lhs60 = _391_globalState;
              _lhs58.f26 = _rhs79;
              _lhs59.f26 = _rhs80;
              _lhs60.f0 = _rhs81;
              _629_v185 = _rhs82;
              _386_v9 = _rhs83;
            }
            let _rhs84 = (((_module.D6.create_DC12(_381_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(721)), ((_756_v4) => function (_757_i31) {
  return _dafny.Seq.of(_756_v4, false);
})(_381_v4)), _378_v1)).dtor_cf20) ? ((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]) : ((_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))]));
            let _rhs85 = (_377_v0)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_377_v0).length))];
            let _rhs86 = _616_v173;
            let _lhs61 = _391_globalState;
            let _lhs62 = _391_globalState;
            _lhs61.f24 = _rhs84;
            _lhs62.f18 = _rhs85;
            _616_v173 = _rhs86;
          }
        }
      }
      process.stdout.write(_dafny.toString((_377_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v0)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_378_v1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_379_v2).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("yn")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_380_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_381_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_382_v5).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_383_v6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v7).equals(_dafny.Set.fromElements(_dafny.ONE, new BigNumber(491)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_385_v8).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_386_v9, _dafny.Seq.of(new BigNumber(491)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_387_v10)[_dafny.ZERO], _dafny.Seq.of(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_387_v10)[_dafny.ONE], _dafny.Seq.of(new BigNumber(491)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_387_v10)[new BigNumber(2)], _dafny.Seq.of(new BigNumber(491)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_388_v11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_388_v11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_388_v11)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_389_v12).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_390_v13).equals(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("yn")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_391_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_391_globalState).f1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_391_globalState).f2)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_391_globalState).f2)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_391_globalState).f2)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_391_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_globalState.f8).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("yn")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_391_globalState).f10)[_dafny.ZERO], _dafny.Seq.of(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_391_globalState).f10)[_dafny.ONE], _dafny.Seq.of(new BigNumber(491)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_391_globalState).f10)[new BigNumber(2)], _dafny.Seq.of(new BigNumber(491)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_391_globalState.f13).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_391_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_391_globalState).f17).equals(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("yn")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_391_globalState.f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_globalState).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_globalState).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_391_globalState.f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_391_globalState).f22).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_391_globalState).f23)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_391_globalState).f23)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_391_globalState).f23)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_391_globalState.f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_globalState).f25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_391_globalState.f26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_453_v61).dtor_cf45).dtor_cf42).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_dafny.ONE, new BigNumber(491))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_454_v62).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_dafny.ONE, new BigNumber(491))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_455_v63).dtor_cf42).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_dafny.ONE, new BigNumber(491))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_599_v159).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,new BigNumber(-491)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_600_v160).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(491),new BigNumber(-491)),new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_601_v161).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_602_v162)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_602_v162)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_602_v162)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_602_v162)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_602_v162)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_603_v163).f38).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(491),new BigNumber(-491)),new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_615_v172).dtor_cf16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_615_v172).dtor_cf17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_616_v173).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_618_v175).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_619_v176).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_620_v177).equals(_dafny.MultiSet.fromElements(new BigNumber(-491), new BigNumber(-491)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_629_v185, _dafny.Seq.of(false, false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_743_i29));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + this.cf0.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.UnicodeFromString("");
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
    static create_DC2(cf2, cf3) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO);
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
    static create_DC4(cf5, cf6, cf7, cf8, cf9) {
      let $dt = new D2(0);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC3(cf4) {
      let $dt = new D2(1);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC5(cf10) {
      let $dt = new D2(2);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + this.cf9.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.Seq.UnicodeFromString(""));
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
    static create_DC6(cf11) {
      let $dt = new D3(0);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf11 === other.cf11;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return false;
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
    static create_DC8(cf13, cf14, cf15) {
      let $dt = new D4(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC9(cf16, cf17) {
      let $dt = new D4(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC7(cf12) {
      let $dt = new D4(2);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16) && this.cf17 === other.cf17;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf12 === other.cf12;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC8(new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, false);
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
    static create_DC10(cf18) {
      let $dt = new D5(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC10" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
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
    static create_DC12(cf20, cf21, cf22) {
      let $dt = new D6(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC11(cf19) {
      let $dt = new D6(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + this.cf22.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC11" + "(" + _dafny.toString(this.cf19) + ")";
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
        return other.$tag === 1 && this.cf19 === other.cf19;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC12(false, _dafny.Seq.of(), _dafny.Seq.UnicodeFromString(""));
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
    static create_DC14(cf24, cf25, cf26, cf27) {
      let $dt = new D7(0);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC13(cf23) {
      let $dt = new D7(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC13" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf24, other.cf24) && this.cf25 === other.cf25 && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC14(_dafny.ZERO, null, _dafny.Seq.UnicodeFromString(""), _dafny.MultiSet.Empty);
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
    static create_DC16(cf29, cf30) {
      let $dt = new D8(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC15(cf28) {
      let $dt = new D8(1);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC16" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC15" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29) && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC16(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC18(cf32) {
      let $dt = new D9(0);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC19(cf33) {
      let $dt = new D9(1);
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC20(cf34, cf35, cf36, cf37, cf38) {
      let $dt = new D9(2);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC17(cf31) {
      let $dt = new D9(3);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get is_DC17() { return this.$tag === 3; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC18" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC19" + "(" + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC17" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf32 === other.cf32;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf33 === other.cf33;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf34, other.cf34) && this.cf35 === other.cf35 && this.cf36 === other.cf36 && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC18(false);
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
    static create_DC22(cf40, cf41) {
      let $dt = new D10(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC21(cf39) {
      let $dt = new D10(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC22" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC21" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC22(_dafny.ZERO, _dafny.MultiSet.Empty);
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
    static create_DC24(cf43, cf44) {
      let $dt = new D11(0);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC23(cf42) {
      let $dt = new D11(1);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC25(cf45) {
      let $dt = new D11(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC24" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC23" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf43, other.cf43) && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC24(_dafny.ZERO, _module.D4.Default());
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
    static create_DC27(cf47, cf48, cf49, cf50) {
      let $dt = new D12(0);
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC28(cf51, cf52, cf53) {
      let $dt = new D12(1);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC26(cf46) {
      let $dt = new D12(2);
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC29(cf54) {
      let $dt = new D12(3);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get is_DC29() { return this.$tag === 3; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC27" + "(" + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC28" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC26" + "(" + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC29" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf47 === other.cf47 && _dafny.areEqual(this.cf48, other.cf48) && _dafny.areEqual(this.cf49, other.cf49) && this.cf50 === other.cf50;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf51, other.cf51) && _dafny.areEqual(this.cf52, other.cf52) && this.cf53 === other.cf53;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf46 === other.cf46;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC27(false, _dafny.MultiSet.Empty, _dafny.ZERO, []);
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
    static create_DC31(cf56) {
      let $dt = new D13(0);
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC32(cf57, cf58, cf59, cf60) {
      let $dt = new D13(1);
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC30(cf55) {
      let $dt = new D13(2);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC33(cf61) {
      let $dt = new D13(3);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get is_DC33() { return this.$tag === 3; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC31" + "(" + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC32" + "(" + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC30" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC33" + "(" + _dafny.toString(this.cf61) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf56, other.cf56);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf57, other.cf57) && _dafny.areEqual(this.cf58, other.cf58) && _dafny.areEqual(this.cf59, other.cf59) && this.cf60 === other.cf60;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf55 === other.cf55;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf61, other.cf61);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC31(_dafny.ZERO);
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
    static create_DC34(cf62) {
      let $dt = new D14(0);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC34" + "(" + _dafny.toString(this.cf62) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf62, other.cf62);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
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
    static create_DC36(cf64, cf65) {
      let $dt = new D15(0);
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC37(cf66, cf67, cf68, cf69) {
      let $dt = new D15(1);
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC38(cf70, cf71, cf72) {
      let $dt = new D15(2);
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC35(cf63) {
      let $dt = new D15(3);
      $dt.cf63 = cf63;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get is_DC38() { return this.$tag === 2; }
    get is_DC35() { return this.$tag === 3; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf63() { return this.cf63; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC36" + "(" + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC37" + "(" + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ", " + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC38" + "(" + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 3) {
        return "D15.DC35" + "(" + _dafny.toString(this.cf63) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf64, other.cf64) && _dafny.areEqual(this.cf65, other.cf65);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf66 === other.cf66 && _dafny.areEqual(this.cf67, other.cf67) && this.cf68 === other.cf68 && this.cf69 === other.cf69;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf70, other.cf70) && this.cf71 === other.cf71 && _dafny.areEqual(this.cf72, other.cf72);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf63, other.cf63);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC36(new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO);
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
    static create_DC40(cf74, cf75, cf76) {
      let $dt = new D16(0);
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC41(cf77) {
      let $dt = new D16(1);
      $dt.cf77 = cf77;
      return $dt;
    }
    static create_DC42(cf78, cf79, cf80, cf81) {
      let $dt = new D16(2);
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC39(cf73) {
      let $dt = new D16(3);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get is_DC41() { return this.$tag === 1; }
    get is_DC42() { return this.$tag === 2; }
    get is_DC39() { return this.$tag === 3; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC40" + "(" + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC41" + "(" + _dafny.toString(this.cf77) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC42" + "(" + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ", " + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 3) {
        return "D16.DC39" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf74 === other.cf74 && _dafny.areEqual(this.cf75, other.cf75) && _dafny.areEqual(this.cf76, other.cf76);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf77, other.cf77);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf78, other.cf78) && this.cf79 === other.cf79 && this.cf80 === other.cf80 && _dafny.areEqual(this.cf81, other.cf81);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf73 === other.cf73;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC40(false, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC44(cf83) {
      let $dt = new D17(0);
      $dt.cf83 = cf83;
      return $dt;
    }
    static create_DC43(cf82) {
      let $dt = new D17(1);
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC45(cf84) {
      let $dt = new D17(2);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC44() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get is_DC45() { return this.$tag === 2; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC44" + "(" + _dafny.toString(this.cf83) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC43" + "(" + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC45" + "(" + _dafny.toString(this.cf84) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf83, other.cf83);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf82 === other.cf82;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf84, other.cf84);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC44(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC47(cf86, cf87) {
      let $dt = new D18(0);
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      return $dt;
    }
    static create_DC46(cf85) {
      let $dt = new D18(1);
      $dt.cf85 = cf85;
      return $dt;
    }
    get is_DC47() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf85() { return this.cf85; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC47" + "(" + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC46" + "(" + _dafny.toString(this.cf85) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf86, other.cf86) && _dafny.areEqual(this.cf87, other.cf87);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf85, other.cf85);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC47(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC49(cf89, cf90, cf91) {
      let $dt = new D19(0);
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      return $dt;
    }
    static create_DC48(cf88) {
      let $dt = new D19(1);
      $dt.cf88 = cf88;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf88() { return this.cf88; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC49" + "(" + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC48" + "(" + _dafny.toString(this.cf88) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf89 === other.cf89 && this.cf90 === other.cf90 && _dafny.areEqual(this.cf91, other.cf91);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf88, other.cf88);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC49(false, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC51() {
      let $dt = new D20(0);
      return $dt;
    }
    static create_DC50(cf92) {
      let $dt = new D20(1);
      $dt.cf92 = cf92;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get dtor_cf92() { return this.cf92; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC51";
      } else if (this.$tag === 1) {
        return "D20.DC50" + "(" + _dafny.toString(this.cf92) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf92, other.cf92);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC51();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.D21 = class D21 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC52(cf93) {
      let $dt = new D21(0);
      $dt.cf93 = cf93;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get dtor_cf93() { return this.cf93; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC52" + "(" + _dafny.toString(this.cf93) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf93 === other.cf93;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
    }
    static Rtd() {
      return class {
        static get Default() {
          return D21.Default();
        }
      };
    }
  }

  $module.D22 = class D22 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC54(cf95, cf96, cf97, cf98) {
      let $dt = new D22(0);
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      $dt.cf97 = cf97;
      $dt.cf98 = cf98;
      return $dt;
    }
    static create_DC55(cf99, cf100, cf101, cf102, cf103) {
      let $dt = new D22(1);
      $dt.cf99 = cf99;
      $dt.cf100 = cf100;
      $dt.cf101 = cf101;
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      return $dt;
    }
    static create_DC53(cf94) {
      let $dt = new D22(2);
      $dt.cf94 = cf94;
      return $dt;
    }
    get is_DC54() { return this.$tag === 0; }
    get is_DC55() { return this.$tag === 1; }
    get is_DC53() { return this.$tag === 2; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf94() { return this.cf94; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC54" + "(" + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ", " + _dafny.toString(this.cf97) + ", " + _dafny.toString(this.cf98) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC55" + "(" + _dafny.toString(this.cf99) + ", " + _dafny.toString(this.cf100) + ", " + this.cf101.toVerbatimString(true) + ", " + this.cf102.toVerbatimString(true) + ", " + _dafny.toString(this.cf103) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC53" + "(" + _dafny.toString(this.cf94) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf95, other.cf95) && this.cf96 === other.cf96 && this.cf97 === other.cf97 && _dafny.areEqual(this.cf98, other.cf98);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf99, other.cf99) && this.cf100 === other.cf100 && _dafny.areEqual(this.cf101, other.cf101) && _dafny.areEqual(this.cf102, other.cf102) && _dafny.areEqual(this.cf103, other.cf103);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf94, other.cf94);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC54(new _dafny.CodePoint('D'.codePointAt(0)), null, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
        }
      };
    }
  }

  $module.D23 = class D23 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC57(cf105, cf106, cf107, cf108) {
      let $dt = new D23(0);
      $dt.cf105 = cf105;
      $dt.cf106 = cf106;
      $dt.cf107 = cf107;
      $dt.cf108 = cf108;
      return $dt;
    }
    static create_DC56(cf104) {
      let $dt = new D23(1);
      $dt.cf104 = cf104;
      return $dt;
    }
    get is_DC57() { return this.$tag === 0; }
    get is_DC56() { return this.$tag === 1; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf108() { return this.cf108; }
    get dtor_cf104() { return this.cf104; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC57" + "(" + _dafny.toString(this.cf105) + ", " + _dafny.toString(this.cf106) + ", " + _dafny.toString(this.cf107) + ", " + _dafny.toString(this.cf108) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC56" + "(" + _dafny.toString(this.cf104) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf105 === other.cf105 && this.cf106 === other.cf106 && _dafny.areEqual(this.cf107, other.cf107) && _dafny.areEqual(this.cf108, other.cf108);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf104, other.cf104);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC57(false, [], _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D23.Default();
        }
      };
    }
  }

  $module.D24 = class D24 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC59() {
      let $dt = new D24(0);
      return $dt;
    }
    static create_DC58(cf109) {
      let $dt = new D24(1);
      $dt.cf109 = cf109;
      return $dt;
    }
    get is_DC59() { return this.$tag === 0; }
    get is_DC58() { return this.$tag === 1; }
    get dtor_cf109() { return this.cf109; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC59";
      } else if (this.$tag === 1) {
        return "D24.DC58" + "(" + _dafny.toString(this.cf109) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf109, other.cf109);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC59();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D24.Default();
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

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.ZERO;
      this.f3 = _dafny.ZERO;
      this.f8 = _dafny.MultiSet.Empty;
      this.f13 = _dafny.Map.Empty;
      this.f15 = _dafny.ZERO;
      this.f18 = _dafny.ZERO;
      this.f21 = _dafny.ZERO;
      this.f24 = _dafny.ZERO;
      this.f26 = false;
      this._f1 = _dafny.Seq.UnicodeFromString("");
      this._f2 = [];
      this._f4 = _dafny.ZERO;
      this._f5 = _dafny.ZERO;
      this._f6 = false;
      this._f7 = false;
      this._f9 = false;
      this._f10 = [];
      this._f11 = _dafny.ZERO;
      this._f12 = _dafny.ZERO;
      this._f14 = false;
      this._f16 = false;
      this._f17 = _dafny.Set.Empty;
      this._f19 = _dafny.ZERO;
      this._f20 = false;
      this._f22 = _dafny.Map.Empty;
      this._f23 = [];
      this._f25 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this).f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this).f21 = f21;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
      (_this).f24 = f24;
      (_this)._f25 = f25;
      (_this).f26 = f26;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f36 = _dafny.ZERO;
      this._f37 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f36, f37) {
      let _this = this;
      (_this)._f36 = f36;
      (_this)._f37 = f37;
      return;
    }
    get f36() {
      let _this = this;
      return _this._f36;
    };
    get f37() {
      let _this = this;
      return _this._f37;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f29 = _dafny.MultiSet.Empty;
      this._f30 = _dafny.Seq.of();
      this._f27 = _dafny.ZERO;
      this._f28 = [];
      this.f35 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f35, f29, f30, f27, f28) {
      let _this = this;
      (_this).f35 = f35;
      (_this)._f29 = f29;
      (_this)._f30 = f30;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    fm3(p0, globalState) {
      let _this = this;
      return !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(784)), function (_758_i0) {
        return (_this).f27;
      }), (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)))).length)), new BigNumber(((_this).f30).length))));
    };
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((_this).f27, (_this).f27);
    };
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f27)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f35,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("tcfvt")).length),_this.f35)).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f35,new BigNumber((_dafny.Seq.UnicodeFromString("istggpsd")).length)));
    };
    fm11(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _this.f35;
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _759_v0;
      let _nw127 = Array((new BigNumber(22)).toNumber()).fill(false);
      _759_v0 = _nw127;
      _759_v0 = _759_v0;
      let _760_v1;
      _760_v1 = _dafny.Seq.UnicodeFromString("ltinafpk");
      let _761_v2;
      _761_v2 = new _dafny.CodePoint('j'.codePointAt(0));
      let _762_v3;
      _762_v3 = _dafny.Set.fromElements(_760_v1, _dafny.Seq.of(_761_v2), _760_v1, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(595)), ((_763_v2) => function (_764_i0) {
        return _763_v2;
      })(_761_v2)), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f27), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(595)), ((_765_v2) => function (_766_i0) {
        return _765_v2;
      })(_761_v2))).length)), _761_v2), _760_v1);
      let _767_v4;
      _767_v4 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,_this.f35);
      let _768_v5;
      _768_v5 = _dafny.Seq.of(_this.f35);
      let _769_v6;
      _769_v6 = _dafny.Seq.of(_this.f35, _this.f35, (((_767_v4).contains(_this.f35)) ? ((_767_v4).get(_this.f35)) : (_this.f35)), _this.f35, (_768_v5)[_module.__default.safeIndex((_this).f27, new BigNumber((_768_v5).length))]);
      let _rhs87 = _760_v1;
      let _rhs88 = !(((_this).f27).isLessThanOrEqualTo((_this).f27));
      let _rhs89 = _this.f35;
      let _rhs90 = _module.__default.safeDivisionInt(new BigNumber((_760_v1).length), (_this).f27);
      let _rhs91 = ((_this).f27).isLessThan((_this).fm4(_762_v3, _this.f35, _768_v5, globalState));
      let _lhs63 = globalState;
      let _lhs64 = _this;
      let _lhs65 = globalState;
      _760_v1 = _rhs87;
      _lhs63.f26 = _rhs88;
      _lhs64.f35 = _rhs89;
      r0 = _rhs90;
      _lhs65.f26 = _rhs91;
      let _770_v7;
      let _nw128 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
      _770_v7 = _nw128;
      let _771_v8;
      _771_v8 = _dafny.Set.fromElements(_759_v0, _759_v0);
      let _index71 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_759_v0).length));
      (_759_v0)[_index71] = (_dafny.Set.fromElements(_759_v0)).IsDisjointFrom(_771_v8);
      let _772_v9;
      let _nw129 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
      _772_v9 = _nw129;
      let _index72 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_772_v9).length));
      (_772_v9)[_index72] = (_this).f27;
      let _index73 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_759_v0).length));
      let _index74 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_772_v9).length));
      let _rhs92 = _770_v7;
      let _rhs93 = ((_this.f35) ? (_this.f35) : (((_this.f35) ? (_this.f35) : (_this.f35))));
      let _rhs94 = _761_v2;
      let _rhs95 = new BigNumber(-114);
      let _rhs96 = (_this).f27;
      let _lhs66 = _759_v0;
      let _lhs67 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_759_v0).length));
      let _lhs68 = _772_v9;
      let _lhs69 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_772_v9).length));
      _770_v7 = _rhs92;
      _lhs66[_lhs67] = _rhs93;
      _761_v2 = _rhs94;
      _lhs68[_lhs69] = _rhs95;
      r2 = _rhs96;
      let _773_v10;
      _773_v10 = _dafny.Set.fromElements((_this).f27);
      let _774_v11;
      _774_v11 = _dafny.Map.Empty.slice().updateUnsafe((_772_v9)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_772_v9).length))],_773_v10);
      let _775_v12;
      _775_v12 = _dafny.Map.Empty.slice().updateUnsafe((_759_v0)[_module.__default.safeIndex(new BigNumber(498), new BigNumber((_759_v0).length))],new BigNumber(145));
      _774_v11 = (_774_v11).update(((_this).f27).plus(new BigNumber((_775_v12).length)), _dafny.Set.fromElements((_772_v9)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_772_v9).length))], (_772_v9)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_772_v9).length))], (_this).f27));
      let _776_v13;
      _776_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_760_v1);
      let _777_v14;
      _777_v14 = (_759_v0)[_module.__default.safeIndex(new BigNumber(498), new BigNumber((_759_v0).length))];
      let _index75 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_772_v9).length));
      let _index76 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_759_v0).length));
      let _rhs97 = (_this).f27;
      let _rhs98 = _module.__default.fm12(_module.__default.safeDivisionInt((_this).f27, (_772_v9)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_772_v9).length))]), (_759_v0)[_module.__default.safeIndex(new BigNumber(498), new BigNumber((_759_v0).length))], globalState);
      let _rhs99 = _776_v13;
      let _rhs100 = _dafny.areEqual(_760_v1, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-42)), ((_778_v1) => function (_779_i1) {
        return (_778_v1)[_module.__default.safeIndex((_this).f27, new BigNumber((_778_v1).length))];
      })(_760_v1)), _760_v1));
      let _rhs101 = (_777_v14);
      let _lhs70 = _772_v9;
      let _lhs71 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_772_v9).length));
      let _lhs72 = _this;
      let _lhs73 = _759_v0;
      let _lhs74 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_759_v0).length));
      _lhs70[_lhs71] = _rhs97;
      _761_v2 = _rhs98;
      _776_v13 = _rhs99;
      _lhs72.f35 = _rhs100;
      _lhs73[_lhs74] = _rhs101;
      let _780_v15;
      let _nw130 = new _module.C0();
      _nw130.__ctor((_772_v9)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_772_v9).length))], _760_v1);
      _780_v15 = _nw130;
      r0 = ((_dafny.ZERO).minus(((_this).f30)[_module.__default.safeIndex((_772_v9)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_772_v9).length))], new BigNumber(((_this).f30).length))])).multipliedBy(new BigNumber((_760_v1).length));
      r1 = (_772_v9)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_772_v9).length))];
      r2 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_772_v9)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_772_v9).length))], (_780_v15).f36));
      return [r0, r1, r2];
    }
    m5(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.Map.Empty;
      let _781_v0;
      _781_v0 = _dafny.Seq.of(p1, p0);
      let _782_v1;
      _782_v1 = _dafny.Set.fromElements((_this).f27, (_this).f27);
      (globalState).f21 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_781_v0, _781_v0)).length), new BigNumber((_782_v1).length));
      let _783_i0;
      _783_i0 = _dafny.ZERO;
      L3: {
        while (p0) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_783_i0)) {
              break L3;
            }
            _783_i0 = (_783_i0).plus(_dafny.ONE);
            if (true) {
              r0 = !(p0);
              let _784_v2;
              _784_v2 = _dafny.Seq.UnicodeFromString("pbfgu");
              _784_v2 = _784_v2;
              let _785_v3;
              _785_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,p1);
              let _786_v4;
              _786_v4 = _784_v2;
              _785_v3 = (_785_v3).update(_module.__default.fm0((_this).f27, globalState), !((_this).fm11(p1, (_this).f27, _786_v4, (_this).f27, globalState)));
              (globalState).f26 = ((!(_dafny.Seq.IsPrefixOf((_this).f30, (_this).f30))) ? (_this.f35) : (_this.f35));
              (globalState).f24 = ((_this).f27).multipliedBy((_this).f27);
            } else {
              let _787_v5;
              _787_v5 = _dafny.Seq.UnicodeFromString("ddutqmb");
              let _788_v6;
              _788_v6 = new _dafny.CodePoint('m'.codePointAt(0));
              let _789_v7;
              let _nw131 = new _module.C0();
              _nw131.__ctor(new BigNumber(-849), _dafny.Seq.update(_787_v5, _module.__default.safeIndex((_this).f27, new BigNumber((_787_v5).length)), _788_v6));
              _789_v7 = _nw131;
              (globalState).f3 = ((_this).f27).minus(_module.__default.fm1(globalState));
              _787_v5 = _dafny.Seq.UnicodeFromString("o");
              let _790_v8;
              let _nw132 = new _module.C0();
              _nw132.__ctor(_module.__default.safeDivisionInt(new BigNumber(735), (_dafny.ZERO).minus((_this).f27)), (_789_v7).f37);
              _790_v8 = _nw132;
              _781_v0 = _module.__default.fm13((_this).f27, ((p0) ? ((_dafny.ZERO).minus((_790_v8).f36)) : ((_789_v7).f36)), p0, new BigNumber(-249), globalState);
            }
            let _791_v9;
            let _nw133 = new _module.C0();
            _nw133.__ctor((_this).f27, _dafny.Seq.UnicodeFromString("fyxj"));
            _791_v9 = _nw133;
            let _792_v10;
            _792_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f27);
            (globalState).f21 = _module.__default.safeModuloInt((_this).f27, (new BigNumber(301)).plus((((_792_v10).contains(p0)) ? ((_792_v10).get(p0)) : ((_this).f27))));
            let _793_v11;
            _793_v11 = _dafny.Seq.UnicodeFromString("yxypbryxn");
            let _794_v12;
            let _nw134 = Array((new BigNumber(6)).toNumber());
            _nw134[(_dafny.ZERO).toNumber()] = ((_791_v9).f36).minus((_this).f27);
            _nw134[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(373)), ((_795_v9) => function (_796_i1) {
              return (_795_v9).f36;
            })(_791_v9)), (_this).f30)).length);
            _nw134[(new BigNumber(2)).toNumber()] = new BigNumber((_781_v0).length);
            _nw134[(new BigNumber(3)).toNumber()] = new BigNumber(-645);
            _nw134[(new BigNumber(4)).toNumber()] = (_this).f27;
            _nw134[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f27, (_791_v9).f36));
            _794_v12 = _nw134;
            let _index77 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_794_v12).length));
            (_794_v12)[_index77] = (_this).f27;
            let _797_v13;
            _797_v13 = _dafny.Seq.of(_793_v11);
            let _index78 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_794_v12).length));
            let _rhs102 = _dafny.Seq.Concat((_791_v9).f37, _dafny.Seq.Concat((_791_v9).f37, _module.__default.fm14(_this.f35, p0, (_781_v0)[_module.__default.safeIndex((_791_v9).f36, new BigNumber((_781_v0).length))], (_this).f27, globalState)));
            let _rhs103 = (_this).fm4((function () {
              let _coll31 = new _dafny.Set();
              for (const _compr_31 of (_797_v13).Elements) {
                let _798_v14 = _compr_31;
                if (_dafny.Seq.contains(_797_v13, _798_v14)) {
                  _coll31.add(_798_v14);
                }
              }
              return _coll31;
            }()).Intersect(function () {
              let _coll32 = new _dafny.Set();
              for (const _compr_32 of (_797_v13).Elements) {
                let _799_v15 = _compr_32;
                if (_dafny.Seq.contains(_797_v13, _799_v15)) {
                  _coll32.add(_799_v15);
                }
              }
              return _coll32;
            }()), false, _781_v0, globalState);
            let _rhs104 = (_this).f27;
            let _rhs105 = !(true) || (!(_module.__default.fm15(new BigNumber((function () {
              let _coll33 = new _dafny.Set();
              for (const _compr_33 of _dafny.IntegerRange(new BigNumber(993), new BigNumber(216))) {
                let _800_v16 = _compr_33;
                if (((new BigNumber(993)).isLessThanOrEqualTo(_800_v16)) && ((_800_v16).isLessThan(new BigNumber(216)))) {
                  _coll33.add((_800_v16).minus((_791_v9).f36));
                }
              }
              return _coll33;
            }()).length), (_791_v9).f36, globalState)).contains((_this).f27));
            let _rhs106 = new BigNumber((_module.__default.fm16((_this).f27, globalState)).length);
            let _lhs75 = _794_v12;
            let _lhs76 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_794_v12).length));
            let _lhs77 = globalState;
            let _lhs78 = _this;
            let _lhs79 = globalState;
            _793_v11 = _rhs102;
            _lhs75[_lhs76] = _rhs103;
            _lhs77.f18 = _rhs104;
            _lhs78.f35 = _rhs105;
            _lhs79.f0 = _rhs106;
          }
        }
      }
      let _801_v17;
      _801_v17 = _dafny.Seq.UnicodeFromString("elhuc");
      let _802_v18;
      _802_v18 = _801_v17;
      let _803_v19;
      _803_v19 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber(((_this).f30).length), globalState),(_this).fm11(p1, (_this).f27, _802_v18, (_this).f27, globalState));
      (globalState).f18 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_782_v1).length), _module.__default.safeModuloInt(new BigNumber((_803_v19).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f27,p0)).length))));
      let _hi3 = (_this).f27;
      for (let _804_i2 = new BigNumber(36); _804_i2.isLessThan(_hi3); _804_i2 = _804_i2.plus(_dafny.ONE)) {
        if (p1) {
          let _805_v20;
          let _nw135 = Array((new BigNumber(4)).toNumber());
          _nw135[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_801_v17, _801_v17);
          _nw135[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(402)), function (_806_i3) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          });
          _nw135[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(743)), function (_807_i4) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          });
          _nw135[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("teeic");
          _805_v20 = _nw135;
          _805_v20 = (_this).f28;
          let _808_v21;
          let _nw136 = Array((new BigNumber(25)).toNumber()).fill(false);
          _808_v21 = _nw136;
          let _index79 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_808_v21).length));
          (_808_v21)[_index79] = p0;
          let _809_v22;
          _809_v22 = _module.D4.create_DC7(_808_v21);
          let _810_v23;
          _810_v23 = _dafny.Seq.of(_808_v21, (_809_v22).dtor_cf12, _808_v21);
          let _811_v24;
          _811_v24 = new _dafny.CodePoint('t'.codePointAt(0));
          let _812_v25;
          let _nw137 = Array((new BigNumber(25)).toNumber());
          _nw137[(_dafny.ZERO).toNumber()] = ((_this.f35) ? (_804_i2) : (new BigNumber((_module.__default.fm14(p1, p1, p0, (_dafny.ZERO).minus((_this).f27), globalState)).length)));
          _nw137[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_this).f27);
          _nw137[(new BigNumber(2)).toNumber()] = (_this).f27;
          _nw137[(new BigNumber(3)).toNumber()] = ((_this).f30)[_module.__default.safeIndex((_this).f27, new BigNumber(((_this).f30).length))];
          _nw137[(new BigNumber(4)).toNumber()] = _804_i2;
          _nw137[(new BigNumber(5)).toNumber()] = new BigNumber((_781_v0).length);
          _nw137[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_804_i2, _804_i2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(351)), function (_813_i5) {
            return new _dafny.CodePoint('q'.codePointAt(0));
          })).length))).length);
          _nw137[(new BigNumber(7)).toNumber()] = _804_i2;
          _nw137[(new BigNumber(8)).toNumber()] = _804_i2;
          _nw137[(new BigNumber(9)).toNumber()] = (_this).f27;
          _nw137[(new BigNumber(10)).toNumber()] = new BigNumber(-750);
          _nw137[(new BigNumber(11)).toNumber()] = (_this).f27;
          _nw137[(new BigNumber(12)).toNumber()] = new BigNumber((((_this).f29).update(false, _module.__default.abs(_804_i2))).cardinality());
          _nw137[(new BigNumber(13)).toNumber()] = new BigNumber((_module.__default.fm17(p1, _this.f35, p0, globalState)).length);
          _nw137[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f27),p0)).length);
          _nw137[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_810_v23, _810_v23)).length);
          _nw137[(new BigNumber(16)).toNumber()] = ((_this.f35) ? (new BigNumber(180)) : (_804_i2));
          _nw137[(new BigNumber(17)).toNumber()] = _804_i2;
          _nw137[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_801_v17, _dafny.Seq.UnicodeFromString("lylquqbbn")), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(651)), function (_814_i6) {
            return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-837)), function (_815_i7) {
              return new _dafny.CodePoint('f'.codePointAt(0));
            })).length);
          })).length), new BigNumber((_dafny.Seq.Concat(_801_v17, _dafny.Seq.UnicodeFromString("lylquqbbn"))).length)), _811_v24)).length));
          _nw137[(new BigNumber(19)).toNumber()] = ((_this).f27).minus(new BigNumber(((_this).f30).length));
          _nw137[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus((_this).f27);
          _nw137[(new BigNumber(21)).toNumber()] = new BigNumber(((_this).f29).cardinality());
          _nw137[(new BigNumber(22)).toNumber()] = _804_i2;
          _nw137[(new BigNumber(23)).toNumber()] = new BigNumber(669);
          _nw137[(new BigNumber(24)).toNumber()] = (_this).f27;
          _812_v25 = _nw137;
          let _816_v26;
          _816_v26 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_781_v0, _781_v0),_this.f35);
          let _index80 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_808_v21).length));
          let _rhs107 = new BigNumber((_816_v26).length);
          let _rhs108 = p1;
          let _rhs109 = _812_v25;
          let _lhs80 = globalState;
          let _lhs81 = _808_v21;
          let _lhs82 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_808_v21).length));
          _lhs80.f0 = _rhs107;
          _lhs81[_lhs82] = _rhs108;
          _812_v25 = _rhs109;
          _811_v24 = _811_v24;
          let _817_v27;
          let _818_v28;
          let _out31;
          let _out32;
          let _outcollector13 = _module.__default.m0(!(!(!(true))), new BigNumber(470), (_this).f27, new BigNumber(253), globalState);
          _out31 = _outcollector13[0];
          _out32 = _outcollector13[1];
          _817_v27 = _out31;
          _818_v28 = _out32;
          let _819_v29;
          _819_v29 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? (_817_v27) : (p1)),_802_v18);
          _819_v29 = (_819_v29).update(_this.f35, _802_v18);
        } else {
          r0 = p1;
          let _820_v30;
          _820_v30 = new _dafny.CodePoint('w'.codePointAt(0));
          let _821_v31;
          _821_v31 = _dafny.MultiSet.fromElements((_this).f27, (_dafny.ZERO).minus(_804_i2), (_this).f27);
          let _822_v32;
          _822_v32 = _dafny.Map.Empty.slice().updateUnsafe(_821_v31,_804_i2);
          let _rhs110 = new _dafny.CodePoint('e'.codePointAt(0));
          let _rhs111 = _dafny.Seq.IsPrefixOf(_801_v17, _801_v17);
          let _rhs112 = (((_781_v0)[_module.__default.safeIndex(_804_i2, new BigNumber((_781_v0).length))]) ? (!(((_dafny.ZERO).minus((_module.D1.create_DC2(_820_v30, new BigNumber(168))).dtor_cf3)).isEqualTo((_this).f27))) : (p0));
          let _rhs113 = ((_dafny.Seq.IsProperPrefixOf(_801_v17, _801_v17)) ? (_804_i2) : ((_this).f27));
          let _rhs114 = _822_v32;
          let _lhs83 = globalState;
          let _lhs84 = globalState;
          let _lhs85 = globalState;
          _820_v30 = _rhs110;
          _lhs83.f26 = _rhs111;
          _lhs84.f26 = _rhs112;
          _lhs85.f18 = _rhs113;
          _822_v32 = _rhs114;
          (globalState).f26 = (p1) || (p0);
          (globalState).f21 = _module.__default.safeModuloInt((_this).f27, new BigNumber((function () {
            let _coll34 = new _dafny.Map();
            for (const _compr_34 of _dafny.IntegerRange(new BigNumber(750), new BigNumber(86))) {
              let _823_v33 = _compr_34;
              if (((new BigNumber(750)).isLessThanOrEqualTo(_823_v33)) && ((_823_v33).isLessThan(new BigNumber(86)))) {
                _coll34.push([(_823_v33).minus(_804_i2),_dafny.Map.Empty.slice().updateUnsafe(_this.f35,_804_i2)]);
              }
            }
            return _coll34;
          }()).length));
          (globalState).f15 = (new BigNumber((_801_v17).length)).plus((_804_i2).multipliedBy(new BigNumber(-170)));
        }
        let _824_v34;
        _824_v34 = new _dafny.CodePoint('c'.codePointAt(0));
        let _825_v35;
        _825_v35 = _dafny.Map.Empty.slice().updateUnsafe(_824_v34,p1);
        let _826_v36;
        let _nw138 = Array((new BigNumber(4)).toNumber());
        _nw138[(_dafny.ZERO).toNumber()] = !(p0);
        _nw138[(_dafny.ONE).toNumber()] = (((_825_v35).contains(_824_v34)) ? ((_825_v35).get(_824_v34)) : (!(_this.f35)));
        _nw138[(new BigNumber(2)).toNumber()] = (new BigNumber((_dafny.Seq.UnicodeFromString("ihcoyjmp")).length)).isLessThanOrEqualTo((_this).f27);
        _nw138[(new BigNumber(3)).toNumber()] = false;
        _826_v36 = _nw138;
        r1 = _826_v36;
        let _index81 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_826_v36).length));
        (_826_v36)[_index81] = _this.f35;
        let _827_v37;
        _827_v37 = _dafny.MultiSet.fromElements(_module.__default.fm16((_this).f27, globalState), (_this).f30);
        let _index82 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_826_v36).length));
        (_826_v36)[_index82] = (_module.__default.fm18(globalState)).IsSubsetOf(_827_v37);
        let _828_v38;
        _828_v38 = _dafny.Map.Empty.slice().updateUnsafe((_826_v36)[_module.__default.safeIndex(new BigNumber(806), new BigNumber((_826_v36).length))],(_this).f27);
        let _829_v39;
        _829_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_dafny.ZERO).minus((_804_i2).minus((((_828_v38).contains(_this.f35)) ? ((_828_v38).get(_this.f35)) : ((_this).f27)))));
        _829_v39 = (_829_v39).update((_this).f28, (_this).f27);
      }
      if ((((_this).f30)[_module.__default.safeIndex(new BigNumber(2), new BigNumber(((_this).f30).length))]).isLessThanOrEqualTo(_module.__default.safeModuloInt((_this).f27, new BigNumber(604)))) {
        let _830_v40;
        _830_v40 = _module.D4.create_DC9((_this).f27, p0);
        let _831_v41;
        _831_v41 = _dafny.Set.fromElements((_830_v40).dtor_cf17);
        let _832_v42;
        _832_v42 = _dafny.Seq.of(_dafny.Set.fromElements(p1));
        let _833_v43;
        _833_v43 = _dafny.MultiSet.fromElements((_this).f27);
        let _834_v44;
        _834_v44 = _dafny.MultiSet.fromElements((_this).f27, new BigNumber((_782_v1).length), new BigNumber((_833_v43).cardinality()), (_this).f27, (_this).f27);
        let _835_v45;
        _835_v45 = _833_v43;
        let _836_v46;
        _836_v46 = _dafny.Set.fromElements((_835_v45));
        let _rhs115 = _dafny.Set.fromElements(p0, (new BigNumber(314)).isLessThan((_dafny.ZERO).minus(new BigNumber((_832_v42).length))), !((_836_v46).IsDisjointFrom(_836_v46)), _module.__default.fm0(new BigNumber((_831_v41).length), globalState), p0);
        let _rhs116 = (_this).f27;
        let _lhs86 = globalState;
        _831_v41 = _rhs115;
        _lhs86.f3 = _rhs116;
        let _837_v47;
        let _nw139 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _837_v47 = _nw139;
        _837_v47 = _837_v47;
        let _838_v48;
        let _nw140 = Array((new BigNumber(24)).toNumber()).fill(false);
        _838_v48 = _nw140;
        r1 = _838_v48;
        let _839_v49;
        let _nw141 = Array((new BigNumber(27)).toNumber()).fill(_module.D2.Default());
        _839_v49 = _nw141;
        let _840_v50;
        _840_v50 = new _dafny.CodePoint('v'.codePointAt(0));
        let _841_v51;
        _841_v51 = _dafny.Map.Empty.slice().updateUnsafe(_840_v50,p1);
        let _842_v52;
        _842_v52 = _module.D2.create_DC3(_841_v51);
        let _index83 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_839_v49).length));
        (_839_v49)[_index83] = _842_v52;
        let _index84 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_839_v49).length));
        (_839_v49)[_index84] = _842_v52;
        r0 = _module.__default.fm0(((((_this).f29).contains(p1)) ? (((_this).f29).get(p1)) : ((_this).f27)), globalState);
      } else {
        let _843_v53;
        let _nw142 = Array((new BigNumber(17)).toNumber()).fill(false);
        _843_v53 = _nw142;
        let _index85 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_843_v53).length));
        (_843_v53)[_index85] = (_781_v0)[_module.__default.safeIndex((_this).f27, new BigNumber((_781_v0).length))];
        let _index86 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_843_v53).length));
        (_843_v53)[_index86] = false;
        if (((_this).f27).isEqualTo((_this).f27)) {
          let _844_v54;
          _844_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,new BigNumber(962));
          let _845_v55;
          _845_v55 = new _dafny.CodePoint('o'.codePointAt(0));
          let _846_v56;
          _846_v56 = _dafny.Set.fromElements(_801_v17, _dafny.Seq.update(_801_v17, _module.__default.safeIndex(new BigNumber((_844_v54).length), new BigNumber((_801_v17).length)), _845_v55));
          let _847_v57;
          let _nw143 = Array((new BigNumber(8)).toNumber());
          _nw143[(_dafny.ZERO).toNumber()] = (_this).f27;
          _nw143[(_dafny.ONE).toNumber()] = new BigNumber(((_this).f30).length);
          _nw143[(new BigNumber(2)).toNumber()] = new BigNumber(818);
          _nw143[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f27)), (_this).f27);
          _nw143[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_this).f27);
          _nw143[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_this).f27);
          _nw143[(new BigNumber(6)).toNumber()] = (_this).fm4(_846_v56, p0, _781_v0, globalState);
          _nw143[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(((_this).f27).multipliedBy((_this).f27));
          _847_v57 = _nw143;
          let _index87 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_847_v57).length));
          (_847_v57)[_index87] = (_dafny.ZERO).minus((_this).f27);
          let _index88 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_847_v57).length));
          (_847_v57)[_index88] = _module.__default.safeModuloInt((_this).f27, (_this).f27);
          (globalState).f26 = !(_this.f35);
          let _index89 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_847_v57).length));
          let _rhs117 = new BigNumber((_dafny.Seq.of(((_this).f27).isLessThan((_this).f27), p1)).length);
          let _rhs118 = (((_this).f30)[_module.__default.safeIndex((_847_v57)[_module.__default.safeIndex(new BigNumber(767), new BigNumber((_847_v57).length))], new BigNumber(((_this).f30).length))]).multipliedBy((_847_v57)[_module.__default.safeIndex(new BigNumber(767), new BigNumber((_847_v57).length))]);
          let _rhs119 = _dafny.Seq.UnicodeFromString("vvpnx");
          let _lhs87 = globalState;
          let _lhs88 = _847_v57;
          let _lhs89 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_847_v57).length));
          _lhs87.f24 = _rhs117;
          _lhs88[_lhs89] = _rhs118;
          _801_v17 = _rhs119;
          (globalState).f3 = (_847_v57)[_module.__default.safeIndex(new BigNumber(767), new BigNumber((_847_v57).length))];
          let _848_v58;
          _848_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_803_v19);
          let _rhs120 = !((_module.__default.fm1(globalState)).isLessThanOrEqualTo((new BigNumber((_801_v17).length)).plus((_847_v57)[_module.__default.safeIndex(new BigNumber(767), new BigNumber((_847_v57).length))])));
          let _rhs121 = ((((_848_v58).contains((_this).f27)) ? ((_848_v58).get((_this).f27)) : (_803_v19))).Merge(_803_v19);
          let _lhs90 = globalState;
          _lhs90.f26 = _rhs120;
          _803_v19 = _rhs121;
        } else {
          let _849_v59;
          _849_v59 = _dafny.Set.fromElements(p1);
          let _850_v60;
          _850_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.Set.fromElements(p0)).Difference(_849_v59));
          _850_v60 = _850_v60;
          (globalState).f24 = _module.__default.safeModuloInt((_this).f27, (_this).f27);
          r1 = _843_v53;
          let _851_v61;
          let _nw144 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Set.Empty);
          _851_v61 = _nw144;
          let _index90 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_851_v61).length));
          (_851_v61)[_index90] = _782_v1;
          let _852_v62;
          _852_v62 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f27);
          let _index91 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_851_v61).length));
          (_851_v61)[_index91] = function () {
            let _coll35 = new _dafny.Set();
            for (const _compr_35 of (_852_v62).Keys.Elements) {
              let _853_v63 = _compr_35;
              if ((_852_v62).contains(_853_v63)) {
                _coll35.add((_853_v63).minus((new BigNumber(782)).multipliedBy(new BigNumber(566))));
              }
            }
            return _coll35;
          }();
          let _854_v64;
          _854_v64 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,(_this).f27);
          let _855_v65;
          _855_v65 = _dafny.Seq.of(_854_v64, (_854_v64).Merge(_854_v64), _dafny.Map.Empty.slice().updateUnsafe(_this.f35,new BigNumber((_801_v17).length)), _854_v64);
          _855_v65 = _dafny.Seq.of(_854_v64);
        }
        let _856_v66;
        _856_v66 = new _dafny.CodePoint('y'.codePointAt(0));
        let _857_v67;
        let _nw145 = new _module.C0();
        _nw145.__ctor((_this).f27, _dafny.Seq.update(_801_v17, _module.__default.safeIndex((_this).f27, new BigNumber((_801_v17).length)), _856_v66));
        _857_v67 = _nw145;
        let _858_v68;
        _858_v68 = _module.D6.create_DC11(_857_v67);
        let _859_v69;
        let _nw146 = Array((new BigNumber(23)).toNumber());
        _nw146[(_dafny.ZERO).toNumber()] = _857_v67;
        _nw146[(_dafny.ONE).toNumber()] = _857_v67;
        _nw146[(new BigNumber(2)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(3)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(4)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(5)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(6)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(7)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(8)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(9)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(10)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(11)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(12)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(13)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(14)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(15)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(16)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(17)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(18)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(19)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(20)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(21)).toNumber()] = _857_v67;
        _nw146[(new BigNumber(22)).toNumber()] = (_858_v68).dtor_cf19;
        _859_v69 = _nw146;
        let _rhs122 = _859_v69;
        let _rhs123 = _dafny.Seq.Concat(_dafny.Seq.update((_857_v67).f37, _module.__default.safeIndex((_dafny.ZERO).minus((_857_v67).f36), new BigNumber(((_857_v67).f37).length)), new _dafny.CodePoint('b'.codePointAt(0))), (_857_v67).f37);
        let _rhs124 = (_857_v67).f36;
        let _lhs91 = globalState;
        _859_v69 = _rhs122;
        _801_v17 = _rhs123;
        _lhs91.f24 = _rhs124;
        let _860_v70;
        _860_v70 = _dafny.Map.Empty.slice().updateUnsafe((_857_v67).f37,(_dafny.ZERO).minus((_this).f27));
        _860_v70 = (_860_v70).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(935)), ((_861_v66) => function (_862_i8) {
          return _861_v66;
        })(_856_v66)), (_857_v67).f36);
        let _863_v71;
        _863_v71 = _dafny.Map.Empty.slice().updateUnsafe((_843_v53)[_module.__default.safeIndex(new BigNumber(937), new BigNumber((_843_v53).length))],(_857_v67).f36);
        _863_v71 = (_863_v71).update((_843_v53)[_module.__default.safeIndex(new BigNumber(937), new BigNumber((_843_v53).length))], (new BigNumber(-461)).multipliedBy(new BigNumber((_803_v19).length)));
      }
      if ((_781_v0)[_module.__default.safeIndex((_this).f27, new BigNumber((_781_v0).length))]) {
        (globalState).f26 = p1;
        (globalState).f3 = new BigNumber((((p0) ? ((_module.D7.create_DC13(_dafny.Seq.of(_this.f35, true))).dtor_cf23) : (_781_v0))).length);
        (globalState).f0 = (_this).f27;
        if (_dafny.areEqual(_module.__default.fm14(_this.f35, _this.f35, p1, (_this).f27, globalState), _dafny.Seq.Concat(_801_v17, _dafny.Seq.UnicodeFromString("qk")))) {
          (globalState).f15 = (_this).f27;
          let _864_v72;
          let _nw147 = new _module.C0();
          _nw147.__ctor((_this).f27, _801_v17);
          _864_v72 = _nw147;
          (globalState).f26 = _this.f35;
          let _865_v73;
          let _nw148 = Array((new BigNumber(22)).toNumber()).fill(false);
          _865_v73 = _nw148;
          let _index92 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_865_v73).length));
          (_865_v73)[_index92] = p0;
          let _866_v74;
          _866_v74 = _module.D6.create_DC11(_864_v72);
          let _867_v75;
          _867_v75 = _dafny.Map.Empty.slice().updateUnsafe(p1,_866_v74);
          let _868_v76;
          _868_v76 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_867_v75).length)).multipliedBy((_this).f27),true);
          let _index93 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_865_v73).length));
          (_865_v73)[_index93] = (((_868_v76).contains(new BigNumber(341))) ? ((_868_v76).get(new BigNumber(341))) : (_this.f35));
          let _869_v77;
          _869_v77 = _dafny.Set.fromElements(_dafny.MultiSet.FromArray(_781_v0));
          let _index94 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_865_v73).length));
          let _rhs125 = p0;
          let _rhs126 = !((_869_v77).IsSubsetOf(_869_v77));
          let _rhs127 = ((_this).f27).multipliedBy((_864_v72).f36);
          let _lhs92 = _865_v73;
          let _lhs93 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_865_v73).length));
          let _lhs94 = globalState;
          let _lhs95 = globalState;
          _lhs92[_lhs93] = _rhs125;
          _lhs94.f26 = _rhs126;
          _lhs95.f24 = _rhs127;
        } else {
          let _870_v78;
          _870_v78 = _module.D8.create_DC15(_803_v19);
          let _871_v79;
          _871_v79 = _dafny.Map.Empty.slice().updateUnsafe((_870_v78).dtor_cf28,p0);
          _781_v0 = _dafny.Seq.Concat(_dafny.Seq.of((((_871_v79).contains(_dafny.Map.Empty.slice().updateUnsafe(_this.f35,p0))) ? ((_871_v79).get(_dafny.Map.Empty.slice().updateUnsafe(_this.f35,p0))) : (p0))), _dafny.Seq.update(_781_v0, _module.__default.safeIndex((_this).f27, new BigNumber((_781_v0).length)), p1));
          let _872_v80;
          _872_v80 = new _dafny.CodePoint('u'.codePointAt(0));
          let _873_v81;
          _873_v81 = _dafny.MultiSet.fromElements(((p0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-284)), function (_874_i9) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          })) : (_dafny.Seq.update(_801_v17, _module.__default.safeIndex((_this).f27, new BigNumber((_801_v17).length)), _872_v80))));
          (globalState).f3 = (((_873_v81).contains(_dafny.Seq.Concat(_801_v17, _801_v17))) ? ((_873_v81).get(_dafny.Seq.Concat(_801_v17, _801_v17))) : ((_dafny.ZERO).minus(((!(_this.f35)) ? ((_this).f27) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(702)), ((_875_v0) => function (_876_i10) {
            return new BigNumber((_875_v0).length);
          })(_781_v0))).length))))));
          (globalState).f18 = (_this).f27;
          (_this).f35 = !(_this.f35);
          let _877_v82;
          _877_v82 = _dafny.MultiSet.fromElements((_this).f27, new BigNumber((_801_v17).length), (_this).f27);
          _877_v82 = _877_v82;
        }
        let _878_v83;
        let _nw149 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
        _878_v83 = _nw149;
        let _index95 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_878_v83).length));
        (_878_v83)[_index95] = function () {
          let _coll36 = new _dafny.Map();
          for (const _compr_36 of (_module.__default.fm19(globalState)).Elements) {
            let _879_v84 = _compr_36;
            if ((_module.__default.fm19(globalState)).contains(_879_v84)) {
              _coll36.push([_879_v84,(_this).f27]);
            }
          }
          return _coll36;
        }();
        let _880_v85;
        _880_v85 = new _dafny.CodePoint('t'.codePointAt(0));
        let _881_v86;
        _881_v86 = _dafny.Map.Empty.slice().updateUnsafe(_880_v85,(_this).f27);
        let _index96 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_878_v83).length));
        (_878_v83)[_index96] = _881_v86;
      } else {
        (globalState).f0 = (_this).f27;
        (globalState).f18 = (_this).f27;
        let _882_v87;
        _882_v87 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f27);
        let _883_v88;
        _883_v88 = new _dafny.CodePoint('k'.codePointAt(0));
        let _884_v89;
        _884_v89 = _module.D4.create_DC8(_883_v88, new BigNumber(57), p1);
        _882_v87 = (_882_v87).update((_884_v89).dtor_cf14, (_this).f27);
        let _885_v90;
        let _nw150 = Array((new BigNumber(12)).toNumber()).fill([]);
        _885_v90 = _nw150;
        let _886_v91;
        _886_v91 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(208)), function (_887_i11) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        })).length));
        let _rhs128 = _885_v90;
        let _rhs129 = (_this).f27;
        let _rhs130 = (_this).f30;
        let _rhs131 = _module.__default.safeDivisionInt((_this).f27, new BigNumber((_801_v17).length));
        let _rhs132 = (_this).f27;
        let _lhs96 = globalState;
        let _lhs97 = globalState;
        let _lhs98 = globalState;
        _885_v90 = _rhs128;
        _lhs96.f21 = _rhs129;
        _886_v91 = _rhs130;
        _lhs97.f0 = _rhs131;
        _lhs98.f0 = _rhs132;
        _801_v17 = _dafny.Seq.Concat(_801_v17, _dafny.Seq.Concat(_801_v17, _module.__default.fm14(_this.f35, _this.f35, p1, (_this).f27, globalState)));
      }
      r0 = p0;
      let _888_v92;
      _888_v92 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f27),(_this).f27);
      let _889_v93;
      _889_v93 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(globalState),_888_v92);
      let _890_v94;
      _890_v94 = new _dafny.CodePoint('r'.codePointAt(0));
      let _891_v95;
      _891_v95 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,_890_v94);
      let _nw151 = Array((new BigNumber(21)).toNumber());
      _nw151[(_dafny.ZERO).toNumber()] = _this.f35;
      _nw151[(_dafny.ONE).toNumber()] = p1;
      _nw151[(new BigNumber(2)).toNumber()] = (p1) && (p0);
      _nw151[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsPrefixOf(_781_v0, _781_v0);
      _nw151[(new BigNumber(4)).toNumber()] = p0;
      _nw151[(new BigNumber(5)).toNumber()] = p0;
      _nw151[(new BigNumber(6)).toNumber()] = p0;
      _nw151[(new BigNumber(7)).toNumber()] = p1;
      _nw151[(new BigNumber(8)).toNumber()] = ((_this).f27).isLessThanOrEqualTo((_this).f27);
      _nw151[(new BigNumber(9)).toNumber()] = p0;
      _nw151[(new BigNumber(10)).toNumber()] = p0;
      _nw151[(new BigNumber(11)).toNumber()] = p1;
      _nw151[(new BigNumber(12)).toNumber()] = ((_this).f27).isLessThanOrEqualTo(new BigNumber(((_this).f30).length));
      _nw151[(new BigNumber(13)).toNumber()] = (_781_v0)[_module.__default.safeIndex((_this).f27, new BigNumber((_781_v0).length))];
      _nw151[(new BigNumber(14)).toNumber()] = p1;
      _nw151[(new BigNumber(15)).toNumber()] = (_this).fm11(p1, new BigNumber((_889_v93).length), _802_v18, new BigNumber((_891_v95).length), globalState);
      _nw151[(new BigNumber(16)).toNumber()] = (_this).fm11(p1, (_this).f27, _802_v18, (_this).f27, globalState);
      _nw151[(new BigNumber(17)).toNumber()] = (_this).fm11(p0, new BigNumber((_801_v17).length), _802_v18, (_this).f27, globalState);
      _nw151[(new BigNumber(18)).toNumber()] = p1;
      _nw151[(new BigNumber(19)).toNumber()] = (_this).fm11(p1, (_this).f27, _802_v18, new BigNumber((_801_v17).length), globalState);
      _nw151[(new BigNumber(20)).toNumber()] = _this.f35;
      r1 = _nw151;
      let _892_v96;
      _892_v96 = _dafny.MultiSet.fromElements((_this).f27);
      let _893_v97;
      _893_v97 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((((p1) ? (_801_v17) : (_801_v17))).length),!((_892_v96).IsProperSubsetOf(_892_v96)));
      r2 = _893_v97;
      return [r0, r1, r2];
    }
    m6(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _894_v0;
      let _nw152 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
      _894_v0 = _nw152;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_894_v0).length))) {
        let _895_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_895_i0)) && ((_895_i0).isLessThan(new BigNumber((_894_v0).length))))) {
          (_894_v0)[(_895_i0)] = (_895_i0).plus(new BigNumber((_dafny.Set.fromElements(p0)).length));
        }
      }
      (globalState).f24 = ((_this).f30)[_module.__default.safeIndex(p0, new BigNumber(((_this).f30).length))];
      let _896_v1;
      let _nw153 = Array((new BigNumber(4)).toNumber());
      _nw153[(_dafny.ZERO).toNumber()] = _this.f35;
      _nw153[(_dafny.ONE).toNumber()] = _module.__default.fm0((_this).f27, globalState);
      _nw153[(new BigNumber(2)).toNumber()] = _this.f35;
      _nw153[(new BigNumber(3)).toNumber()] = _this.f35;
      _896_v1 = _nw153;
      let _897_v2;
      _897_v2 = _module.D4.create_DC7(_896_v1);
      let _pat_let_tv7 = _896_v1;
      let _898_v3;
      _898_v3 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let14_0) {
        return function (_899_dt__update__tmp_h0) {
          return function (_pat_let15_0) {
            return function (_900_dt__update_hcf12_h0) {
              return _module.D4.create_DC7(_900_dt__update_hcf12_h0);
            }(_pat_let15_0);
          }(_pat_let_tv7);
        }(_pat_let14_0);
      }(_897_v2),_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f35, false, _this.f35, _this.f35, _this.f35)));
      let _901_v4;
      _901_v4 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,new BigNumber((_898_v3).length));
      let _902_v5;
      _902_v5 = _dafny.Seq.UnicodeFromString("it");
      let _903_v6;
      _903_v6 = _dafny.Map.Empty.slice().updateUnsafe(_902_v5,(_this).f27);
      let _904_v7;
      _904_v7 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_module.__default.fm16((((_901_v4).contains(_this.f35)) ? ((_901_v4).get(_this.f35)) : (p0)), globalState), (_this).f30),_module.__default.safeDivisionInt((_this).f27, (((_903_v6).contains(_902_v5)) ? ((_903_v6).get(_902_v5)) : ((_this).f27))));
      _904_v7 = (_904_v7).update(_dafny.Seq.Concat((_this).f30, (_this).f30), (_this).f27);
      _902_v5 = _dafny.Seq.Concat(_module.__default.fm14(true, _this.f35, _this.f35, new BigNumber((_902_v5).length), globalState), _module.__default.fm14(_this.f35, _this.f35, _this.f35, new BigNumber((_module.__default.fm14(_this.f35, _this.f35, _this.f35, (_this).f27, globalState)).length), globalState));
      if (!(((_this.f35) ? (_this.f35) : (_this.f35)))) {
        let _905_v8;
        _905_v8 = new _dafny.CodePoint('w'.codePointAt(0));
        let _906_v9;
        _906_v9 = _module.D2.create_DC4(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f35,false)).length),_this.f35)).length), ((_this).f27).minus(p0), (_dafny.ZERO).minus(p0), _905_v8, _dafny.Seq.Concat(_902_v5, _902_v5));
        let _907_v10;
        _907_v10 = _dafny.Map.Empty.slice().updateUnsafe(_905_v8,_this.f35);
        let _908_v11;
        _908_v11 = _module.D2.create_DC3(_907_v10);
        let _pat_let_tv8 = _907_v10;
        let _909_v12;
        _909_v12 = _module.D2.create_DC3((function (_pat_let16_0) {
  return function (_910_dt__update__tmp_h1) {
    return function (_pat_let17_0) {
      return function (_911_dt__update_hcf4_h0) {
        return _module.D2.create_DC3(_911_dt__update_hcf4_h0);
      }(_pat_let17_0);
    }(_pat_let_tv8);
  }(_pat_let16_0);
}(_908_v11)).dtor_cf4);
        let _912_v13;
        _912_v13 = _module.D2.create_DC5(_909_v12);
        let _rhs133 = _dafny.Seq.Concat(_902_v5, _902_v5);
        let _rhs134 = _906_v9;
        let _rhs135 = (_this).f27;
        let _rhs136 = _912_v13;
        let _rhs137 = false;
        let _lhs99 = globalState;
        let _lhs100 = globalState;
        _902_v5 = _rhs133;
        _906_v9 = _rhs134;
        _lhs99.f21 = _rhs135;
        _912_v13 = _rhs136;
        _lhs100.f26 = _rhs137;
        let _913_v14;
        _913_v14 = _dafny.Set.fromElements(_this.f35, false);
        let _914_v15;
        _914_v15 = _dafny.Seq.of(_913_v14);
        let _rhs138 = _this.f35;
        let _rhs139 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_914_v15, _dafny.Seq.of(_913_v14)), _914_v15), _module.__default.safeIndex((_this).f27, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_914_v15, _dafny.Seq.of(_913_v14)), _914_v15)).length)), _913_v14);
        let _lhs101 = _this;
        _lhs101.f35 = _rhs138;
        _914_v15 = _rhs139;
        (globalState).f26 = _this.f35;
        let _915_v16;
        _915_v16 = _dafny.Seq.of(_module.__default.fm14(_this.f35, _this.f35, _this.f35, p0, globalState), _902_v5);
        let _916_v17;
        _916_v17 = _module.D6.create_DC12(_this.f35, _module.__default.fm20(_this.f35, globalState), (_915_v16)[_module.__default.safeIndex(p0, new BigNumber((_915_v16).length))]);
        let _source14 = _916_v17;
        if (_source14.is_DC12) {
          let _917___mcc_h0 = (_source14).cf20;
          let _918___mcc_h1 = (_source14).cf21;
          let _919___mcc_h2 = (_source14).cf22;
          let _920_cf22 = _919___mcc_h2;
          let _921_cf21 = _918___mcc_h1;
          let _922_cf20 = _917___mcc_h0;
          let _923_v18;
          let _init16 = ((_924_p0) => function (_925_i1) {
            return (_dafny.Set.fromElements(_924_p0)).Difference((_module.D9.create_DC17(_dafny.Set.fromElements(new BigNumber(((_this).f29).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("vlmvh")).length)))).dtor_cf31);
          })(p0);
          let _nw154 = Array((new BigNumber(11)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw154.length); _i0_16++) {
            _nw154[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _923_v18 = _nw154;
          let _rhs140 = ((_this.f35) ? (_923_v18) : (_923_v18));
          let _rhs141 = _905_v8;
          _923_v18 = _rhs140;
          _905_v8 = _rhs141;
          let _926_v19;
          _926_v19 = _dafny.Set.fromElements(_905_v8);
          let _index97 = _module.__default.safeIndex(new BigNumber(350), new BigNumber((_894_v0).length));
          (_894_v0)[_index97] = new BigNumber((_926_v19).length);
          let _index98 = _module.__default.safeIndex(new BigNumber(350), new BigNumber((_894_v0).length));
          (_894_v0)[_index98] = (p0).minus((p0).minus((_this).f27));
          let _927_v20;
          _927_v20 = _dafny.MultiSet.fromElements((_this).f27, p0, (_894_v0)[_module.__default.safeIndex(new BigNumber(350), new BigNumber((_894_v0).length))], (_this).f27, (_this).f27);
          let _928_v21;
          _928_v21 = _module.D8.create_DC16((_this).f27, (_this).f27);
          let _929_v22;
          _929_v22 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f30)[_module.__default.safeIndex(new BigNumber((_module.__default.fm13(p0, new BigNumber(612), !(_922_cf20), new BigNumber((_920_cf22).length), globalState)).length), new BigNumber(((_this).f30).length))],_928_v21);
          let _930_v23;
          _930_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_905_v8,(_this).f27)).length),(_this).f27);
          let _931_v24;
          let _932_v25;
          let _out33;
          let _out34;
          let _outcollector14 = _module.__default.m0((_927_v20).IsDisjointFrom(_dafny.MultiSet.fromElements((_894_v0)[_module.__default.safeIndex(new BigNumber(350), new BigNumber((_894_v0).length))])), new BigNumber((_929_v22).length), (_this).f27, (((_930_v23).contains((_this).f27)) ? ((_930_v23).get((_this).f27)) : (new BigNumber((_902_v5).length))), globalState);
          _out33 = _outcollector14[0];
          _out34 = _outcollector14[1];
          _931_v24 = _out33;
          _932_v25 = _out34;
          let _index99 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_896_v1).length));
          (_896_v1)[_index99] = _922_cf20;
          let _index100 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_896_v1).length));
          (_896_v1)[_index100] = _this.f35;
        } else {
          let _933___mcc_h3 = (_source14).cf19;
          let _934_cf19 = _933___mcc_h3;
          (globalState).f26 = _this.f35;
          let _935_v26;
          _935_v26 = _dafny.MultiSet.fromElements((_934_cf19).f36, new BigNumber(-805));
          let _936_v27;
          _936_v27 = _dafny.Seq.of(_this.f35, _this.f35);
          let _937_v28;
          _937_v28 = _module.D1.create_DC2(_905_v8, (_dafny.ZERO).minus(new BigNumber((_936_v27).length)));
          let _938_v29;
          _938_v29 = _dafny.Set.fromElements(new BigNumber((_935_v26).cardinality()), (_937_v28).dtor_cf3);
          let _939_v30;
          let _940_v31;
          let _941_v32;
          let _out35;
          let _out36;
          let _out37;
          let _outcollector15 = (_this).m5(_this.f35, (_938_v29).IsDisjointFrom(_938_v29), globalState);
          _out35 = _outcollector15[0];
          _out36 = _outcollector15[1];
          _out37 = _outcollector15[2];
          _939_v30 = _out35;
          _940_v31 = _out36;
          _941_v32 = _out37;
          (globalState).f24 = (_this).f27;
          let _index101 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_894_v0).length));
          (_894_v0)[_index101] = ((_this).f27).plus(p0);
          let _index102 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_894_v0).length));
          (_894_v0)[_index102] = _module.__default.safeDivisionInt(new BigNumber((_902_v5).length), p0);
        }
        let _942_v33;
        let _nw155 = Array((new BigNumber(26)).toNumber()).fill([]);
        _942_v33 = _nw155;
        let _index103 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_942_v33).length));
        (_942_v33)[_index103] = _896_v1;
        let _index104 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_942_v33).length));
        (_942_v33)[_index104] = _896_v1;
      } else {
        let _943_v34;
        _943_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(767),_this.f35);
        (globalState).f21 = ((p0).minus((_dafny.ZERO).minus(new BigNumber((_943_v34).length)))).plus((_this).f27);
        (globalState).f26 = (p0).isEqualTo(_module.__default.fm1(globalState));
        let _944_v35;
        let _nw156 = new _module.C0();
        _nw156.__ctor((_this).f27, _902_v5);
        _944_v35 = _nw156;
        let _945_v36;
        _945_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(364),_944_v35);
        let _946_v37;
        _946_v37 = _module.D6.create_DC12(_this.f35, _dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), function (_947_i2) {
  return _dafny.Seq.of(_this.f35);
}), (_944_v35).f37);
        let _948_v38;
        _948_v38 = _module.D6.create_DC11((((_945_v36).contains(new BigNumber(((_946_v37).dtor_cf22).length))) ? ((_945_v36).get(new BigNumber(((_946_v37).dtor_cf22).length))) : (_944_v35)));
        _948_v38 = _948_v38;
        _943_v34 = (_943_v34).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f27,true));
        (globalState).f3 = (_944_v35).f36;
      }
      let _949_v39;
      _949_v39 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,_902_v5);
      let _950_v40;
      let _nw157 = new _module.C0();
      _nw157.__ctor((_this).f27, (((_949_v39).contains(true)) ? ((_949_v39).get(true)) : (_902_v5)));
      _950_v40 = _nw157;
      let _951_v41;
      _951_v41 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,_this.f35);
      let _952_v42;
      _952_v42 = _module.D8.create_DC15(_951_v41);
      let _pat_let_tv9 = p0;
      r0 = function (_source15) {
        if (_source15.is_DC16) {
          let _953___mcc_h4 = (_source15).cf29;
          let _954___mcc_h5 = (_source15).cf30;
          let _955_cf30 = _954___mcc_h5;
          let _956_cf29 = _953___mcc_h4;
          return _956_cf29;
        } else {
          let _957___mcc_h6 = (_source15).cf28;
          let _958_cf28 = _957___mcc_h6;
          return _pat_let_tv9;
        }
      }(_952_v42);
      return r0;
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f31 = _dafny.MultiSet.Empty;
      this._f32 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T2];
    }
    get f31() {
      let _this = this;
      return _this._f31;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
    __ctor(f31, f32) {
      let _this = this;
      (_this)._f31 = f31;
      (_this)._f32 = f32;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return !(function (_source16) {
        if (_source16.is_DC18) {
          let _959___mcc_h0 = (_source16).cf32;
          let _960_cf32 = _959___mcc_h0;
          return (_960_cf32) || (_960_cf32);
        } else if (_source16.is_DC19) {
          let _961___mcc_h1 = (_source16).cf33;
          let _962_cf33 = _961___mcc_h1;
          return (_module.D9.create_DC20(new BigNumber((_dafny.Seq.UnicodeFromString("okdff")).length), _962_cf33, _962_cf33, _962_cf33, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_962_cf33, _962_cf33)).length),new BigNumber((_dafny.Set.fromElements(_962_cf33)).length)))).dtor_cf36;
        } else if (_source16.is_DC20) {
          let _963___mcc_h2 = (_source16).cf34;
          let _964___mcc_h3 = (_source16).cf35;
          let _965___mcc_h4 = (_source16).cf36;
          let _966___mcc_h5 = (_source16).cf37;
          let _967___mcc_h6 = (_source16).cf38;
          let _968_cf38 = _967___mcc_h6;
          let _969_cf37 = _966___mcc_h5;
          let _970_cf36 = _965___mcc_h4;
          let _971_cf35 = _964___mcc_h3;
          let _972_cf34 = _963___mcc_h2;
          return !(_970_cf36);
        } else {
          let _973___mcc_h7 = (_source16).cf31;
          let _974_cf31 = _973___mcc_h7;
          return _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-888)), function (_975_i0) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          }), _dafny.Seq.UnicodeFromString("im"));
        }
      }(_module.D9.create_DC20((_this).f32, !(!(true)), true, false, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-15),new BigNumber(((_dafny.MultiSet.fromElements((_this).f32, (_this).f32))).cardinality())))));
    };
    fm6(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-875)), function (_976_i0) {
        return _dafny.Seq.of(true);
      }), _dafny.Seq.of(_dafny.Seq.of(true, !(true)))), _dafny.Seq.of(_dafny.Seq.of(!(false), true)));
    };
    fm21(globalState) {
      let _this = this;
      return !(!(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f32,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f32,true))).length)).isEqualTo(new BigNumber(((_this).f31).cardinality())));
    };
    fm22(p0, p1, p2, globalState) {
      let _this = this;
      let _source17 = _module.D1.create_DC1(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f32),!(true)));
      if (_source17.is_DC2) {
        let _977___mcc_h0 = (_source17).cf2;
        let _978___mcc_h1 = (_source17).cf3;
        let _979_cf3 = _978___mcc_h1;
        let _980_cf2 = _977___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rdr"), _dafny.Seq.UnicodeFromString("gxqk"));
      } else {
        let _981___mcc_h2 = (_source17).cf1;
        let _982_cf1 = _981___mcc_h2;
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kgvy"), _dafny.Seq.UnicodeFromString("tp"));
      }
    };
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let _983_v0;
      _983_v0 = new _dafny.CodePoint('i'.codePointAt(0));
      _983_v0 = _983_v0;
      let _984_v1;
      let _nw158 = Array((new BigNumber(14)).toNumber()).fill(false);
      _984_v1 = _nw158;
      let _985_v2;
      _985_v2 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
      let _986_v3;
      _986_v3 = _dafny.Map.Empty.slice().updateUnsafe(_984_v1,_985_v2);
      _986_v3 = (_986_v3).update(_984_v1, _985_v2);
      if (p1) {
        (globalState).f26 = _module.__default.fm0(p2, globalState);
        let _987_v4;
        _987_v4 = _dafny.MultiSet.fromElements(_984_v1, _984_v1);
        if ((_987_v4).IsSubsetOf(_987_v4)) {
          (globalState).f24 = new BigNumber(308);
          let _988_v5;
          _988_v5 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f31);
          let _989_v6;
          _989_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
          let _index105 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_984_v1).length));
          (_984_v1)[_index105] = ((_this).f31).IsDisjointFrom((((_988_v5).contains(new BigNumber((_989_v6).length))) ? ((_988_v5).get(new BigNumber((_989_v6).length))) : ((_this).f31)));
          let _990_v7;
          _990_v7 = _dafny.Seq.UnicodeFromString("ojlonqco");
          let _991_v8;
          _991_v8 = _module.D4.create_DC8(_983_v0, new BigNumber((_990_v7).length), p1);
          let _992_v9;
          _992_v9 = _dafny.Set.fromElements(_991_v8);
          let _993_v10;
          _993_v10 = _dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_module.__default.fm23(new BigNumber((_989_v6).length), !(true), globalState)).length)));
          let _994_v11;
          _994_v11 = _dafny.Seq.of((new BigNumber((_990_v7).length)).minus(new BigNumber((_992_v9).length)), (_dafny.ZERO).minus(p3), p3, _module.__default.safeModuloInt(new BigNumber((_990_v7).length), new BigNumber((_993_v10).length)));
          let _995_v12;
          _995_v12 = _dafny.MultiSet.fromElements(_993_v10);
          let _index106 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_984_v1).length));
          let _rhs142 = !_dafny.areEqual(_990_v7, _990_v7);
          let _rhs143 = _dafny.Seq.of(p3, new BigNumber((((p1) ? (_dafny.Seq.of(p2)) : (_994_v11))).length));
          let _rhs144 = (_dafny.ZERO).minus((((_995_v12).contains(_dafny.Map.Empty.slice().updateUnsafe(!(p0),(_dafny.ZERO).minus(p3)))) ? ((_995_v12).get(_dafny.Map.Empty.slice().updateUnsafe(!(p0),(_dafny.ZERO).minus(p3)))) : (_module.__default.safeDivisionInt((_this).f32, p3))));
          let _rhs145 = _984_v1;
          let _rhs146 = (_this).fm21(globalState);
          let _lhs102 = _984_v1;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_984_v1).length));
          let _lhs104 = globalState;
          let _lhs105 = globalState;
          _lhs102[_lhs103] = _rhs142;
          _994_v11 = _rhs143;
          _lhs104.f15 = _rhs144;
          _984_v1 = _rhs145;
          _lhs105.f26 = _rhs146;
          let _996_v13;
          _996_v13 = _990_v7;
          let _997_v14;
          let _nw159 = Array((new BigNumber(14)).toNumber()).fill(false);
          _997_v14 = _nw159;
          let _998_v15;
          _998_v15 = (_984_v1)[_module.__default.safeIndex(new BigNumber(766), new BigNumber((_984_v1).length))];
          let _index107 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_997_v14).length));
          (_997_v14)[_index107] = _998_v15;
          let _index108 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_997_v14).length));
          let _rhs147 = p1;
          let _rhs148 = _996_v13;
          let _rhs149 = p1;
          let _rhs150 = _998_v15;
          let _rhs151 = p0;
          let _lhs106 = globalState;
          let _lhs107 = globalState;
          let _lhs108 = _997_v14;
          let _lhs109 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_997_v14).length));
          let _lhs110 = globalState;
          _lhs106.f26 = _rhs147;
          _996_v13 = _rhs148;
          _lhs107.f26 = _rhs149;
          _lhs108[_lhs109] = _rhs150;
          _lhs110.f26 = _rhs151;
          let _999_v16;
          _999_v16 = _dafny.Map.Empty.slice().updateUnsafe(((false) ? (_990_v7) : (_990_v7)),((_this).f31).Difference((_this).f31));
          _999_v16 = _999_v16;
          let _rhs152 = (_984_v1)[_module.__default.safeIndex(new BigNumber(766), new BigNumber((_984_v1).length))];
          let _rhs153 = (_this).fm22((_993_v10).Merge(_module.__default.fm24(p3, globalState)), false, (((_993_v10).contains((_984_v1)[_module.__default.safeIndex(new BigNumber(766), new BigNumber((_984_v1).length))])) ? ((_993_v10).get((_984_v1)[_module.__default.safeIndex(new BigNumber(766), new BigNumber((_984_v1).length))])) : (new BigNumber(((_this).f31).cardinality()))), globalState);
          let _rhs154 = !(p1);
          let _lhs111 = globalState;
          let _lhs112 = globalState;
          _lhs111.f26 = _rhs152;
          _990_v7 = _rhs153;
          _lhs112.f26 = _rhs154;
        } else {
          let _1000_v17;
          _1000_v17 = _dafny.Seq.UnicodeFromString("vkvarnrk");
          let _1001_v18;
          let _nw160 = new _module.C0();
          _nw160.__ctor((_dafny.ZERO).minus(p3), _1000_v17);
          _1001_v18 = _nw160;
          let _1002_v19;
          _1002_v19 = _module.D6.create_DC11(_1001_v18);
          let _1003_v20;
          _1003_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1002_v19,_dafny.Seq.update(_dafny.Seq.Concat(_1000_v17, _1000_v17), _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_983_v0)).length), new BigNumber((_dafny.Seq.Concat(_1000_v17, _1000_v17)).length)), _983_v0));
          _1003_v20 = (_1003_v20).update(_1002_v19, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gvq"), _1000_v17));
          let _1004_v21;
          _1004_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(886)), ((_1005_v0) => function (_1006_i0) {
            return _1005_v0;
          })(_983_v0)),p0);
          let _1007_v22;
          let _nw161 = new _module.C0();
          _nw161.__ctor((new BigNumber((_1004_v21).length)).plus((_this).f32), (_1001_v18).f37);
          _1007_v22 = _nw161;
          let _1008_v23;
          _1008_v23 = _dafny.Seq.of(p0);
          _1008_v23 = _dafny.Seq.of(false, p1, _dafny.Seq.IsPrefixOf((_1001_v18).f37, _dafny.Seq.UnicodeFromString("oxvupb")), p0, p1);
          let _index109 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_984_v1).length));
          (_984_v1)[_index109] = p1;
          let _1009_v24;
          _1009_v24 = _dafny.Set.fromElements(_984_v1);
          let _index110 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_984_v1).length));
          (_984_v1)[_index110] = (_dafny.Set.fromElements(_984_v1)).IsDisjointFrom(_1009_v24);
          (globalState).f24 = new BigNumber(315);
        }
        let _1010_v25;
        _1010_v25 = _dafny.Seq.of(_984_v1);
        let _1011_v26;
        _1011_v26 = _dafny.Seq.UnicodeFromString("pbbsuc");
        let _1012_v27;
        _1012_v27 = _dafny.MultiSet.fromElements(_1011_v26, _1011_v26);
        let _rhs155 = ((new BigNumber((_1010_v25).length)).minus(p3)).isLessThan((_this).f32);
        let _rhs156 = _1012_v27;
        let _lhs113 = globalState;
        let _lhs114 = globalState;
        _lhs113.f26 = _rhs155;
        _lhs114.f8 = _rhs156;
        (globalState).f26 = !(p0);
        (globalState).f26 = !((_dafny.ZERO).minus(p3)).isEqualTo((_this).f32);
      } else {
        (globalState).f26 = (p2).isLessThanOrEqualTo(p2);
        (globalState).f26 = p1;
        let _1013_v28;
        let _nw162 = Array((new BigNumber(2)).toNumber());
        _nw162[(_dafny.ZERO).toNumber()] = p2;
        _nw162[(_dafny.ONE).toNumber()] = _module.__default.fm1(globalState);
        _1013_v28 = _nw162;
        _1013_v28 = _1013_v28;
        let _1014_v29;
        let _nw163 = Array((new BigNumber(6)).toNumber());
        _nw163[(_dafny.ZERO).toNumber()] = _984_v1;
        _nw163[(_dafny.ONE).toNumber()] = _984_v1;
        _nw163[(new BigNumber(2)).toNumber()] = _984_v1;
        _nw163[(new BigNumber(3)).toNumber()] = _984_v1;
        _nw163[(new BigNumber(4)).toNumber()] = _984_v1;
        _nw163[(new BigNumber(5)).toNumber()] = _984_v1;
        _1014_v29 = _nw163;
        _1014_v29 = _1014_v29;
        let _1015_v30;
        let _nw164 = Array((new BigNumber(13)).toNumber()).fill(_module.D8.Default());
        _1015_v30 = _nw164;
        let _1016_v31;
        _1016_v31 = _module.D8.create_DC16((_this).f32, p3);
        let _index111 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_1015_v30).length));
        (_1015_v30)[_index111] = _1016_v31;
        let _index112 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_1015_v30).length));
        let _rhs157 = _1014_v29;
        let _rhs158 = _1016_v31;
        let _rhs159 = (_dafny.ZERO).minus(((_dafny.ZERO).minus(_module.__default.fm1(globalState))).plus((_dafny.ZERO).minus(p3)));
        let _lhs115 = _1015_v30;
        let _lhs116 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_1015_v30).length));
        _1014_v29 = _rhs157;
        _lhs115[_lhs116] = _rhs158;
        r1 = _rhs159;
      }
      let _1017_v32;
      _1017_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm5(globalState),(_this).f32);
      let _1018_v33;
      _1018_v33 = _dafny.Seq.UnicodeFromString("efdnwcf");
      let _1019_v34;
      _1019_v34 = _module.D2.create_DC4(new BigNumber((_1017_v32).length), new BigNumber(343), p3, _983_v0, _1018_v33);
      let _1020_v35;
      _1020_v35 = _dafny.MultiSet.fromElements(_module.D2.create_DC5(_1019_v34));
      let _1021_v36;
      _1021_v36 = _module.D2.create_DC5(_1019_v34);
      let _1022_v37;
      _1022_v37 = _dafny.Map.Empty.slice().updateUnsafe(p3,false);
      let _1023_v38;
      _1023_v38 = _dafny.Seq.of((((_1022_v37).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).length))) ? ((_1022_v37).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).length))) : (p1)), true);
      let _1024_v39;
      _1024_v39 = _dafny.Seq.of(p2);
      let _1025_v40;
      _1025_v40 = _module.D10.create_DC21(_1024_v39);
      let _1026_v41;
      let _nw165 = Array((new BigNumber(9)).toNumber());
      _nw165[(_dafny.ZERO).toNumber()] = _1018_v33;
      _nw165[(_dafny.ONE).toNumber()] = _1018_v33;
      _nw165[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("n");
      _nw165[(new BigNumber(3)).toNumber()] = _1018_v33;
      _nw165[(new BigNumber(4)).toNumber()] = _1018_v33;
      _nw165[(new BigNumber(5)).toNumber()] = _1018_v33;
      _nw165[(new BigNumber(6)).toNumber()] = _1018_v33;
      _nw165[(new BigNumber(7)).toNumber()] = _1018_v33;
      _nw165[(new BigNumber(8)).toNumber()] = _1018_v33;
      _1026_v41 = _nw165;
      let _1027_v42;
      let _nw166 = new _module.C1();
      _nw166.__ctor((_dafny.MultiSet.fromElements(_1021_v36)).IsSubsetOf(_1020_v35), (_dafny.MultiSet.FromArray(_dafny.Seq.of(p0))).Difference(_dafny.MultiSet.FromArray(_1023_v38)), (_1025_v40).dtor_cf39, _module.__default.safeDivisionInt(new BigNumber(-160), (_this).f32), ((p0) ? (_1026_v41) : (_1026_v41)));
      _1027_v42 = _nw166;
      let _index113 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_984_v1).length));
      (_984_v1)[_index113] = !(_dafny.Set.fromElements(p0)).contains(p0);
      let _pat_let_tv10 = _1027_v42;
      let _pat_let_tv11 = _1027_v42;
      let _index114 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_984_v1).length));
      (_984_v1)[_index114] = function (_source18) {
        if (_source18.is_DC22) {
          let _1028___mcc_h0 = (_source18).cf40;
          let _1029___mcc_h1 = (_source18).cf41;
          let _1030_cf41 = _1029___mcc_h1;
          let _1031_cf40 = _1028___mcc_h0;
          return _pat_let_tv10.f35;
        } else {
          let _1032___mcc_h2 = (_source18).cf39;
          let _1033_cf39 = _1032___mcc_h2;
          return _pat_let_tv11.f35;
        }
      }(_1025_v40);
      let _1034_v43;
      let _nw167 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.of());
      _1034_v43 = _nw167;
      let _index115 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((_1034_v43).length));
      (_1034_v43)[_index115] = _1023_v38;
      let _1035_v44;
      _1035_v44 = _dafny.Set.fromElements(true);
      let _1036_v45;
      _1036_v45 = _dafny.Seq.of(_1035_v44);
      let _index116 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((_1034_v43).length));
      let _index117 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_984_v1).length));
      let _rhs160 = _dafny.Seq.Concat(_1023_v38, _1023_v38);
      let _rhs161 = (p1) === (!(false));
      let _rhs162 = !((_1035_v44).IsProperSubsetOf((_1036_v45)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f32), new BigNumber((_1036_v45).length))]));
      let _lhs117 = _1034_v43;
      let _lhs118 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((_1034_v43).length));
      let _lhs119 = _984_v1;
      let _lhs120 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_984_v1).length));
      let _lhs121 = globalState;
      _lhs117[_lhs118] = _rhs160;
      _lhs119[_lhs120] = _rhs161;
      _lhs121.f26 = _rhs162;
      r0 = _984_v1;
      r1 = p2;
      return [r0, r1];
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _1037_v0;
      let _nw168 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _1037_v0 = _nw168;
      let _index118 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length));
      (_1037_v0)[_index118] = ((p0) ? (new BigNumber((p2).length)) : (new BigNumber((p2).length)));
      let _1038_v1;
      _1038_v1 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _1039_v2;
      _1039_v2 = _module.D8.create_DC15(_1038_v1);
      let _pat_let_tv12 = p3;
      let _pat_let_tv13 = p3;
      let _index119 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length));
      (_1037_v0)[_index119] = function (_source19) {
        if (_source19.is_DC16) {
          let _1040___mcc_h0 = (_source19).cf29;
          let _1041___mcc_h1 = (_source19).cf30;
          let _1042_cf30 = _1041___mcc_h1;
          let _1043_cf29 = _1040___mcc_h0;
          return new BigNumber((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(true), _dafny.Set.fromElements(_pat_let_tv12, _pat_let_tv13))).cardinality());
        } else {
          let _1044___mcc_h2 = (_source19).cf28;
          let _1045_cf28 = _1044___mcc_h2;
          return new BigNumber(-118);
        }
      }(_1039_v2);
      if (false) {
        let _1046_v3;
        _1046_v3 = _dafny.Seq.of((_this).f32);
        let _1047_v4;
        _1047_v4 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
        let _1048_v5;
        _1048_v5 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber(680));
        let _1049_v6;
        _1049_v6 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(250)), ((_1050_p1) => function (_1051_i1) {
          return _1050_p1;
        })(p1)));
        let _1052_v7;
        let _nw169 = Array((new BigNumber(9)).toNumber());
        _nw169[(_dafny.ZERO).toNumber()] = p2;
        _nw169[(_dafny.ONE).toNumber()] = p2;
        _nw169[(new BigNumber(2)).toNumber()] = (((_1047_v4).contains(p1)) ? ((_1047_v4).get(p1)) : (_dafny.Seq.UnicodeFromString("qnpxuk")));
        _nw169[(new BigNumber(3)).toNumber()] = p2;
        _nw169[(new BigNumber(4)).toNumber()] = (_this).fm22(_1048_v5, p3, (_this).f32, globalState);
        _nw169[(new BigNumber(5)).toNumber()] = (_1049_v6)[_module.__default.safeIndex((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], new BigNumber((_1049_v6).length))];
        _nw169[(new BigNumber(6)).toNumber()] = p2;
        _nw169[(new BigNumber(7)).toNumber()] = p2;
        _nw169[(new BigNumber(8)).toNumber()] = p2;
        _1052_v7 = _nw169;
        let _1053_v8;
        let _nw170 = new _module.C1();
        _nw170.__ctor(p0, (_this).f31, _dafny.Seq.Concat(_1046_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(192)), function (_1054_i0) {
          return (_this).f32;
        })), new BigNumber(798), _1052_v7);
        _1053_v8 = _nw170;
        let _1055_v9;
        let _init17 = ((_1056_v8) => function (_1057_i2) {
          return _1056_v8.f35;
        })(_1053_v8);
        let _nw171 = Array((new BigNumber(15)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw171.length); _i0_17++) {
          _nw171[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _1055_v9 = _nw171;
        let _1058_v10;
        _1058_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1055_v9,p3);
        let _1059_v11;
        let _nw172 = new _module.C0();
        _nw172.__ctor((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], p2);
        _1059_v11 = _nw172;
        let _1060_v12;
        _1060_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1053_v8.f35,_1059_v11);
        let _1061_v13;
        _1061_v13 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(186)), ((_1062_p1) => function (_1063_i4) {
          return _1062_p1;
        })(p1));
        let _1064_v14;
        let _nw173 = Array((new BigNumber(22)).toNumber());
        _nw173[(_dafny.ZERO).toNumber()] = _1059_v11;
        _nw173[(_dafny.ONE).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(2)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(3)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(4)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(5)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(6)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(7)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(8)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(9)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(10)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(11)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(12)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(13)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(14)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(15)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(16)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(17)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(18)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(19)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(20)).toNumber()] = _1059_v11;
        _nw173[(new BigNumber(21)).toNumber()] = (((_1060_v12).contains((_1053_v8).fm11(p3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-262)), ((_1067_v11) => function (_1068_i3) {
          return (_1067_v11).f36;
        })(_1059_v11))).length), _1061_v13, (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], globalState))) ? ((_1060_v12).get((_1053_v8).fm11(p3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-262)), ((_1065_v11) => function (_1066_i3) {
          return (_1065_v11).f36;
        })(_1059_v11))).length), _1061_v13, (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], globalState))) : (_1059_v11));
        _1064_v14 = _nw173;
        let _index120 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1064_v14).length));
        (_1064_v14)[_index120] = _1059_v11;
        let _1069_v15;
        _1069_v15 = _dafny.Seq.UnicodeFromString("vhmyrad");
        let _index121 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1064_v14).length));
        let _rhs163 = _1058_v10;
        let _rhs164 = _1059_v11;
        let _rhs165 = p2;
        let _lhs122 = _1064_v14;
        let _lhs123 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1064_v14).length));
        _1058_v10 = _rhs163;
        _lhs122[_lhs123] = _rhs164;
        _1069_v15 = _rhs165;
        (globalState).f26 = !((_1059_v11).f36).isEqualTo((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))]);
        (globalState).f26 = !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(140)), ((_1070_v11) => function (_1071_i5) {
          return (_1070_v11).f36;
        })(_1059_v11)), _module.__default.safeDivisionInt(new BigNumber(-20), new BigNumber((_1046_v3).length)));
        _1069_v15 = _1069_v15;
      } else {
        _1037_v0 = _1037_v0;
        let _1072_v16;
        _1072_v16 = _dafny.Seq.of((_this).f32, (_dafny.ZERO).minus(new BigNumber((_module.__default.fm25(p0, globalState)).length)), (_this).f32);
        let _1073_v17;
        let _nw174 = Array((new BigNumber(2)).toNumber());
        _nw174[(_dafny.ZERO).toNumber()] = p2;
        _nw174[(_dafny.ONE).toNumber()] = p2;
        _1073_v17 = _nw174;
        let _1074_v18;
        let _nw175 = new _module.C1();
        _nw175.__ctor(p3, (_this).f31, _1072_v16, (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], _1073_v17);
        _1074_v18 = _nw175;
        let _1075_v19;
        _1075_v19 = _module.D4.create_DC8(new _dafny.CodePoint('x'.codePointAt(0)), (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], p0);
        let _1076_v20;
        _1076_v20 = p2;
        let _1077_v21;
        _1077_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1075_v19,_1076_v20);
        if (!(_1077_v21).contains(_1075_v19)) {
          let _1078_v22;
          let _nw176 = new _module.C1();
          _nw176.__ctor(((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))]).isEqualTo(new BigNumber(-977)), _dafny.MultiSet.fromElements(p3, !(p0)), _dafny.Seq.update(_dafny.Seq.Concat(_1072_v16, _1072_v16), _module.__default.safeIndex((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], new BigNumber((_dafny.Seq.Concat(_1072_v16, _1072_v16)).length)), new BigNumber((p2).length)), (_this).f32, _1073_v17);
          _1078_v22 = _nw176;
          let _1079_v23;
          _1079_v23 = _dafny.Set.fromElements(p3);
          let _1080_v24;
          _1080_v24 = _dafny.MultiSet.fromElements(new BigNumber((_1079_v23).length), (_this).f32, new BigNumber(419));
          let _index122 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length));
          let _index123 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length));
          let _rhs166 = (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))];
          let _rhs167 = (((p3) ? ((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))]) : (new BigNumber((function () {
            let _coll37 = new _dafny.Set();
            for (const _compr_37 of (function () {
              let _coll38 = new _dafny.Set();
              for (const _compr_38 of (_1080_v24).Elements) {
                let _1081_v25 = _compr_38;
                if ((_1080_v24).contains(_1081_v25)) {
                  _coll38.add((_1081_v25).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
                    let _coll39 = new _dafny.Set();
                    for (const _compr_39 of _dafny.IntegerRange(new BigNumber(-421), new BigNumber(239))) {
                      let _1082_v26 = _compr_39;
                      if (((new BigNumber(-421)).isLessThanOrEqualTo(_1082_v26)) && ((_1082_v26).isLessThan(new BigNumber(239)))) {
                        _coll39.add((_1082_v26).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("pdfc"))).length)));
                      }
                    }
                    return _coll39;
                  }()).length), new BigNumber(543))).length))));
                }
              }
              return _coll38;
            }()).Elements) {
              let _1083_v27 = _compr_37;
              if ((function () {
                let _coll40 = new _dafny.Set();
                for (const _compr_40 of (_1080_v24).Elements) {
                  let _1084_v25 = _compr_40;
                  if ((_1080_v24).contains(_1084_v25)) {
                    _coll40.add((_1084_v25).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
                      let _coll41 = new _dafny.Set();
                      for (const _compr_41 of _dafny.IntegerRange(new BigNumber(-421), new BigNumber(239))) {
                        let _1085_v26 = _compr_41;
                        if (((new BigNumber(-421)).isLessThanOrEqualTo(_1085_v26)) && ((_1085_v26).isLessThan(new BigNumber(239)))) {
                          _coll41.add((_1085_v26).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("pdfc"))).length)));
                        }
                      }
                      return _coll41;
                    }()).length), new BigNumber(543))).length))));
                  }
                }
                return _coll40;
              }()).contains(_1083_v27)) {
                _coll37.add(_module.__default.safeModuloInt(_1083_v27, new BigNumber(891)));
              }
            }
            return _coll37;
          }()).length)))).plus((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))]);
          let _rhs168 = new BigNumber(135);
          let _rhs169 = (_1078_v22).fm4(_dafny.Set.fromElements(p2), p0, _dafny.Seq.of(p0, _1078_v22.f35), globalState);
          let _rhs170 = (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))];
          let _lhs124 = globalState;
          let _lhs125 = _1037_v0;
          let _lhs126 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length));
          let _lhs127 = globalState;
          let _lhs128 = globalState;
          let _lhs129 = _1037_v0;
          let _lhs130 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length));
          _lhs124.f3 = _rhs166;
          _lhs125[_lhs126] = _rhs167;
          _lhs127.f18 = _rhs168;
          _lhs128.f24 = _rhs169;
          _lhs129[_lhs130] = _rhs170;
          let _1086_v28;
          let _nw177 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Set.Empty);
          _1086_v28 = _nw177;
          let _1087_v29;
          _1087_v29 = _dafny.Seq.of(_1074_v18);
          let _1088_v30;
          let _nw178 = Array((new BigNumber(24)).toNumber());
          _nw178[(_dafny.ZERO).toNumber()] = _1078_v22;
          _nw178[(_dafny.ONE).toNumber()] = _1074_v18;
          _nw178[(new BigNumber(2)).toNumber()] = _1074_v18;
          _nw178[(new BigNumber(3)).toNumber()] = _1074_v18;
          _nw178[(new BigNumber(4)).toNumber()] = _1078_v22;
          _nw178[(new BigNumber(5)).toNumber()] = _1078_v22;
          _nw178[(new BigNumber(6)).toNumber()] = _1078_v22;
          _nw178[(new BigNumber(7)).toNumber()] = _1078_v22;
          _nw178[(new BigNumber(8)).toNumber()] = _1074_v18;
          _nw178[(new BigNumber(9)).toNumber()] = _1078_v22;
          _nw178[(new BigNumber(10)).toNumber()] = _1078_v22;
          _nw178[(new BigNumber(11)).toNumber()] = _1078_v22;
          _nw178[(new BigNumber(12)).toNumber()] = _1078_v22;
          _nw178[(new BigNumber(13)).toNumber()] = _1074_v18;
          _nw178[(new BigNumber(14)).toNumber()] = _1078_v22;
          _nw178[(new BigNumber(15)).toNumber()] = _1074_v18;
          _nw178[(new BigNumber(16)).toNumber()] = _1074_v18;
          _nw178[(new BigNumber(17)).toNumber()] = (_1087_v29)[_module.__default.safeIndex((_this).f32, new BigNumber((_1087_v29).length))];
          _nw178[(new BigNumber(18)).toNumber()] = _1074_v18;
          _nw178[(new BigNumber(19)).toNumber()] = _1078_v22;
          _nw178[(new BigNumber(20)).toNumber()] = _1078_v22;
          _nw178[(new BigNumber(21)).toNumber()] = _1074_v18;
          _nw178[(new BigNumber(22)).toNumber()] = _1078_v22;
          _nw178[(new BigNumber(23)).toNumber()] = _1074_v18;
          _1088_v30 = _nw178;
          let _1089_v31;
          _1089_v31 = _dafny.Set.fromElements(_1088_v30);
          let _index124 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_1086_v28).length));
          (_1086_v28)[_index124] = (_1089_v31).Difference(_1089_v31);
          let _index125 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_1086_v28).length));
          (_1086_v28)[_index125] = ((_1074_v18.f35) ? (_dafny.Set.fromElements(_1088_v30)) : (_dafny.Set.fromElements(_1088_v30)));
          let _1090_v32;
          _1090_v32 = _dafny.Seq.of(!(p3), ((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))]).isLessThanOrEqualTo(new BigNumber(-645)), p0, _1078_v22.f35, (_this).fm5(globalState));
          _1090_v32 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_1074_v18.f35, _1074_v18.f35), _dafny.Seq.update(_dafny.Seq.of(p3), _module.__default.safeIndex((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], new BigNumber((_dafny.Seq.of(p3)).length)), p0)), _1090_v32);
          let _1091_v33;
          let _init18 = ((_1092_p3) => function (_1093_i6) {
            return _1092_p3;
          })(p3);
          let _nw179 = Array((new BigNumber(8)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw179.length); _i0_18++) {
            _nw179[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _1091_v33 = _nw179;
          let _1094_v34;
          _1094_v34 = _module.D4.create_DC7(_1091_v33);
          let _1095_v35;
          _1095_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,_1091_v33);
          let _1096_v36;
          let _nw180 = Array((new BigNumber(6)).toNumber());
          _nw180[(_dafny.ZERO).toNumber()] = ((false) ? (_1091_v33) : (_1091_v33));
          _nw180[(_dafny.ONE).toNumber()] = _1091_v33;
          _nw180[(new BigNumber(2)).toNumber()] = _1091_v33;
          _nw180[(new BigNumber(3)).toNumber()] = (_1094_v34).dtor_cf12;
          _nw180[(new BigNumber(4)).toNumber()] = (((_1095_v35).contains((_this).f32)) ? ((_1095_v35).get((_this).f32)) : (_1091_v33));
          _nw180[(new BigNumber(5)).toNumber()] = _1091_v33;
          _1096_v36 = _nw180;
          let _index126 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_1096_v36).length));
          (_1096_v36)[_index126] = _1091_v33;
          let _index127 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_1096_v36).length));
          let _nw181 = Array((new BigNumber(12)).toNumber()).fill(false);
          (_1096_v36)[_index127] = _nw181;
        } else {
          let _1097_v37;
          let _nw182 = new _module.C1();
          _nw182.__ctor(false, (_this).f31, _1072_v16, (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], _1073_v17);
          _1097_v37 = _nw182;
          let _1098_v38;
          _1098_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1097_v37,p0);
          _1098_v38 = (_1098_v38).update(_1097_v37, (new BigNumber(-231)).isLessThan(((_1097_v37).f30)[_module.__default.safeIndex((_this).f32, new BigNumber(((_1097_v37).f30).length))]));
          let _1099_v39;
          _1099_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,p3);
          let _1100_v40;
          let _nw183 = new _module.C1();
          _nw183.__ctor(p0, (_this).f31, _dafny.Seq.of((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], (_module.D2.create_DC4((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], (_dafny.ZERO).minus(new BigNumber((_1099_v39).length)), (_this).f32, p1, p2)).dtor_cf6), (new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality())).multipliedBy((_1097_v37).f27), (_1097_v37).f28);
          _1100_v40 = _nw183;
          let _1101_v41;
          _1101_v41 = _dafny.Seq.of(p0);
          let _1102_v42;
          _1102_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1074_v18.f35,(_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))]);
          let _1103_v43;
          let _nw184 = new _module.C0();
          _nw184.__ctor(new BigNumber((_dafny.Seq.of(new BigNumber((_1101_v41).length), new BigNumber(((_this).fm22(_1102_v42, _1074_v18.f35, _module.__default.fm1(globalState), globalState)).length))).length), p2);
          _1103_v43 = _nw184;
          (globalState).f26 = !((_this).fm5(globalState)) || (!((_1074_v18.f35) || (_1074_v18.f35)));
          (_1074_v18).f35 = p3;
        }
        let _1104_v44;
        _1104_v44 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm26(p1, (_this).f32, (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], globalState),(_this).f32);
        (globalState).f18 = _module.__default.safeModuloInt((((_1104_v44).contains(p1)) ? ((_1104_v44).get(p1)) : ((_this).f32)), (_this).f32);
        let _1105_v45;
        let _nw185 = new _module.C0();
        _nw185.__ctor(new BigNumber((_1072_v16).length), p2);
        _1105_v45 = _nw185;
      }
      let _hi4 = new BigNumber(220);
      for (let _1106_i7 = (_dafny.ZERO).minus((_this).f32); _1106_i7.isLessThan(_hi4); _1106_i7 = _1106_i7.plus(_dafny.ONE)) {
        if ((_1106_i7).isLessThan(_1106_i7)) {
          let _1107_v46;
          let _nw186 = new _module.C0();
          _nw186.__ctor((_this).f32, p2);
          _1107_v46 = _nw186;
          let _1108_v47;
          _1108_v47 = _dafny.Set.fromElements(_1107_v46, _1107_v46, _1107_v46);
          (globalState).f26 = (_1108_v47).IsProperSubsetOf(_1108_v47);
          let _1109_v48;
          _1109_v48 = new _dafny.CodePoint('k'.codePointAt(0));
          let _1110_v49;
          _1110_v49 = _dafny.Seq.of(p3);
          let _rhs171 = !(_module.__default.fm0(((p3) ? (new BigNumber((_1110_v49).length)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(90)), function (_1111_i8) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          })).length))), globalState));
          let _rhs172 = ((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))]).minus(_module.__default.safeModuloInt((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], (_this).f32));
          let _rhs173 = ((_1107_v46).f36).minus(_1106_i7);
          let _rhs174 = p1;
          let _lhs131 = globalState;
          let _lhs132 = globalState;
          let _lhs133 = globalState;
          _lhs131.f26 = _rhs171;
          _lhs132.f15 = _rhs172;
          _lhs133.f18 = _rhs173;
          _1109_v48 = _rhs174;
          let _1112_v50;
          _1112_v50 = _dafny.Map.Empty.slice().updateUnsafe((_1107_v46).f36,p0);
          let _1113_v51;
          let _nw187 = Array((new BigNumber(23)).toNumber());
          _nw187[(_dafny.ZERO).toNumber()] = (_1107_v46).f37;
          _nw187[(_dafny.ONE).toNumber()] = (_1107_v46).f37;
          _nw187[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(p2, _module.__default.safeIndex((_1107_v46).f36, new BigNumber((p2).length)), p1), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1112_v50).length)), new BigNumber((_dafny.Seq.update(p2, _module.__default.safeIndex((_1107_v46).f36, new BigNumber((p2).length)), p1)).length)), p1);
          _nw187[(new BigNumber(3)).toNumber()] = (_1107_v46).f37;
          _nw187[(new BigNumber(4)).toNumber()] = (_this).fm22(_module.__default.fm24((_this).f32, globalState), p0, new BigNumber((_dafny.Seq.UnicodeFromString("ngpnnpqao")).length), globalState);
          _nw187[(new BigNumber(5)).toNumber()] = p2;
          _nw187[(new BigNumber(6)).toNumber()] = p2;
          _nw187[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(p2, p2);
          _nw187[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vkphmjmvi"), p2);
          _nw187[(new BigNumber(9)).toNumber()] = p2;
          _nw187[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("tmodjtc");
          _nw187[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("baincv"), (_1107_v46).f37);
          _nw187[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hehm"), _dafny.Seq.UnicodeFromString("fgcyfxm"));
          _nw187[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(p2, p2);
          _nw187[(new BigNumber(14)).toNumber()] = (_1107_v46).f37;
          _nw187[(new BigNumber(15)).toNumber()] = p2;
          _nw187[(new BigNumber(16)).toNumber()] = (_1107_v46).f37;
          _nw187[(new BigNumber(17)).toNumber()] = (_1107_v46).f37;
          _nw187[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("bwkywvb");
          _nw187[(new BigNumber(19)).toNumber()] = (_1107_v46).f37;
          _nw187[(new BigNumber(20)).toNumber()] = p2;
          _nw187[(new BigNumber(21)).toNumber()] = p2;
          _nw187[(new BigNumber(22)).toNumber()] = p2;
          _1113_v51 = _nw187;
          let _index128 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_1113_v51).length));
          (_1113_v51)[_index128] = (_1107_v46).f37;
          let _index129 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_1113_v51).length));
          (_1113_v51)[_index129] = (_1107_v46).f37;
          let _1114_v53;
          _1114_v53 = _dafny.MultiSet.fromElements(function () {
            let _coll42 = new _dafny.Set();
            for (const _compr_42 of _dafny.IntegerRange(new BigNumber(274), new BigNumber(326))) {
              let _1115_v52 = _compr_42;
              if (((new BigNumber(274)).isLessThanOrEqualTo(_1115_v52)) && ((_1115_v52).isLessThan(new BigNumber(326)))) {
                _coll42.add((_1115_v52).multipliedBy(new BigNumber((_1038_v1).length)));
              }
            }
            return _coll42;
          }(), _dafny.Set.fromElements(((((_this).f31).contains(p3)) ? (((_this).f31).get(p3)) : ((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))])), (_1107_v46).f36, (_1107_v46).f36, (_1107_v46).f36, (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))]));
          let _1116_v54;
          _1116_v54 = _module.D11.create_DC23(_1114_v53);
          _1112_v50 = (_1112_v50).update(((_1107_v46).f36).plus(new BigNumber(((_1116_v54).dtor_cf42).cardinality())), !(p0));
          let _1117_v55;
          let _nw188 = Array((new BigNumber(14)).toNumber()).fill(false);
          _1117_v55 = _nw188;
          let _index130 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1117_v55).length));
          (_1117_v55)[_index130] = !(true);
          let _1118_v56;
          _1118_v56 = _module.D2.create_DC4((_dafny.ZERO).minus((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))]), (_1107_v46).f36, new BigNumber(986), new _dafny.CodePoint('y'.codePointAt(0)), (_1107_v46).f37);
          let _1119_v57;
          _1119_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1118_v56,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p3,p3)).length));
          let _1120_v58;
          _1120_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1119_v57,_dafny.Map.Empty.slice().updateUnsafe(p0,_1106_i7));
          let _1121_v59;
          _1121_v59 = true;
          let _1122_v60;
          _1122_v60 = _dafny.Map.Empty.slice().updateUnsafe((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))],_1121_v59);
          let _1123_v61;
          _1123_v61 = _dafny.Seq.of(_1107_v46);
          let _1124_v62;
          _1124_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1122_v60,_dafny.Seq.of(_1107_v46, (_1123_v61)[_module.__default.safeIndex((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], new BigNumber((_1123_v61).length))]));
          let _1125_v63;
          _1125_v63 = _dafny.Seq.of(_1124_v62);
          let _1126_v64;
          _1126_v64 = _dafny.Seq.of((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))]);
          let _1127_v65;
          _1127_v65 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1125_v63)[_module.__default.safeIndex((_this).f32, new BigNumber((_1125_v63).length))]).length),_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1126_v64).length)));
          let _1128_v66;
          _1128_v66 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f32);
          let _1129_v67;
          let _nw189 = Array((new BigNumber(12)).toNumber());
          _nw189[(_dafny.ZERO).toNumber()] = _1120_v58;
          _nw189[(_dafny.ONE).toNumber()] = _1120_v58;
          _nw189[(new BigNumber(2)).toNumber()] = (_1120_v58).Merge(_1120_v58);
          _nw189[(new BigNumber(3)).toNumber()] = _1120_v58;
          _nw189[(new BigNumber(4)).toNumber()] = _1120_v58;
          _nw189[(new BigNumber(5)).toNumber()] = _1120_v58;
          _nw189[(new BigNumber(6)).toNumber()] = (_1120_v58).Merge(_1120_v58);
          _nw189[(new BigNumber(7)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_1118_v56,(_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))]),(_dafny.Map.Empty.slice().updateUnsafe(p3,(_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))])).update(p0, _1106_i7))).update(_1119_v57, (((_1127_v65).contains((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))])) ? ((_1127_v65).get((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))])) : (_1128_v66)));
          _nw189[(new BigNumber(8)).toNumber()] = _1120_v58;
          _nw189[(new BigNumber(9)).toNumber()] = _1120_v58;
          _nw189[(new BigNumber(10)).toNumber()] = _1120_v58;
          _nw189[(new BigNumber(11)).toNumber()] = (_1120_v58).Merge(_1120_v58);
          _1129_v67 = _nw189;
          let _index131 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_1129_v67).length));
          (_1129_v67)[_index131] = _1120_v58;
          let _1130_v68;
          _1130_v68 = _module.D4.create_DC8(p1, new BigNumber(802), false);
          let _index132 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_1113_v51).length));
          let _index133 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1117_v55).length));
          let _index134 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_1129_v67).length));
          let _rhs175 = (_module.D11.create_DC24((((_1128_v66).contains(p0)) ? ((_1128_v66).get(p0)) : ((_1107_v46).f36)), _1130_v68)).dtor_cf43;
          let _rhs176 = (_1107_v46).f37;
          let _rhs177 = true;
          let _rhs178 = (_1107_v46).f36;
          let _rhs179 = _1120_v58;
          let _lhs134 = globalState;
          let _lhs135 = _1113_v51;
          let _lhs136 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_1113_v51).length));
          let _lhs137 = _1117_v55;
          let _lhs138 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1117_v55).length));
          let _lhs139 = globalState;
          let _lhs140 = _1129_v67;
          let _lhs141 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_1129_v67).length));
          _lhs134.f24 = _rhs175;
          _lhs135[_lhs136] = _rhs176;
          _lhs137[_lhs138] = _rhs177;
          _lhs139.f21 = _rhs178;
          _lhs140[_lhs141] = _rhs179;
        } else {
          let _1131_v70;
          _1131_v70 = _dafny.Seq.of((_this).f32);
          (globalState).f24 = new BigNumber((function () {
            let _coll43 = new _dafny.Map();
            for (const _compr_43 of (_1131_v70).Elements) {
              let _1132_v69 = _compr_43;
              if (_dafny.Seq.contains(_1131_v70, _1132_v69)) {
                _coll43.push([_module.__default.safeDivisionInt(_1132_v69, _1106_i7),p1]);
              }
            }
            return _coll43;
          }()).length);
          let _1133_v71;
          _1133_v71 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
          let _1134_v72;
          _1134_v72 = _module.D10.create_DC22(new BigNumber(((((_1133_v71).contains(p3)) ? ((_1133_v71).get(p3)) : (p2))).length), _dafny.MultiSet.fromElements(p0));
          (_this).m7(_1038_v1, (_1134_v72).dtor_cf40, globalState);
          (globalState).f26 = !(((_this).f32).isLessThan(_module.__default.fm1(globalState)));
          let _1135_v73;
          let _init19 = ((_1136_p0) => function (_1137_i9) {
            return !(!(true)) || (_1136_p0);
          })(p0);
          let _nw190 = Array((new BigNumber(14)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw190.length); _i0_19++) {
            _nw190[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _1135_v73 = _nw190;
          let _index135 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_1135_v73).length));
          (_1135_v73)[_index135] = p3;
          let _1138_v74;
          _1138_v74 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))]);
          let _1139_v75;
          _1139_v75 = _dafny.Seq.of(p0);
          let _1140_v76;
          _1140_v76 = _dafny.Seq.of(_1139_v75, _1139_v75);
          let _index136 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_1135_v73).length));
          (_1135_v73)[_index136] = (_module.__default.fm27(new BigNumber((_1140_v76).length), globalState)).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_1138_v74).length))));
          let _1141_v77;
          _1141_v77 = _dafny.Map.Empty.slice().updateUnsafe((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))],p3);
          let _1142_v78;
          _1142_v78 = _module.D1.create_DC1(_1141_v77);
          let _1143_v79;
          _1143_v79 = _dafny.Seq.UnicodeFromString("tqpsjmtb");
          let _1144_v80;
          _1144_v80 = _dafny.Seq.of(_1135_v73, _1135_v73);
          let _1145_v81;
          _1145_v81 = _dafny.MultiSet.fromElements(_1106_i7, (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], _1106_i7, (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], (_this).f32);
          let _pat_let_tv14 = _1037_v0;
          let _pat_let_tv15 = _1037_v0;
          let _pat_let_tv16 = _1135_v73;
          let _pat_let_tv17 = _1135_v73;
          let _pat_let_tv18 = _1141_v77;
          let _rhs180 = (_1144_v80)[_module.__default.safeIndex(new BigNumber((_1145_v81).cardinality()), new BigNumber((_1144_v80).length))];
          let _rhs181 = ((p0) ? (function (_pat_let18_0) {
            return function (_1146_dt__update__tmp_h0) {
              return function (_pat_let19_0) {
                return function (_1147_dt__update_hcf1_h0) {
                  return _module.D1.create_DC1(_1147_dt__update_hcf1_h0);
                }(_pat_let19_0);
              }(_dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv15)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_pat_let_tv14).length))],(_pat_let_tv17)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_pat_let_tv16).length))]));
            }(_pat_let18_0);
          }(_1142_v78)) : (function (_pat_let20_0) {
            return function (_1148_dt__update__tmp_h1) {
              return function (_pat_let21_0) {
                return function (_1149_dt__update_hcf1_h1) {
                  return _module.D1.create_DC1(_1149_dt__update_hcf1_h1);
                }(_pat_let21_0);
              }(_pat_let_tv18);
            }(_pat_let20_0);
          }(_module.__default.fm28(new _dafny.CodePoint('k'.codePointAt(0)), (_dafny.ZERO).minus((_dafny.ZERO).minus(_1106_i7)), (_1135_v73)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_1135_v73).length))], (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], globalState))));
          let _rhs182 = _1143_v79;
          _1135_v73 = _rhs180;
          _1142_v78 = _rhs181;
          _1143_v79 = _rhs182;
        }
        let _index137 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length));
        (_1037_v0)[_index137] = _module.__default.safeModuloInt(_1106_i7, new BigNumber(408));
        (globalState).f26 = p3;
        if (p0) {
          let _1150_v82;
          _1150_v82 = _dafny.Seq.of(_1106_i7);
          (globalState).f18 = ((_this).f32).plus((_1150_v82)[_module.__default.safeIndex((_this).f32, new BigNumber((_1150_v82).length))]);
          (globalState).f3 = _1106_i7;
          let _1151_v83;
          _1151_v83 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1106_i7);
          let _1152_v84;
          _1152_v84 = _dafny.Set.fromElements(p3);
          (globalState).f0 = _module.__default.safeModuloInt((_this).f32, _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber(((_this).fm22(_1151_v83, p0, new BigNumber((_1152_v84).length), globalState)).length)), _1106_i7));
          (globalState).f26 = !(p3) || (true);
          let _1153_v85;
          let _nw191 = new _module.C0();
          _nw191.__ctor(_module.__default.safeDivisionInt((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], _1106_i7), _dafny.Seq.Concat((_this).fm22(_1151_v83, true, (_this).f32, globalState), (_this).fm22(_1151_v83, p0, new BigNumber(306), globalState)));
          _1153_v85 = _nw191;
        } else {
          (globalState).f26 = p0;
          let _1154_v86;
          let _nw192 = Array((new BigNumber(25)).toNumber()).fill(false);
          _1154_v86 = _nw192;
          let _index138 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_1154_v86).length));
          (_1154_v86)[_index138] = p0;
          let _1155_v87;
          _1155_v87 = _dafny.Seq.of(!(p0));
          let _index139 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_1154_v86).length));
          (_1154_v86)[_index139] = (_1155_v87)[_module.__default.safeIndex(_1106_i7, new BigNumber((_1155_v87).length))];
          let _1156_v88;
          _1156_v88 = _dafny.MultiSet.fromElements(p2);
          (globalState).f8 = _1156_v88;
          (globalState).f26 = p0;
          let _index140 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_1154_v86).length));
          (_1154_v86)[_index140] = p0;
        }
      }
      let _1157_v89;
      _1157_v89 = new _dafny.CodePoint('g'.codePointAt(0));
      _1157_v89 = p1;
      let _ingredients0 = [];
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1037_v0).length))) {
        let _1158_i10 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1158_i10)) && ((_1158_i10).isLessThan(new BigNumber((_1037_v0).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_1037_v0, (_1158_i10).toNumber(), (_1158_i10).multipliedBy((_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))])));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      if (!(p0) || (!(p3))) {
        (globalState).f15 = (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))];
        let _1159_v90;
        _1159_v90 = _dafny.Seq.of(new BigNumber((p2).length), (_this).f32);
        let _1160_v91;
        let _nw193 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1160_v91 = _nw193;
        let _1161_v92;
        let _nw194 = new _module.C1();
        _nw194.__ctor(p3, (_this).f31, _1159_v90, (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))], _1160_v91);
        _1161_v92 = _nw194;
        let _1162_v93;
        let _nw195 = Array((new BigNumber(5)).toNumber());
        _nw195[(_dafny.ZERO).toNumber()] = _1161_v92;
        _nw195[(_dafny.ONE).toNumber()] = _1161_v92;
        _nw195[(new BigNumber(2)).toNumber()] = _1161_v92;
        _nw195[(new BigNumber(3)).toNumber()] = _1161_v92;
        _nw195[(new BigNumber(4)).toNumber()] = _1161_v92;
        _1162_v93 = _nw195;
        _1162_v93 = _1162_v93;
        let _1163_v94;
        _1163_v94 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm1(globalState)).plus((_this).f32),!(p0));
        _1163_v94 = (_1163_v94).update(new BigNumber((_1163_v94).length), !(p0));
        let _1164_v95;
        _1164_v95 = _dafny.MultiSet.fromElements((_this).f32);
        let _1165_v96;
        _1165_v96 = _module.D4.create_DC9(new BigNumber((_1164_v95).cardinality()), p3);
        _1165_v96 = _1165_v96;
      } else {
        let _index141 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length));
        (_1037_v0)[_index141] = (_1037_v0)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1037_v0).length))];
        (globalState).f26 = p0;
        let _1166_v97;
        let _nw196 = Array((new BigNumber(3)).toNumber());
        _1166_v97 = _nw196;
        _1166_v97 = _1166_v97;
        let _1167_v98;
        let _nw197 = Array((new BigNumber(11)).toNumber()).fill(false);
        _1167_v98 = _nw197;
        let _index142 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1167_v98).length));
        (_1167_v98)[_index142] = false;
        let _index143 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1167_v98).length));
        (_1167_v98)[_index143] = p0;
        let _1168_v99;
        _1168_v99 = _dafny.Set.fromElements(p0, (_this).fm5(globalState));
        (globalState).f26 = (_1168_v99).IsSubsetOf(_1168_v99);
      }
      return;
    }
    m7(p0, p1, globalState) {
      let _this = this;
      (globalState).f3 = (p1).multipliedBy(p1);
      if (!(p1).isEqualTo(p1)) {
        let _1169_v0;
        let _nw198 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1169_v0 = _nw198;
        let _1170_v1;
        _1170_v1 = new _dafny.CodePoint('d'.codePointAt(0));
        let _index144 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_1169_v0).length));
        (_1169_v0)[_index144] = _1170_v1;
        let _1171_v2;
        _1171_v2 = true;
        let _1172_v3;
        _1172_v3 = _dafny.Set.fromElements(_1171_v2);
        let _index145 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_1169_v0).length));
        let _rhs183 = _module.__default.safeModuloInt(new BigNumber((((_1171_v2) ? (_1172_v3) : (_1172_v3))).length), (_this).f32);
        let _rhs184 = !(_1171_v2);
        let _rhs185 = _1170_v1;
        let _lhs142 = globalState;
        let _lhs143 = globalState;
        let _lhs144 = _1169_v0;
        let _lhs145 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_1169_v0).length));
        _lhs142.f21 = _rhs183;
        _lhs143.f26 = _rhs184;
        _lhs144[_lhs145] = _rhs185;
        (globalState).f24 = (_this).f32;
        let _1173_v4;
        let _init20 = function (_1174_i0) {
          return (_1174_i0).minus((_this).f32);
        };
        let _nw199 = Array((new BigNumber(7)).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw199.length); _i0_20++) {
          _nw199[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _1173_v4 = _nw199;
        let _rhs186 = (_module.__default.fm29(_1171_v2, true, _1171_v2, globalState));
        let _rhs187 = _1171_v2;
        let _rhs188 = _1173_v4;
        let _lhs146 = globalState;
        let _lhs147 = globalState;
        _lhs146.f26 = _rhs186;
        _lhs147.f26 = _rhs187;
        _1173_v4 = _rhs188;
        let _1175_v5;
        let _nw200 = Array((new BigNumber(13)).toNumber()).fill(false);
        _1175_v5 = _nw200;
        let _index146 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_1175_v5).length));
        (_1175_v5)[_index146] = (_this).fm21(globalState);
        let _1176_v6;
        _1176_v6 = _dafny.Seq.UnicodeFromString("wxwtlr");
        let _index147 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_1175_v5).length));
        (_1175_v5)[_index147] = _dafny.Seq.IsPrefixOf(_1176_v6, _1176_v6);
        let _index148 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_1173_v4).length));
        (_1173_v4)[_index148] = new BigNumber(-949);
        let _index149 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_1173_v4).length));
        (_1173_v4)[_index149] = _module.__default.fm1(globalState);
      } else {
        let _1177_v7;
        _1177_v7 = false;
        let _1178_v8;
        let _1179_v9;
        let _out38;
        let _out39;
        let _outcollector16 = _module.__default.m0(!(_1177_v7), (_dafny.ZERO).minus((p1).minus(p1)), (_this).f32, (p1).plus(p1), globalState);
        _out38 = _outcollector16[0];
        _out39 = _outcollector16[1];
        _1178_v8 = _out38;
        _1179_v9 = _out39;
        let _1180_v10;
        _1180_v10 = _dafny.Seq.UnicodeFromString("lmvkrjqv");
        let _1181_v11;
        let _nw201 = new _module.C0();
        _nw201.__ctor((_1179_v9).multipliedBy(_module.__default.fm1(globalState)), _1180_v10);
        _1181_v11 = _nw201;
        let _1182_v12;
        _1182_v12 = _dafny.Set.fromElements((_dafny.ZERO).minus(p1));
        let _1183_v13;
        _1183_v13 = _dafny.MultiSet.fromElements((_this).f32);
        let _1184_v14;
        _1184_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1177_v7,_1183_v13);
        let _1185_v15;
        _1185_v15 = _dafny.Set.fromElements(_1182_v12, _dafny.Set.fromElements(_1179_v9, (_1181_v11).f36, (_dafny.ZERO).minus(new BigNumber((_1184_v14).length)), (_1181_v11).f36), _dafny.Set.fromElements((_this).f32, new BigNumber(322)), _1182_v12, (_1182_v12).Difference(_1182_v12));
        let _1186_v16;
        _1186_v16 = _dafny.Seq.of(_1178_v8, _1177_v7);
        let _rhs189 = _1185_v15;
        let _rhs190 = _1177_v7;
        let _rhs191 = (_1179_v9).minus(new BigNumber((_1186_v16).length));
        let _lhs148 = globalState;
        _1185_v15 = _rhs189;
        _1178_v8 = _rhs190;
        _lhs148.f21 = _rhs191;
        let _rhs192 = _1186_v16;
        let _rhs193 = ((!(_1177_v7)) ? (_dafny.Seq.UnicodeFromString("coh")) : (_dafny.Seq.UnicodeFromString("vbd")));
        _1186_v16 = _rhs192;
        _1180_v10 = _rhs193;
        if (((!((_1186_v16)[_module.__default.safeIndex((_this).f32, new BigNumber((_1186_v16).length))])) ? (_1177_v7) : (_1177_v7))) {
          (globalState).f26 = _1178_v8;
          let _1187_v17;
          _1187_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1179_v9,(_this).f32);
          _1187_v17 = (_1187_v17).update(new BigNumber((_dafny.Seq.Concat(_1180_v10, (_1181_v11).f37)).length), _module.__default.fm1(globalState));
          let _1188_v18;
          _1188_v18 = _dafny.Seq.of((_1181_v11).f36);
          let _1189_v19;
          _1189_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1178_v8,_1188_v18);
          let _1190_v20;
          _1190_v20 = _dafny.Map.Empty.slice().updateUnsafe(!(_1178_v8),new BigNumber((_1189_v19).length));
          let _rhs194 = new BigNumber(((_1190_v20).update(_1178_v8, p1)).length);
          let _rhs195 = !(_1178_v8);
          let _rhs196 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1179_v9));
          let _rhs197 = ((_1179_v9).plus(p1)).multipliedBy((_1181_v11).f36);
          let _lhs149 = globalState;
          let _lhs150 = globalState;
          let _lhs151 = globalState;
          let _lhs152 = globalState;
          _lhs149.f21 = _rhs194;
          _lhs150.f26 = _rhs195;
          _lhs151.f21 = _rhs196;
          _lhs152.f18 = _rhs197;
          let _1191_v21;
          _1191_v21 = new _dafny.CodePoint('g'.codePointAt(0));
          let _1192_v22;
          let _nw202 = new _module.C0();
          _nw202.__ctor(_module.__default.fm1(globalState), _dafny.Seq.update((_1181_v11).f37, _module.__default.safeIndex(new BigNumber(605), new BigNumber(((_1181_v11).f37).length)), _1191_v21));
          _1192_v22 = _nw202;
          (globalState).f26 = ((_1178_v8) ? (_1177_v7) : (_1177_v7));
        } else {
          _1178_v8 = _1178_v8;
          _1180_v10 = (_1181_v11).f37;
          (globalState).f18 = _1179_v9;
          (globalState).f26 = !(((_this).f31).Union((_this).f31)).equals((_this).f31);
          let _1193_v24;
          _1193_v24 = _dafny.Seq.of(function () {
            let _coll44 = new _dafny.Map();
            for (const _compr_44 of _dafny.IntegerRange(new BigNumber(454), new BigNumber(314))) {
              let _1194_v23 = _compr_44;
              if (((new BigNumber(454)).isLessThanOrEqualTo(_1194_v23)) && ((_1194_v23).isLessThan(new BigNumber(314)))) {
                _coll44.push([(_1194_v23).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1177_v7,_1179_v9)).length)),(_1181_v11).f36]);
              }
            }
            return _coll44;
          }());
          let _1195_v25;
          _1195_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1181_v11,_1179_v9);
          let _1196_v26;
          _1196_v26 = _dafny.Map.Empty.slice().updateUnsafe((_1181_v11).f36,(((_1195_v25).contains(_1181_v11)) ? ((_1195_v25).get(_1181_v11)) : (new BigNumber(((_1181_v11).f37).length))));
          let _1197_v27;
          _1197_v27 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_1181_v11).f36,new BigNumber(63)), ((_1193_v24)[_module.__default.safeIndex((_this).f32, new BigNumber((_1193_v24).length))]).update(new BigNumber(371), _1179_v9), _module.__default.fm30(globalState), _1196_v26, _1196_v26);
          let _1198_v28;
          _1198_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1197_v27)[_module.__default.safeIndex((_1181_v11).f36, new BigNumber((_1197_v27).length))]).length),!(!(_1178_v8)) || ((_this).fm5(globalState)));
          _1198_v28 = _1198_v28;
        }
      }
      let _1199_v29;
      let _nw203 = Array((new BigNumber(14)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1199_v29 = _nw203;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1199_v29).length))) {
        let _1200_i1 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1200_i1)) && ((_1200_i1).isLessThan(new BigNumber((_1199_v29).length))))) {
          (_1199_v29)[(_1200_i1)] = ((_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("ic"), _dafny.Seq.UnicodeFromString("hksuqeycu"))) ? (new _dafny.CodePoint('u'.codePointAt(0))) : (new _dafny.CodePoint('r'.codePointAt(0))));
        }
      }
      let _1201_v30;
      _1201_v30 = true;
      (globalState).f26 = _1201_v30;
      let _1202_v31;
      _1202_v31 = _dafny.MultiSet.fromElements(_1201_v30);
      _1202_v31 = (_this).f31;
      let _1203_v32;
      _1203_v32 = _dafny.Seq.of(false);
      let _1204_v33;
      _1204_v33 = _dafny.MultiSet.fromElements(_1203_v32);
      let _1205_v34;
      let _nw204 = Array((_dafny.ONE).toNumber());
      _nw204[(_dafny.ZERO).toNumber()] = new BigNumber((_1204_v33).cardinality());
      _1205_v34 = _nw204;
      _1205_v34 = _1205_v34;
      return;
    }
    m8(p0, p1, globalState) {
      let _this = this;
      let _1206_v0;
      let _nw205 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1206_v0 = _nw205;
      let _index150 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_1206_v0).length));
      (_1206_v0)[_index150] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(155)), function (_1207_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      });
      let _1208_v1;
      _1208_v1 = _dafny.Seq.UnicodeFromString("rwtdknsgw");
      let _1209_v2;
      _1209_v2 = false;
      let _1210_v3;
      _1210_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1208_v1).length),_1209_v2);
      let _1211_v4;
      _1211_v4 = _dafny.Map.Empty.slice().updateUnsafe((((_1210_v3).contains(p1)) ? ((_1210_v3).get(p1)) : (_1209_v2)),_1208_v1);
      let _1212_v5;
      _1212_v5 = _dafny.Set.fromElements(!(_1209_v2), _1209_v2);
      let _1213_v6;
      _1213_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1212_v5,!(_1209_v2));
      let _index151 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_1206_v0).length));
      (_1206_v0)[_index151] = (((_1211_v4).contains((((_1213_v6).contains(_1212_v5)) ? ((_1213_v6).get(_1212_v5)) : (_1209_v2)))) ? ((_1211_v4).get((((_1213_v6).contains(_1212_v5)) ? ((_1213_v6).get(_1212_v5)) : (_1209_v2)))) : (_1208_v1));
      let _1214_v7;
      _1214_v7 = _dafny.MultiSet.fromElements(_1209_v2);
      let _index152 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_1206_v0).length));
      let _index153 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_1206_v0).length));
      let _rhs198 = ((_this).f32).minus(new BigNumber(978));
      let _rhs199 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(228)), function (_1215_i1) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      });
      let _rhs200 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(547)), function (_1216_i2) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      });
      let _rhs201 = _dafny.Seq.UnicodeFromString("hucvcbfcp");
      let _rhs202 = _1214_v7;
      let _lhs153 = globalState;
      let _lhs154 = _1206_v0;
      let _lhs155 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_1206_v0).length));
      let _lhs156 = _1206_v0;
      let _lhs157 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_1206_v0).length));
      _lhs153.f18 = _rhs198;
      _lhs154[_lhs155] = _rhs199;
      _lhs156[_lhs157] = _rhs200;
      _1208_v1 = _rhs201;
      _1214_v7 = _rhs202;
      let _1217_v8;
      _1217_v8 = _dafny.Set.fromElements((_this).f32, (_this).f32);
      let _1218_v9;
      _1218_v9 = _dafny.Seq.of(_1217_v8, _1217_v8);
      let _1219_v10;
      _1219_v10 = _dafny.MultiSet.fromElements(p1, p1, p1);
      let _1220_v11;
      _1220_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1209_v2,(_this).f32);
      let _1221_v12;
      _1221_v12 = _module.D8.create_DC16((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1218_v9).length))), (((_1219_v10).contains(new BigNumber((_1220_v11).length))) ? ((_1219_v10).get(new BigNumber((_1220_v11).length))) : ((_this).f32)));
      let _pat_let_tv19 = _1209_v2;
      (globalState).f26 = function (_source20) {
        if (_source20.is_DC16) {
          let _1222___mcc_h0 = (_source20).cf29;
          let _1223___mcc_h1 = (_source20).cf30;
          let _1224_cf30 = _1223___mcc_h1;
          let _1225_cf29 = _1222___mcc_h0;
          return _pat_let_tv19;
        } else {
          let _1226___mcc_h2 = (_source20).cf28;
          let _1227_cf28 = _1226___mcc_h2;
          return true;
        }
      }(_1221_v12);
      let _1228_v13;
      _1228_v13 = _dafny.Map.Empty.slice().updateUnsafe(!(_1209_v2),_1209_v2);
      let _1229_v14;
      _1229_v14 = _dafny.Seq.of(p1);
      let _1230_v15;
      _1230_v15 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_1229_v14),false);
      let _1231_v16;
      _1231_v16 = _module.D8.create_DC15(_1228_v13);
      let _1232_v17;
      let _nw206 = Array((new BigNumber(19)).toNumber());
      _nw206[(_dafny.ZERO).toNumber()] = _1228_v13;
      _nw206[(_dafny.ONE).toNumber()] = (_1228_v13).update(_1209_v2, (((_1230_v15).contains(_dafny.MultiSet.FromArray(_1229_v14))) ? ((_1230_v15).get(_dafny.MultiSet.FromArray(_1229_v14))) : (!(_1209_v2))));
      _nw206[(new BigNumber(2)).toNumber()] = _1228_v13;
      _nw206[(new BigNumber(3)).toNumber()] = _1228_v13;
      _nw206[(new BigNumber(4)).toNumber()] = (_1228_v13).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1209_v2,!(_1209_v2)));
      _nw206[(new BigNumber(5)).toNumber()] = _1228_v13;
      _nw206[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(false,_1209_v2);
      _nw206[(new BigNumber(7)).toNumber()] = _1228_v13;
      _nw206[(new BigNumber(8)).toNumber()] = (_1228_v13).update(_1209_v2, _1209_v2);
      _nw206[(new BigNumber(9)).toNumber()] = _1228_v13;
      _nw206[(new BigNumber(10)).toNumber()] = _1228_v13;
      _nw206[(new BigNumber(11)).toNumber()] = (_1228_v13).Merge(_module.__default.fm23(p1, _1209_v2, globalState));
      _nw206[(new BigNumber(12)).toNumber()] = (_1231_v16).dtor_cf28;
      _nw206[(new BigNumber(13)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1209_v2,_1209_v2)).Merge(_1228_v13);
      _nw206[(new BigNumber(14)).toNumber()] = _1228_v13;
      _nw206[(new BigNumber(15)).toNumber()] = _1228_v13;
      _nw206[(new BigNumber(16)).toNumber()] = (_1228_v13).Merge(_1228_v13);
      _nw206[(new BigNumber(17)).toNumber()] = _1228_v13;
      _nw206[(new BigNumber(18)).toNumber()] = _1228_v13;
      _1232_v17 = _nw206;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1232_v17).length))) {
        let _1233_i3 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1233_i3)) && ((_1233_i3).isLessThan(new BigNumber((_1232_v17).length))))) {
          (_1232_v17)[(_1233_i3)] = (_1228_v13).Merge(_1228_v13);
        }
      }
      let _1234_i4;
      _1234_i4 = _dafny.ZERO;
      L4: {
        while ((new BigNumber((((_1214_v7).update(_1209_v2, _module.__default.abs(p1))).Intersect((_this).f31)).cardinality())).isEqualTo(p1)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1234_i4)) {
              break L4;
            }
            _1234_i4 = (_1234_i4).plus(_dafny.ONE);
            let _1235_v18;
            let _nw207 = new _module.C0();
            _nw207.__ctor((_this).f32, _dafny.Seq.Create(_module.__default.abs(new BigNumber(521)), function (_1236_i5) {
              return new _dafny.CodePoint('m'.codePointAt(0));
            }));
            _1235_v18 = _nw207;
            let _nw208 = new _module.C0();
            _nw208.__ctor(_module.__default.safeDivisionInt(new BigNumber((_1217_v8).length), p1), _dafny.Seq.UnicodeFromString("lwjbo"));
            _1235_v18 = _nw208;
            let _1237_v19;
            let _nw209 = Array((new BigNumber(20)).toNumber()).fill(false);
            _1237_v19 = _nw209;
            let _1238_v20;
            _1238_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v19,new BigNumber((_dafny.MultiSet.fromElements(_1214_v7, (_this).f31)).cardinality()));
            let _1239_v21;
            _1239_v21 = new _dafny.CodePoint('g'.codePointAt(0));
            let _1240_v22;
            _1240_v22 = _module.D2.create_DC4((_this).f32, new BigNumber(((_1235_v18).f37).length), (((_1219_v10).contains((((_1238_v20).contains(_1237_v19)) ? ((_1238_v20).get(_1237_v19)) : (p1)))) ? ((_1219_v10).get((((_1238_v20).contains(_1237_v19)) ? ((_1238_v20).get(_1237_v19)) : (p1)))) : ((_this).f32)), _1239_v21, _1208_v1);
            let _source21 = _1240_v22;
            if (_source21.is_DC4) {
              let _1241___mcc_h3 = (_source21).cf5;
              let _1242___mcc_h4 = (_source21).cf6;
              let _1243___mcc_h5 = (_source21).cf7;
              let _1244___mcc_h6 = (_source21).cf8;
              let _1245___mcc_h7 = (_source21).cf9;
              let _1246_cf9 = _1245___mcc_h7;
              let _1247_cf8 = _1244___mcc_h6;
              let _1248_cf7 = _1243___mcc_h5;
              let _1249_cf6 = _1242___mcc_h4;
              let _1250_cf5 = _1241___mcc_h3;
              let _1251_v23;
              _1251_v23 = _dafny.Seq.of((_1235_v18).f37);
              _1249_cf6 = ((_1209_v2) ? ((new BigNumber((_dafny.MultiSet.FromArray(_1251_v23)).cardinality())).plus((_1235_v18).f36)) : ((_1235_v18).f36));
              _1249_cf6 = new BigNumber(((_dafny.Set.fromElements(false, _1209_v2)).Difference((_dafny.Set.fromElements(_1209_v2)).Difference(_1212_v5))).length);
              let _1252_v24;
              _1252_v24 = _1209_v2;
              let _1253_v25;
              _1253_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v19,_1209_v2);
              let _1254_v26;
              _1254_v26 = _dafny.Seq.of((_1249_cf6).isLessThanOrEqualTo(p1), (_1209_v2), (((_1253_v25).contains(_1237_v19)) ? ((_1253_v25).get(_1237_v19)) : (_1209_v2)), _1209_v2, _1209_v2);
              let _1255_v27;
              _1255_v27 = _dafny.MultiSet.fromElements(_1237_v19);
              let _1256_v28;
              let _init21 = ((_1257_v10) => function (_1258_i6) {
                return _1257_v10;
              })(_1219_v10);
              let _nw210 = Array((new BigNumber(27)).toNumber());
              for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw210.length); _i0_21++) {
                _nw210[_i0_21] = _init21(new BigNumber(_i0_21));
              }
              _1256_v28 = _nw210;
              let _index154 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_1256_v28).length));
              (_1256_v28)[_index154] = ((_1209_v2) ? (_1219_v10) : (_1219_v10));
              let _index155 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_1256_v28).length));
              let _rhs203 = !(_1209_v2);
              let _rhs204 = _1254_v26;
              let _rhs205 = (_1255_v27).Difference(((_1255_v27).update(_1237_v19, _module.__default.abs((_1235_v18).f36))).Union(_1255_v27));
              let _rhs206 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1229_v14, _1229_v14))).Difference((_1219_v10).Intersect(_1219_v10));
              let _rhs207 = (_dafny.ZERO).minus(((new BigNumber(163)).multipliedBy(new BigNumber(-469))).minus(new BigNumber(883)));
              let _lhs158 = globalState;
              let _lhs159 = _1256_v28;
              let _lhs160 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_1256_v28).length));
              let _lhs161 = globalState;
              _lhs158.f26 = _rhs203;
              _1254_v26 = _rhs204;
              _1255_v27 = _rhs205;
              _lhs159[_lhs160] = _rhs206;
              _lhs161.f24 = _rhs207;
              let _1259_v29;
              let _nw211 = new _module.C0();
              _nw211.__ctor(new BigNumber((_dafny.Seq.UnicodeFromString("jlhbrjm")).length), _dafny.Seq.update(((_1209_v2) ? (_1208_v1) : (_dafny.Seq.UnicodeFromString("fjfolha"))), _module.__default.safeIndex((_1235_v18).f36, new BigNumber((((_1209_v2) ? (_1208_v1) : (_dafny.Seq.UnicodeFromString("fjfolha")))).length)), ((_1206_v0)[_module.__default.safeIndex(new BigNumber(509), new BigNumber((_1206_v0).length))])[_module.__default.safeIndex((_1235_v18).f36, new BigNumber(((_1206_v0)[_module.__default.safeIndex(new BigNumber(509), new BigNumber((_1206_v0).length))]).length))]));
              _1259_v29 = _nw211;
            } else if (_source21.is_DC3) {
              let _1260___mcc_h8 = (_source21).cf4;
              let _1261_cf4 = _1260___mcc_h8;
              (globalState).f18 = p1;
              _1239_v21 = ((_1209_v2) ? (((_1209_v2) ? (_1239_v21) : (_1239_v21))) : (_1239_v21));
              let _1262_v30;
              _1262_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1209_v2,_1228_v13);
              (globalState).f24 = new BigNumber((_1262_v30).length);
              let _1263_v31;
              let _1264_v32;
              let _out40;
              let _out41;
              let _outcollector17 = _module.__default.m0(_1209_v2, new BigNumber(935), p1, new BigNumber(299), globalState);
              _out40 = _outcollector17[0];
              _out41 = _outcollector17[1];
              _1263_v31 = _out40;
              _1264_v32 = _out41;
            } else {
              let _1265___mcc_h9 = (_source21).cf10;
              let _1266_cf10 = _1265___mcc_h9;
              _1229_v14 = _1229_v14;
              (globalState).f26 = !(_1209_v2);
              let _rhs208 = (_1235_v18).f36;
              let _rhs209 = !(_1209_v2) || (_1209_v2);
              let _rhs210 = !(((_1235_v18).f36).isLessThan(((_1235_v18).f36).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1209_v2,true)).length))));
              let _lhs162 = globalState;
              let _lhs163 = globalState;
              _lhs162.f24 = _rhs208;
              _1209_v2 = _rhs209;
              _lhs163.f26 = _rhs210;
              (globalState).f26 = _1209_v2;
            }
            let _1267_v33;
            _1267_v33 = _dafny.MultiSet.fromElements((_1206_v0)[_module.__default.safeIndex(new BigNumber(509), new BigNumber((_1206_v0).length))], _dafny.Seq.UnicodeFromString("mlmcu"), (_1235_v18).f37);
            (globalState).f8 = _1267_v33;
            let _1268_v34;
            _1268_v34 = _dafny.MultiSet.fromElements(_1217_v8, _1217_v8);
            let _1269_v35;
            _1269_v35 = _module.D11.create_DC23(_1268_v34);
            let _source22 = _1269_v35;
            if (_source22.is_DC24) {
              let _1270___mcc_h10 = (_source22).cf43;
              let _1271___mcc_h11 = (_source22).cf44;
              let _1272_cf44 = _1271___mcc_h11;
              let _1273_cf43 = _1270___mcc_h10;
              (globalState).f26 = _1209_v2;
              let _1274_v36;
              _1274_v36 = _dafny.Seq.of(_1209_v2);
              (globalState).f3 = (new BigNumber((_1274_v36).length)).multipliedBy((_1235_v18).f36);
              let _index156 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_1237_v19).length));
              (_1237_v19)[_index156] = _1209_v2;
              let _1275_v37;
              _1275_v37 = _module.D4.create_DC9(p1, _1209_v2);
              let _index157 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_1237_v19).length));
              (_1237_v19)[_index157] = (((_1228_v13).contains(!((_1275_v37).dtor_cf17) || (true))) ? ((_1228_v13).get(!((_1275_v37).dtor_cf17) || (true))) : (!(_1209_v2)));
              let _1276_v38;
              let _init22 = ((_1277_p1) => function (_1278_i7) {
                return (_1278_i7).plus(_1277_p1);
              })(p1);
              let _nw212 = Array((new BigNumber(29)).toNumber());
              for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw212.length); _i0_22++) {
                _nw212[_i0_22] = _init22(new BigNumber(_i0_22));
              }
              _1276_v38 = _nw212;
              let _index158 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_1276_v38).length));
              (_1276_v38)[_index158] = _1273_cf43;
              let _1279_v39;
              let _nw213 = Array((new BigNumber(23)).toNumber()).fill([]);
              _1279_v39 = _nw213;
              let _1280_v40;
              let _nw214 = Array((new BigNumber(10)).toNumber());
              _nw214[(_dafny.ZERO).toNumber()] = _1235_v18;
              _nw214[(_dafny.ONE).toNumber()] = _1235_v18;
              _nw214[(new BigNumber(2)).toNumber()] = _1235_v18;
              _nw214[(new BigNumber(3)).toNumber()] = _1235_v18;
              _nw214[(new BigNumber(4)).toNumber()] = _1235_v18;
              _nw214[(new BigNumber(5)).toNumber()] = _1235_v18;
              _nw214[(new BigNumber(6)).toNumber()] = _1235_v18;
              _nw214[(new BigNumber(7)).toNumber()] = _1235_v18;
              _nw214[(new BigNumber(8)).toNumber()] = _1235_v18;
              _nw214[(new BigNumber(9)).toNumber()] = _1235_v18;
              _1280_v40 = _nw214;
              let _1281_v41;
              _1281_v41 = _dafny.Seq.of(_1280_v40, _1280_v40);
              let _index159 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1279_v39).length));
              (_1279_v39)[_index159] = (_1281_v41)[_module.__default.safeIndex(new BigNumber(683), new BigNumber((_1281_v41).length))];
              let _index160 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_1276_v38).length));
              let _index161 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1279_v39).length));
              let _rhs211 = _1217_v8;
              let _rhs212 = _module.__default.safeModuloInt(_1273_cf43, _1273_cf43);
              let _rhs213 = ((((_1237_v19)[_module.__default.safeIndex(new BigNumber(844), new BigNumber((_1237_v19).length))]) ? ((_1237_v19)[_module.__default.safeIndex(new BigNumber(844), new BigNumber((_1237_v19).length))]) : ((_1237_v19)[_module.__default.safeIndex(new BigNumber(844), new BigNumber((_1237_v19).length))]))) === (_1209_v2);
              let _rhs214 = (_this).f32;
              let _rhs215 = _1280_v40;
              let _lhs164 = _1276_v38;
              let _lhs165 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_1276_v38).length));
              let _lhs166 = globalState;
              let _lhs167 = globalState;
              let _lhs168 = _1279_v39;
              let _lhs169 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1279_v39).length));
              _1217_v8 = _rhs211;
              _lhs164[_lhs165] = _rhs212;
              _lhs166.f26 = _rhs213;
              _lhs167.f15 = _rhs214;
              _lhs168[_lhs169] = _rhs215;
            } else if (_source22.is_DC23) {
              let _1282___mcc_h12 = (_source22).cf42;
              let _1283_cf42 = _1282___mcc_h12;
              _1208_v1 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(713)), ((_1284_v21) => function (_1285_i8) {
                return _1284_v21;
              })(_1239_v21)), (_1235_v18).f37);
              let _1286_v42;
              let _nw215 = new _module.C1();
              _nw215.__ctor(_1209_v2, (_this).f31, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(9)), ((_1287_v18) => function (_1288_i9) {
                return (_1287_v18).f36;
              })(_1235_v18)), _1229_v14), (_this).f32, _1206_v0);
              _1286_v42 = _nw215;
              let _1289_v43;
              _1289_v43 = _dafny.Seq.of(_1286_v42.f35);
              (globalState).f26 = (((true) ? ((_1235_v18).f36) : ((_1235_v18).f36))).isLessThan(new BigNumber((_1289_v43).length));
              let _1290_v44;
              _1290_v44 = _dafny.Seq.of(_1289_v43);
              let _1291_v45;
              _1291_v45 = _module.D6.create_DC12(!(_1286_v42.f35), _1290_v44, (((_1211_v4).contains(_1286_v42.f35)) ? ((_1211_v4).get(_1286_v42.f35)) : (_1208_v1)));
              (_1286_v42).f35 = _dafny.Seq.contains((_1291_v45).dtor_cf22, _1239_v21);
            } else {
              let _1292___mcc_h13 = (_source22).cf45;
              let _1293_cf45 = _1292___mcc_h13;
              (globalState).f26 = ((_1214_v7).Union((_1214_v7).update(_1209_v2, _module.__default.abs((_1235_v18).f36)))).IsDisjointFrom(_1214_v7);
              (globalState).f0 = (_dafny.ZERO).minus((_1235_v18).f36);
              let _1294_v46;
              let _nw216 = new _module.C1();
              _nw216.__ctor(_1209_v2, (_this).f31, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-32)), ((_1295_v18) => function (_1296_i10) {
                return (_1295_v18).f36;
              })(_1235_v18)), new BigNumber((_1212_v5).length), _1206_v0);
              _1294_v46 = _nw216;
              let _1297_v47;
              _1297_v47 = _dafny.Set.fromElements(_1294_v46);
              let _1298_v48;
              let _nw217 = Array((new BigNumber(18)).toNumber());
              _nw217[(_dafny.ZERO).toNumber()] = (_this).f32;
              _nw217[(_dafny.ONE).toNumber()] = (_1235_v18).f36;
              _nw217[(new BigNumber(2)).toNumber()] = (_this).f32;
              _nw217[(new BigNumber(3)).toNumber()] = (_1235_v18).f36;
              _nw217[(new BigNumber(4)).toNumber()] = new BigNumber(((_1238_v20).update(_1237_v19, p1)).length);
              _nw217[(new BigNumber(5)).toNumber()] = (_1235_v18).f36;
              _nw217[(new BigNumber(6)).toNumber()] = p1;
              _nw217[(new BigNumber(7)).toNumber()] = p1;
              _nw217[(new BigNumber(8)).toNumber()] = (_1235_v18).f36;
              _nw217[(new BigNumber(9)).toNumber()] = ((_1235_v18).f36).minus((_dafny.ZERO).minus(p1));
              _nw217[(new BigNumber(10)).toNumber()] = new BigNumber(204);
              _nw217[(new BigNumber(11)).toNumber()] = p1;
              _nw217[(new BigNumber(12)).toNumber()] = (_1235_v18).f36;
              _nw217[(new BigNumber(13)).toNumber()] = new BigNumber((_1208_v1).length);
              _nw217[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1297_v47).length));
              _nw217[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_1208_v1, _module.__default.safeIndex(p1, new BigNumber((_1208_v1).length)), _1239_v21)).length));
              _nw217[(new BigNumber(16)).toNumber()] = (_1235_v18).f36;
              _nw217[(new BigNumber(17)).toNumber()] = (_1235_v18).f36;
              _1298_v48 = _nw217;
              let _index162 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_1298_v48).length));
              (_1298_v48)[_index162] = (_dafny.ZERO).minus((_1235_v18).f36);
              let _pat_let_tv20 = _1239_v21;
              let _pat_let_tv21 = _1235_v18;
              let _pat_let_tv22 = globalState;
              let _index163 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_1298_v48).length));
              (_1298_v48)[_index163] = (function (_pat_let22_0) {
                return function (_1299_dt__update__tmp_h1) {
                  return function (_pat_let23_0) {
                    return function (_1300_dt__update_hcf8_h0) {
                      return _module.D2.create_DC4((_1299_dt__update__tmp_h1).dtor_cf5, (_1299_dt__update__tmp_h1).dtor_cf6, (_1299_dt__update__tmp_h1).dtor_cf7, _1300_dt__update_hcf8_h0, (_1299_dt__update__tmp_h1).dtor_cf9);
                    }(_pat_let23_0);
                  }(_module.__default.fm26(_pat_let_tv20, new BigNumber(677), (_pat_let_tv21).f36, _pat_let_tv22));
                }(_pat_let22_0);
              }(_1240_v22)).dtor_cf7;
              let _1301_v49;
              _1301_v49 = _dafny.Seq.of((((((_1228_v13).contains(_1209_v2)) ? ((_1228_v13).get(_1209_v2)) : (_1209_v2))) ? (_1298_v48) : (_1298_v48)));
              let _1302_v50;
              _1302_v50 = _module.D10.create_DC21(_1229_v14);
              _1298_v48 = (_1301_v49)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_1229_v14, (_1302_v50).dtor_cf39)).length), new BigNumber((_1301_v49).length))];
            }
          }
        }
      }
      let _1303_v51;
      _1303_v51 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements((_this).f32));
      _1303_v51 = _1303_v51;
      _1209_v2 = _1209_v2;
      return;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f27 = _dafny.ZERO;
      this._f28 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f27, f28) {
      let _this = this;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f27);
    };
    fm45(globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus((_this).f27)).minus((_this).f27);
    };
    fm46(p0, p1, globalState) {
      let _this = this;
      return ((!(false)) || (false)) === (!((((_module.D16.create_DC40(false, new BigNumber(418), (_dafny.ZERO).minus((_this).f27))).dtor_cf74) ? (true) : (true))));
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f27 = _dafny.ZERO;
      this._f28 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f27, f28) {
      let _this = this;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      if (!_dafny.Seq.contains(_dafny.Seq.of(false, true), false)) {
        return _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f27);
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f27);
      }
    };
    fm34(p0, p1, globalState) {
      let _this = this;
      return ((_this).f27).minus((_this).f27);
    };
    fm35(p0, p1, globalState) {
      let _this = this;
      return !(true);
    };
    m11(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1304_v0;
      _1304_v0 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      let _1305_v1;
      _1305_v1 = _dafny.Seq.of(p1, p1);
      let _1306_v2;
      _1306_v2 = _dafny.Seq.of(_1305_v1);
      let _1307_v3;
      _1307_v3 = _module.D6.create_DC12(p1, _1306_v2, (p0).f37);
      let _1308_v4;
      _1308_v4 = new _dafny.CodePoint('k'.codePointAt(0));
      (globalState).f18 = new BigNumber(((((((_1304_v0).contains(p1)) ? ((_1304_v0).get(p1)) : (p1))) ? ((_1307_v3).dtor_cf22) : (_dafny.Seq.update(_dafny.Seq.Concat((p0).f37, (p0).f37), _module.__default.safeIndex((_this).f27, new BigNumber((_dafny.Seq.Concat((p0).f37, (p0).f37)).length)), _1308_v4)))).length);
      let _1309_i0;
      _1309_i0 = _dafny.ZERO;
      L5: {
        while (!(!((p1) && (p1)) || (p1))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1309_i0)) {
              break L5;
            }
            _1309_i0 = (_1309_i0).plus(_dafny.ONE);
            let _1310_v5;
            _1310_v5 = _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(p1)).length), new BigNumber(277));
            let _1311_v6;
            _1311_v6 = _dafny.Set.fromElements(new BigNumber(-76));
            let _1312_v7;
            _1312_v7 = _module.D2.create_DC4((_this).f27, (p0).f36, (p0).f36, _1308_v4, _dafny.Seq.UnicodeFromString("njlrvlaga"));
            let _1313_v8;
            let _nw218 = new _module.C1();
            _nw218.__ctor(p1, _dafny.MultiSet.FromArray(_1305_v1), _dafny.Seq.Concat(_1310_v5, _1310_v5), _module.__default.safeDivisionInt((_this).fm34(_1311_v6, _1312_v7, globalState), new BigNumber((_1311_v6).length)), (_this).f28);
            _1313_v8 = _nw218;
            let _1314_v9;
            let _init23 = ((_1315_p1, _1316_v8) => function (_1317_i1) {
              return _dafny.Seq.of(_1315_p1, _1316_v8.f35);
            })(p1, _1313_v8);
            let _nw219 = Array((new BigNumber(8)).toNumber());
            for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw219.length); _i0_23++) {
              _nw219[_i0_23] = _init23(new BigNumber(_i0_23));
            }
            _1314_v9 = _nw219;
            let _index164 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_1314_v9).length));
            (_1314_v9)[_index164] = (_1306_v2)[_module.__default.safeIndex(new BigNumber((_1311_v6).length), new BigNumber((_1306_v2).length))];
            let _1318_v10;
            _1318_v10 = _dafny.MultiSet.fromElements(p1);
            let _pat_let_tv23 = _1313_v8;
            let _index165 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_1314_v9).length));
            let _rhs216 = (function (_pat_let24_0) {
              return function (_1319_dt__update__tmp_h0) {
                return function (_pat_let25_0) {
                  return function (_1320_dt__update_hcf55_h0) {
                    return _module.D13.create_DC30(_1320_dt__update_hcf55_h0);
                  }(_pat_let25_0);
                }(_pat_let_tv23);
              }(_pat_let24_0);
            }(_module.D13.create_DC30(_1313_v8))).dtor_cf55;
            let _rhs217 = true;
            let _rhs218 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.update(_1305_v1, _module.__default.safeIndex((p0).f36, new BigNumber((_1305_v1).length)), _1313_v8.f35), _module.__default.safeIndex((p0).f36, new BigNumber((_dafny.Seq.update(_1305_v1, _module.__default.safeIndex((p0).f36, new BigNumber((_1305_v1).length)), _1313_v8.f35)).length)), p1), _module.__default.safeIndex(_module.__default.safeModuloInt((p0).f36, new BigNumber(688)), new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_1305_v1, _module.__default.safeIndex((p0).f36, new BigNumber((_1305_v1).length)), _1313_v8.f35), _module.__default.safeIndex((p0).f36, new BigNumber((_dafny.Seq.update(_1305_v1, _module.__default.safeIndex((p0).f36, new BigNumber((_1305_v1).length)), _1313_v8.f35)).length)), p1)).length)), _dafny.areEqual(_dafny.Seq.of(_dafny.Seq.of(p1, (_1313_v8).fm3(_1318_v10, globalState)), _1305_v1, _1305_v1), _1306_v2));
            let _lhs170 = globalState;
            let _lhs171 = _1314_v9;
            let _lhs172 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_1314_v9).length));
            _1313_v8 = _rhs216;
            _lhs170.f26 = _rhs217;
            _lhs171[_lhs172] = _rhs218;
            (globalState).f3 = (p0).f36;
            (globalState).f26 = (p1) && (!(_1313_v8.f35) || ((_1305_v1)[_module.__default.safeIndex((_this).f27, new BigNumber((_1305_v1).length))]));
            if (p1) {
              let _1321_v11;
              _1321_v11 = _dafny.Seq.of(_1312_v7, _module.__default.fm36(!(p1), globalState));
              let _pat_let_tv24 = p0;
              let _pat_let_tv25 = globalState;
              _1321_v11 = _dafny.Seq.of(function (_pat_let26_0) {
                return function (_1322_dt__update__tmp_h1) {
                  return function (_pat_let27_0) {
                    return function (_1323_dt__update_hcf6_h0) {
                      return function (_pat_let28_0) {
                        return function (_1324_dt__update_hcf5_h0) {
                          return _module.D2.create_DC4(_1324_dt__update_hcf5_h0, _1323_dt__update_hcf6_h0, (_1322_dt__update__tmp_h1).dtor_cf7, (_1322_dt__update__tmp_h1).dtor_cf8, (_1322_dt__update__tmp_h1).dtor_cf9);
                        }(_pat_let28_0);
                      }(_module.__default.fm1(_pat_let_tv25));
                    }(_pat_let27_0);
                  }((_pat_let_tv24).f36);
                }(_pat_let26_0);
              }(_1312_v7), _1312_v7, _1312_v7);
              let _1325_v12;
              _1325_v12 = _dafny.Set.fromElements(_1313_v8.f35);
              let _1326_v13;
              let _nw220 = new _module.C2();
              _nw220.__ctor(_1318_v10, (p0).f36);
              _1326_v13 = _nw220;
              let _1327_v14;
              _1327_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1313_v8,_1326_v13);
              let _1328_v15;
              _1328_v15 = _dafny.Seq.of((p0).f37);
              let _1329_v16;
              let _nw221 = Array((new BigNumber(11)).toNumber());
              _nw221[(_dafny.ZERO).toNumber()] = new BigNumber(((p0).f37).length);
              _nw221[(_dafny.ONE).toNumber()] = (p0).f36;
              _nw221[(new BigNumber(2)).toNumber()] = new BigNumber(((_1314_v9)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_1314_v9).length))]).length);
              _nw221[(new BigNumber(3)).toNumber()] = new BigNumber(((p0).f37).length);
              _nw221[(new BigNumber(4)).toNumber()] = (_this).f27;
              _nw221[(new BigNumber(5)).toNumber()] = (p0).f36;
              _nw221[(new BigNumber(6)).toNumber()] = (_1326_v13).f32;
              _nw221[(new BigNumber(7)).toNumber()] = (p0).f36;
              _nw221[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(557)), ((_1330_v4) => function (_1331_i2) {
                return _1330_v4;
              })(_1308_v4))).length);
              _nw221[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(693)), ((_1332_p0) => function (_1333_i3) {
                return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(106)), ((_1334_p0) => function (_1335_i4) {
                  return ((_1334_p0).f37)[_module.__default.safeIndex((_this).f27, new BigNumber(((_1334_p0).f37).length))];
                })(_1332_p0))).length);
              })(p0))).length);
              _nw221[(new BigNumber(10)).toNumber()] = (p0).f36;
              _1329_v16 = _nw221;
              let _1336_v17;
              _1336_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1305_v1,_1329_v16);
              let _1337_v18;
              _1337_v18 = _dafny.Seq.of((((_1336_v17).contains((_1314_v9)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_1314_v9).length))])) ? ((_1336_v17).get((_1314_v9)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_1314_v9).length))])) : (_1329_v16)), _1329_v16, _1329_v16);
              let _1338_v19;
              _1338_v19 = _dafny.Map.Empty.slice().updateUnsafe((p0).f36,(_1326_v13).f32);
              let _1339_v20;
              _1339_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1313_v8.f35,new BigNumber((_1338_v19).length));
              let _1340_v21;
              _1340_v21 = _dafny.MultiSet.fromElements((_1337_v18)[_module.__default.safeIndex(new BigNumber((_1339_v20).length), new BigNumber((_1337_v18).length))], _1329_v16);
              let _1341_v22;
              _1341_v22 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_1329_v16, _1329_v16, _1329_v16), _1340_v21, _1340_v21);
              let _1342_v23;
              _1342_v23 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm37((p0).f36, (_this).fm34(_1311_v6, _module.D2.create_DC4(new BigNumber((_1325_v12).length), new BigNumber((_1327_v14).length), new BigNumber(662), _1308_v4, (_1328_v15)[_module.__default.safeIndex((p0).f36, new BigNumber((_1328_v15).length))]), globalState), globalState),(((_1341_v22).contains(_dafny.MultiSet.fromElements(_1329_v16))) ? ((_1341_v22).get(_dafny.MultiSet.fromElements(_1329_v16))) : ((_dafny.ZERO).minus((_this).f27))));
              (globalState).f24 = (_dafny.ZERO).minus((((_1342_v23).contains((p0).f37)) ? ((_1342_v23).get((p0).f37)) : ((p0).f36)));
              let _1343_v24;
              _1343_v24 = _dafny.Seq.UnicodeFromString("tqatqn");
              _1343_v24 = _dafny.Seq.Concat(_dafny.Seq.Concat((p0).f37, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("jqpon"), _module.__default.safeIndex((_1326_v13).f32, new BigNumber((_dafny.Seq.UnicodeFromString("jqpon")).length)), _1308_v4)), _dafny.Seq.UnicodeFromString("qq"));
              r0 = new BigNumber(695);
              (globalState).f26 = _1313_v8.f35;
            } else {
              let _1344_v25;
              _1344_v25 = _dafny.Seq.UnicodeFromString("kjg");
              let _1345_v26;
              _1345_v26 = _dafny.Seq.of(_dafny.Seq.of(_1308_v4, new _dafny.CodePoint('a'.codePointAt(0))));
              let _1346_v27;
              _1346_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1308_v4,new BigNumber(985));
              let _1347_v28;
              _1347_v28 = _dafny.MultiSet.fromElements(new BigNumber((_1346_v27).length));
              let _1348_v29;
              _1348_v29 = _module.D9.create_DC19(p1);
              let _1349_v30;
              let _nw222 = Array((new BigNumber(17)).toNumber());
              _nw222[(_dafny.ZERO).toNumber()] = _1313_v8.f35;
              _nw222[(_dafny.ONE).toNumber()] = ((p1) ? (false) : (_1313_v8.f35));
              _nw222[(new BigNumber(2)).toNumber()] = true;
              _nw222[(new BigNumber(3)).toNumber()] = false;
              _nw222[(new BigNumber(4)).toNumber()] = _1313_v8.f35;
              _nw222[(new BigNumber(5)).toNumber()] = _dafny.Seq.contains(_dafny.Seq.update(_1345_v26, _module.__default.safeIndex((_this).f27, new BigNumber((_1345_v26).length)), (p0).f37), _dafny.Seq.of(_1308_v4));
              _nw222[(new BigNumber(6)).toNumber()] = ((p1) ? ((_this).fm35(_1308_v4, _module.__default.fm1(globalState), globalState)) : (_1313_v8.f35));
              _nw222[(new BigNumber(7)).toNumber()] = true;
              _nw222[(new BigNumber(8)).toNumber()] = !((p0).f36).isEqualTo((p0).f36);
              _nw222[(new BigNumber(9)).toNumber()] = (_1347_v28).IsProperSubsetOf(_dafny.MultiSet.FromArray(_1310_v5));
              _nw222[(new BigNumber(10)).toNumber()] = ((p0).f36).isLessThanOrEqualTo((p0).f36);
              _nw222[(new BigNumber(11)).toNumber()] = (_1318_v10).equals(_dafny.MultiSet.fromElements(_1313_v8.f35, true));
              _nw222[(new BigNumber(12)).toNumber()] = !(_1313_v8.f35);
              _nw222[(new BigNumber(13)).toNumber()] = p1;
              _nw222[(new BigNumber(14)).toNumber()] = _1313_v8.f35;
              _nw222[(new BigNumber(15)).toNumber()] = (_1348_v29).dtor_cf33;
              _nw222[(new BigNumber(16)).toNumber()] = false;
              _1349_v30 = _nw222;
              let _index166 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1349_v30).length));
              (_1349_v30)[_index166] = !(p1);
              let _index167 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1349_v30).length));
              let _rhs219 = _dafny.Seq.UnicodeFromString("gpkm");
              let _rhs220 = !(!(p1) || (_1313_v8.f35));
              let _lhs173 = _1349_v30;
              let _lhs174 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1349_v30).length));
              _1344_v25 = _rhs219;
              _lhs173[_lhs174] = _rhs220;
              let _1350_v31;
              let _nw223 = new _module.C0();
              _nw223.__ctor(new BigNumber((_1344_v25).length), (p0).f37);
              _1350_v31 = _nw223;
              (globalState).f0 = _module.__default.fm1(globalState);
              let _1351_v32;
              _1351_v32 = _module.D4.create_DC8(_1308_v4, new BigNumber(444), p1);
              let _1352_v33;
              _1352_v33 = _module.D11.create_DC24((_1350_v31).f36, _1351_v32);
              _1352_v33 = _1352_v33;
              let _1353_v34;
              _1353_v34 = _module.D13.create_DC31((p0).f36);
              let _1354_v35;
              _1354_v35 = _module.D13.create_DC33(_1353_v34);
              _1354_v35 = _1354_v35;
            }
          }
        }
      }
      (globalState).f3 = (_this).f27;
      (globalState).f21 = ((p0).f36).minus(new BigNumber(-794));
      let _rhs221 = !(p1);
      let _rhs222 = p1;
      let _rhs223 = !(p1);
      let _lhs175 = globalState;
      let _lhs176 = globalState;
      let _lhs177 = globalState;
      _lhs175.f26 = _rhs221;
      _lhs176.f26 = _rhs222;
      _lhs177.f26 = _rhs223;
      let _1355_v36;
      _1355_v36 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f27);
      let _1356_v37;
      _1356_v37 = _dafny.Set.fromElements(_1355_v36, (_1355_v36).update(true, new BigNumber(677)), (_1355_v36).update(_module.__default.fm0(new BigNumber((_1305_v1).length), globalState), (_this).f27), _1355_v36);
      let _hi5 = new BigNumber(((_1356_v37).Difference(_1356_v37)).length);
      for (let _1357_i5 = _module.__default.safeModuloInt((_this).f27, (_this).f27); _1357_i5.isLessThan(_hi5); _1357_i5 = _1357_i5.plus(_dafny.ONE)) {
        if (!(((p0).f36).isEqualTo(new BigNumber(431)))) {
          let _1358_v38;
          _1358_v38 = _dafny.Map.Empty.slice().updateUnsafe((p0).f36,_1305_v1);
          _1358_v38 = (_1358_v38).update(((_this).f27).minus(_1357_i5), _dafny.Seq.of(p1));
          let _1359_v39;
          _1359_v39 = _dafny.Map.Empty.slice().updateUnsafe((p0).f36,_module.__default.fm0((p0).f36, globalState));
          let _1360_v40;
          _1360_v40 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1359_v39);
          let _1361_v41;
          let _nw224 = Array((new BigNumber(4)).toNumber());
          _nw224[(_dafny.ZERO).toNumber()] = _1360_v40;
          _nw224[(_dafny.ONE).toNumber()] = (_1360_v40).Merge(_1360_v40);
          _nw224[(new BigNumber(2)).toNumber()] = _1360_v40;
          _nw224[(new BigNumber(3)).toNumber()] = (_1360_v40).update(p1, _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1357_i5),false));
          _1361_v41 = _nw224;
          let _index168 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_1361_v41).length));
          (_1361_v41)[_index168] = _1360_v40;
          let _index169 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_1361_v41).length));
          (_1361_v41)[_index169] = ((_1360_v40).Merge(_1360_v40)).Merge(_1360_v40);
          let _1362_v42;
          _1362_v42 = _dafny.MultiSet.fromElements(new BigNumber(422));
          let _1363_v43;
          let _nw225 = new _module.C0();
          _nw225.__ctor(((p1) ? ((_dafny.ZERO).minus((p0).f36)) : (new BigNumber((_1362_v42).cardinality()))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(46)), function (_1364_i6) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          }));
          _1363_v43 = _nw225;
          let _1365_v44;
          let _nw226 = Array((new BigNumber(24)).toNumber()).fill([]);
          _1365_v44 = _nw226;
          let _nw227 = Array((new BigNumber(29)).toNumber()).fill([]);
          _1365_v44 = _nw227;
          let _1366_v45;
          let _nw228 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1366_v45 = _nw228;
          let _index170 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_1366_v45).length));
          (_1366_v45)[_index170] = new _dafny.CodePoint('c'.codePointAt(0));
          let _index171 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_1366_v45).length));
          (_1366_v45)[_index171] = _module.__default.fm38(_dafny.Map.Empty.slice().updateUnsafe(true,false), globalState);
        } else {
          let _1367_v46;
          let _init24 = ((_1368_v4, _1369_i5, _1370_p1) => function (_1371_i7) {
            return _module.D4.create_DC8(_1368_v4, _1369_i5, _1370_p1);
          })(_1308_v4, _1357_i5, p1);
          let _nw229 = Array((new BigNumber(4)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw229.length); _i0_24++) {
            _nw229[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _1367_v46 = _nw229;
          let _1372_v47;
          let _nw230 = Array((new BigNumber(26)).toNumber()).fill(_module.D10.Default());
          _1372_v47 = _nw230;
          let _1373_v48;
          _1373_v48 = _dafny.Seq.of(new BigNumber(320));
          let _1374_v49;
          _1374_v49 = _module.D10.create_DC21(_1373_v48);
          let _index172 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_1372_v47).length));
          (_1372_v47)[_index172] = _1374_v49;
          let _index173 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_1372_v47).length));
          let _rhs224 = (_module.__default.safeModuloInt(new BigNumber(691), (p0).f36)).minus((p0).f36);
          let _rhs225 = p1;
          let _rhs226 = _1304_v0;
          let _rhs227 = _1367_v46;
          let _rhs228 = _1374_v49;
          let _lhs178 = globalState;
          let _lhs179 = globalState;
          let _lhs180 = _1372_v47;
          let _lhs181 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_1372_v47).length));
          _lhs178.f21 = _rhs224;
          _lhs179.f26 = _rhs225;
          _1304_v0 = _rhs226;
          _1367_v46 = _rhs227;
          _lhs180[_lhs181] = _rhs228;
          let _1375_v51;
          _1375_v51 = _dafny.Set.fromElements((p0).f36);
          let _1376_v52;
          _1376_v52 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
          let _1377_v53;
          _1377_v53 = _module.D2.create_DC4((_dafny.ZERO).minus((_this).f27), new BigNumber((_1376_v52).length), new BigNumber(-286), _1308_v4, (p0).f37);
          let _1378_v54;
          _1378_v54 = _dafny.Map.Empty.slice().updateUnsafe((p0).f36,p1);
          let _1379_v55;
          _1379_v55 = _dafny.MultiSet.fromElements(_module.__default.fm0((_this).f27, globalState));
          let _1380_v56;
          let _nw231 = Array((new BigNumber(28)).toNumber());
          _nw231[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(_1357_i5, _1357_i5);
          _nw231[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_1357_i5);
          _nw231[(new BigNumber(2)).toNumber()] = (_1357_i5).multipliedBy((p0).f36);
          _nw231[(new BigNumber(3)).toNumber()] = _1357_i5;
          _nw231[(new BigNumber(4)).toNumber()] = (_1373_v48)[_module.__default.safeIndex(_1357_i5, new BigNumber((_1373_v48).length))];
          _nw231[(new BigNumber(5)).toNumber()] = new BigNumber((_1305_v1).length);
          _nw231[(new BigNumber(6)).toNumber()] = (p0).f36;
          _nw231[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt((p0).f36, new BigNumber((function () {
            let _coll45 = new _dafny.Set();
            for (const _compr_45 of _dafny.IntegerRange(new BigNumber(865), new BigNumber(146))) {
              let _1381_v50 = _compr_45;
              if (((new BigNumber(865)).isLessThanOrEqualTo(_1381_v50)) && ((_1381_v50).isLessThan(new BigNumber(146)))) {
                _coll45.add((_1381_v50).multipliedBy((p0).f36));
              }
            }
            return _coll45;
          }()).length));
          _nw231[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus((_this).f27);
          _nw231[(new BigNumber(9)).toNumber()] = (p0).f36;
          _nw231[(new BigNumber(10)).toNumber()] = (_this).f27;
          _nw231[(new BigNumber(11)).toNumber()] = (p0).f36;
          _nw231[(new BigNumber(12)).toNumber()] = ((_this).fm34(_1375_v51, _1377_v53, globalState)).minus((_dafny.ZERO).minus(new BigNumber((_1378_v54).length)));
          _nw231[(new BigNumber(13)).toNumber()] = (p0).f36;
          _nw231[(new BigNumber(14)).toNumber()] = (((_1355_v36).contains(p1)) ? ((_1355_v36).get(p1)) : ((_this).f27));
          _nw231[(new BigNumber(15)).toNumber()] = (((_1379_v55).contains(p1)) ? ((_1379_v55).get(p1)) : ((p0).f36));
          _nw231[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus((p0).f36);
          _nw231[(new BigNumber(17)).toNumber()] = _1357_i5;
          _nw231[(new BigNumber(18)).toNumber()] = (_this).f27;
          _nw231[(new BigNumber(19)).toNumber()] = (p0).f36;
          _nw231[(new BigNumber(20)).toNumber()] = (p0).f36;
          _nw231[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(300)), ((_1382_v4) => function (_1383_i8) {
            return _1382_v4;
          })(_1308_v4))).length);
          _nw231[(new BigNumber(22)).toNumber()] = ((p0).f36).plus((p0).f36);
          _nw231[(new BigNumber(23)).toNumber()] = _1357_i5;
          _nw231[(new BigNumber(24)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(594), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(884)), function (_1384_i9) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          })).length));
          _nw231[(new BigNumber(25)).toNumber()] = (_this).f27;
          _nw231[(new BigNumber(26)).toNumber()] = _1357_i5;
          _nw231[(new BigNumber(27)).toNumber()] = (p0).f36;
          _1380_v56 = _nw231;
          let _index174 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_1380_v56).length));
          (_1380_v56)[_index174] = _module.__default.safeDivisionInt((p0).f36, new BigNumber((_1373_v48).length));
          let _index175 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_1380_v56).length));
          (_1380_v56)[_index175] = _1357_i5;
          (globalState).f3 = _1357_i5;
          let _1385_v57;
          let _nw232 = Array((new BigNumber(29)).toNumber()).fill(false);
          _1385_v57 = _nw232;
          _1385_v57 = (((_dafny.MultiSet.fromElements(p1, p1, !(_module.__default.fm0(new BigNumber(736), globalState)))).IsSubsetOf(_1379_v55)) ? (_1385_v57) : (_1385_v57));
          let _1386_v58;
          _1386_v58 = _dafny.Seq.UnicodeFromString("efdl");
          _1386_v58 = _module.__default.fm37((_this).f27, _1357_i5, globalState);
        }
        (globalState).f26 = !(p1);
        let _1387_v59;
        let _nw233 = Array((new BigNumber(20)).toNumber()).fill(_module.D13.Default());
        _1387_v59 = _nw233;
        let _1388_v60;
        _1388_v60 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_1305_v1).length)), (_this).f27);
        let _1389_v61;
        _1389_v61 = _dafny.MultiSet.fromElements(p1, p1);
        let _1390_v62;
        _1390_v62 = _module.D13.create_DC32((((_1388_v60).contains((p0).f36)) ? ((_1388_v60).get((p0).f36)) : (_1357_i5)), new BigNumber((_1389_v61).cardinality()), _dafny.Seq.of(_1357_i5, (_dafny.ZERO).minus(new BigNumber((_1304_v0).length))), p1);
        let _index176 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_1387_v59).length));
        (_1387_v59)[_index176] = ((p1) ? (_1390_v62) : (_1390_v62));
        let _1391_v63;
        _1391_v63 = _dafny.Set.fromElements(_1305_v1);
        let _1392_v64;
        _1392_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_1391_v63);
        let _1393_v65;
        _1393_v65 = _dafny.Set.fromElements(_1305_v1);
        let _index177 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_1387_v59).length));
        let _rhs229 = ((p0).f36).minus((p0).f36);
        let _rhs230 = (p0).f36;
        let _rhs231 = ((_1393_v65)).IsSubsetOf((((_1392_v64).contains((p0).f36)) ? ((_1392_v64).get((p0).f36)) : (_dafny.Set.fromElements(_1305_v1))));
        let _rhs232 = _1390_v62;
        let _lhs182 = globalState;
        let _lhs183 = globalState;
        let _lhs184 = globalState;
        let _lhs185 = _1387_v59;
        let _lhs186 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_1387_v59).length));
        _lhs182.f0 = _rhs229;
        _lhs183.f21 = _rhs230;
        _lhs184.f26 = _rhs231;
        _lhs185[_lhs186] = _rhs232;
        _1308_v4 = _1308_v4;
      }
      r0 = (_this).f27;
      return r0;
    }
    m12(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let r2 = [];
      let r3 = false;
      if (!(p2)) {
        let _1394_v0;
        _1394_v0 = _dafny.Seq.UnicodeFromString("t");
        let _1395_v1;
        _1395_v1 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qppgnht"), _1394_v0), _module.__default.fm37((_this).f27, new BigNumber((_1394_v0).length), globalState));
        let _1396_v2;
        _1396_v2 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
        let _1397_v3;
        _1397_v3 = _dafny.MultiSet.fromElements(p3);
        let _1398_v4;
        _1398_v4 = _dafny.Seq.of(p0, p2, false, p0);
        let _1399_v5;
        _1399_v5 = _dafny.Seq.of(_1398_v4);
        let _1400_v6;
        _1400_v6 = _module.D6.create_DC12(p0, _1399_v5, _1394_v0);
        let _rhs233 = _1395_v1;
        let _rhs234 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1400_v6).dtor_cf20);
        let _rhs235 = _1397_v3;
        let _rhs236 = _module.__default.safeDivisionInt((_this).f27, (_module.D4.create_DC8(p3, new BigNumber((_dafny.Set.fromElements(p2)).length), p2)).dtor_cf14);
        let _lhs187 = globalState;
        _1395_v1 = _rhs233;
        _1396_v2 = _rhs234;
        _1397_v3 = _rhs235;
        _lhs187.f0 = _rhs236;
        _1394_v0 = _dafny.Seq.Concat(_1394_v0, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("m"), _1394_v0));
        _1395_v1 = _1395_v1;
        (globalState).f24 = ((!(p0)) ? ((_this).f27) : ((_this).f27));
        (globalState).f21 = (_this).f27;
      } else {
        let _1401_v7;
        let _init25 = ((_1402_p0) => function (_1403_i0) {
          return _1402_p0;
        })(p0);
        let _nw234 = Array((new BigNumber(2)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw234.length); _i0_25++) {
          _nw234[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _1401_v7 = _nw234;
        let _index178 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length));
        (_1401_v7)[_index178] = (true) && (p0);
        let _1404_v8;
        _1404_v8 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
        let _index179 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length));
        (_1401_v7)[_index179] = !(!(((p2) ? (p0) : ((_this).fm35(_module.__default.fm38(_1404_v8, globalState), (_this).f27, globalState)))) || (false));
        let _1405_v11;
        _1405_v11 = _dafny.Seq.UnicodeFromString("g");
        let _1406_v12;
        _1406_v12 = _module.D2.create_DC4(new BigNumber((_1405_v11).length), (_this).f27, (_this).f27, p3, _1405_v11);
        r3 = (((_this).f27).plus((_this).f27)).isEqualTo((_this).fm34(function () {
          let _coll46 = new _dafny.Set();
          for (const _compr_46 of _dafny.IntegerRange(new BigNumber(234), new BigNumber(-710))) {
            let _1407_v9 = _compr_46;
            if (((new BigNumber(234)).isLessThanOrEqualTo(_1407_v9)) && ((_1407_v9).isLessThan(new BigNumber(-710)))) {
              _coll46.add(_module.__default.safeModuloInt(_1407_v9, new BigNumber((function () {
                let _coll47 = new _dafny.Map();
                for (const _compr_47 of _dafny.IntegerRange(new BigNumber(-807), new BigNumber(827))) {
                  let _1408_v10 = _compr_47;
                  if (((new BigNumber(-807)).isLessThanOrEqualTo(_1408_v10)) && ((_1408_v10).isLessThan(new BigNumber(827)))) {
                    _coll47.push([_module.__default.safeDivisionInt(_1408_v10, (_this).f27),p3]);
                  }
                }
                return _coll47;
              }()).length)));
            }
          }
          return _coll46;
        }(), _1406_v12, globalState));
        if (p2) {
          let _1409_v13;
          _1409_v13 = _module.D11.create_DC25(_module.__default.fm39(globalState));
          let _1410_v14;
          let _nw235 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
          _1410_v14 = _nw235;
          let _1411_v15;
          _1411_v15 = _dafny.Seq.of(p0, p2);
          let _1412_v16;
          _1412_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1405_v11,_1411_v15);
          let _index180 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_1410_v14).length));
          (_1410_v14)[_index180] = (_1412_v16).update(_1405_v11, _1411_v15);
          let _1413_v17;
          _1413_v17 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f27);
          let _1414_v18;
          _1414_v18 = _module.D11.create_DC24(new BigNumber(166), _module.D4.create_DC8(p3, new BigNumber((_1413_v17).length), p2));
          let _1415_v19;
          _1415_v19 = _dafny.MultiSet.fromElements(p0);
          let _1416_v20;
          let _nw236 = new _module.C2();
          _nw236.__ctor(_1415_v19, (new BigNumber((_1405_v11).length)).minus((_this).f27));
          _1416_v20 = _nw236;
          let _1417_v21;
          _1417_v21 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f27,new BigNumber(437))).length));
          let _1418_v22;
          _1418_v22 = _dafny.Set.fromElements((((_1417_v21).contains(new BigNumber((_1405_v11).length))) ? ((_1417_v21).get(new BigNumber((_1405_v11).length))) : ((_this).f27)), (_this).f27);
          let _1419_v23;
          _1419_v23 = _dafny.MultiSet.fromElements(_1418_v22, _1418_v22);
          let _1420_v24;
          _1420_v24 = _module.D11.create_DC23(_1419_v23);
          let _1421_v25;
          _1421_v25 = _module.D11.create_DC25(_1420_v24);
          let _1422_v26;
          _1422_v26 = _module.D12.create_DC26(_1416_v20);
          let _1423_v27;
          _1423_v27 = _dafny.Seq.of((_1422_v26).dtor_cf46, _1416_v20, _1416_v20);
          let _index181 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_1410_v14).length));
          let _rhs237 = _module.D11.create_DC25(_1421_v25);
          let _rhs238 = _module.__default.fm40(new BigNumber(154), p0, !((_1401_v7)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length))]) || (p2), _module.__default.fm41((_this).f27, (_this).f27, p2, globalState), globalState);
          let _rhs239 = _module.__default.fm39(globalState);
          let _rhs240 = new BigNumber(-619);
          let _rhs241 = (_1423_v27)[_module.__default.safeIndex((_this).f27, new BigNumber((_1423_v27).length))];
          let _lhs188 = _1410_v14;
          let _lhs189 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_1410_v14).length));
          let _lhs190 = globalState;
          _1409_v13 = _rhs237;
          _lhs188[_lhs189] = _rhs238;
          _1414_v18 = _rhs239;
          _lhs190.f3 = _rhs240;
          _1416_v20 = _rhs241;
          let _index182 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length));
          (_1401_v7)[_index182] = p2;
          let _1424_v28;
          _1424_v28 = _dafny.Seq.of(_1405_v11, _1405_v11);
          let _index183 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length));
          let _rhs242 = (_1424_v28)[_module.__default.safeIndex((_this).f27, new BigNumber((_1424_v28).length))];
          let _rhs243 = (_this).f27;
          let _rhs244 = new BigNumber(-720);
          let _rhs245 = !((new BigNumber(848)).isLessThan((_this).f27));
          let _lhs191 = globalState;
          let _lhs192 = globalState;
          let _lhs193 = _1401_v7;
          let _lhs194 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length));
          _1405_v11 = _rhs242;
          _lhs191.f15 = _rhs243;
          _lhs192.f0 = _rhs244;
          _lhs193[_lhs194] = _rhs245;
          let _1425_v29;
          let _nw237 = new _module.C0();
          _nw237.__ctor(new BigNumber((_dafny.Seq.Concat(_1411_v15, _1411_v15)).length), _1405_v11);
          _1425_v29 = _nw237;
          let _1426_v30;
          _1426_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(754),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(478),true)).length));
          let _1427_v31;
          _1427_v31 = _dafny.Map.Empty.slice().updateUnsafe((_1425_v29).f36,p2);
          let _1428_v32;
          _1428_v32 = _module.D9.create_DC20(new BigNumber((_1411_v15).length), (((_1427_v31).contains((_1425_v29).f36)) ? ((_1427_v31).get((_1425_v29).f36)) : ((_1401_v7)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length))])), p0, (_1401_v7)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length))], _1426_v30);
          let _1429_v34;
          _1429_v34 = _dafny.Seq.of((_1425_v29).f36, _module.__default.fm1(globalState));
          let _1430_v37;
          _1430_v37 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll48 = new _dafny.Set();
            for (const _compr_48 of (_1429_v34).Elements) {
              let _1431_v36 = _compr_48;
              if (_dafny.Seq.contains(_1429_v34, _1431_v36)) {
                _coll48.add(_module.__default.safeDivisionInt(_1431_v36, new BigNumber(29)));
              }
            }
            return _coll48;
          }()).length),(_1425_v29).f36), _1426_v30);
          let _1432_v38;
          _1432_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1411_v15).length),p3);
          let _1433_v39;
          _1433_v39 = _dafny.Seq.of(_1432_v38);
          let _1434_v40;
          _1434_v40 = _module.D15.create_DC35(_1433_v39);
          let _1435_v41;
          let _nw238 = Array((new BigNumber(29)).toNumber());
          _nw238[(_dafny.ZERO).toNumber()] = (_1426_v30).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1425_v29).f37).length),(_1425_v29).f36));
          _nw238[(_dafny.ONE).toNumber()] = (_1428_v32).dtor_cf38;
          _nw238[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f27);
          _nw238[(new BigNumber(3)).toNumber()] = (_1426_v30).Merge(function () {
            let _coll49 = new _dafny.Map();
            for (const _compr_49 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1429_v34).length),(_1425_v29).f36)).Keys.Elements) {
              let _1436_v33 = _compr_49;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1429_v34).length),(_1425_v29).f36)).contains(_1436_v33)) {
                _coll49.push([(_1436_v33).plus((_1425_v29).f36),new BigNumber(319)]);
              }
            }
            return _coll49;
          }());
          _nw238[(new BigNumber(4)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(5)).toNumber()] = ((false) ? (_1426_v30) : (_1426_v30));
          _nw238[(new BigNumber(6)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(7)).toNumber()] = (_1426_v30).Merge(function () {
            let _coll50 = new _dafny.Map();
            for (const _compr_50 of _dafny.IntegerRange(new BigNumber(32), new BigNumber(657))) {
              let _1437_v35 = _compr_50;
              if (((new BigNumber(32)).isLessThanOrEqualTo(_1437_v35)) && ((_1437_v35).isLessThan(new BigNumber(657)))) {
                _coll50.push([_module.__default.safeModuloInt(_1437_v35, (_1425_v29).f36),(_1425_v29).f36]);
              }
            }
            return _coll50;
          }());
          _nw238[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f27);
          _nw238[(new BigNumber(9)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(10)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(11)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(12)).toNumber()] = _module.__default.fm42((((_1415_v19).contains(p2)) ? ((_1415_v19).get(p2)) : ((_this).f27)), true, (_1425_v29).f36, globalState);
          _nw238[(new BigNumber(13)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(14)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(15)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(16)).toNumber()] = (_1426_v30).Merge(_1426_v30);
          _nw238[(new BigNumber(17)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(18)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(19)).toNumber()] = (_1430_v37)[_module.__default.safeIndex(new BigNumber(((_dafny.MultiSet.FromArray((_1434_v40).dtor_cf63)).update(_1432_v38, _module.__default.abs((_1425_v29).f36))).cardinality()), new BigNumber((_1430_v37).length))];
          _nw238[(new BigNumber(20)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(21)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(22)).toNumber()] = (_1426_v30).Merge((_1426_v30).update((_this).fm34(_1418_v22, _1406_v12, globalState), (_1425_v29).f36));
          _nw238[(new BigNumber(23)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(24)).toNumber()] = (_1426_v30).Merge(((_1428_v32).dtor_cf38).update((_1425_v29).f36, (_this).f27));
          _nw238[(new BigNumber(25)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(26)).toNumber()] = _1426_v30;
          _nw238[(new BigNumber(27)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1411_v15).length),(_1425_v29).f36);
          _nw238[(new BigNumber(28)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_1425_v29).f36);
          _1435_v41 = _nw238;
          let _1438_v42;
          _1438_v42 = _module.D16.create_DC39(_1435_v41);
          let _1439_v43;
          _1439_v43 = (_1401_v7)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length))];
          let _1440_v44;
          _1440_v44 = _module.D9.create_DC18(true);
          let _rhs246 = (_1425_v29).f36;
          let _rhs247 = (_1438_v42).dtor_cf73;
          let _rhs248 = true;
          let _rhs249 = _dafny.Seq.Concat(_module.__default.fm43(p2, p2, p0, globalState), _1429_v34);
          let _rhs250 = !(false) || ((_1440_v44).dtor_cf32);
          let _lhs195 = globalState;
          let _lhs196 = globalState;
          _lhs195.f21 = _rhs246;
          _1435_v41 = _rhs247;
          _lhs196.f26 = _rhs248;
          _1429_v34 = _rhs249;
          r1 = _rhs250;
        } else {
          let _1441_v45;
          _1441_v45 = _module.D9.create_DC19(p0);
          _1441_v45 = (((_1401_v7)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length))]) ? (_1441_v45) : (_1441_v45));
          let _1442_v46;
          let _nw239 = new _module.C0();
          _nw239.__ctor((_this).f27, _1405_v11);
          _1442_v46 = _nw239;
          let _1443_v47;
          let _init26 = ((_1444_v46) => function (_1445_i1) {
            return (_1445_i1).minus((_1444_v46).f36);
          })(_1442_v46);
          let _nw240 = Array((new BigNumber(20)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw240.length); _i0_26++) {
            _nw240[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _1443_v47 = _nw240;
          let _1446_v48;
          _1446_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,p2);
          let _1447_v49;
          _1447_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1443_v47,new BigNumber((_1446_v48).length));
          let _1448_v50;
          _1448_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1447_v49,p0);
          (globalState).f21 = (((((_1448_v50).contains(_1447_v49)) ? ((_1448_v50).get(_1447_v49)) : (p2))) ? ((_this).f27) : (new BigNumber(729)));
          (globalState).f26 = (new BigNumber((_1405_v11).length)).isLessThan((_1442_v46).f36);
          let _index184 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length));
          (_1401_v7)[_index184] = p0;
        }
        r3 = (_1401_v7)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length))];
        let _index185 = _module.__default.safeIndex(new BigNumber(748), new BigNumber(((_this).f28).length));
        ((_this).f28)[_index185] = _dafny.Seq.UnicodeFromString("naxcrpkh");
        let _1449_v51;
        _1449_v51 = _dafny.MultiSet.fromElements((_1401_v7)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length))]);
        let _1450_v52;
        _1450_v52 = _module.D15.create_DC38((_this).f27, true, new BigNumber(245));
        let _1451_v53;
        _1451_v53 = (_1401_v7)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length))];
        let _1452_v55;
        _1452_v55 = _module.D16.create_DC40(!((_1449_v51).IsProperSubsetOf(_1449_v51)), (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).fm34(function () {
  let _coll51 = new _dafny.Set();
  for (const _compr_51 of (_module.__default.fm43((_1450_v52).dtor_cf71, _1451_v53, (_1401_v7)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length))], globalState)).Elements) {
    let _1453_v54 = _compr_51;
    if (_dafny.Seq.contains(_module.__default.fm43((_1450_v52).dtor_cf71, _1451_v53, (_1401_v7)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1401_v7).length))], globalState), _1453_v54)) {
      _coll51.add((_1453_v54).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("dp")).length),false)).length)));
    }
  }
  return _coll51;
}(), _1406_v12, globalState), (_dafny.ZERO).minus((_this).f27))), (_this).f27);
        let _1454_v56;
        let _init27 = ((_1455_p0) => function (_1456_i2) {
          return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(false, _1455_p0)), _dafny.Seq.of(_dafny.Seq.of(_1455_p0)));
        })(p0);
        let _nw241 = Array((new BigNumber(23)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw241.length); _i0_27++) {
          _nw241[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1454_v56 = _nw241;
        let _index186 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_1454_v56).length));
        (_1454_v56)[_index186] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(18)), ((_1457_p0) => function (_1458_i3) {
          return _dafny.Seq.of(_1457_p0);
        })(p0)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(558)), ((_1459_v7) => function (_1460_i4) {
          return _dafny.Seq.of((_1459_v7)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1459_v7).length))]);
        })(_1401_v7)));
        let _1461_v57;
        _1461_v57 = _module.D4.create_DC7(_1401_v7);
        let _pat_let_tv26 = p2;
        let _index187 = _module.__default.safeIndex(new BigNumber(748), new BigNumber(((_this).f28).length));
        let _index188 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_1454_v56).length));
        let _rhs251 = _1405_v11;
        let _rhs252 = function (_pat_let29_0) {
          return function (_1462_dt__update__tmp_h1) {
            return function (_pat_let30_0) {
              return function (_1463_dt__update_hcf74_h0) {
                return function (_pat_let31_0) {
                  return function (_1464_dt__update_hcf76_h0) {
                    return _module.D16.create_DC40(_1463_dt__update_hcf74_h0, (_1462_dt__update__tmp_h1).dtor_cf75, _1464_dt__update_hcf76_h0);
                  }(_pat_let31_0);
                }(_module.__default.safeDivisionInt((_this).f27, new BigNumber(527)));
              }(_pat_let30_0);
            }(_pat_let_tv26);
          }(_pat_let29_0);
        }(_1452_v55);
        let _rhs253 = _module.__default.fm44(globalState);
        let _rhs254 = (_module.__default.safeDivisionInt((_this).f27, (_this).f27)).plus((new BigNumber(155)).plus((_this).f27));
        let _rhs255 = (_1461_v57).dtor_cf12;
        let _lhs197 = (_this).f28;
        let _lhs198 = _module.__default.safeIndex(new BigNumber(748), new BigNumber(((_this).f28).length));
        let _lhs199 = _1454_v56;
        let _lhs200 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_1454_v56).length));
        let _lhs201 = globalState;
        _lhs197[_lhs198] = _rhs251;
        _1452_v55 = _rhs252;
        _lhs199[_lhs200] = _rhs253;
        _lhs201.f15 = _rhs254;
        _1401_v7 = _rhs255;
      }
      let _1465_v58;
      _1465_v58 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _1466_v59;
      let _nw242 = Array((new BigNumber(25)).toNumber());
      _nw242[(_dafny.ZERO).toNumber()] = _module.__default.fm38(_1465_v58, globalState);
      _nw242[(_dafny.ONE).toNumber()] = p3;
      _nw242[(new BigNumber(2)).toNumber()] = _module.__default.fm38(_dafny.Map.Empty.slice().updateUnsafe(p0,p0), globalState);
      _nw242[(new BigNumber(3)).toNumber()] = p3;
      _nw242[(new BigNumber(4)).toNumber()] = p3;
      _nw242[(new BigNumber(5)).toNumber()] = p3;
      _nw242[(new BigNumber(6)).toNumber()] = p3;
      _nw242[(new BigNumber(7)).toNumber()] = p3;
      _nw242[(new BigNumber(8)).toNumber()] = p3;
      _nw242[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
      _nw242[(new BigNumber(10)).toNumber()] = p3;
      _nw242[(new BigNumber(11)).toNumber()] = p3;
      _nw242[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
      _nw242[(new BigNumber(13)).toNumber()] = p3;
      _nw242[(new BigNumber(14)).toNumber()] = p3;
      _nw242[(new BigNumber(15)).toNumber()] = p3;
      _nw242[(new BigNumber(16)).toNumber()] = p3;
      _nw242[(new BigNumber(17)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
      _nw242[(new BigNumber(18)).toNumber()] = p3;
      _nw242[(new BigNumber(19)).toNumber()] = p3;
      _nw242[(new BigNumber(20)).toNumber()] = p3;
      _nw242[(new BigNumber(21)).toNumber()] = p3;
      _nw242[(new BigNumber(22)).toNumber()] = p3;
      _nw242[(new BigNumber(23)).toNumber()] = _module.__default.fm38(_1465_v58, globalState);
      _nw242[(new BigNumber(24)).toNumber()] = p3;
      _1466_v59 = _nw242;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1466_v59).length))) {
        let _1467_i5 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1467_i5)) && ((_1467_i5).isLessThan(new BigNumber((_1466_v59).length))))) {
          (_1466_v59)[(_1467_i5)] = new _dafny.CodePoint('y'.codePointAt(0));
        }
      }
      let _1468_v60;
      _1468_v60 = _dafny.Seq.of(p0, p0);
      let _1469_v61;
      let _nw243 = new _module.C3();
      _nw243.__ctor((new BigNumber((_1468_v60).length)).plus((_this).f27), (_this).f28);
      _1469_v61 = _nw243;
      _1469_v61 = _1469_v61;
      let _1470_v62;
      _1470_v62 = _module.D12.create_DC28((_this).f27, (_dafny.ZERO).minus((_1469_v61).f27), p0);
      let _1471_i6;
      _1471_i6 = _dafny.ZERO;
      L6: {
        while (_module.__default.fm0(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_1468_v60, _module.__default.fm41((_1469_v61).f27, (_1469_v61).f27, (_1470_v62).dtor_cf53, globalState)), _module.__default.safeIndex((_1469_v61).f27, new BigNumber((_dafny.Seq.Concat(_1468_v60, _module.__default.fm41((_1469_v61).f27, (_1469_v61).f27, (_1470_v62).dtor_cf53, globalState))).length)), p0)).length), globalState)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1471_i6)) {
              break L6;
            }
            _1471_i6 = (_1471_i6).plus(_dafny.ONE);
            if (p2) {
              let _1472_v63;
              let _nw244 = new _module.C3();
              _nw244.__ctor(new BigNumber(342), (_1469_v61).f28);
              _1472_v63 = _nw244;
              let _1473_v64;
              _1473_v64 = _dafny.Seq.of(_dafny.Seq.of((_dafny.ZERO).minus((_1469_v61).f27)));
              _1473_v64 = _1473_v64;
              let _1474_v65;
              _1474_v65 = _dafny.Seq.UnicodeFromString("r");
              let _1475_v66;
              _1475_v66 = _dafny.Set.fromElements((_1469_v61).f27, _module.__default.fm1(globalState), (_this).f27, (_1469_v61).f27);
              let _1476_v67;
              _1476_v67 = _dafny.Set.fromElements(_dafny.Seq.update(_1468_v60, _module.__default.safeIndex(new BigNumber((_1475_v66).length), new BigNumber((_1468_v60).length)), p2));
              let _1477_v68;
              let _nw245 = Array((new BigNumber(10)).toNumber());
              _nw245[(_dafny.ZERO).toNumber()] = (_1469_v61).f27;
              _nw245[(_dafny.ONE).toNumber()] = (_1469_v61).f27;
              _nw245[(new BigNumber(2)).toNumber()] = (_this).f27;
              _nw245[(new BigNumber(3)).toNumber()] = (_1469_v61).f27;
              _nw245[(new BigNumber(4)).toNumber()] = (_1469_v61).f27;
              _nw245[(new BigNumber(5)).toNumber()] = (_1469_v61).f27;
              _nw245[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_1474_v65).length), (_this).f27);
              _nw245[(new BigNumber(7)).toNumber()] = ((_1469_v61).f27).plus(new BigNumber((_1476_v67).length));
              _nw245[(new BigNumber(8)).toNumber()] = new BigNumber((_module.__default.fm47(globalState)).length);
              _nw245[(new BigNumber(9)).toNumber()] = (_1469_v61).f27;
              _1477_v68 = _nw245;
              let _index189 = _module.__default.safeIndex(new BigNumber(786), new BigNumber((_1477_v68).length));
              (_1477_v68)[_index189] = _module.__default.safeDivisionInt((_1469_v61).f27, new BigNumber(595));
              let _1478_v69;
              let _nw246 = Array((new BigNumber(4)).toNumber()).fill(false);
              _1478_v69 = _nw246;
              let _index190 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_1478_v69).length));
              (_1478_v69)[_index190] = !(p0) || (p0);
              let _1479_v70;
              _1479_v70 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1469_v61).f27);
              let _1480_v71;
              _1480_v71 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1479_v70).length));
              let _1481_v72;
              _1481_v72 = _dafny.Map.Empty.slice().updateUnsafe((_1469_v61).f27,p2);
              let _index191 = _module.__default.safeIndex(new BigNumber(786), new BigNumber((_1477_v68).length));
              let _index192 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_1478_v69).length));
              let _rhs256 = (_1469_v61).f27;
              let _rhs257 = (new BigNumber((((_1479_v70).update(false, (_1469_v61).f27)).update(p0, new BigNumber((_1481_v72).length))).length)).isLessThan((_dafny.ZERO).minus((_this).f27));
              let _rhs258 = p0;
              let _rhs259 = _module.__default.fm0(((p0) ? (new BigNumber((_dafny.Seq.UnicodeFromString("nngcbxfxb")).length)) : ((_1469_v61).f27)), globalState);
              let _rhs260 = ((_1469_v61).f27).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus((_1469_v61).f27))).length));
              let _lhs202 = _1477_v68;
              let _lhs203 = _module.__default.safeIndex(new BigNumber(786), new BigNumber((_1477_v68).length));
              let _lhs204 = globalState;
              let _lhs205 = globalState;
              let _lhs206 = _1478_v69;
              let _lhs207 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_1478_v69).length));
              _lhs202[_lhs203] = _rhs256;
              _lhs204.f26 = _rhs257;
              r1 = _rhs258;
              _lhs205.f26 = _rhs259;
              _lhs206[_lhs207] = _rhs260;
              _1477_v68 = _1477_v68;
              (globalState).f0 = _module.__default.safeModuloInt((_1477_v68)[_module.__default.safeIndex(new BigNumber(786), new BigNumber((_1477_v68).length))], new BigNumber(477));
            } else {
              let _1482_v73;
              _1482_v73 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.fm1(globalState), (_this).f27)),p0);
              _1482_v73 = (_module.__default.fm48(p2, p2, new BigNumber(((_dafny.MultiSet.fromElements((_1469_v61).f27, (_dafny.ZERO).minus((_1469_v61).f27))).update((_1469_v61).f27, _module.__default.abs(new BigNumber(-216)))).cardinality()), globalState)).Merge(_1482_v73);
              let _1483_v74;
              _1483_v74 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(33)), ((_1484_v60) => function (_1485_i7) {
                return new BigNumber((_1484_v60).length);
              })(_1468_v60)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(942)), ((_1486_v61) => function (_1487_i8) {
                return (_1486_v61).f27;
              })(_1469_v61)));
              let _rhs261 = ((p0) ? (_module.__default.fm1(globalState)) : (new BigNumber((_dafny.MultiSet.FromArray((_1483_v74)[_module.__default.safeIndex((_this).f27, new BigNumber((_1483_v74).length))])).cardinality())));
              let _rhs262 = p2;
              let _lhs208 = globalState;
              let _lhs209 = globalState;
              _lhs208.f15 = _rhs261;
              _lhs209.f26 = _rhs262;
              let _1488_v75;
              _1488_v75 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("cscw"));
              let _1489_v76;
              _1489_v76 = _dafny.Seq.of((_1469_v61).f27, new BigNumber(892), (_1469_v61).f27);
              let _1490_v77;
              _1490_v77 = _dafny.MultiSet.fromElements(p2);
              let _1491_v78;
              _1491_v78 = _dafny.Set.fromElements((_this).f27);
              let _1492_v79;
              _1492_v79 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_1469_v61).f27);
              let _rhs263 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(84)), function (_1493_i9) {
                return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qwi"), _dafny.Seq.UnicodeFromString("aqpplff"));
              });
              let _rhs264 = ((_1490_v77).Union(_1490_v77)).IsDisjointFrom((_1490_v77).update(p0, _module.__default.abs((_1469_v61).f27)));
              let _rhs265 = _1489_v76;
              let _rhs266 = !((_1491_v78).IsDisjointFrom(_1491_v78));
              let _rhs267 = (new BigNumber((_1492_v79).length)).minus((_this).f27);
              let _lhs210 = globalState;
              _1488_v75 = _rhs263;
              r1 = _rhs264;
              _1489_v76 = _rhs265;
              r1 = _rhs266;
              _lhs210.f0 = _rhs267;
              (globalState).f18 = ((_dafny.ZERO).minus((new BigNumber(866)).minus((_dafny.ZERO).minus(_module.__default.fm1(globalState))))).multipliedBy((_1469_v61).f27);
              let _1494_v80;
              let _1495_v81;
              let _out42;
              let _out43;
              let _outcollector18 = _module.__default.m0(p0, (_1469_v61).f27, (_1469_v61).f27, (_1469_v61).f27, globalState);
              _out42 = _outcollector18[0];
              _out43 = _outcollector18[1];
              _1494_v80 = _out42;
              _1495_v81 = _out43;
            }
            let _1496_v82;
            _1496_v82 = _dafny.Seq.of((_1469_v61).f27, (_1469_v61).f27);
            let _1497_v83;
            _1497_v83 = _dafny.Seq.of(_1496_v82);
            (globalState).f3 = (new BigNumber(125)).multipliedBy(((_1469_v61).f27).multipliedBy(new BigNumber((_1497_v83).length)));
            let _1498_v84;
            _1498_v84 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,true);
            let _1499_v85;
            _1499_v85 = _dafny.Seq.of(_1498_v84, _1498_v84);
            let _1500_v86;
            _1500_v86 = _dafny.Seq.UnicodeFromString("cyisvag");
            let _1501_v87;
            _1501_v87 = _dafny.MultiSet.fromElements(p2);
            let _1502_v88;
            _1502_v88 = _dafny.MultiSet.fromElements((_this).f27, new BigNumber((_1468_v60).length));
            let _1503_v89;
            let _nw247 = Array((new BigNumber(27)).toNumber());
            _nw247[(_dafny.ZERO).toNumber()] = (_1469_v61).f27;
            _nw247[(_dafny.ONE).toNumber()] = (new BigNumber(((_1499_v85)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_1499_v85).length))]).length)).multipliedBy((_1469_v61).f27);
            _nw247[(new BigNumber(2)).toNumber()] = (_1469_v61).f27;
            _nw247[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_module.__default.fm41(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p2)).length), (_this).f27, p0, globalState)).length), (_dafny.ZERO).minus((_this).f27));
            _nw247[(new BigNumber(4)).toNumber()] = (_1469_v61).f27;
            _nw247[(new BigNumber(5)).toNumber()] = (_1469_v61).f27;
            _nw247[(new BigNumber(6)).toNumber()] = new BigNumber(-541);
            _nw247[(new BigNumber(7)).toNumber()] = (_1469_v61).f27;
            _nw247[(new BigNumber(8)).toNumber()] = (_1469_v61).f27;
            _nw247[(new BigNumber(9)).toNumber()] = (_this).f27;
            _nw247[(new BigNumber(10)).toNumber()] = (_1469_v61).f27;
            _nw247[(new BigNumber(11)).toNumber()] = (_this).f27;
            _nw247[(new BigNumber(12)).toNumber()] = (_this).f27;
            _nw247[(new BigNumber(13)).toNumber()] = (_1469_v61).f27;
            _nw247[(new BigNumber(14)).toNumber()] = (_1469_v61).f27;
            _nw247[(new BigNumber(15)).toNumber()] = (_this).f27;
            _nw247[(new BigNumber(16)).toNumber()] = (_this).f27;
            _nw247[(new BigNumber(17)).toNumber()] = (_module.__default.fm1(globalState)).multipliedBy(new BigNumber((_1500_v86).length));
            _nw247[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Set.fromElements((_1468_v60)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements((_1468_v60)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_1468_v60).length))], p0)).length), new BigNumber((_1468_v60).length))], p2)).length);
            _nw247[(new BigNumber(19)).toNumber()] = ((_1469_v61).f27).plus(new BigNumber((_1496_v82).length));
            _nw247[(new BigNumber(20)).toNumber()] = (_this).f27;
            _nw247[(new BigNumber(21)).toNumber()] = (_1469_v61).f27;
            _nw247[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus((_1469_v61).f27);
            _nw247[(new BigNumber(23)).toNumber()] = ((_this).f27).multipliedBy((_1469_v61).f27);
            _nw247[(new BigNumber(24)).toNumber()] = new BigNumber(-743);
            _nw247[(new BigNumber(25)).toNumber()] = (((_1501_v87).contains(p2)) ? ((_1501_v87).get(p2)) : ((((_1502_v88).contains((_1469_v61).f27)) ? ((_1502_v88).get((_1469_v61).f27)) : (new BigNumber((_1501_v87).cardinality())))));
            _nw247[(new BigNumber(26)).toNumber()] = (_1469_v61).f27;
            _1503_v89 = _nw247;
            let _index193 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_1503_v89).length));
            (_1503_v89)[_index193] = (_this).f27;
            let _index194 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_1503_v89).length));
            (_1503_v89)[_index194] = new BigNumber((_1496_v82).length);
            let _1504_v90;
            _1504_v90 = _dafny.Set.fromElements(p2, p0, p0);
            (globalState).f0 = _module.__default.safeDivisionInt((_1469_v61).f27, (((_1502_v88).contains((_1503_v89)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_1503_v89).length))])) ? ((_1502_v88).get((_1503_v89)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_1503_v89).length))])) : (new BigNumber((_1504_v90).length))));
          }
        }
      }
      let _1505_v91;
      _1505_v91 = _dafny.Seq.UnicodeFromString("cd");
      let _1506_v92;
      _1506_v92 = _module.D4.create_DC8(p3, new BigNumber((_1505_v91).length), p0);
      (globalState).f24 = (_1506_v92).dtor_cf14;
      let _1507_v93;
      let _nw248 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
      _1507_v93 = _nw248;
      let _1508_v94;
      _1508_v94 = _module.D16.create_DC39(_1507_v93);
      let _1509_v95;
      let _nw249 = Array((new BigNumber(3)).toNumber()).fill(false);
      _1509_v95 = _nw249;
      let _1510_v96;
      _1510_v96 = _dafny.Map.Empty.slice().updateUnsafe(_1508_v94,_1509_v95);
      _1510_v96 = _1510_v96;
      r0 = _1468_v60;
      r1 = (_module.__default.safeDivisionInt((_this).f27, (_dafny.ZERO).minus((_1469_v61).f27))).isLessThan(new BigNumber(987));
      let _1511_v97;
      let _nw250 = Array((new BigNumber(23)).toNumber()).fill([]);
      _1511_v97 = _nw250;
      let _1512_v98;
      _1512_v98 = _dafny.Seq.of(_1511_v97, _1511_v97, _1511_v97, _1511_v97);
      let _1513_v99;
      _1513_v99 = _dafny.Seq.of(new BigNumber((_1505_v91).length));
      r2 = (_1512_v98)[_module.__default.safeIndex(new BigNumber((_1513_v99).length), new BigNumber((_1512_v98).length))];
      let _1514_v100;
      _1514_v100 = _dafny.MultiSet.fromElements(new BigNumber((_1505_v91).length));
      let _1515_v101;
      _1515_v101 = _1514_v100;
      let _pat_let_tv27 = p2;
      r3 = function (_source23) {
        let _1516___mcc_h0 = _source23;
        let _1517_cf18 = _1516___mcc_h0;
        return _pat_let_tv27;
      }(_1515_v101);
      return [r0, r1, r2, r3];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f29 = _dafny.MultiSet.Empty;
      this._f30 = _dafny.Seq.of();
      this._f27 = _dafny.ZERO;
      this._f28 = [];
      this._f38 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f38, f29, f30, f27, f28) {
      let _this = this;
      (_this)._f38 = f38;
      (_this)._f29 = f29;
      (_this)._f30 = f30;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    fm3(p0, globalState) {
      let _this = this;
      return !(false) || (true);
    };
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      if (((_this).f29).contains(true)) {
        return ((_this).f29).get(true);
      } else {
        return (_dafny.ZERO).minus(((_this).f27).multipliedBy((_this).f27));
      }
    };
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      let _source24 = _module.D8.create_DC16((_this).f27, (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(205))).cardinality())));
      if (_source24.is_DC16) {
        let _1518___mcc_h0 = (_source24).cf29;
        let _1519___mcc_h1 = (_source24).cf30;
        let _1520_cf30 = _1519___mcc_h1;
        let _1521_cf29 = _1518___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(true,_1521_cf29)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),(_this).f27));
      } else {
        let _1522___mcc_h2 = (_source24).cf28;
        let _1523_cf28 = _1522___mcc_h2;
        return _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f27);
      }
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _1524_v0;
      _1524_v0 = true;
      let _1525_v1;
      _1525_v1 = _dafny.Seq.of(!((((_this).f29).update(!(_1524_v0), _module.__default.abs((_this).f27))).IsDisjointFrom((_this).f29)));
      if ((_1525_v1)[_module.__default.safeIndex(_module.__default.fm1(globalState), new BigNumber((_1525_v1).length))]) {
        let _1526_v2;
        _1526_v2 = _dafny.MultiSet.fromElements((_this).f27);
        let _1527_v3;
        _1527_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1526_v2).cardinality()),(_this).f27);
        let _1528_v4;
        _1528_v4 = _module.D9.create_DC20((_this).f27, _1524_v0, _1524_v0, _1524_v0, _1527_v3);
        let _pat_let_tv28 = _1524_v0;
        let _pat_let_tv29 = _1524_v0;
        let _pat_let_tv30 = _1524_v0;
        let _pat_let_tv31 = _1524_v0;
        let _source25 = function (_pat_let32_0) {
          return function (_1531_dt__update__tmp_h1) {
            return function (_pat_let35_0) {
              return function (_1532_dt__update_hcf37_h0) {
                return function (_pat_let36_0) {
                  return function (_1533_dt__update_hcf34_h0) {
                    return function (_pat_let37_0) {
                      return function (_1534_dt__update_hcf35_h0) {
                        return _module.D9.create_DC20(_1533_dt__update_hcf34_h0, _1534_dt__update_hcf35_h0, (_1531_dt__update__tmp_h1).dtor_cf36, _1532_dt__update_hcf37_h0, (_1531_dt__update__tmp_h1).dtor_cf38);
                      }(_pat_let37_0);
                    }(((_pat_let_tv31) ? (_pat_let_tv29) : (_pat_let_tv30)));
                  }(_pat_let36_0);
                }((_this).f27);
              }(_pat_let35_0);
            }(!(false));
          }(_pat_let32_0);
        }(function (_pat_let33_0) {
          return function (_1529_dt__update__tmp_h0) {
            return function (_pat_let34_0) {
              return function (_1530_dt__update_hcf36_h0) {
                return _module.D9.create_DC20((_1529_dt__update__tmp_h0).dtor_cf34, (_1529_dt__update__tmp_h0).dtor_cf35, _1530_dt__update_hcf36_h0, (_1529_dt__update__tmp_h0).dtor_cf37, (_1529_dt__update__tmp_h0).dtor_cf38);
              }(_pat_let34_0);
            }(_pat_let_tv28);
          }(_pat_let33_0);
        }(_1528_v4));
        if (_source25.is_DC18) {
          let _1535___mcc_h0 = (_source25).cf32;
          let _1536_cf32 = _1535___mcc_h0;
          let _1537_v5;
          _1537_v5 = _dafny.Seq.UnicodeFromString("fq");
          let _1538_v6;
          let _nw251 = new _module.C0();
          _nw251.__ctor((_this).f27, _1537_v5);
          _1538_v6 = _nw251;
          let _1539_v7;
          _1539_v7 = _module.D6.create_DC11(_1538_v6);
          let _1540_v8;
          _1540_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1536_cf32,_1539_v7);
          let _1541_v9;
          _1541_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_1540_v8);
          let _rhs268 = (_this).f27;
          let _rhs269 = _module.__default.safeDivisionInt((_this).f27, ((_this).f27).plus((_this).f27));
          let _rhs270 = ((_this).f27).isLessThan((_this).f27);
          let _rhs271 = (_1541_v9).contains(((_this).f30)[_module.__default.safeIndex((_this).f27, new BigNumber(((_this).f30).length))]);
          let _rhs272 = (_1538_v6).f36;
          let _lhs211 = globalState;
          let _lhs212 = globalState;
          let _lhs213 = globalState;
          let _lhs214 = globalState;
          _lhs211.f18 = _rhs268;
          _lhs212.f0 = _rhs269;
          _lhs213.f26 = _rhs270;
          _1524_v0 = _rhs271;
          _lhs214.f18 = _rhs272;
          let _1542_v10;
          let _init28 = ((_1543_v0) => function (_1544_i0) {
            return _1543_v0;
          })(_1524_v0);
          let _nw252 = Array((new BigNumber(21)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw252.length); _i0_28++) {
            _nw252[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _1542_v10 = _nw252;
          let _1545_v11;
          _1545_v11 = _dafny.Set.fromElements((_1538_v6).f36);
          let _1546_v12;
          _1546_v12 = _dafny.MultiSet.fromElements(_1545_v11);
          let _1547_v13;
          _1547_v13 = _module.D11.create_DC23(_1546_v12);
          let _1548_v14;
          _1548_v14 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_1547_v13, _1547_v13, _1547_v13),_1542_v10);
          let _1549_v15;
          _1549_v15 = _dafny.Set.fromElements(_1547_v13);
          _1542_v10 = (((_1548_v14).contains((_1549_v15).Union(_1549_v15))) ? ((_1548_v14).get((_1549_v15).Union(_1549_v15))) : (_1542_v10));
          _1527_v3 = _1527_v3;
          _1542_v10 = _1542_v10;
        } else if (_source25.is_DC19) {
          let _1550___mcc_h1 = (_source25).cf33;
          let _1551_cf33 = _1550___mcc_h1;
          let _1552_v16;
          _1552_v16 = new _dafny.CodePoint('h'.codePointAt(0));
          let _1553_v17;
          _1553_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1552_v16,(_this).f27);
          let _1554_v19;
          _1554_v19 = _dafny.MultiSet.fromElements(_1553_v17, _dafny.Map.Empty.slice().updateUnsafe(_1552_v16,(_this).f27), _1553_v17, function () {
            let _coll52 = new _dafny.Map();
            for (const _compr_52 of (_dafny.MultiSet.fromElements(_1552_v16, new _dafny.CodePoint('o'.codePointAt(0)), _1552_v16)).Elements) {
              let _1555_v18 = _compr_52;
              if ((_dafny.MultiSet.fromElements(_1552_v16, new _dafny.CodePoint('o'.codePointAt(0)), _1552_v16)).contains(_1555_v18)) {
                _coll52.push([_1555_v18,(((_1527_v3).contains((_this).f27)) ? ((_1527_v3).get((_this).f27)) : ((_this).f27))]);
              }
            }
            return _coll52;
          }());
          let _1556_v20;
          _1556_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1524_v0,!(_1524_v0));
          let _1557_v21;
          _1557_v21 = _dafny.Seq.of(_1556_v20);
          let _1558_v22;
          _1558_v22 = _dafny.Set.fromElements((_this).f27, (_this).f27, (_this).f27, (_this).f27, new BigNumber((_module.__default.fm31(_1524_v0, _1551_cf33, globalState)).length));
          let _1559_v23;
          _1559_v23 = _dafny.Map.Empty.slice().updateUnsafe((_1525_v1)[_module.__default.safeIndex((_this).f27, new BigNumber((_1525_v1).length))],new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1552_v16,(_this).f27)).length));
          let _1560_v24;
          _1560_v24 = _dafny.Seq.of(_1552_v16);
          let _1561_v25;
          let _nw253 = Array((new BigNumber(19)).toNumber());
          _nw253[(_dafny.ZERO).toNumber()] = (((_1554_v19).contains(_dafny.Map.Empty.slice().updateUnsafe(_1552_v16,(_dafny.ZERO).minus((_this).f27)))) ? ((_1554_v19).get(_dafny.Map.Empty.slice().updateUnsafe(_1552_v16,(_dafny.ZERO).minus((_this).f27)))) : ((_this).f27));
          _nw253[(_dafny.ONE).toNumber()] = (_this).f27;
          _nw253[(new BigNumber(2)).toNumber()] = (_this).f27;
          _nw253[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.of(_1551_cf33, _1551_cf33)).length);
          _nw253[(new BigNumber(4)).toNumber()] = ((_1524_v0) ? ((_this).f27) : ((_dafny.ZERO).minus((_this).f27)));
          _nw253[(new BigNumber(5)).toNumber()] = new BigNumber(19);
          _nw253[(new BigNumber(6)).toNumber()] = new BigNumber(((_1557_v21)[_module.__default.safeIndex((_this).f27, new BigNumber((_1557_v21).length))]).length);
          _nw253[(new BigNumber(7)).toNumber()] = (_module.D10.create_DC22(new BigNumber((_1556_v20).length), (_this).f29)).dtor_cf40;
          _nw253[(new BigNumber(8)).toNumber()] = (_this).f27;
          _nw253[(new BigNumber(9)).toNumber()] = (_this).f27;
          _nw253[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(967), (_this).f27);
          _nw253[(new BigNumber(11)).toNumber()] = (_this).f27;
          _nw253[(new BigNumber(12)).toNumber()] = (_this).f27;
          _nw253[(new BigNumber(13)).toNumber()] = (((_this).f30)[_module.__default.safeIndex((_this).f27, new BigNumber(((_this).f30).length))]).plus(new BigNumber(-925));
          _nw253[(new BigNumber(14)).toNumber()] = (_this).f27;
          _nw253[(new BigNumber(15)).toNumber()] = ((_1551_cf33) ? ((_this).f27) : (new BigNumber((_1558_v22).length)));
          _nw253[(new BigNumber(16)).toNumber()] = new BigNumber((_1559_v23).length);
          _nw253[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(891)), ((_1562_v16) => function (_1563_i1) {
            return _1562_v16;
          })(_1552_v16)), _1560_v24)).length));
          _nw253[(new BigNumber(18)).toNumber()] = (_this).f27;
          _1561_v25 = _nw253;
          let _index195 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1561_v25).length));
          (_1561_v25)[_index195] = ((_this).f27).minus((_this).f27);
          let _1564_v27;
          _1564_v27 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f29).IsProperSubsetOf((_this).f29),function () {
            let _coll53 = new _dafny.Map();
            for (const _compr_53 of _dafny.IntegerRange(new BigNumber(593), new BigNumber(-287))) {
              let _1565_v26 = _compr_53;
              if (((new BigNumber(593)).isLessThanOrEqualTo(_1565_v26)) && ((_1565_v26).isLessThan(new BigNumber(-287)))) {
                _coll53.push([_module.__default.safeDivisionInt(_1565_v26, (_this).f27),_1524_v0]);
              }
            }
            return _coll53;
          }());
          let _index196 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1561_v25).length));
          let _rhs273 = (_this).f27;
          let _rhs274 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1560_v24, _module.__default.safeIndex((_this).f27, new BigNumber((_1560_v24).length)), _1552_v16), _1560_v24)).length);
          let _rhs275 = new BigNumber((_1564_v27).length);
          let _rhs276 = (((_1556_v20).contains(_1524_v0)) ? ((_1556_v20).get(_1524_v0)) : (true));
          let _lhs215 = globalState;
          let _lhs216 = _1561_v25;
          let _lhs217 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1561_v25).length));
          let _lhs218 = globalState;
          r0 = _rhs273;
          _lhs215.f24 = _rhs274;
          _lhs216[_lhs217] = _rhs275;
          _lhs218.f26 = _rhs276;
          let _1566_v28;
          let _nw254 = new _module.C1();
          _nw254.__ctor(_1551_cf33, (_this).f29, (_this).f30, (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f27)), (_this).f28);
          _1566_v28 = _nw254;
          let _1567_v29;
          _1567_v29 = _dafny.Seq.of(_1560_v24, _1560_v24);
          let _1568_v30;
          _1568_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1551_cf33,_dafny.Seq.update(_1567_v29, _module.__default.safeIndex((_1561_v25)[_module.__default.safeIndex(new BigNumber(757), new BigNumber((_1561_v25).length))], new BigNumber((_1567_v29).length)), _dafny.Seq.of(_1552_v16, _1552_v16, new _dafny.CodePoint('u'.codePointAt(0)))));
          (_1566_v28).f35 = _dafny.areEqual((((_1568_v30).contains(_1524_v0)) ? ((_1568_v30).get(_1524_v0)) : (_1567_v29)), _dafny.Seq.Concat(_1567_v29, _1567_v29));
          let _1569_v31;
          let _nw255 = Array((new BigNumber(22)).toNumber()).fill(false);
          _1569_v31 = _nw255;
          let _1570_v32;
          _1570_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1569_v31,false);
          _1524_v0 = !(_1570_v32).contains(_1569_v31);
        } else if (_source25.is_DC20) {
          let _1571___mcc_h2 = (_source25).cf34;
          let _1572___mcc_h3 = (_source25).cf35;
          let _1573___mcc_h4 = (_source25).cf36;
          let _1574___mcc_h5 = (_source25).cf37;
          let _1575___mcc_h6 = (_source25).cf38;
          let _1576_cf38 = _1575___mcc_h6;
          let _1577_cf37 = _1574___mcc_h5;
          let _1578_cf36 = _1573___mcc_h4;
          let _1579_cf35 = _1572___mcc_h3;
          let _1580_cf34 = _1571___mcc_h2;
          let _1581_v33;
          _1581_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1578_cf36,_1524_v0);
          let _1582_v34;
          _1582_v34 = _dafny.Map.Empty.slice().updateUnsafe((((_1581_v33).contains(_1524_v0)) ? ((_1581_v33).get(_1524_v0)) : (_1577_cf37)),(_this).f27);
          let _1583_v35;
          let _nw256 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _1583_v35 = _nw256;
          let _1584_v36;
          _1584_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1582_v34,_1583_v35);
          let _1585_v37;
          let _nw257 = Array((new BigNumber(17)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1585_v37 = _nw257;
          let _index197 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_1585_v37).length));
          (_1585_v37)[_index197] = (_this).f29;
          let _1586_v38;
          _1586_v38 = _dafny.Set.fromElements(_1580_cf34, _1580_cf34, (_this).f27, (_this).f27, _1580_cf34);
          let _index198 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_1585_v37).length));
          let _rhs277 = _dafny.Seq.update(_dafny.Seq.Concat(_1525_v1, _1525_v1), _module.__default.safeIndex((_this).f27, new BigNumber((_dafny.Seq.Concat(_1525_v1, _1525_v1)).length)), _1577_cf37);
          let _rhs278 = (_1580_cf34).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Concat(_1525_v1, _1525_v1)).length));
          let _rhs279 = ((_1584_v36).update(_1582_v34, _1583_v35)).Merge(_1584_v36);
          let _rhs280 = (_1586_v38).IsSubsetOf(_1586_v38);
          let _rhs281 = (((_this).f29).Difference(_module.__default.fm32(globalState))).Union(_dafny.MultiSet.fromElements(false));
          let _lhs219 = globalState;
          let _lhs220 = _1585_v37;
          let _lhs221 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_1585_v37).length));
          _1525_v1 = _rhs277;
          _lhs219.f26 = _rhs278;
          _1584_v36 = _rhs279;
          _1577_cf37 = _rhs280;
          _lhs220[_lhs221] = _rhs281;
          let _1587_v39;
          _1587_v39 = _dafny.Set.fromElements(_1578_cf36, _1579_cf35, _1579_cf35);
          _1579_cf35 = !((_1587_v39).IsProperSubsetOf(_1587_v39));
          let _1588_v40;
          _1588_v40 = new _dafny.CodePoint('y'.codePointAt(0));
          let _1589_v41;
          _1589_v41 = _module.D4.create_DC9(_1580_cf34, _1579_cf35);
          let _1590_v42;
          _1590_v42 = _dafny.Set.fromElements(_1589_v41, _1589_v41);
          let _1591_v43;
          _1591_v43 = _module.D1.create_DC2(_1588_v40, new BigNumber((_1590_v42).length));
          (globalState).f18 = ((((_1585_v37)[_module.__default.safeIndex(new BigNumber(333), new BigNumber((_1585_v37).length))]).contains(_1578_cf36)) ? (((_1585_v37)[_module.__default.safeIndex(new BigNumber(333), new BigNumber((_1585_v37).length))]).get(_1578_cf36)) : ((_1591_v43).dtor_cf3));
          (globalState).f18 = (_1580_cf34).minus((_this).f27);
        } else {
          let _1592___mcc_h7 = (_source25).cf31;
          let _1593_cf31 = _1592___mcc_h7;
          let _1594_v44;
          let _nw258 = new _module.C0();
          _nw258.__ctor((_dafny.ZERO).minus(new BigNumber(((_this).f30).length)), _dafny.Seq.UnicodeFromString("tidvowswu"));
          _1594_v44 = _nw258;
          let _1595_v45;
          _1595_v45 = _module.D6.create_DC11(_1594_v44);
          let _pat_let_tv32 = _1594_v44;
          let _1596_v46;
          let _nw259 = Array((new BigNumber(4)).toNumber());
          _nw259[(_dafny.ZERO).toNumber()] = _1595_v45;
          _nw259[(_dafny.ONE).toNumber()] = _1595_v45;
          _nw259[(new BigNumber(2)).toNumber()] = _module.D6.create_DC11(_1594_v44);
          _nw259[(new BigNumber(3)).toNumber()] = function (_pat_let38_0) {
            return function (_1597_dt__update__tmp_h2) {
              return function (_pat_let39_0) {
                return function (_1598_dt__update_hcf19_h0) {
                  return _module.D6.create_DC11(_1598_dt__update_hcf19_h0);
                }(_pat_let39_0);
              }(_pat_let_tv32);
            }(_pat_let38_0);
          }(_module.D6.create_DC11(_1594_v44));
          _1596_v46 = _nw259;
          let _pat_let_tv33 = _1594_v44;
          let _index199 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_1596_v46).length));
          (_1596_v46)[_index199] = function (_pat_let40_0) {
            return function (_1599_dt__update__tmp_h3) {
              return function (_pat_let41_0) {
                return function (_1600_dt__update_hcf19_h1) {
                  return _module.D6.create_DC11(_1600_dt__update_hcf19_h1);
                }(_pat_let41_0);
              }(_pat_let_tv33);
            }(_pat_let40_0);
          }(_1595_v45);
          let _pat_let_tv34 = _1594_v44;
          let _index200 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_1596_v46).length));
          (_1596_v46)[_index200] = function (_pat_let42_0) {
            return function (_1601_dt__update__tmp_h4) {
              return function (_pat_let43_0) {
                return function (_1602_dt__update_hcf19_h2) {
                  return _module.D6.create_DC11(_1602_dt__update_hcf19_h2);
                }(_pat_let43_0);
              }(_pat_let_tv34);
            }(_pat_let42_0);
          }(_1595_v45);
          (globalState).f26 = !(_1524_v0);
          (globalState).f0 = ((_this).f27).multipliedBy((_this).f27);
          let _1603_v47;
          let _nw260 = new _module.C0();
          _nw260.__ctor((_1594_v44).f36, ((!(_1524_v0)) ? ((_1594_v44).f37) : ((_1594_v44).f37)));
          _1603_v47 = _nw260;
        }
        if (_1524_v0) {
          let _1604_v48;
          _1604_v48 = _dafny.Seq.UnicodeFromString("aljehft");
          let _1605_v49;
          let _nw261 = new _module.C0();
          _nw261.__ctor(new BigNumber((_1604_v48).length), _1604_v48);
          _1605_v49 = _nw261;
          let _1606_v50;
          _1606_v50 = _module.D6.create_DC11(_1605_v49);
          let _pat_let_tv35 = _1605_v49;
          _1606_v50 = function (_pat_let44_0) {
            return function (_1607_dt__update__tmp_h5) {
              return function (_pat_let45_0) {
                return function (_1608_dt__update_hcf19_h3) {
                  return _module.D6.create_DC11(_1608_dt__update_hcf19_h3);
                }(_pat_let45_0);
              }(_pat_let_tv35);
            }(_pat_let44_0);
          }(_1606_v50);
          (globalState).f26 = !(_1524_v0);
          let _1609_v51;
          let _init29 = ((_1610_v49) => function (_1611_i2) {
            return _module.__default.safeDivisionInt(_1611_i2, (_1610_v49).f36);
          })(_1605_v49);
          let _nw262 = Array((new BigNumber(5)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw262.length); _i0_29++) {
            _nw262[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1609_v51 = _nw262;
          let _1612_v52;
          _1612_v52 = _dafny.Map.Empty.slice().updateUnsafe((_1528_v4).dtor_cf34,_1609_v51);
          let _1613_v53;
          _1613_v53 = _dafny.Seq.of(_1609_v51, _1609_v51, _1609_v51, _1609_v51, _1609_v51);
          _1612_v52 = (_1612_v52).update((_1605_v49).f36, (_1613_v53)[_module.__default.safeIndex((_1605_v49).f36, new BigNumber((_1613_v53).length))]);
          let _1614_v54;
          _1614_v54 = new _dafny.CodePoint('r'.codePointAt(0));
          let _1615_v55;
          _1615_v55 = _module.D4.create_DC8(_1614_v54, (_this).f27, false);
          let _1616_v56;
          _1616_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1524_v0,(_this).f27);
          let _1617_v57;
          _1617_v57 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements((_1615_v55).dtor_cf15, _1524_v0)).IsDisjointFrom((_this).f29),new BigNumber((_1616_v56).length));
          _1616_v56 = (_1616_v56).update(_1524_v0, (_dafny.ZERO).minus((_this).f27));
          (globalState).f26 = (_this).fm3((_this).f29, globalState);
        } else {
          let _1618_v58;
          let _nw263 = Array((new BigNumber(28)).toNumber()).fill(false);
          _1618_v58 = _nw263;
          let _1619_v59;
          _1619_v59 = _dafny.Seq.of(_1618_v58);
          let _1620_v60;
          let _nw264 = new _module.C0();
          _nw264.__ctor(new BigNumber((_1619_v59).length), _dafny.Seq.UnicodeFromString("mwnfm"));
          _1620_v60 = _nw264;
          let _1621_v61;
          _1621_v61 = new _dafny.CodePoint('k'.codePointAt(0));
          _1621_v61 = _1621_v61;
          _1524_v0 = _1524_v0;
          let _1622_v62;
          _1622_v62 = _dafny.Set.fromElements(_1524_v0);
          let _1623_v63;
          _1623_v63 = _dafny.Set.fromElements(_1622_v62);
          (globalState).f26 = !(!(((_1620_v60).f36).isLessThan(new BigNumber((_1623_v63).length))));
          (globalState).f21 = (_1620_v60).f36;
        }
        let _1624_v64;
        _1624_v64 = _module.D10.create_DC22(((_this).f27).minus(new BigNumber(((_this).f29).cardinality())), ((_this).f29).Union((_this).f29));
        _1624_v64 = _1624_v64;
        let _1625_v65;
        _1625_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_1524_v0);
        _1625_v65 = ((_1625_v65).Merge(_1625_v65)).Merge((_1625_v65).Merge(_1625_v65));
        let _1626_v66;
        _1626_v66 = _dafny.Seq.UnicodeFromString("lxncklgcj");
        let _1627_v67;
        _1627_v67 = _module.D6.create_DC12(true, _dafny.Seq.of(_1525_v1, _1525_v1, _1525_v1), _1626_v66);
        let _1628_v68;
        _1628_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1627_v67,true);
        let _1629_v69;
        let _init30 = function (_1630_i3) {
          return false;
        };
        let _nw265 = Array((new BigNumber(4)).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw265.length); _i0_30++) {
          _nw265[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1629_v69 = _nw265;
        let _1631_v70;
        _1631_v70 = _dafny.Map.Empty.slice().updateUnsafe(true,_1524_v0);
        let _rhs282 = (_this).f27;
        let _rhs283 = _1628_v68;
        let _rhs284 = _1629_v69;
        let _rhs285 = (!(_1631_v70).contains(_1524_v0)) || (_1524_v0);
        let _rhs286 = _1524_v0;
        let _lhs222 = globalState;
        let _lhs223 = globalState;
        r2 = _rhs282;
        _1628_v68 = _rhs283;
        _1629_v69 = _rhs284;
        _lhs222.f26 = _rhs285;
        _lhs223.f26 = _rhs286;
      } else {
        (globalState).f26 = _1524_v0;
        _1525_v1 = _1525_v1;
        let _1632_v71;
        _1632_v71 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).f27, globalState),_1524_v0);
        let _1633_v72;
        _1633_v72 = _dafny.Seq.UnicodeFromString("ukkh");
        let _1634_v73;
        let _nw266 = new _module.C0();
        _nw266.__ctor(new BigNumber((_1632_v71).length), _1633_v72);
        _1634_v73 = _nw266;
        let _1635_v74;
        _1635_v74 = _dafny.Map.Empty.slice().updateUnsafe((_1634_v73).f36,_1634_v73);
        let _1636_v75;
        let _nw267 = Array((new BigNumber(22)).toNumber());
        _nw267[(_dafny.ZERO).toNumber()] = _1634_v73;
        _nw267[(_dafny.ONE).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(2)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(3)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(4)).toNumber()] = (((_1635_v74).contains((_dafny.ZERO).minus((_this).f27))) ? ((_1635_v74).get((_dafny.ZERO).minus((_this).f27))) : (_1634_v73));
        _nw267[(new BigNumber(5)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(6)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(7)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(8)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(9)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(10)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(11)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(12)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(13)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(14)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(15)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(16)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(17)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(18)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(19)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(20)).toNumber()] = _1634_v73;
        _nw267[(new BigNumber(21)).toNumber()] = _1634_v73;
        _1636_v75 = _nw267;
        let _index201 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_1636_v75).length));
        (_1636_v75)[_index201] = _1634_v73;
        let _1637_v76;
        _1637_v76 = _module.D2.create_DC4((_this).f27, new BigNumber((_dafny.Set.fromElements(_1524_v0)).length), (_this).f27, new _dafny.CodePoint('s'.codePointAt(0)), _dafny.Seq.UnicodeFromString("iga"));
        let _index202 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_1636_v75).length));
        let _rhs287 = ((_this).f27).plus((_dafny.ZERO).minus((_1634_v73).f36));
        let _rhs288 = _1634_v73;
        let _rhs289 = new BigNumber(((_1637_v76).dtor_cf9).length);
        let _lhs224 = globalState;
        let _lhs225 = _1636_v75;
        let _lhs226 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_1636_v75).length));
        let _lhs227 = globalState;
        _lhs224.f15 = _rhs287;
        _lhs225[_lhs226] = _rhs288;
        _lhs227.f3 = _rhs289;
        let _1638_v77;
        _1638_v77 = new _dafny.CodePoint('e'.codePointAt(0));
        _1638_v77 = _1638_v77;
        (globalState).f21 = (_this).f27;
      }
      let _1639_v78;
      _1639_v78 = _dafny.Seq.UnicodeFromString("odaravso");
      let _1640_v79;
      _1640_v79 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f27),_1524_v0);
      let _1641_v80;
      let _nw268 = new _module.C2();
      _nw268.__ctor((_this).f29, (_this).f27);
      _1641_v80 = _nw268;
      let _1642_v81;
      _1642_v81 = _module.D12.create_DC26(_1641_v80);
      let _1643_v82;
      _1643_v82 = _dafny.Set.fromElements(_1641_v80, _1641_v80, _1641_v80, (_1642_v81).dtor_cf46, _1641_v80);
      let _1644_v83;
      _1644_v83 = _dafny.MultiSet.fromElements((_this).f27, _module.__default.fm1(globalState));
      let _1645_v84;
      let _nw269 = Array((new BigNumber(13)).toNumber());
      _nw269[(_dafny.ZERO).toNumber()] = (_1525_v1)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("yspuqcihq")).length), new BigNumber((_1525_v1).length))];
      _nw269[(_dafny.ONE).toNumber()] = _1524_v0;
      _nw269[(new BigNumber(2)).toNumber()] = _1524_v0;
      _nw269[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1639_v78, _1639_v78);
      _nw269[(new BigNumber(4)).toNumber()] = _1524_v0;
      _nw269[(new BigNumber(5)).toNumber()] = _1524_v0;
      _nw269[(new BigNumber(6)).toNumber()] = _1524_v0;
      _nw269[(new BigNumber(7)).toNumber()] = false;
      _nw269[(new BigNumber(8)).toNumber()] = ((_this).f27).isLessThan((_this).f27);
      _nw269[(new BigNumber(9)).toNumber()] = ((_this).f27).isLessThan((_this).f27);
      _nw269[(new BigNumber(10)).toNumber()] = _dafny.areEqual(_module.__default.fm33(_1640_v79, new BigNumber((_1643_v82).length), new BigNumber((_1525_v1).length), new BigNumber((_1644_v83).cardinality()), globalState), _1525_v1);
      _nw269[(new BigNumber(11)).toNumber()] = _1524_v0;
      _nw269[(new BigNumber(12)).toNumber()] = true;
      _1645_v84 = _nw269;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1645_v84).length))) {
        let _1646_i4 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1646_i4)) && ((_1646_i4).isLessThan(new BigNumber((_1645_v84).length))))) {
          (_1645_v84)[(_1646_i4)] = !((new BigNumber(871)).isLessThan(new BigNumber(175)));
        }
      }
      let _1647_v85;
      _1647_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_1645_v84);
      let _1648_v86;
      _1648_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1524_v0,(_this).f27);
      _1647_v85 = (_1647_v85).update((((_1648_v86).contains(_module.__default.fm0((_this).f27, globalState))) ? ((_1648_v86).get(_module.__default.fm0((_this).f27, globalState))) : (new BigNumber((_1640_v79).length))), _1645_v84);
      let _1649_v87;
      _1649_v87 = new _dafny.CodePoint('k'.codePointAt(0));
      let _1650_v88;
      _1650_v88 = _dafny.Map.Empty.slice().updateUnsafe(_1524_v0,_1649_v87);
      _1650_v88 = (_1650_v88).update(_1524_v0, _1649_v87);
      let _hi6 = new BigNumber(18);
      for (let _1651_i5 = (_this).f27; _1651_i5.isLessThan(_hi6); _1651_i5 = _1651_i5.plus(_dafny.ONE)) {
        (globalState).f18 = (_dafny.ZERO).minus(new BigNumber(-374));
        (globalState).f26 = !(new BigNumber(420)).isEqualTo(new BigNumber(871));
        if ((((_this).fm3((_this).f29, globalState)) ? (!((_dafny.MultiSet.fromElements(_1524_v0)).IsProperSubsetOf((_dafny.MultiSet.fromElements(_1524_v0, _1524_v0)).update(false, _module.__default.abs(_1651_i5))))) : (_1524_v0))) {
          _1524_v0 = !(true);
          let _rhs290 = _1524_v0;
          let _rhs291 = _1639_v78;
          _1524_v0 = _rhs290;
          _1639_v78 = _rhs291;
          let _1652_v89;
          let _nw270 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _1652_v89 = _nw270;
          let _index203 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_1652_v89).length));
          (_1652_v89)[_index203] = _1651_i5;
          let _index204 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_1645_v84).length));
          (_1645_v84)[_index204] = _1524_v0;
          let _index205 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_1652_v89).length));
          let _index206 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_1645_v84).length));
          let _rhs292 = _module.__default.safeModuloInt((_this).f27, new BigNumber(316));
          let _rhs293 = (_1651_i5).isEqualTo((_this).f27);
          let _rhs294 = _1645_v84;
          let _rhs295 = _module.__default.safeModuloInt(_1651_i5, new BigNumber((_dafny.Seq.of((_this).f27, (_this).f27)).length));
          let _rhs296 = (_dafny.MultiSet.fromElements((_this).f27, (_this).f27, _1651_i5)).IsProperSubsetOf(_dafny.MultiSet.fromElements((_this).f27, _1651_i5));
          let _lhs228 = _1652_v89;
          let _lhs229 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_1652_v89).length));
          let _lhs230 = _1645_v84;
          let _lhs231 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_1645_v84).length));
          _lhs228[_lhs229] = _rhs292;
          _1524_v0 = _rhs293;
          _1645_v84 = _rhs294;
          r2 = _rhs295;
          _lhs230[_lhs231] = _rhs296;
          let _1653_v90;
          _1653_v90 = _1644_v83;
          let _1654_v91;
          _1654_v91 = _dafny.Map.Empty.slice().updateUnsafe(_1653_v90,(_1645_v84)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_1645_v84).length))]);
          let _1655_v92;
          _1655_v92 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1653_v90,(_1645_v84)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_1645_v84).length))]), _dafny.Map.Empty.slice().updateUnsafe(_1653_v90,(_1645_v84)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_1645_v84).length))]));
          let _1656_v93;
          _1656_v93 = _dafny.Set.fromElements((_1645_v84)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_1645_v84).length))]);
          _1654_v91 = (_1655_v92)[_module.__default.safeIndex((new BigNumber((_1656_v93).length)).plus(_1651_i5), new BigNumber((_1655_v92).length))];
          (globalState).f3 = (_this).f27;
        } else {
          (globalState).f0 = _1651_i5;
          let _1657_v94;
          let _init31 = ((_1658_v1) => function (_1659_i6) {
            return (_1659_i6).minus(new BigNumber((_1658_v1).length));
          })(_1525_v1);
          let _nw271 = Array((new BigNumber(21)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw271.length); _i0_31++) {
            _nw271[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1657_v94 = _nw271;
          let _index207 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_1657_v94).length));
          (_1657_v94)[_index207] = new BigNumber(311);
          let _index208 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_1657_v94).length));
          (_1657_v94)[_index208] = _1651_i5;
          let _1660_v95;
          _1660_v95 = _dafny.Seq.UnicodeFromString("hfpe");
          let _1661_v96;
          let _nw272 = new _module.C0();
          _nw272.__ctor(new BigNumber(((_this).f29).cardinality()), (_1660_v95));
          _1661_v96 = _nw272;
          let _1662_v97;
          _1662_v97 = _dafny.Map.Empty.slice().updateUnsafe(_1524_v0,_1661_v96);
          let _1663_v98;
          let _nw273 = new _module.C2();
          _nw273.__ctor((_this).f29, new BigNumber((_1662_v97).length));
          _1663_v98 = _nw273;
          _1639_v78 = (_1661_v96).f37;
          let _1664_v99;
          _1664_v99 = _dafny.Map.Empty.slice().updateUnsafe((_1657_v94)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1657_v94).length))],(_1657_v94)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1657_v94).length))]);
          let _1665_v100;
          _1665_v100 = _dafny.Set.fromElements((_1641_v80).fm21(globalState), _1524_v0);
          let _1666_v101;
          let _nw274 = new _module.C3();
          _nw274.__ctor(new BigNumber((_1665_v100).length), (_this).f28);
          _1666_v101 = _nw274;
          let _1667_v102;
          _1667_v102 = _dafny.Set.fromElements(_1666_v101, _1666_v101);
          let _1668_v103;
          _1668_v103 = _dafny.Map.Empty.slice().updateUnsafe(_1651_i5,_1648_v86);
          _1664_v99 = (_1664_v99).update(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1651_i5,_dafny.Map.Empty.slice().updateUnsafe(_1524_v0,new BigNumber((_1667_v102).length)))).Merge(_1668_v103)).length), _1651_i5);
        }
        let _1669_v104;
        _1669_v104 = _module.D4.create_DC7(_1645_v84);
        _1669_v104 = _1669_v104;
      }
      let _hi7 = (_this).f27;
      for (let _1670_i7 = (_dafny.ZERO).minus((_this).f27); _1670_i7.isLessThan(_hi7); _1670_i7 = _1670_i7.plus(_dafny.ONE)) {
        let _1671_v105;
        let _nw275 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
        _1671_v105 = _nw275;
        let _index209 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_1671_v105).length));
        (_1671_v105)[_index209] = (_1640_v79).Merge(_1640_v79);
        let _index210 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_1671_v105).length));
        (_1671_v105)[_index210] = _1640_v79;
        let _1672_v106;
        let _init32 = ((_1673_v0) => function (_1674_i8) {
          return ((_1673_v0) ? (_dafny.Seq.of((_this).f30)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-529)), function (_1675_i9) {
            return (_this).f30;
          })));
        })(_1524_v0);
        let _nw276 = Array((new BigNumber(27)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw276.length); _i0_32++) {
          _nw276[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1672_v106 = _nw276;
        let _1676_v107;
        _1676_v107 = _dafny.Seq.of((_this).f30, (_this).f30);
        let _index211 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_1672_v106).length));
        (_1672_v106)[_index211] = _1676_v107;
        let _index212 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_1672_v106).length));
        (_1672_v106)[_index212] = _dafny.Seq.Concat(_1676_v107, _1676_v107);
        let _1677_v108;
        _1677_v108 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1670_i7,_1524_v0));
        let _index213 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_1671_v105).length));
        (_1671_v105)[_index213] = (_1677_v108)[_module.__default.safeIndex(_1670_i7, new BigNumber((_1677_v108).length))];
        r0 = (_this).f27;
      }
      let _1678_v109;
      _1678_v109 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f27).plus((_this).f27),new BigNumber(53));
      r0 = (((_1678_v109).contains((_this).f27)) ? ((_1678_v109).get((_this).f27)) : ((_dafny.ZERO).minus((_this).f27)));
      let _1679_v110;
      _1679_v110 = _module.D9.create_DC20((_this).f27, true, _1524_v0, false, _1678_v109);
      let _1680_v111;
      let _init33 = function (_1681_i10) {
        return (_1681_i10).multipliedBy((_dafny.ZERO).minus((_this).f27));
      };
      let _nw277 = Array((new BigNumber(13)).toNumber());
      for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw277.length); _i0_33++) {
        _nw277[_i0_33] = _init33(new BigNumber(_i0_33));
      }
      _1680_v111 = _nw277;
      let _1682_v112;
      _1682_v112 = _module.D12.create_DC27(_1524_v0, _1644_v83, new BigNumber((_1639_v78).length), _1680_v111);
      let _pat_let_tv36 = _1524_v0;
      let _1683_v113;
      _1683_v113 = _dafny.Seq.of(function (_pat_let46_0) {
        return function (_1684_dt__update__tmp_h6) {
          return function (_pat_let47_0) {
            return function (_1685_dt__update_hcf47_h0) {
              return _module.D12.create_DC27(_1685_dt__update_hcf47_h0, (_1684_dt__update__tmp_h6).dtor_cf48, (_1684_dt__update__tmp_h6).dtor_cf49, (_1684_dt__update__tmp_h6).dtor_cf50);
            }(_pat_let47_0);
          }(_pat_let_tv36);
        }(_pat_let46_0);
      }(_module.D12.create_DC27(_1524_v0, _1644_v83, new BigNumber(468), _1680_v111)), _1682_v112);
      r1 = _module.__default.safeDivisionInt(((_this).f30)[_module.__default.safeIndex((_1679_v110).dtor_cf34, new BigNumber(((_this).f30).length))], _module.__default.safeDivisionInt((_this).f27, new BigNumber(((_dafny.MultiSet.FromArray(_1683_v113)).update(_1682_v112, _module.__default.abs((_this).f27))).cardinality())));
      r2 = ((_1524_v0) ? ((_this).f27) : (new BigNumber(((_this).f30).length)));
      return [r0, r1, r2];
    }
    m9(globalState) {
      let _this = this;
      let r0 = [];
      let r1 = false;
      let r2 = false;
      let _1686_v0;
      _1686_v0 = false;
      if (((_this).f27).isLessThanOrEqualTo(((_this).f27).plus(new BigNumber(((_this).fm2(_1686_v0, _1686_v0, _1686_v0, globalState)).length)))) {
        let _1687_v1;
        _1687_v1 = _dafny.Seq.UnicodeFromString("a");
        let _1688_v2;
        let _1689_v3;
        let _out44;
        let _out45;
        let _outcollector19 = _module.__default.m0(true, new BigNumber((_1687_v1).length), ((_this).f27).multipliedBy((_this).f27), (_this).f27, globalState);
        _out44 = _outcollector19[0];
        _out45 = _outcollector19[1];
        _1688_v2 = _out44;
        _1689_v3 = _out45;
        let _1690_v4;
        let _nw278 = new _module.C2();
        _nw278.__ctor((_this).f29, (_this).f27);
        _1690_v4 = _nw278;
        let _1691_v5;
        _1691_v5 = _module.D12.create_DC29(_module.D12.create_DC26(_1690_v4));
        let _1692_v6;
        _1692_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1691_v5,_1688_v2);
        let _1693_v7;
        _1693_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1688_v2,_1686_v0);
        _1692_v6 = (_1692_v6).update(_1691_v5, (((_1693_v7).contains(_1688_v2)) ? ((_1693_v7).get(_1688_v2)) : (_1686_v0)));
        let _1694_v8;
        let _nw279 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _1694_v8 = _nw279;
        let _1695_v9;
        _1695_v9 = _dafny.Seq.of(_1694_v8);
        if (_dafny.Seq.IsProperPrefixOf(_1695_v9, _1695_v9)) {
          let _1696_v10;
          let _init34 = ((_1697_v0, _1698_v2) => function (_1699_i0) {
            return _dafny.Seq.of(_1697_v0, _1697_v0, _1698_v2);
          })(_1686_v0, _1688_v2);
          let _nw280 = Array((new BigNumber(26)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw280.length); _i0_34++) {
            _nw280[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1696_v10 = _nw280;
          let _1700_v11;
          _1700_v11 = _dafny.Seq.of(!((_1690_v4).fm21(globalState)));
          let _index214 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1696_v10).length));
          (_1696_v10)[_index214] = _1700_v11;
          let _index215 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1696_v10).length));
          (_1696_v10)[_index215] = _1700_v11;
          let _1701_v12;
          let _nw281 = new _module.C3();
          _nw281.__ctor((_this).f27, (_this).f28);
          _1701_v12 = _nw281;
          let _1702_v13;
          _1702_v13 = _dafny.Map.Empty.slice().updateUnsafe(((false) ? (!(_1686_v0)) : (_1686_v0)),_module.__default.safeDivisionInt((_this).f27, (_this).f27));
          _1702_v13 = _1702_v13;
          let _index216 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_1694_v8).length));
          (_1694_v8)[_index216] = new BigNumber((_1700_v11).length);
          let _1703_v14;
          let _nw282 = Array((_dafny.ONE).toNumber());
          _nw282[(_dafny.ZERO).toNumber()] = _1686_v0;
          _1703_v14 = _nw282;
          let _index217 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_1694_v8).length));
          let _rhs297 = _1686_v0;
          let _rhs298 = _1703_v14;
          let _rhs299 = _1689_v3;
          let _rhs300 = ((new BigNumber((_1687_v1).length)).isLessThanOrEqualTo(new BigNumber(520))) || (_dafny.Seq.IsPrefixOf(_1700_v11, (_1696_v10)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_1696_v10).length))]));
          let _lhs232 = globalState;
          let _lhs233 = _1694_v8;
          let _lhs234 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_1694_v8).length));
          let _lhs235 = globalState;
          _lhs232.f26 = _rhs297;
          r0 = _rhs298;
          _lhs233[_lhs234] = _rhs299;
          _lhs235.f26 = _rhs300;
          let _1704_v15;
          _1704_v15 = _dafny.Set.fromElements(_1703_v14);
          let _1705_v16;
          _1705_v16 = _module.D12.create_DC28(new BigNumber(507), new BigNumber((_1704_v15).length), _1688_v2);
          let _1706_v17;
          _1706_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1705_v16,_1687_v1);
          _1706_v17 = (_1706_v17).update(_1705_v16, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fcnwjnotv"), _1687_v1));
        } else {
          let _1707_v18;
          let _nw283 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1707_v18 = _nw283;
          _1707_v18 = _1707_v18;
          (globalState).f24 = (_this).f27;
          let _nw284 = Array((new BigNumber(25)).toNumber()).fill(false);
          r0 = _nw284;
          r2 = _1686_v0;
          r1 = _1688_v2;
        }
        _1693_v7 = (_1693_v7).update(_1686_v0, _1688_v2);
        let _1708_v19;
        _1708_v19 = _dafny.Seq.of(_1688_v2, true, false, _1688_v2);
        if ((_1708_v19)[_module.__default.safeIndex((_this).f27, new BigNumber((_1708_v19).length))]) {
          let _1709_v20;
          _1709_v20 = new _dafny.CodePoint('x'.codePointAt(0));
          let _1710_v21;
          let _nw285 = Array((new BigNumber(13)).toNumber());
          _nw285[(_dafny.ZERO).toNumber()] = _1709_v20;
          _nw285[(_dafny.ONE).toNumber()] = _1709_v20;
          _nw285[(new BigNumber(2)).toNumber()] = _1709_v20;
          _nw285[(new BigNumber(3)).toNumber()] = _1709_v20;
          _nw285[(new BigNumber(4)).toNumber()] = (_1687_v1)[_module.__default.safeIndex((_this).f27, new BigNumber((_1687_v1).length))];
          _nw285[(new BigNumber(5)).toNumber()] = _1709_v20;
          _nw285[(new BigNumber(6)).toNumber()] = _1709_v20;
          _nw285[(new BigNumber(7)).toNumber()] = _1709_v20;
          _nw285[(new BigNumber(8)).toNumber()] = _1709_v20;
          _nw285[(new BigNumber(9)).toNumber()] = _module.__default.fm38(_1693_v7, globalState);
          _nw285[(new BigNumber(10)).toNumber()] = _1709_v20;
          _nw285[(new BigNumber(11)).toNumber()] = _1709_v20;
          _nw285[(new BigNumber(12)).toNumber()] = (_module.D1.create_DC2(_1709_v20, _1689_v3)).dtor_cf2;
          _1710_v21 = _nw285;
          let _1711_v22;
          _1711_v22 = _dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC31(_1689_v3),_1710_v21);
          let _1712_v23;
          _1712_v23 = _module.D13.create_DC31(_1689_v3);
          let _1713_v24;
          let _nw286 = Array((new BigNumber(12)).toNumber());
          _nw286[(_dafny.ZERO).toNumber()] = _1710_v21;
          _nw286[(_dafny.ONE).toNumber()] = _1710_v21;
          _nw286[(new BigNumber(2)).toNumber()] = _1710_v21;
          _nw286[(new BigNumber(3)).toNumber()] = _1710_v21;
          _nw286[(new BigNumber(4)).toNumber()] = _1710_v21;
          _nw286[(new BigNumber(5)).toNumber()] = (((_1711_v22).contains(_1712_v23)) ? ((_1711_v22).get(_1712_v23)) : (_1710_v21));
          _nw286[(new BigNumber(6)).toNumber()] = _1710_v21;
          _nw286[(new BigNumber(7)).toNumber()] = _1710_v21;
          _nw286[(new BigNumber(8)).toNumber()] = _1710_v21;
          _nw286[(new BigNumber(9)).toNumber()] = _1710_v21;
          _nw286[(new BigNumber(10)).toNumber()] = _1710_v21;
          _nw286[(new BigNumber(11)).toNumber()] = _1710_v21;
          _1713_v24 = _nw286;
          let _1714_v25;
          _1714_v25 = _dafny.Seq.of(_dafny.Seq.of(false, _1688_v2, _1686_v0), _1708_v19, _1708_v19);
          let _1715_v26;
          _1715_v26 = _module.D6.create_DC12(_1688_v2, _1714_v25, _1687_v1);
          let _1716_v27;
          _1716_v27 = _dafny.Seq.of(_1715_v26);
          let _1717_v28;
          _1717_v28 = _dafny.Set.fromElements(_1687_v1, _1687_v1, _1687_v1);
          let _rhs301 = ((_this).fm4(_1717_v28, _1688_v2, _dafny.Seq.of(_1686_v0), globalState)).multipliedBy((_1689_v3).plus(_1689_v3));
          let _rhs302 = _1713_v24;
          let _rhs303 = _1688_v2;
          let _rhs304 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(906)), ((_1718_v26) => function (_1719_i1) {
            return _1718_v26;
          })(_1715_v26));
          let _rhs305 = new BigNumber((_1687_v1).length);
          let _lhs236 = globalState;
          let _lhs237 = globalState;
          _lhs236.f0 = _rhs301;
          _1713_v24 = _rhs302;
          _1686_v0 = _rhs303;
          _1716_v27 = _rhs304;
          _lhs237.f21 = _rhs305;
          let _1720_v29;
          _1720_v29 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("wtdcmrxwf"));
          (globalState).f3 = new BigNumber(((_1720_v29).update(_1687_v1, _module.__default.abs((_this).f27))).cardinality());
          let _1721_v30;
          let _nw287 = new _module.C3();
          _nw287.__ctor(new BigNumber(918), (_this).f28);
          _1721_v30 = _nw287;
          let _1722_v31;
          let _nw288 = new _module.C3();
          _nw288.__ctor(_1689_v3, (_this).f28);
          _1722_v31 = _nw288;
          let _1723_v32;
          _1723_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1722_v31,new _dafny.CodePoint('k'.codePointAt(0)));
          _1723_v32 = _1723_v32;
          let _1724_v33;
          let _nw289 = new _module.C3();
          _nw289.__ctor((_this).f27, (_1722_v31).f28);
          _1724_v33 = _nw289;
        } else {
          (globalState).f21 = _1689_v3;
          let _1725_v34;
          _1725_v34 = _dafny.Seq.of(_1708_v19);
          let _1726_v35;
          _1726_v35 = _module.D6.create_DC12(_1686_v0, _1725_v34, _1687_v1);
          (globalState).f26 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat((_1726_v35).dtor_cf21, _1725_v34), _1725_v34);
          let _1727_v36;
          let _nw290 = new _module.C3();
          _nw290.__ctor((_dafny.ZERO).minus(_1689_v3), (_this).f28);
          _1727_v36 = _nw290;
          let _1728_v37;
          _1728_v37 = new _dafny.CodePoint('k'.codePointAt(0));
          let _rhs306 = _1686_v0;
          let _rhs307 = _1728_v37;
          let _rhs308 = (_this).f27;
          let _lhs238 = globalState;
          _1688_v2 = _rhs306;
          _1728_v37 = _rhs307;
          _lhs238.f0 = _rhs308;
          let _1729_v38;
          _1729_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1728_v37,_1686_v0);
          let _1730_v39;
          _1730_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1689_v3,_1728_v37);
          let _1731_v40;
          _1731_v40 = _dafny.Seq.of(((_this).f29).update(_1688_v2, _module.__default.abs((_this).f27)));
          _1729_v38 = (_1729_v38).update((((_1730_v39).contains(_1689_v3)) ? ((_1730_v39).get(_1689_v3)) : (_1728_v37)), ((_1731_v40)[_module.__default.safeIndex((_this).f27, new BigNumber((_1731_v40).length))]).IsDisjointFrom((_this).f29));
        }
      } else {
        let _1732_v41;
        let _nw291 = Array((new BigNumber(23)).toNumber());
        _nw291[(_dafny.ZERO).toNumber()] = _1686_v0;
        _nw291[(_dafny.ONE).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(2)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(3)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(4)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(5)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(6)).toNumber()] = !(_1686_v0);
        _nw291[(new BigNumber(7)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(8)).toNumber()] = !(!(_1686_v0));
        _nw291[(new BigNumber(9)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(10)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(11)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(12)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(13)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(14)).toNumber()] = !(_1686_v0);
        _nw291[(new BigNumber(15)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(16)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(17)).toNumber()] = true;
        _nw291[(new BigNumber(18)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(19)).toNumber()] = true;
        _nw291[(new BigNumber(20)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(21)).toNumber()] = _1686_v0;
        _nw291[(new BigNumber(22)).toNumber()] = _1686_v0;
        _1732_v41 = _nw291;
        let _1733_v42;
        _1733_v42 = _dafny.MultiSet.fromElements(_1732_v41, _1732_v41, _1732_v41);
        let _1734_v43;
        let _nw292 = new _module.C4();
        _nw292.__ctor(new BigNumber(((_1733_v42).Union(_1733_v42)).cardinality()), (_this).f28);
        _1734_v43 = _nw292;
        (globalState).f0 = _module.__default.safeModuloInt((new BigNumber((_dafny.Seq.UnicodeFromString("wibmqbfun")).length)).multipliedBy((_this).f27), (_this).f27);
        (globalState).f21 = (_this).f27;
        if (_1686_v0) {
          let _1735_v44;
          let _nw293 = new _module.C3();
          _nw293.__ctor((_this).f27, (function (_pat_let48_0) {
            return function (_1736_dt__update__tmp_h0) {
              return function (_pat_let49_0) {
                return function (_1737_dt__update_hcf82_h0) {
                  return _module.D17.create_DC43(_1737_dt__update_hcf82_h0);
                }(_pat_let49_0);
              }((_this).f28);
            }(_pat_let48_0);
          }(_module.D17.create_DC43((_this).f28))).dtor_cf82);
          _1735_v44 = _nw293;
          (globalState).f26 = _1686_v0;
          _1686_v0 = _1686_v0;
          let _1738_v45;
          _1738_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1686_v0,_1686_v0);
          let _1739_v46;
          _1739_v46 = _dafny.Seq.UnicodeFromString("ikcash");
          _1738_v45 = (_1738_v45).update(_1686_v0, ((_this).f27).isLessThan(new BigNumber((_1739_v46).length)));
          let _1740_v47;
          let _init35 = ((_1741_v0) => function (_1742_i2) {
            return _1741_v0;
          })(_1686_v0);
          let _nw294 = Array((new BigNumber(18)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw294.length); _i0_35++) {
            _nw294[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1740_v47 = _nw294;
          let _1743_v48;
          _1743_v48 = _dafny.Seq.of(_1740_v47);
          r2 = !_dafny.areEqual(_1743_v48, _1743_v48);
        } else {
          let _index218 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_1732_v41).length));
          (_1732_v41)[_index218] = _1686_v0;
          let _1744_v49;
          let _init36 = function (_1745_i3) {
            return _module.__default.safeModuloInt(_1745_i3, (_this).f27);
          };
          let _nw295 = Array((new BigNumber(3)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw295.length); _i0_36++) {
            _nw295[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1744_v49 = _nw295;
          let _index219 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_1744_v49).length));
          (_1744_v49)[_index219] = ((_module.__default.fm0((_dafny.ZERO).minus((_this).f27), globalState)) ? ((_this).f27) : (new BigNumber(-75)));
          let _1746_v50;
          _1746_v50 = _module.D17.create_DC44((_this).f27);
          let _1747_v51;
          _1747_v51 = _module.D17.create_DC45(_1746_v50);
          let _1748_v52;
          _1748_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1686_v0,_1686_v0);
          let _1749_v53;
          _1749_v53 = _dafny.Set.fromElements(_1748_v52, _1748_v52, _1748_v52);
          let _1750_v54;
          _1750_v54 = _dafny.Seq.of(_1749_v53);
          let _1751_v55;
          _1751_v55 = _module.D18.create_DC46((_1750_v54)[_module.__default.safeIndex((_this).f27, new BigNumber((_1750_v54).length))]);
          let _index220 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_1732_v41).length));
          let _index221 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_1744_v49).length));
          let _rhs309 = !_dafny.Seq.contains((_this).f30, (_dafny.ZERO).minus((_this).f27));
          let _rhs310 = (_this).f27;
          let _rhs311 = _1747_v51;
          let _rhs312 = _module.__default.safeDivisionInt((_this).f27, new BigNumber((((_1751_v55).dtor_cf85).Union(_dafny.Set.fromElements(_1748_v52, _dafny.Map.Empty.slice().updateUnsafe(_1686_v0,false), _1748_v52))).length));
          let _rhs313 = (_this).f27;
          let _lhs239 = _1732_v41;
          let _lhs240 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_1732_v41).length));
          let _lhs241 = _1744_v49;
          let _lhs242 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_1744_v49).length));
          let _lhs243 = globalState;
          let _lhs244 = globalState;
          _lhs239[_lhs240] = _rhs309;
          _lhs241[_lhs242] = _rhs310;
          _1747_v51 = _rhs311;
          _lhs243.f3 = _rhs312;
          _lhs244.f15 = _rhs313;
          let _1752_v56;
          let _nw296 = new _module.C2();
          _nw296.__ctor(_dafny.MultiSet.fromElements((_1732_v41)[_module.__default.safeIndex(new BigNumber(603), new BigNumber((_1732_v41).length))], _1686_v0, !((_1732_v41)[_module.__default.safeIndex(new BigNumber(603), new BigNumber((_1732_v41).length))])), new BigNumber(268));
          _1752_v56 = _nw296;
          let _1753_v57;
          let _nw297 = Array((new BigNumber(9)).toNumber()).fill(false);
          _1753_v57 = _nw297;
          let _1754_v58;
          _1754_v58 = _dafny.Map.Empty.slice().updateUnsafe((_1744_v49)[_module.__default.safeIndex(new BigNumber(629), new BigNumber((_1744_v49).length))],_1753_v57);
          _1754_v58 = _1754_v58;
          let _1755_v59;
          _1755_v59 = _dafny.MultiSet.fromElements((_this).f27);
          let _1756_v60;
          _1756_v60 = _dafny.Set.fromElements(_1755_v59, _1755_v59);
          let _1757_v61;
          _1757_v61 = _dafny.Map.Empty.slice().updateUnsafe(_module.D17.create_DC44((_this).f27),_1756_v60);
          let _1758_v62;
          _1758_v62 = _dafny.Seq.of(_1686_v0);
          let _1759_v63;
          _1759_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1758_v62,new BigNumber((_1754_v58).length));
          let _1760_v64;
          _1760_v64 = _module.D17.create_DC44(new BigNumber((_1759_v63).length));
          _1757_v61 = (_1757_v61).update(((_1686_v0) ? (function (_pat_let50_0) {
            return function (_1761_dt__update__tmp_h1) {
              return function (_pat_let51_0) {
                return function (_1762_dt__update_hcf83_h0) {
                  return _module.D17.create_DC44(_1762_dt__update_hcf83_h0);
                }(_pat_let51_0);
              }(new BigNumber(((_this).f30).length));
            }(_pat_let50_0);
          }(_1760_v64)) : (_1760_v64)), _1756_v60);
          let _1763_v65;
          _1763_v65 = _dafny.Set.fromElements(_1744_v49);
          let _index222 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_1732_v41).length));
          let _rhs314 = ((_1763_v65).Difference(_dafny.Set.fromElements(_1744_v49))).IsSubsetOf(_1763_v65);
          let _rhs315 = !(_1686_v0);
          let _rhs316 = _1744_v49;
          let _rhs317 = (_1732_v41)[_module.__default.safeIndex(new BigNumber(603), new BigNumber((_1732_v41).length))];
          let _lhs245 = _1732_v41;
          let _lhs246 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_1732_v41).length));
          let _lhs247 = globalState;
          let _lhs248 = globalState;
          _lhs245[_lhs246] = _rhs314;
          _lhs247.f26 = _rhs315;
          _1744_v49 = _rhs316;
          _lhs248.f26 = _rhs317;
        }
        let _1764_v66;
        _1764_v66 = _dafny.Set.fromElements((_this).f27);
        let _1765_v67;
        _1765_v67 = _module.D9.create_DC17(_1764_v66);
        let _1766_v68;
        _1766_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1765_v67,(_this).f27);
        let _1767_v69;
        _1767_v69 = _dafny.Seq.UnicodeFromString("heeuww");
        let _1768_v70;
        let _nw298 = new _module.C0();
        _nw298.__ctor(new BigNumber((_1766_v68).length), _1767_v69);
        _1768_v70 = _nw298;
        let _1769_v71;
        _1769_v71 = _dafny.Seq.of(_1768_v70);
        let _1770_v72;
        let _out46;
        _out46 = (_1734_v43).m11((_1769_v71)[_module.__default.safeIndex((_1768_v70).f36, new BigNumber((_1769_v71).length))], _1686_v0, globalState);
        _1770_v72 = _out46;
      }
      let _1771_v73;
      _1771_v73 = new _dafny.CodePoint('g'.codePointAt(0));
      let _1772_v74;
      _1772_v74 = _dafny.Seq.of(_1771_v73);
      let _1773_v75;
      _1773_v75 = _dafny.MultiSet.fromElements((_this).f27, (_this).f27);
      let _1774_v76;
      _1774_v76 = _dafny.Set.fromElements(_1772_v74, _dafny.Seq.update(_1772_v74, _module.__default.safeIndex(new BigNumber((_1773_v75).cardinality()), new BigNumber((_1772_v74).length)), _1771_v73), _1772_v74, _1772_v74);
      let _1775_v77;
      _1775_v77 = _dafny.Map.Empty.slice().updateUnsafe(_1771_v73,_1686_v0);
      let _1776_v78;
      _1776_v78 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,false);
      let _1777_v79;
      _1777_v79 = _dafny.Seq.of((((_1776_v78).contains(new BigNumber((_dafny.Seq.update(_1772_v74, _module.__default.safeIndex((_this).f27, new BigNumber((_1772_v74).length)), _1771_v73)).length))) ? ((_1776_v78).get(new BigNumber((_dafny.Seq.update(_1772_v74, _module.__default.safeIndex((_this).f27, new BigNumber((_1772_v74).length)), _1771_v73)).length))) : (_1686_v0)));
      let _1778_v80;
      _1778_v80 = _module.D15.create_DC36(((_1686_v0) ? (_1771_v73) : (_1771_v73)), (_this).fm4(_1774_v76, (((_1775_v77).contains(_1771_v73)) ? ((_1775_v77).get(_1771_v73)) : (_1686_v0)), _1777_v79, globalState));
      _1778_v80 = _1778_v80;
      _1772_v74 = _dafny.Seq.Concat(_module.__default.fm31(_1686_v0, _1686_v0, globalState), _1772_v74);
      let _1779_i4;
      _1779_i4 = _dafny.ZERO;
      L7: {
        while (_1686_v0) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1779_i4)) {
              break L7;
            }
            _1779_i4 = (_1779_i4).plus(_dafny.ONE);
            let _1780_v81;
            let _nw299 = new _module.C1();
            _nw299.__ctor((new BigNumber(((_this).f30).length)).isLessThanOrEqualTo((_this).f27), (_this).f29, (_this).f30, (_this).f27, (_this).f28);
            _1780_v81 = _nw299;
            let _1781_v82;
            _1781_v82 = _dafny.Seq.of(_1780_v81);
            _1780_v81 = ((_1686_v0) ? ((_1781_v82)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((_this).f30).length)), new BigNumber((_1781_v82).length))]) : (_1780_v81));
            _1772_v74 = _1772_v74;
            (globalState).f21 = (_this).f27;
            let _1782_v83;
            _1782_v83 = _dafny.MultiSet.fromElements(_1771_v73);
            let _1783_v84;
            _1783_v84 = _dafny.Map.Empty.slice().updateUnsafe((_1686_v0) && (_1780_v81.f35),_1782_v83);
            _1783_v84 = (_1783_v84).update(_dafny.Seq.IsProperPrefixOf(_1777_v79, _1777_v79), _1782_v83);
          }
        }
      }
      let _1784_v85;
      let _nw300 = new _module.C3();
      _nw300.__ctor((_this).f27, (_this).f28);
      _1784_v85 = _nw300;
      if (_1686_v0) {
        _1686_v0 = _module.__default.fm0(new BigNumber(860), globalState);
        let _index223 = _module.__default.safeIndex(new BigNumber(243), new BigNumber(((_this).f28).length));
        ((_this).f28)[_index223] = _1772_v74;
        let _1785_v86;
        let _init37 = ((_1786_v0) => function (_1787_i5) {
          return _1786_v0;
        })(_1686_v0);
        let _nw301 = Array((new BigNumber(28)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw301.length); _i0_37++) {
          _nw301[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1785_v86 = _nw301;
        let _1788_v87;
        _1788_v87 = _dafny.Set.fromElements((_this).f27);
        let _1789_v88;
        _1789_v88 = _module.D4.create_DC8(_1771_v73, new BigNumber((_1788_v87).length), _1686_v0);
        let _index224 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1785_v86).length));
        (_1785_v86)[_index224] = (_1777_v79)[_module.__default.safeIndex((_1789_v88).dtor_cf14, new BigNumber((_1777_v79).length))];
        let _1790_v89;
        _1790_v89 = _module.D6.create_DC12(_module.__default.fm0((_this).f27, globalState), _module.__default.fm44(globalState), _dafny.Seq.update(_module.__default.fm31(_1686_v0, _1686_v0, globalState), _module.__default.safeIndex((_this).f27, new BigNumber((_module.__default.fm31(_1686_v0, _1686_v0, globalState)).length)), _module.__default.fm38(_dafny.Map.Empty.slice().updateUnsafe(_1686_v0,_1686_v0), globalState)));
        let _index225 = _module.__default.safeIndex(new BigNumber(243), new BigNumber(((_this).f28).length));
        let _index226 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1785_v86).length));
        let _rhs318 = (_1790_v89).dtor_cf22;
        let _rhs319 = (_1686_v0) === (false);
        let _lhs249 = (_this).f28;
        let _lhs250 = _module.__default.safeIndex(new BigNumber(243), new BigNumber(((_this).f28).length));
        let _lhs251 = _1785_v86;
        let _lhs252 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1785_v86).length));
        _lhs249[_lhs250] = _rhs318;
        _lhs251[_lhs252] = _rhs319;
        let _1791_v90;
        let _init38 = ((_1792_v73, _1793_v0) => function (_1794_i6) {
          return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(882)), ((_1795_v73) => function (_1796_i7) {
            return _module.D1.create_DC2(_1795_v73, (_this).f27);
          })(_1792_v73)),_1793_v0);
        })(_1771_v73, _1686_v0);
        let _nw302 = Array((_dafny.ONE).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw302.length); _i0_38++) {
          _nw302[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1791_v90 = _nw302;
        let _1797_v91;
        _1797_v91 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,new BigNumber(-418));
        let _1798_v92;
        _1798_v92 = _dafny.Map.Empty.slice().updateUnsafe((_1785_v86)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1785_v86).length))],_dafny.Set.fromElements(_1686_v0));
        let _1799_v93;
        _1799_v93 = _module.D1.create_DC2(_1771_v73, (((_1797_v91).contains(new BigNumber((_1777_v79).length))) ? ((_1797_v91).get(new BigNumber((_1777_v79).length))) : (new BigNumber((_1798_v92).length))));
        let _1800_v94;
        _1800_v94 = _dafny.Seq.of(_1799_v93);
        let _1801_v95;
        _1801_v95 = _dafny.Map.Empty.slice().updateUnsafe(_1800_v94,(_1785_v86)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1785_v86).length))]);
        let _index227 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1791_v90).length));
        (_1791_v90)[_index227] = (_dafny.Map.Empty.slice().updateUnsafe(_1800_v94,(_1784_v85).fm46((_this).f27, _1771_v73, globalState))).Merge(_1801_v95);
        let _index228 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1791_v90).length));
        (_1791_v90)[_index228] = _1801_v95;
        let _index229 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1785_v86).length));
        (_1785_v86)[_index229] = (_1785_v86)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1785_v86).length))];
        let _1802_v96;
        _1802_v96 = _dafny.Seq.of((_this).f30);
        let _1803_v97;
        _1803_v97 = _dafny.Map.Empty.slice().updateUnsafe(_1686_v0,_1771_v73);
        let _1804_v98;
        _1804_v98 = _dafny.Map.Empty.slice().updateUnsafe(!(new BigNumber((_1802_v96).length)).isEqualTo(new BigNumber((_1803_v97).length)),(_this).f27);
        (globalState).f21 = (((_1804_v98).contains((_1785_v86)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1785_v86).length))])) ? ((_1804_v98).get((_1785_v86)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1785_v86).length))])) : ((_this).f27));
      } else {
        _1686_v0 = _1686_v0;
        let _1805_v99;
        _1805_v99 = _dafny.Set.fromElements(_1771_v73, _1771_v73);
        let _1806_v100;
        _1806_v100 = _module.D19.create_DC48(_1805_v99);
        (globalState).f3 = (new BigNumber(363)).minus(new BigNumber(((function () {
          let _coll54 = new _dafny.Set();
          for (const _compr_54 of (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm37((_this).f27, new BigNumber(((_1806_v100).dtor_cf88).length), globalState),_1686_v0)).Keys.Elements) {
            let _1807_v101 = _compr_54;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm37((_this).f27, new BigNumber(((_1806_v100).dtor_cf88).length), globalState),_1686_v0)).contains(_1807_v101)) {
              _coll54.add(_1807_v101);
            }
          }
          return _coll54;
        }()).Union(_1774_v76)).length));
        let _1808_v102;
        _1808_v102 = _dafny.Map.Empty.slice().updateUnsafe(_1686_v0,_1686_v0);
        _1808_v102 = (_1808_v102).update(((_this).f27).isEqualTo(new BigNumber((_1772_v74).length)), _1686_v0);
        let _1809_v103;
        _1809_v103 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f27);
        let _1810_v104;
        let _nw303 = new _module.C3();
        _nw303.__ctor((_this).f27, (_this).f28);
        _1810_v104 = _nw303;
        let _1811_v105;
        let _nw304 = Array((new BigNumber(29)).toNumber());
        _nw304[(_dafny.ZERO).toNumber()] = (_this).f27;
        _nw304[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_1809_v103).contains((_this).f27)) ? ((_1809_v103).get((_this).f27)) : (new BigNumber((_module.__default.fm48(_1686_v0, _1686_v0, new BigNumber((_dafny.Set.fromElements(_1810_v104)).length), globalState)).length))),(_this).f27)).length);
        _nw304[(new BigNumber(2)).toNumber()] = new BigNumber(477);
        _nw304[(new BigNumber(3)).toNumber()] = (_1810_v104).f27;
        _nw304[(new BigNumber(4)).toNumber()] = (_this).f27;
        _nw304[(new BigNumber(5)).toNumber()] = (_1810_v104).f27;
        _nw304[(new BigNumber(6)).toNumber()] = new BigNumber((_1772_v74).length);
        _nw304[(new BigNumber(7)).toNumber()] = (_this).f27;
        _nw304[(new BigNumber(8)).toNumber()] = (_this).f27;
        _nw304[(new BigNumber(9)).toNumber()] = (_this).f27;
        _nw304[(new BigNumber(10)).toNumber()] = (_1810_v104).f27;
        _nw304[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((_this).f27);
        _nw304[(new BigNumber(12)).toNumber()] = new BigNumber((_1776_v78).length);
        _nw304[(new BigNumber(13)).toNumber()] = (_this).f27;
        _nw304[(new BigNumber(14)).toNumber()] = (_this).f27;
        _nw304[(new BigNumber(15)).toNumber()] = (_this).f27;
        _nw304[(new BigNumber(16)).toNumber()] = new BigNumber((_1772_v74).length);
        _nw304[(new BigNumber(17)).toNumber()] = (_1810_v104).f27;
        _nw304[(new BigNumber(18)).toNumber()] = (_this).f27;
        _nw304[(new BigNumber(19)).toNumber()] = (_1810_v104).f27;
        _nw304[(new BigNumber(20)).toNumber()] = new BigNumber(386);
        _nw304[(new BigNumber(21)).toNumber()] = (_1810_v104).f27;
        _nw304[(new BigNumber(22)).toNumber()] = (_1810_v104).f27;
        _nw304[(new BigNumber(23)).toNumber()] = (_this).f27;
        _nw304[(new BigNumber(24)).toNumber()] = (_this).f27;
        _nw304[(new BigNumber(25)).toNumber()] = (_this).fm4(_1774_v76, false, _1777_v79, globalState);
        _nw304[(new BigNumber(26)).toNumber()] = (_1810_v104).f27;
        _nw304[(new BigNumber(27)).toNumber()] = (_1810_v104).f27;
        _nw304[(new BigNumber(28)).toNumber()] = new BigNumber(235);
        _1811_v105 = _nw304;
        let _1812_v106;
        _1812_v106 = _dafny.Set.fromElements(_1811_v105, _1811_v105, _1811_v105, _1811_v105);
        let _1813_v107;
        _1813_v107 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).fm3((_this).f29, globalState)),_1809_v103);
        let _1814_v108;
        let _nw305 = Array((new BigNumber(17)).toNumber());
        _nw305[(_dafny.ZERO).toNumber()] = _1809_v103;
        _nw305[(_dafny.ONE).toNumber()] = _1809_v103;
        _nw305[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f27);
        _nw305[(new BigNumber(3)).toNumber()] = (_1809_v103).Merge(_1809_v103);
        _nw305[(new BigNumber(4)).toNumber()] = _1809_v103;
        _nw305[(new BigNumber(5)).toNumber()] = (_1809_v103).update((_this).f27, (_this).f27);
        _nw305[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f27);
        _nw305[(new BigNumber(7)).toNumber()] = _1809_v103;
        _nw305[(new BigNumber(8)).toNumber()] = _1809_v103;
        _nw305[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-591),new BigNumber((_1812_v106).length));
        _nw305[(new BigNumber(10)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_1810_v104).f27)).update(new BigNumber((_1777_v79).length), (_this).f27);
        _nw305[(new BigNumber(11)).toNumber()] = _1809_v103;
        _nw305[(new BigNumber(12)).toNumber()] = ((_1686_v0) ? (_dafny.Map.Empty.slice().updateUnsafe((_1810_v104).f27,(_this).f27)) : (_1809_v103));
        _nw305[(new BigNumber(13)).toNumber()] = _1809_v103;
        _nw305[(new BigNumber(14)).toNumber()] = (((_1813_v107).contains((((_1808_v102).contains((_1777_v79)[_module.__default.safeIndex((_1810_v104).f27, new BigNumber((_1777_v79).length))])) ? ((_1808_v102).get((_1777_v79)[_module.__default.safeIndex((_1810_v104).f27, new BigNumber((_1777_v79).length))])) : (_1686_v0)))) ? ((_1813_v107).get((((_1808_v102).contains((_1777_v79)[_module.__default.safeIndex((_1810_v104).f27, new BigNumber((_1777_v79).length))])) ? ((_1808_v102).get((_1777_v79)[_module.__default.safeIndex((_1810_v104).f27, new BigNumber((_1777_v79).length))])) : (_1686_v0)))) : (_1809_v103));
        _nw305[(new BigNumber(15)).toNumber()] = _1809_v103;
        _nw305[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(230),new BigNumber((_1772_v74).length));
        _1814_v108 = _nw305;
        let _index230 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_1814_v108).length));
        (_1814_v108)[_index230] = (_1809_v103).Merge(_1809_v103);
        let _1815_v110;
        _1815_v110 = _dafny.Map.Empty.slice().updateUnsafe((_1810_v104).f27,_1771_v73);
        let _1816_v111;
        _1816_v111 = _dafny.Map.Empty.slice().updateUnsafe((_1810_v104).f27,_1815_v110);
        let _1817_v112;
        _1817_v112 = _dafny.Seq.of(function () {
          let _coll55 = new _dafny.Map();
          for (const _compr_55 of _dafny.IntegerRange(new BigNumber(124), new BigNumber(382))) {
            let _1818_v109 = _compr_55;
            if (((new BigNumber(124)).isLessThanOrEqualTo(_1818_v109)) && ((_1818_v109).isLessThan(new BigNumber(382)))) {
              _coll55.push([_module.__default.safeModuloInt(_1818_v109, (_1810_v104).f27),new _dafny.CodePoint('t'.codePointAt(0))]);
            }
          }
          return _coll55;
        }(), (_1815_v110).update(new BigNumber((_dafny.Seq.of(!(_1686_v0))).length), _1771_v73), _1815_v110, (((_1816_v111).contains((_this).f27)) ? ((_1816_v111).get((_this).f27)) : (_1815_v110)));
        let _1819_v113;
        _1819_v113 = _module.D15.create_DC35(_1817_v112);
        let _1820_v114;
        _1820_v114 = _dafny.Seq.of(_module.D15.create_DC35(_1817_v112), _1819_v113);
        let _1821_v115;
        _1821_v115 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(362)), ((_1822_v104, _1823_v73) => function (_1824_i8) {
          return _module.D15.create_DC35(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_1822_v104).f27,_1823_v73)));
        })(_1810_v104, _1771_v73)), _module.__default.safeIndex(new BigNumber(-918), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(362)), ((_1825_v104, _1826_v73) => function (_1827_i8) {
          return _module.D15.create_DC35(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_1825_v104).f27,_1826_v73)));
        })(_1810_v104, _1771_v73))).length)), _1819_v113), _dafny.Seq.Concat(_1820_v114, _1820_v114), _dafny.Seq.Concat(_1820_v114, _1820_v114));
        let _1828_v116;
        let _nw306 = new _module.C0();
        _nw306.__ctor((_this).f27, _1772_v74);
        _1828_v116 = _nw306;
        let _index231 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_1814_v108).length));
        let _rhs320 = ((_1809_v103).Merge(_1809_v103)).Merge(_1809_v103);
        let _rhs321 = _1686_v0;
        let _rhs322 = _1821_v115;
        let _rhs323 = _1828_v116;
        let _lhs253 = _1814_v108;
        let _lhs254 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_1814_v108).length));
        _lhs253[_lhs254] = _rhs320;
        r2 = _rhs321;
        _1821_v115 = _rhs322;
        _1828_v116 = _rhs323;
        let _1829_v117;
        _1829_v117 = _dafny.Seq.of(_1811_v105);
        _1829_v117 = _dafny.Seq.Concat(_1829_v117, _1829_v117);
      }
      let _1830_v118;
      _1830_v118 = _module.D9.create_DC19(!(_1686_v0));
      let _1831_v119;
      _1831_v119 = _dafny.Set.fromElements(true, true, _1686_v0, _1686_v0, _1686_v0);
      let _nw307 = Array((new BigNumber(3)).toNumber());
      _nw307[(_dafny.ZERO).toNumber()] = !((_1830_v118).dtor_cf33) || (_module.__default.fm0((_this).f27, globalState));
      _nw307[(_dafny.ONE).toNumber()] = _1686_v0;
      _nw307[(new BigNumber(2)).toNumber()] = !((_1831_v119).IsProperSubsetOf(_1831_v119));
      r0 = _nw307;
      let _1832_v120;
      _1832_v120 = _module.D7.create_DC13(_1777_v79);
      let _pat_let_tv37 = _1686_v0;
      let _pat_let_tv38 = _1686_v0;
      r1 = function (_source26) {
        if (_source26.is_DC14) {
          let _1833___mcc_h0 = (_source26).cf24;
          let _1834___mcc_h1 = (_source26).cf25;
          let _1835___mcc_h2 = (_source26).cf26;
          let _1836___mcc_h3 = (_source26).cf27;
          let _1837_cf27 = _1836___mcc_h3;
          let _1838_cf26 = _1835___mcc_h2;
          let _1839_cf25 = _1834___mcc_h1;
          let _1840_cf24 = _1833___mcc_h0;
          return _pat_let_tv37;
        } else {
          let _1841___mcc_h4 = (_source26).cf23;
          let _1842_cf23 = _1841___mcc_h4;
          return _pat_let_tv38;
        }
      }(_1832_v120);
      r2 = _dafny.Seq.IsPrefixOf(_dafny.Seq.of((_this).f27, (_this).f27), _dafny.Seq.update(_dafny.Seq.Concat((_this).f30, _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(((_this).f30)[_module.__default.safeIndex(new BigNumber(58), new BigNumber(((_this).f30).length))],_1686_v0)).length))), _module.__default.safeIndex(new BigNumber(((_dafny.MultiSet.FromArray((_this).f30)).update((_this).f27, _module.__default.abs((_this).f27))).cardinality()), new BigNumber((_dafny.Seq.Concat((_this).f30, _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(((_this).f30)[_module.__default.safeIndex(new BigNumber(58), new BigNumber(((_this).f30).length))],_1686_v0)).length)))).length)), new BigNumber(669)));
      return [r0, r1, r2];
    }
    m10(p0, p1, globalState) {
      let _this = this;
      let _1843_v0;
      _1843_v0 = new _dafny.CodePoint('i'.codePointAt(0));
      let _1844_v1;
      _1844_v1 = _dafny.Seq.UnicodeFromString("xggvd");
      let _1845_v2;
      _1845_v2 = _module.D2.create_DC4((_this).f27, new BigNumber(521), (_dafny.ZERO).minus((_this).f27), _1843_v0, _1844_v1);
      let _1846_v3;
      _1846_v3 = _dafny.Set.fromElements(_1845_v2);
      let _hi8 = ((p0) ? (new BigNumber((_1846_v3).length)) : ((_this).f27));
      for (let _1847_i0 = ((p0) ? ((_this).f27) : ((_this).f27)); _1847_i0.isLessThan(_hi8); _1847_i0 = _1847_i0.plus(_dafny.ONE)) {
        let _1848_v4;
        _1848_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,p0);
        let _1849_v5;
        let _nw308 = new _module.C2();
        _nw308.__ctor(_dafny.MultiSet.fromElements(p0), new BigNumber(((_1848_v4).Merge(_1848_v4)).length));
        _1849_v5 = _nw308;
        (globalState).f26 = p0;
        let _1850_v6;
        _1850_v6 = _dafny.Seq.of(p0);
        _1850_v6 = _1850_v6;
        let _1851_v7;
        let _nw309 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1851_v7 = _nw309;
        let _1852_v8;
        _1852_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1850_v6,_1851_v7);
        let _1853_v9;
        _1853_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1850_v6,(((_1852_v8).contains(_1850_v6)) ? ((_1852_v8).get(_1850_v6)) : (_1851_v7)));
        _1853_v9 = (_1853_v9).update(_dafny.Seq.of(!(p0), p0), _1851_v7);
      }
      let _1854_v10;
      let _nw310 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
      _1854_v10 = _nw310;
      let _1855_v11;
      _1855_v11 = _module.D16.create_DC39(_1854_v10);
      let _pat_let_tv39 = _1854_v10;
      let _pat_let_tv40 = _1854_v10;
      _1855_v11 = function (_pat_let52_0) {
        return function (_1858_dt__update__tmp_h1) {
          return function (_pat_let55_0) {
            return function (_1859_dt__update_hcf73_h1) {
              return _module.D16.create_DC39(_1859_dt__update_hcf73_h1);
            }(_pat_let55_0);
          }(_pat_let_tv40);
        }(_pat_let52_0);
      }(function (_pat_let53_0) {
        return function (_1856_dt__update__tmp_h0) {
          return function (_pat_let54_0) {
            return function (_1857_dt__update_hcf73_h0) {
              return _module.D16.create_DC39(_1857_dt__update_hcf73_h0);
            }(_pat_let54_0);
          }(_pat_let_tv39);
        }(_pat_let53_0);
      }(_1855_v11));
      let _1860_v12;
      _1860_v12 = _dafny.Seq.of(((_this).f27).isLessThanOrEqualTo(new BigNumber(386)));
      let _rhs324 = _dafny.Seq.of(p0, p0);
      let _rhs325 = _module.__default.fm1(globalState);
      let _rhs326 = (_this).f27;
      let _lhs255 = globalState;
      let _lhs256 = globalState;
      _1860_v12 = _rhs324;
      _lhs255.f24 = _rhs325;
      _lhs256.f0 = _rhs326;
      let _1861_v13;
      _1861_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(425),(_dafny.ZERO).minus((_this).f27));
      let _1862_v14;
      _1862_v14 = _dafny.Map.Empty.slice().updateUnsafe((((_1861_v13).contains(((_this).f30)[_module.__default.safeIndex((_this).f27, new BigNumber(((_this).f30).length))])) ? ((_1861_v13).get(((_this).f30)[_module.__default.safeIndex((_this).f27, new BigNumber(((_this).f30).length))])) : ((_dafny.ZERO).minus((_this).f27))),false);
      _1862_v14 = (_1862_v14).update(new BigNumber((_module.__default.fm48(p0, p0, new BigNumber(-276), globalState)).length), p0);
      _1843_v0 = _1843_v0;
      let _1863_v15;
      _1863_v15 = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm0((_this).f27, globalState)),true);
      let _1864_v16;
      _1864_v16 = _module.D8.create_DC15((_1863_v15).Merge(_1863_v15));
      let _source27 = _1864_v16;
      if (_source27.is_DC16) {
        let _1865___mcc_h0 = (_source27).cf29;
        let _1866___mcc_h1 = (_source27).cf30;
        let _1867_cf30 = _1866___mcc_h1;
        let _1868_cf29 = _1865___mcc_h0;
        let _1869_v17;
        _1869_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.UnicodeFromString("jiyhtfc"));
        let _1870_v18;
        let _nw311 = Array((new BigNumber(20)).toNumber());
        _nw311[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("rnsowybbf");
        _nw311[(_dafny.ONE).toNumber()] = _1844_v1;
        _nw311[(new BigNumber(2)).toNumber()] = _1844_v1;
        _nw311[(new BigNumber(3)).toNumber()] = _module.__default.fm31(p0, false, globalState);
        _nw311[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_1844_v1, _1844_v1);
        _nw311[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm37((_this).f27, new BigNumber(422), globalState), _1844_v1);
        _nw311[(new BigNumber(6)).toNumber()] = _1844_v1;
        _nw311[(new BigNumber(7)).toNumber()] = _1844_v1;
        _nw311[(new BigNumber(8)).toNumber()] = _1844_v1;
        _nw311[(new BigNumber(9)).toNumber()] = _1844_v1;
        _nw311[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yo"), _1844_v1), _module.__default.safeIndex(_1867_cf30, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yo"), _1844_v1)).length)), _1843_v0);
        _nw311[(new BigNumber(11)).toNumber()] = _1844_v1;
        _nw311[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_1844_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-646)), ((_1871_v0) => function (_1872_i1) {
          return _1871_v0;
        })(_1843_v0)));
        _nw311[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-14)), ((_1873_v0) => function (_1874_i2) {
          return _1873_v0;
        })(_1843_v0)), _1844_v1);
        _nw311[(new BigNumber(14)).toNumber()] = _1844_v1;
        _nw311[(new BigNumber(15)).toNumber()] = _1844_v1;
        _nw311[(new BigNumber(16)).toNumber()] = (((_1869_v17).contains(true)) ? ((_1869_v17).get(true)) : (_1844_v1));
        _nw311[(new BigNumber(17)).toNumber()] = _1844_v1;
        _nw311[(new BigNumber(18)).toNumber()] = _1844_v1;
        _nw311[(new BigNumber(19)).toNumber()] = _dafny.Seq.UnicodeFromString("ggdic");
        _1870_v18 = _nw311;
        let _1875_v19;
        _1875_v19 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_1844_v1).length)).minus(new BigNumber((_1844_v1).length)),(_this).f28);
        _1870_v18 = (((_1875_v19).contains(_1868_cf29)) ? ((_1875_v19).get(_1868_cf29)) : ((_this).f28));
        let _1876_v20;
        let _nw312 = Array((new BigNumber(16)).toNumber()).fill(_module.D12.Default());
        _1876_v20 = _nw312;
        let _1877_v21;
        _1877_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1876_v20,_1843_v0);
        let _1878_v22;
        _1878_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1868_cf29,_1843_v0);
        let _rhs327 = ((_this).f27).multipliedBy((p1)[_module.__default.safeIndex(_1868_cf29, new BigNumber((p1).length))]);
        let _rhs328 = p0;
        let _rhs329 = _dafny.Map.Empty.slice().updateUnsafe(_1876_v20,(((_1878_v22).contains((_this).f27)) ? ((_1878_v22).get((_this).f27)) : (_1843_v0)));
        let _rhs330 = new BigNumber(-672);
        let _lhs257 = globalState;
        let _lhs258 = globalState;
        let _lhs259 = globalState;
        _lhs257.f0 = _rhs327;
        _lhs258.f26 = _rhs328;
        _1877_v21 = _rhs329;
        _lhs259.f18 = _rhs330;
        let _1879_v23;
        _1879_v23 = _dafny.Seq.of(_1844_v1);
        let _1880_v24;
        let _nw313 = new _module.C4();
        _nw313.__ctor((new BigNumber((_1844_v1).length)).minus(_1868_cf29), (_this).f28);
        _1880_v24 = _nw313;
        let _1881_v25;
        let _init39 = ((_1882_cf29) => function (_1883_i3) {
          return _module.__default.safeDivisionInt(_1883_i3, _1882_cf29);
        })(_1868_cf29);
        let _nw314 = Array((new BigNumber(28)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw314.length); _i0_39++) {
          _nw314[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1881_v25 = _nw314;
        let _1884_v26;
        _1884_v26 = _dafny.MultiSet.fromElements((new BigNumber(737)).isLessThan(new BigNumber((_dafny.Seq.update(_1844_v1, _module.__default.safeIndex((_this).f27, new BigNumber((_1844_v1).length)), _1843_v0)).length)), false, p0, ((p0) ? (p0) : (p0)));
        let _1885_v27;
        _1885_v27 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f27);
        let _rhs331 = _1879_v23;
        let _rhs332 = _1880_v24;
        let _rhs333 = ((p0) ? (_1881_v25) : (_1881_v25));
        let _rhs334 = (_1868_cf29).minus((new BigNumber((_1885_v27).length)).plus((_this).f27));
        let _rhs335 = _dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.Concat(_module.__default.fm41(_1868_cf29, new BigNumber((_1862_v14).length), (_this).fm3(_dafny.MultiSet.fromElements(p0), globalState), globalState), _1860_v12), _module.__default.safeIndex(_1868_cf29, new BigNumber((_dafny.Seq.Concat(_module.__default.fm41(_1868_cf29, new BigNumber((_1862_v14).length), (_this).fm3(_dafny.MultiSet.fromElements(p0), globalState), globalState), _1860_v12)).length)), p0));
        let _lhs260 = globalState;
        _1879_v23 = _rhs331;
        _1880_v24 = _rhs332;
        _1881_v25 = _rhs333;
        _lhs260.f0 = _rhs334;
        _1884_v26 = _rhs335;
        let _1886_v28;
        let _nw315 = new _module.C0();
        _nw315.__ctor(new BigNumber((_1844_v1).length), _1844_v1);
        _1886_v28 = _nw315;
        let _1887_v29;
        _1887_v29 = _module.D6.create_DC11(_1886_v28);
        let _source28 = _1887_v29;
        if (_source28.is_DC12) {
          let _1888___mcc_h3 = (_source28).cf20;
          let _1889___mcc_h4 = (_source28).cf21;
          let _1890___mcc_h5 = (_source28).cf22;
          let _1891_cf22 = _1890___mcc_h5;
          let _1892_cf21 = _1889___mcc_h4;
          let _1893_cf20 = _1888___mcc_h3;
          let _1894_v30;
          let _nw316 = new _module.C0();
          _nw316.__ctor((_1880_v24).f27, (_1886_v28).f37);
          _1894_v30 = _nw316;
          _1868_cf29 = (_1886_v28).f36;
          let _1895_v31;
          _1895_v31 = _module.D4.create_DC8(_1843_v0, _1867_cf30, false);
          let _1896_v32;
          _1896_v32 = _dafny.Set.fromElements((_1894_v30).f36, (_1894_v30).f36, _1868_cf29, _1868_cf29, (_1880_v24).f27);
          let _1897_v33;
          let _nw317 = Array((new BigNumber(15)).toNumber());
          _nw317[(_dafny.ZERO).toNumber()] = p0;
          _nw317[(_dafny.ONE).toNumber()] = p0;
          _nw317[(new BigNumber(2)).toNumber()] = (_this).fm3((_this).f29, globalState);
          _nw317[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsPrefixOf(_1860_v12, _1860_v12);
          _nw317[(new BigNumber(4)).toNumber()] = _1893_cf20;
          _nw317[(new BigNumber(5)).toNumber()] = p0;
          _nw317[(new BigNumber(6)).toNumber()] = p0;
          _nw317[(new BigNumber(7)).toNumber()] = (_1896_v32).contains(_1867_cf30);
          _nw317[(new BigNumber(8)).toNumber()] = _1893_cf20;
          _nw317[(new BigNumber(9)).toNumber()] = (_module.__default.fm49((_this).f27, globalState)).dtor_cf17;
          _nw317[(new BigNumber(10)).toNumber()] = _1893_cf20;
          _nw317[(new BigNumber(11)).toNumber()] = ((_1880_v24).f27).isLessThanOrEqualTo(_1868_cf29);
          _nw317[(new BigNumber(12)).toNumber()] = p0;
          _nw317[(new BigNumber(13)).toNumber()] = false;
          _nw317[(new BigNumber(14)).toNumber()] = _1893_cf20;
          _1897_v33 = _nw317;
          let _index232 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1897_v33).length));
          (_1897_v33)[_index232] = false;
          let _pat_let_tv41 = _1894_v30;
          let _pat_let_tv42 = _1868_cf29;
          let _index233 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1897_v33).length));
          let _rhs336 = _1867_cf30;
          let _rhs337 = function (_pat_let56_0) {
            return function (_1898_dt__update__tmp_h2) {
              return function (_pat_let57_0) {
                return function (_1899_dt__update_hcf15_h0) {
                  return _module.D4.create_DC8((_1898_dt__update__tmp_h2).dtor_cf13, (_1898_dt__update__tmp_h2).dtor_cf14, _1899_dt__update_hcf15_h0);
                }(_pat_let57_0);
              }(((_pat_let_tv41).f36).isEqualTo(_pat_let_tv42));
            }(_pat_let56_0);
          }(_module.D4.create_DC8(_1843_v0, (_1886_v28).f36, _1893_cf20));
          let _rhs338 = !((new BigNumber((_1860_v12).length)).isLessThan((_1886_v28).f36));
          let _lhs261 = globalState;
          let _lhs262 = _1897_v33;
          let _lhs263 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1897_v33).length));
          _lhs261.f0 = _rhs336;
          _1895_v31 = _rhs337;
          _lhs262[_lhs263] = _rhs338;
          (globalState).f24 = (_1894_v30).f36;
        } else {
          let _1900___mcc_h6 = (_source28).cf19;
          let _1901_cf19 = _1900___mcc_h6;
          let _rhs339 = _1881_v25;
          let _rhs340 = (((_1862_v14).contains((_1880_v24).f27)) ? ((_1862_v14).get((_1880_v24).f27)) : (p0));
          let _rhs341 = ((_1880_v24).f27).multipliedBy(_1868_cf29);
          let _rhs342 = ((p0) ? (!(p0)) : (false));
          let _rhs343 = _1861_v13;
          let _lhs264 = globalState;
          let _lhs265 = globalState;
          let _lhs266 = globalState;
          _1881_v25 = _rhs339;
          _lhs264.f26 = _rhs340;
          _lhs265.f24 = _rhs341;
          _lhs266.f26 = _rhs342;
          _1861_v13 = _rhs343;
          _1863_v15 = (_1863_v15).update(p0, p0);
          let _1902_v34;
          let _nw318 = new _module.C1();
          _nw318.__ctor(p0, ((p0) ? (_1884_v26) : ((_this).f29)), (_this).f30, ((_1880_v24).f27).multipliedBy(_1868_cf29), (_1880_v24).f28);
          _1902_v34 = _nw318;
          (globalState).f0 = _module.__default.safeDivisionInt(new BigNumber((_1844_v1).length), (_1886_v28).f36);
        }
      } else {
        let _1903___mcc_h2 = (_source27).cf28;
        let _1904_cf28 = _1903___mcc_h2;
        let _1905_v35;
        let _nw319 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1905_v35 = _nw319;
        let _1906_v36;
        let _nw320 = new _module.C0();
        _nw320.__ctor((_this).f27, _1844_v1);
        _1906_v36 = _nw320;
        let _1907_v37;
        let _nw321 = Array((new BigNumber(20)).toNumber()).fill(false);
        _1907_v37 = _nw321;
        let _index234 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_1907_v37).length));
        (_1907_v37)[_index234] = p0;
        let _index235 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_1907_v37).length));
        let _rhs344 = _1905_v35;
        let _rhs345 = _1906_v36;
        let _rhs346 = !(p0);
        let _lhs267 = _1907_v37;
        let _lhs268 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_1907_v37).length));
        _1905_v35 = _rhs344;
        _1906_v36 = _rhs345;
        _lhs267[_lhs268] = _rhs346;
        let _1908_v38;
        _1908_v38 = _dafny.MultiSet.fromElements((_1907_v37)[_module.__default.safeIndex(new BigNumber(391), new BigNumber((_1907_v37).length))]);
        let _1909_v39;
        let _nw322 = Array((new BigNumber(11)).toNumber()).fill([]);
        _1909_v39 = _nw322;
        let _1910_v40;
        let _nw323 = new _module.C3();
        _nw323.__ctor((_1906_v36).f36, (_this).f28);
        _1910_v40 = _nw323;
        let _1911_v41;
        _1911_v41 = _dafny.Seq.of(_1910_v40, _1910_v40);
        let _1912_v42;
        let _nw324 = Array((new BigNumber(26)).toNumber());
        _nw324[(_dafny.ZERO).toNumber()] = (_1906_v36).f36;
        _nw324[(_dafny.ONE).toNumber()] = (_1906_v36).f36;
        _nw324[(new BigNumber(2)).toNumber()] = (_this).f27;
        _nw324[(new BigNumber(3)).toNumber()] = (_this).f27;
        _nw324[(new BigNumber(4)).toNumber()] = new BigNumber((_1911_v41).length);
        _nw324[(new BigNumber(5)).toNumber()] = new BigNumber((_1863_v15).length);
        _nw324[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("gtwh"), _module.__default.safeIndex((_1906_v36).f36, new BigNumber((_dafny.Seq.UnicodeFromString("gtwh")).length)), _1843_v0)).length);
        _nw324[(new BigNumber(7)).toNumber()] = (_this).f27;
        _nw324[(new BigNumber(8)).toNumber()] = (_1906_v36).f36;
        _nw324[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("qrfkfxp")).length);
        _nw324[(new BigNumber(10)).toNumber()] = _module.__default.fm1(globalState);
        _nw324[(new BigNumber(11)).toNumber()] = (_this).f27;
        _nw324[(new BigNumber(12)).toNumber()] = (_1906_v36).f36;
        _nw324[(new BigNumber(13)).toNumber()] = ((_this).f30)[_module.__default.safeIndex(new BigNumber(896), new BigNumber(((_this).f30).length))];
        _nw324[(new BigNumber(14)).toNumber()] = (_1906_v36).f36;
        _nw324[(new BigNumber(15)).toNumber()] = new BigNumber(943);
        _nw324[(new BigNumber(16)).toNumber()] = (_1906_v36).f36;
        _nw324[(new BigNumber(17)).toNumber()] = (_this).f27;
        _nw324[(new BigNumber(18)).toNumber()] = new BigNumber(719);
        _nw324[(new BigNumber(19)).toNumber()] = (_this).f27;
        _nw324[(new BigNumber(20)).toNumber()] = new BigNumber(306);
        _nw324[(new BigNumber(21)).toNumber()] = (_this).f27;
        _nw324[(new BigNumber(22)).toNumber()] = new BigNumber(947);
        _nw324[(new BigNumber(23)).toNumber()] = new BigNumber(-688);
        _nw324[(new BigNumber(24)).toNumber()] = (_this).f27;
        _nw324[(new BigNumber(25)).toNumber()] = (_this).f27;
        _1912_v42 = _nw324;
        let _index236 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_1909_v39).length));
        (_1909_v39)[_index236] = _1912_v42;
        let _1913_v43;
        _1913_v43 = _dafny.MultiSet.fromElements(_1843_v0);
        let _1914_v44;
        let _nw325 = new _module.C2();
        _nw325.__ctor(_1908_v38, (_this).f27);
        _1914_v44 = _nw325;
        let _1915_v45;
        _1915_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1914_v44,(_1907_v37)[_module.__default.safeIndex(new BigNumber(391), new BigNumber((_1907_v37).length))]);
        let _index237 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_1909_v39).length));
        let _rhs347 = ((((_1913_v43).contains(_1843_v0)) ? ((_1913_v43).get(_1843_v0)) : (new BigNumber((_1915_v45).length)))).minus((_1906_v36).f36);
        let _rhs348 = _1908_v38;
        let _rhs349 = _1912_v42;
        let _rhs350 = p0;
        let _lhs269 = globalState;
        let _lhs270 = _1909_v39;
        let _lhs271 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_1909_v39).length));
        let _lhs272 = globalState;
        _lhs269.f15 = _rhs347;
        _1908_v38 = _rhs348;
        _lhs270[_lhs271] = _rhs349;
        _lhs272.f26 = _rhs350;
        let _rhs351 = _module.__default.safeModuloInt((_this).f27, (_this).f27);
        let _rhs352 = _1907_v37;
        let _lhs273 = globalState;
        _lhs273.f24 = _rhs351;
        _1907_v37 = _rhs352;
        _1862_v14 = _module.__default.fm48(!(true), false, (_dafny.ZERO).minus((_1906_v36).f36), globalState);
      }
      return;
    }
    get f38() {
      let _this = this;
      return _this._f38;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f29 = _dafny.MultiSet.Empty;
      this._f30 = _dafny.Seq.of();
      this._f27 = _dafny.ZERO;
      this._f28 = [];
      this._f31 = _dafny.MultiSet.Empty;
      this._f32 = _dafny.ZERO;
      this.f34 = _dafny.ZERO;
      this._f33 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T2, _module.T0];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    get f31() {
      let _this = this;
      return _this._f31;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
    __ctor(f33, f34, f29, f30, f27, f28, f31, f32) {
      let _this = this;
      (_this)._f33 = f33;
      (_this).f34 = f34;
      (_this)._f29 = f29;
      (_this)._f30 = f30;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      (_this)._f31 = f31;
      (_this)._f32 = f32;
      return;
    }
    fm3(p0, globalState) {
      let _this = this;
      return (_this).f33;
    };
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_this).f33, false, (_this).f33, false)).length),(_this).f33)).length), new BigNumber((_dafny.Seq.UnicodeFromString("rio")).length))).length);
    };
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements((_this).f27, new BigNumber(((_dafny.Seq.UnicodeFromString("spkmp"))).length), new BigNumber(4))).IsDisjointFrom(function () {
        let _coll56 = new _dafny.Set();
        for (const _compr_56 of _dafny.IntegerRange(new BigNumber(-947), new BigNumber(-397))) {
          let _1916_v0 = _compr_56;
          if (((new BigNumber(-947)).isLessThanOrEqualTo(_1916_v0)) && ((_1916_v0).isLessThan(new BigNumber(-397)))) {
            _coll56.add((_1916_v0).minus(new BigNumber((_dafny.Seq.of((_this).f33, (_this).f33)).length)));
          }
        }
        return _coll56;
      }()),(_this).f32);
    };
    fm5(globalState) {
      let _this = this;
      return (_this).f33;
    };
    fm6(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-678)), function (_1917_i0) {
        return _dafny.Seq.of((_this).f33, (_this).f33);
      });
    };
    fm7(globalState) {
      let _this = this;
      return new BigNumber(((_module.D1.create_DC1(_dafny.Map.Empty.slice().updateUnsafe(_this.f34,!((_this).f33)))).dtor_cf1).length);
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _1918_v0;
      _1918_v0 = new _dafny.CodePoint('f'.codePointAt(0));
      let _1919_v1;
      _1919_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1918_v0,(_this).f33);
      let _1920_v2;
      _1920_v2 = _dafny.Seq.of(_1919_v1, _1919_v1);
      let _1921_v3;
      _1921_v3 = _module.D2.create_DC3(_1919_v1);
      let _1922_v6;
      _1922_v6 = _dafny.Set.fromElements(_1918_v0);
      let _1923_v7;
      let _nw326 = Array((new BigNumber(22)).toNumber());
      _nw326[(_dafny.ZERO).toNumber()] = _1919_v1;
      _nw326[(_dafny.ONE).toNumber()] = (_1919_v1).update(_1918_v0, !((_this).fm3((_this).f29, globalState)));
      _nw326[(new BigNumber(2)).toNumber()] = _1919_v1;
      _nw326[(new BigNumber(3)).toNumber()] = (_1920_v2)[_module.__default.safeIndex(_this.f34, new BigNumber((_1920_v2).length))];
      _nw326[(new BigNumber(4)).toNumber()] = (_1921_v3).dtor_cf4;
      _nw326[(new BigNumber(5)).toNumber()] = _1919_v1;
      _nw326[(new BigNumber(6)).toNumber()] = _1919_v1;
      _nw326[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1918_v0,(_this).f33);
      _nw326[(new BigNumber(8)).toNumber()] = (_1919_v1).Merge(function () {
        let _coll57 = new _dafny.Map();
        for (const _compr_57 of (_dafny.Map.Empty.slice().updateUnsafe(_1918_v0,(_this).f33)).Keys.Elements) {
          let _1924_v4 = _compr_57;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_1918_v0,(_this).f33)).contains(_1924_v4)) {
            _coll57.push([_1924_v4,(_this).f33]);
          }
        }
        return _coll57;
      }());
      _nw326[(new BigNumber(9)).toNumber()] = _1919_v1;
      _nw326[(new BigNumber(10)).toNumber()] = _1919_v1;
      _nw326[(new BigNumber(11)).toNumber()] = _1919_v1;
      _nw326[(new BigNumber(12)).toNumber()] = (_1921_v3).dtor_cf4;
      _nw326[(new BigNumber(13)).toNumber()] = _1919_v1;
      _nw326[(new BigNumber(14)).toNumber()] = _1919_v1;
      _nw326[(new BigNumber(15)).toNumber()] = (_1919_v1).Merge(_1919_v1);
      _nw326[(new BigNumber(16)).toNumber()] = (_1920_v2)[_module.__default.safeIndex(_this.f34, new BigNumber((_1920_v2).length))];
      _nw326[(new BigNumber(17)).toNumber()] = (((_this).f33) ? (_1919_v1) : (_1919_v1));
      _nw326[(new BigNumber(18)).toNumber()] = (_1919_v1).Merge(_1919_v1);
      _nw326[(new BigNumber(19)).toNumber()] = _1919_v1;
      _nw326[(new BigNumber(20)).toNumber()] = (function () {
        let _coll58 = new _dafny.Map();
        for (const _compr_58 of (_1922_v6).Elements) {
          let _1925_v5 = _compr_58;
          if ((_1922_v6).contains(_1925_v5)) {
            _coll58.push([_1925_v5,(_this).f33]);
          }
        }
        return _coll58;
      }()).Merge(_1919_v1);
      _nw326[(new BigNumber(21)).toNumber()] = _1919_v1;
      _1923_v7 = _nw326;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1923_v7).length))) {
        let _1926_i0 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1926_i0)) && ((_1926_i0).isLessThan(new BigNumber((_1923_v7).length))))) {
          (_1923_v7)[(_1926_i0)] = _dafny.Map.Empty.slice().updateUnsafe(_1918_v0,(_this).f33);
        }
      }
      let _1927_v8;
      _1927_v8 = _dafny.Map.Empty.slice().updateUnsafe(!(true) || (!(!((_this).f33))),((_this).f27).isLessThan(new BigNumber((_dafny.Seq.of(new BigNumber(110), new BigNumber((_module.__default.fm8(globalState)).length))).length)));
      _1927_v8 = (_1927_v8).update((_this).f33, (_this).f33);
      (globalState).f26 = (_this).f33;
      (globalState).f18 = ((_this).f30)[_module.__default.safeIndex(new BigNumber(((_this).f30).length), new BigNumber(((_this).f30).length))];
      _1918_v0 = _1918_v0;
      let _1928_v9;
      _1928_v9 = _dafny.Set.fromElements((_this).f33);
      let _1929_v10;
      _1929_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f33,_1928_v9);
      let _hi9 = new BigNumber((_dafny.Seq.update(_module.__default.fm9((_this).f33, globalState), _module.__default.safeIndex(_this.f34, new BigNumber((_module.__default.fm9((_this).f33, globalState)).length)), _module.__default.fm1(globalState))).length);
      for (let _1930_i1 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(784)), ((_1955_v0) => function (_1956_i2) {
        return _1955_v0;
      })(_1918_v0)), _module.__default.safeIndex(new BigNumber((_1929_v10).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(784)), ((_1957_v0) => function (_1958_i2) {
        return _1957_v0;
      })(_1918_v0))).length)), _1918_v0)).length); _1930_i1.isLessThan(_hi9); _1930_i1 = _1930_i1.plus(_dafny.ONE)) {
        let _1931_v11;
        _1931_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,(_this).f33);
        let _1932_v12;
        _1932_v12 = _dafny.MultiSet.fromElements(_1931_v11);
        let _1933_v13;
        _1933_v13 = _dafny.Seq.UnicodeFromString("yhq");
        let _1934_v14;
        _1934_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1932_v12).cardinality()),_1933_v13);
        if ((_module.__default.safeModuloInt(new BigNumber(((((_1934_v14).contains((_this).f27)) ? ((_1934_v14).get((_this).f27)) : (_1933_v13))).length), (_this).f32)).isEqualTo((_this).f27)) {
          let _1935_v15;
          _1935_v15 = _dafny.Seq.of((((_1931_v11).contains(new BigNumber(((_this).f30).length))) ? ((_1931_v11).get(new BigNumber(((_this).f30).length))) : ((_this).f33)));
          let _1936_v16;
          let _nw327 = Array((new BigNumber(28)).toNumber());
          _nw327[(_dafny.ZERO).toNumber()] = (_this).f33;
          _nw327[(_dafny.ONE).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(2)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(3)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(4)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(5)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(6)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(7)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(8)).toNumber()] = !((_this).f33);
          _nw327[(new BigNumber(9)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(10)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(11)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(12)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(13)).toNumber()] = !((_this).f33);
          _nw327[(new BigNumber(14)).toNumber()] = true;
          _nw327[(new BigNumber(15)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(16)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(17)).toNumber()] = (_1935_v15)[_module.__default.safeIndex((_this).f27, new BigNumber((_1935_v15).length))];
          _nw327[(new BigNumber(18)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(19)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(20)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(21)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(22)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(23)).toNumber()] = !((_this).f33);
          _nw327[(new BigNumber(24)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(25)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(26)).toNumber()] = (_this).f33;
          _nw327[(new BigNumber(27)).toNumber()] = (_this).f33;
          _1936_v16 = _nw327;
          let _1937_v17;
          _1937_v17 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f32).multipliedBy(new BigNumber((_1931_v11).length)),_1936_v16);
          _1937_v17 = (_1937_v17).update((_this).f32, _1936_v16);
          (globalState).f26 = (((_1927_v8).contains(((_this).f27).isLessThanOrEqualTo(new BigNumber((_1928_v9).length)))) ? ((_1927_v8).get(((_this).f27).isLessThanOrEqualTo(new BigNumber((_1928_v9).length)))) : ((_this).f33));
          let _index238 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_1936_v16).length));
          (_1936_v16)[_index238] = !((_this).f33) || ((_this).f33);
          let _1938_v18;
          let _nw328 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _1938_v18 = _nw328;
          let _index239 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_1938_v18).length));
          (_1938_v18)[_index239] = (_this).f27;
          let _index240 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_1936_v16).length));
          let _index241 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_1938_v18).length));
          let _rhs353 = (_this).fm3(_dafny.MultiSet.fromElements((_this).f33), globalState);
          let _rhs354 = new BigNumber(-998);
          let _rhs355 = _1936_v16;
          let _lhs274 = _1936_v16;
          let _lhs275 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_1936_v16).length));
          let _lhs276 = _1938_v18;
          let _lhs277 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_1938_v18).length));
          _lhs274[_lhs275] = _rhs353;
          _lhs276[_lhs277] = _rhs354;
          _1936_v16 = _rhs355;
          let _1939_v19;
          _1939_v19 = _dafny.MultiSet.fromElements((new BigNumber((_1935_v15).length)).minus(_1930_i1), (_dafny.ZERO).minus(_1930_i1), (_this).f27, new BigNumber((_1933_v13).length), ((_1938_v18)[_module.__default.safeIndex(new BigNumber(312), new BigNumber((_1938_v18).length))]).multipliedBy(_1930_i1));
          _1939_v19 = _1939_v19;
          let _1940_v20;
          _1940_v20 = _dafny.Map.Empty.slice().updateUnsafe((_1936_v16)[_module.__default.safeIndex(new BigNumber(588), new BigNumber((_1936_v16).length))],(_this).fm7(globalState));
          let _1941_v22;
          _1941_v22 = _dafny.Set.fromElements(new BigNumber((function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of _dafny.IntegerRange(new BigNumber(215), new BigNumber(925))) {
              let _1942_v21 = _compr_59;
              if (((new BigNumber(215)).isLessThanOrEqualTo(_1942_v21)) && ((_1942_v21).isLessThan(new BigNumber(925)))) {
                _coll59.push([_module.__default.safeDivisionInt(_1942_v21, new BigNumber((_dafny.Set.fromElements(_this.f34)).length)),(_this).f33]);
              }
            }
            return _coll59;
          }()).length));
          let _1943_v23;
          _1943_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1940_v20,new BigNumber(((((_1936_v16)[_module.__default.safeIndex(new BigNumber(588), new BigNumber((_1936_v16).length))]) ? (_1941_v22) : (_module.__default.fm10(globalState)))).length));
          _1943_v23 = (((_1936_v16)[_module.__default.safeIndex(new BigNumber(588), new BigNumber((_1936_v16).length))]) ? (_1943_v23) : (_1943_v23));
        } else {
          r2 = (_this.f34).plus((_1930_i1).minus((_this).f32));
          r0 = (_this).f32;
          let _1944_v24;
          _1944_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1928_v9).length),_1927_v8);
          let _1945_v25;
          _1945_v25 = _dafny.Seq.of(_1927_v8, _1927_v8, (((_1944_v24).contains(_1930_i1)) ? ((_1944_v24).get(_1930_i1)) : (_1927_v8)));
          let _1946_v26;
          _1946_v26 = _dafny.Seq.of((_1945_v25)[_module.__default.safeIndex(_1930_i1, new BigNumber((_1945_v25).length))], _dafny.Map.Empty.slice().updateUnsafe((_this).f33,(_this).f33), (_1927_v8).Merge(_1927_v8), (_dafny.Map.Empty.slice().updateUnsafe((_this).f33,(_this).f33)).Merge(_1927_v8), _1927_v8);
          _1946_v26 = _1946_v26;
          let _1947_v27;
          _1947_v27 = _dafny.Map.Empty.slice().updateUnsafe((((_1919_v1).contains(_1918_v0)) ? ((_1919_v1).get(_1918_v0)) : ((_this).f33)),(_this).f27);
          let _1948_v28;
          let _nw329 = new _module.C1();
          _nw329.__ctor(false, (_this).f29, _dafny.Seq.update(_dafny.Seq.update((_this).f30, _module.__default.safeIndex(_1930_i1, new BigNumber(((_this).f30).length)), new BigNumber((_1947_v27).length)), _module.__default.safeIndex(_this.f34, new BigNumber((_dafny.Seq.update((_this).f30, _module.__default.safeIndex(_1930_i1, new BigNumber(((_this).f30).length)), new BigNumber((_1947_v27).length))).length)), (_this).f27), (_this).f27, (_this).f28);
          _1948_v28 = _nw329;
          let _1949_v29;
          _1949_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1933_v13,_1948_v28);
          _1949_v29 = (_1949_v29).update(_dafny.Seq.UnicodeFromString("nl"), _1948_v28);
          let _1950_v30;
          _1950_v30 = _dafny.Seq.of(!((_this).f33), (_this).f33, (_this).f33, (_this).f33, !(!((_this).f33)));
          let _1951_v31;
          let _nw330 = Array((_dafny.ONE).toNumber());
          _nw330[(_dafny.ZERO).toNumber()] = _1950_v30;
          _1951_v31 = _nw330;
          let _index242 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1951_v31).length));
          (_1951_v31)[_index242] = _1950_v30;
          let _index243 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1951_v31).length));
          (_1951_v31)[_index243] = _dafny.Seq.update(_1950_v30, _module.__default.safeIndex((_this).f27, new BigNumber((_1950_v30).length)), (_this).f33);
        }
        let _1952_v32;
        let _nw331 = Array((new BigNumber(27)).toNumber());
        _1952_v32 = _nw331;
        let _1953_v33;
        _1953_v33 = _dafny.Seq.of((_this).f33);
        let _1954_v34;
        let _nw332 = new _module.C2();
        _nw332.__ctor(_dafny.MultiSet.FromArray(_1953_v33), (_this).f32);
        _1954_v34 = _nw332;
        let _index244 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_1952_v32).length));
        (_1952_v32)[_index244] = (((_this).f33) ? (_1954_v34) : (_1954_v34));
        let _index245 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_1952_v32).length));
        (_1952_v32)[_index245] = _1954_v34;
        let _rhs356 = _1930_i1;
        let _rhs357 = (_this).f33;
        let _lhs278 = globalState;
        let _lhs279 = globalState;
        _lhs278.f15 = _rhs356;
        _lhs279.f26 = _rhs357;
        _1933_v13 = _1933_v13;
      }
      let _1959_v35;
      _1959_v35 = _dafny.Seq.UnicodeFromString("asxhdikym");
      r0 = _module.__default.safeModuloInt(_module.__default.fm1(globalState), (new BigNumber((_1959_v35).length)).multipliedBy(_this.f34));
      r1 = (_this).f27;
      r2 = (_this).f27;
      return [r0, r1, r2];
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let _1960_v0;
      _1960_v0 = _dafny.Seq.of(p1);
      let _1961_v1;
      _1961_v1 = _dafny.Seq.of(_1960_v0);
      let _1962_v2;
      _1962_v2 = _dafny.Seq.UnicodeFromString("mmsbclr");
      let _pat_let_tv43 = _1961_v1;
      let _1963_v3;
      let _1964_v4;
      let _out47;
      let _out48;
      let _outcollector20 = _module.__default.m0((_this).fm5(globalState), ((_this).f27).plus(new BigNumber(((function (_pat_let58_0) {
        return function (_1965_dt__update__tmp_h0) {
          return function (_pat_let59_0) {
            return function (_1966_dt__update_hcf21_h0) {
              return _module.D6.create_DC12((_1965_dt__update__tmp_h0).dtor_cf20, _1966_dt__update_hcf21_h0, (_1965_dt__update__tmp_h0).dtor_cf22);
            }(_pat_let59_0);
          }(_pat_let_tv43);
        }(_pat_let58_0);
      }(_module.D6.create_DC12((_this).f33, _1961_v1, _1962_v2))).dtor_cf21).length)), (_dafny.ZERO).minus(_this.f34), _this.f34, globalState);
      _out47 = _outcollector20[0];
      _out48 = _outcollector20[1];
      _1963_v3 = _out47;
      _1964_v4 = _out48;
      let _1967_i0;
      _1967_i0 = _dafny.ZERO;
      L8: {
        while (((_this).f29).IsProperSubsetOf((_this).f29)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1967_i0)) {
              break L8;
            }
            _1967_i0 = (_1967_i0).plus(_dafny.ONE);
            let _1968_v5;
            _1968_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1963_v3,((_this).f27).isLessThan((_this).f32));
            _1968_v5 = (_1968_v5).Merge(_1968_v5);
            let _1969_v6;
            _1969_v6 = _dafny.Map.Empty.slice().updateUnsafe(_this.f34,!(_1963_v3));
            (globalState).f15 = ((new BigNumber(((_this).fm2((_this).f33, _1963_v3, p1, globalState)).length)).multipliedBy(new BigNumber((_1969_v6).length))).plus(new BigNumber((_1969_v6).length));
            let _1970_v7;
            _1970_v7 = _dafny.Set.fromElements((_this).fm5(globalState));
            (globalState).f26 = ((_1970_v7).Intersect(_1970_v7)).contains(!(new BigNumber((_1960_v0).length)).isEqualTo((_this).f32));
            let _1971_v8;
            let _nw333 = Array((new BigNumber(24)).toNumber());
            _nw333[(_dafny.ZERO).toNumber()] = _1963_v3;
            _nw333[(_dafny.ONE).toNumber()] = (_this).f33;
            _nw333[(new BigNumber(2)).toNumber()] = !(p0);
            _nw333[(new BigNumber(3)).toNumber()] = false;
            _nw333[(new BigNumber(4)).toNumber()] = (_this).f33;
            _nw333[(new BigNumber(5)).toNumber()] = p1;
            _nw333[(new BigNumber(6)).toNumber()] = (_this).f33;
            _nw333[(new BigNumber(7)).toNumber()] = p1;
            _nw333[(new BigNumber(8)).toNumber()] = _1963_v3;
            _nw333[(new BigNumber(9)).toNumber()] = _1963_v3;
            _nw333[(new BigNumber(10)).toNumber()] = p1;
            _nw333[(new BigNumber(11)).toNumber()] = _1963_v3;
            _nw333[(new BigNumber(12)).toNumber()] = !(p0);
            _nw333[(new BigNumber(13)).toNumber()] = _module.__default.fm0((_this).f27, globalState);
            _nw333[(new BigNumber(14)).toNumber()] = (_this).f33;
            _nw333[(new BigNumber(15)).toNumber()] = true;
            _nw333[(new BigNumber(16)).toNumber()] = p1;
            _nw333[(new BigNumber(17)).toNumber()] = p1;
            _nw333[(new BigNumber(18)).toNumber()] = (_1960_v0)[_module.__default.safeIndex(p3, new BigNumber((_1960_v0).length))];
            _nw333[(new BigNumber(19)).toNumber()] = _1963_v3;
            _nw333[(new BigNumber(20)).toNumber()] = p0;
            _nw333[(new BigNumber(21)).toNumber()] = p1;
            _nw333[(new BigNumber(22)).toNumber()] = p1;
            _nw333[(new BigNumber(23)).toNumber()] = (_this).f33;
            _1971_v8 = _nw333;
            let _1972_v9;
            _1972_v9 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_1971_v8);
            _1972_v9 = (_1972_v9).update(p1, _1971_v8);
          }
        }
      }
      _1963_v3 = (_this).f33;
      (globalState).f15 = ((_this).f32).plus(_1964_v4);
      let _hi10 = (_this).f32;
      for (let _1973_i1 = new BigNumber(395); _1973_i1.isLessThan(_hi10); _1973_i1 = _1973_i1.plus(_dafny.ONE)) {
        (globalState).f18 = (_dafny.ZERO).minus(_this.f34);
        let _1974_v10;
        _1974_v10 = new _dafny.CodePoint('t'.codePointAt(0));
        let _1975_v11;
        _1975_v11 = _dafny.Set.fromElements(_1974_v10);
        if ((_1975_v11).contains(_1974_v10)) {
          let _1976_v12;
          _1976_v12 = _dafny.Map.Empty.slice().updateUnsafe(_this.f34,_1963_v3);
          let _1977_v13;
          _1977_v13 = _module.D1.create_DC1(_1976_v12);
          let _1978_v14;
          _1978_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1977_v13,p0);
          _1978_v14 = (_1978_v14).update(_1977_v13, !(_1963_v3));
          (globalState).f0 = _1973_i1;
          let _1979_v15;
          let _1980_v16;
          let _1981_v17;
          let _out49;
          let _out50;
          let _out51;
          let _outcollector21 = (_this).m4(_module.__default.fm0((_this).f27, globalState), _this.f34, p1, p1, globalState);
          _out49 = _outcollector21[0];
          _out50 = _outcollector21[1];
          _out51 = _outcollector21[2];
          _1979_v15 = _out49;
          _1980_v16 = _out50;
          _1981_v17 = _out51;
          _1962_v2 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(370)), function (_1982_i2) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }), ((p1) ? (_1962_v2) : (_1962_v2)));
          let _1983_v18;
          let _nw334 = new _module.C2();
          _nw334.__ctor((_this).f31, (_dafny.ZERO).minus(new BigNumber((_1960_v0).length)));
          _1983_v18 = _nw334;
        } else {
          let _1984_v19;
          let _nw335 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.of());
          _1984_v19 = _nw335;
          let _index246 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_1984_v19).length));
          (_1984_v19)[_index246] = (_this).f30;
          let _1985_v20;
          _1985_v20 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          let _index247 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_1984_v19).length));
          let _rhs358 = _module.__default.fm14(((_this).f32).isLessThanOrEqualTo((_this).f27), (((_1985_v20).contains((_this).f33)) ? ((_1985_v20).get((_this).f33)) : (p1)), p1, (_this).f27, globalState);
          let _rhs359 = p3;
          let _rhs360 = (_this).f32;
          let _rhs361 = p0;
          let _rhs362 = (_this).f30;
          let _lhs280 = globalState;
          let _lhs281 = globalState;
          let _lhs282 = _1984_v19;
          let _lhs283 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_1984_v19).length));
          _1962_v2 = _rhs358;
          _lhs280.f21 = _rhs359;
          _lhs281.f0 = _rhs360;
          _1963_v3 = _rhs361;
          _lhs282[_lhs283] = _rhs362;
          let _1986_v21;
          let _init40 = ((_1987_p2) => function (_1988_i3) {
            return _dafny.Set.fromElements(_1987_p2);
          })(p2);
          let _nw336 = Array((new BigNumber(9)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw336.length); _i0_40++) {
            _nw336[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1986_v21 = _nw336;
          let _1989_v22;
          let _nw337 = Array((new BigNumber(6)).toNumber());
          _1989_v22 = _nw337;
          let _1990_v23;
          _1990_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1986_v21,_1989_v22);
          _1990_v23 = (_1990_v23).update(_1986_v21, _1989_v22);
          let _1991_v24;
          _1991_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).f27, globalState),(_this).f28);
          let _1992_v25;
          let _nw338 = new _module.C1();
          _nw338.__ctor(p1, (_dafny.MultiSet.FromArray(_module.__default.fm13(p2, (_this).f32, p0, new BigNumber(421), globalState))).Union(_dafny.MultiSet.fromElements(p1)), _dafny.Seq.update(_dafny.Seq.update((_1984_v19)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_1984_v19).length))], _module.__default.safeIndex((_this).f27, new BigNumber(((_1984_v19)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_1984_v19).length))]).length)), p3), _module.__default.safeIndex(_module.__default.fm1(globalState), new BigNumber((_dafny.Seq.update((_1984_v19)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_1984_v19).length))], _module.__default.safeIndex((_this).f27, new BigNumber(((_1984_v19)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_1984_v19).length))]).length)), p3)).length)), _module.__default.fm1(globalState)), p3, (((_1991_v24).contains((_this).f33)) ? ((_1991_v24).get((_this).f33)) : ((_this).f28)));
          _1992_v25 = _nw338;
          _1992_v25 = _1992_v25;
          let _1993_v26;
          let _nw339 = Array((new BigNumber(2)).toNumber());
          _nw339[(_dafny.ZERO).toNumber()] = p1;
          _nw339[(_dafny.ONE).toNumber()] = false;
          _1993_v26 = _nw339;
          let _1994_v27;
          _1994_v27 = _dafny.MultiSet.fromElements(_1993_v26);
          let _1995_v28;
          let _1996_v29;
          let _1997_v30;
          let _out52;
          let _out53;
          let _out54;
          let _outcollector22 = (_1992_v25).m5((_1994_v27).IsDisjointFrom(_dafny.MultiSet.fromElements(_1993_v26)), p1, globalState);
          _out52 = _outcollector22[0];
          _out53 = _outcollector22[1];
          _out54 = _outcollector22[2];
          _1995_v28 = _out52;
          _1996_v29 = _out53;
          _1997_v30 = _out54;
          let _1998_v32;
          _1998_v32 = _dafny.Set.fromElements(_1973_i1);
          let _1999_v33;
          _1999_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1973_i1,new BigNumber((_1998_v32).length));
          let _2000_v34;
          _2000_v34 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1973_i1,new BigNumber(((_this).f31).cardinality())), _1999_v33);
          let _2001_v35;
          _2001_v35 = _module.D20.create_DC50(_2000_v34);
          let _2002_v36;
          let _nw340 = new _module.C5();
          _nw340.__ctor(function () {
            let _coll60 = new _dafny.Map();
            for (const _compr_60 of ((_2001_v35).dtor_cf92).Elements) {
              let _2003_v31 = _compr_60;
              if (((_2001_v35).dtor_cf92).contains(_2003_v31)) {
                _coll60.push([_2003_v31,new BigNumber(4)]);
              }
            }
            return _coll60;
          }(), _dafny.MultiSet.FromArray(_dafny.Seq.of(p1, false, _1992_v25.f35)), (_this).f30, _this.f34, (_this).f28);
          _2002_v36 = _nw340;
          let _2004_v37;
          _2004_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2002_v36,p3);
          let _2005_v38;
          _2005_v38 = _module.D10.create_DC22((((_2004_v37).contains(_2002_v36)) ? ((_2004_v37).get(_2002_v36)) : (new BigNumber(596))), (_this).f31);
          let _pat_let_tv44 = _1960_v0;
          let _pat_let_tv45 = _1960_v0;
          _2005_v38 = function (_pat_let60_0) {
            return function (_2006_dt__update__tmp_h1) {
              return function (_pat_let61_0) {
                return function (_2007_dt__update_hcf41_h0) {
                  return _module.D10.create_DC22((_2006_dt__update__tmp_h1).dtor_cf40, _2007_dt__update_hcf41_h0);
                }(_pat_let61_0);
              }(((_this).f31).Intersect(_dafny.MultiSet.fromElements((_pat_let_tv44)[_module.__default.safeIndex(new BigNumber(-773), new BigNumber((_pat_let_tv45).length))])));
            }(_pat_let60_0);
          }(_2005_v38);
        }
        (globalState).f26 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_1960_v0, _1960_v0), _dafny.Seq.Concat(_1960_v0, _dafny.Seq.of(!(p0))));
        let _2008_v39;
        _2008_v39 = _dafny.Set.fromElements((_this).f30, (_this).f30);
        (globalState).f26 = !(((p1) ? (_2008_v39) : (_2008_v39))).contains((_this).f30);
      }
      let _2009_v40;
      _2009_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(550),new BigNumber((_dafny.MultiSet.FromArray((_this).f30)).cardinality()));
      _1963_v3 = !((new BigNumber((_2009_v40).length)).isLessThanOrEqualTo((_this).f32)) || ((_this).f33);
      let _2010_v41;
      let _nw341 = Array((_dafny.ONE).toNumber()).fill(false);
      _2010_v41 = _nw341;
      r0 = (((_this.f34).isLessThan((_this).f32)) ? (_2010_v41) : (_2010_v41));
      r1 = (new BigNumber((_dafny.MultiSet.fromElements((_this).fm3((_this).f31, globalState), _1963_v3, (_this).f33)).cardinality())).plus(new BigNumber(((_this).f31).cardinality()));
      return [r0, r1];
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _2011_v0;
      let _init41 = function (_2012_i0) {
        return (_this).f33;
      };
      let _nw342 = Array((new BigNumber(27)).toNumber());
      for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw342.length); _i0_41++) {
        _nw342[_i0_41] = _init41(new BigNumber(_i0_41));
      }
      _2011_v0 = _nw342;
      _2011_v0 = _2011_v0;
      let _2013_v1;
      _2013_v1 = _dafny.Seq.of(_2011_v0, _2011_v0, _2011_v0);
      let _2014_v2;
      _2014_v2 = _dafny.Map.Empty.slice().updateUnsafe(p2,((p0) ? (p3) : (p0)));
      let _2015_v3;
      _2015_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f33,(_this).f33);
      let _2016_v4;
      _2016_v4 = _dafny.MultiSet.fromElements((_this).f27);
      let _rhs363 = _2013_v1;
      let _rhs364 = (_this).fm3((_this).f31, globalState);
      let _rhs365 = (((_this).f33) ? ((_2014_v2).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("yskuy"), _module.__default.safeIndex((_this).f27, new BigNumber((_dafny.Seq.UnicodeFromString("yskuy")).length)), p1),p3))) : ((_2014_v2).Merge(_2014_v2)));
      let _rhs366 = !(!((_this).f33)) || ((((_2015_v3).contains(p0)) ? ((_2015_v3).get(p0)) : (p3)));
      let _rhs367 = (_2016_v4).IsSubsetOf(_2016_v4);
      let _lhs284 = globalState;
      let _lhs285 = globalState;
      let _lhs286 = globalState;
      _2013_v1 = _rhs363;
      _lhs284.f26 = _rhs364;
      _2014_v2 = _rhs365;
      _lhs285.f26 = _rhs366;
      _lhs286.f26 = _rhs367;
      let _2017_i1;
      _2017_i1 = _dafny.ZERO;
      L9: {
        while (_dafny.areEqual(_dafny.Seq.UnicodeFromString("msled"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(492)), ((_2027_p1) => function (_2028_i2) {
          return _2027_p1;
        })(p1)))) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2017_i1)) {
              break L9;
            }
            _2017_i1 = (_2017_i1).plus(_dafny.ONE);
            (globalState).f0 = new BigNumber((function () {
              let _coll61 = new _dafny.Set();
              for (const _compr_61 of _dafny.IntegerRange(new BigNumber(-80), new BigNumber(763))) {
                let _2018_v5 = _compr_61;
                if (((new BigNumber(-80)).isLessThanOrEqualTo(_2018_v5)) && ((_2018_v5).isLessThan(new BigNumber(763)))) {
                  _coll61.add(_module.__default.safeModuloInt(_2018_v5, (_this).f32));
                }
              }
              return _coll61;
            }()).length);
            let _2019_v6;
            _2019_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Map.Empty.slice().updateUnsafe(_this.f34,(p2)[_module.__default.safeIndex((_this).f32, new BigNumber((p2).length))]));
            let _2020_v7;
            _2020_v7 = _dafny.Map.Empty.slice().updateUnsafe(_this.f34,new _dafny.CodePoint('j'.codePointAt(0)));
            let _2021_v8;
            let _nw343 = new _module.C2();
            _nw343.__ctor((_this).f31, new BigNumber(((((_2019_v6).contains(p1)) ? ((_2019_v6).get(p1)) : (_2020_v7))).length));
            _2021_v8 = _nw343;
            let _index248 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_2011_v0).length));
            (_2011_v0)[_index248] = p3;
            let _2022_v9;
            let _nw344 = new _module.C1();
            _nw344.__ctor(p3, ((_this).f31).update(p3, _module.__default.abs(_module.__default.fm1(globalState))), _dafny.Seq.of(_this.f34), _this.f34, (_this).f28);
            _2022_v9 = _nw344;
            let _2023_v10;
            _2023_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2022_v9,_this.f34);
            let _index249 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_2011_v0).length));
            (_2011_v0)[_index249] = ((_this.f34).isLessThan(new BigNumber((_2023_v10).length))) || (p0);
            let _2024_v11;
            _2024_v11 = _dafny.Set.fromElements(p2, p2);
            let _2025_v12;
            _2025_v12 = _dafny.Seq.of(false);
            let _2026_v13;
            _2026_v13 = _dafny.Map.Empty.slice().updateUnsafe((_2022_v9).fm4(_2024_v11, (_this).f33, _2025_v12, globalState),p0);
            _2026_v13 = (_2026_v13).update((_this).f32, p3);
          }
        }
      }
      let _2029_v14;
      _2029_v14 = _dafny.Set.fromElements(_dafny.Seq.of(_this.f34, _this.f34), (_this).f30, (_this).f30, (_this).f30, (_this).f30);
      (globalState).f26 = (_2029_v14).IsProperSubsetOf(_module.__default.fm50(globalState));
      let _hi11 = _this.f34;
      for (let _2030_i3 = (_dafny.ZERO).minus((_this).f32); _2030_i3.isLessThan(_hi11); _2030_i3 = _2030_i3.plus(_dafny.ONE)) {
        let _2031_v15;
        _2031_v15 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
        let _2032_v16;
        _2032_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_2031_v15).contains(p0)) ? ((_2031_v15).get(p0)) : (p2))).length),_2030_i3);
        let _2033_v18;
        let _nw345 = Array((new BigNumber(14)).toNumber());
        _nw345[(_dafny.ZERO).toNumber()] = _2032_v16;
        _nw345[(_dafny.ONE).toNumber()] = _module.__default.fm15((_this).f27, _2030_i3, globalState);
        _nw345[(new BigNumber(2)).toNumber()] = function () {
          let _coll62 = new _dafny.Map();
          for (const _compr_62 of _dafny.IntegerRange(new BigNumber(-310), new BigNumber(360))) {
            let _2034_v17 = _compr_62;
            if (((new BigNumber(-310)).isLessThanOrEqualTo(_2034_v17)) && ((_2034_v17).isLessThan(new BigNumber(360)))) {
              _coll62.push([_module.__default.safeModuloInt(_2034_v17, (_this).f27),new BigNumber(-274)]);
            }
          }
          return _coll62;
        }();
        _nw345[(new BigNumber(3)).toNumber()] = _2032_v16;
        _nw345[(new BigNumber(4)).toNumber()] = _module.__default.fm30(globalState);
        _nw345[(new BigNumber(5)).toNumber()] = _2032_v16;
        _nw345[(new BigNumber(6)).toNumber()] = _2032_v16;
        _nw345[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2030_i3,_2030_i3);
        _nw345[(new BigNumber(8)).toNumber()] = _2032_v16;
        _nw345[(new BigNumber(9)).toNumber()] = _2032_v16;
        _nw345[(new BigNumber(10)).toNumber()] = _2032_v16;
        _nw345[(new BigNumber(11)).toNumber()] = _2032_v16;
        _nw345[(new BigNumber(12)).toNumber()] = _2032_v16;
        _nw345[(new BigNumber(13)).toNumber()] = _module.__default.fm42((_this).f32, p3, _2030_i3, globalState);
        _2033_v18 = _nw345;
        let _2035_v19;
        _2035_v19 = _dafny.MultiSet.fromElements(_2033_v18, _2033_v18, _2033_v18);
        (globalState).f0 = (((_2035_v19).contains(_2033_v18)) ? ((_2035_v19).get(_2033_v18)) : ((_this.f34).minus((_this).f32)));
        let _2036_v20;
        _2036_v20 = _dafny.Seq.of(_module.D18.create_DC47(new BigNumber((_dafny.Seq.update((_this).f30, _module.__default.safeIndex(_2030_i3, new BigNumber(((_this).f30).length)), _2030_i3)).length), p1));
        (globalState).f3 = new BigNumber((_module.__default.fm51(_dafny.Seq.Concat(_2036_v20, _dafny.Seq.of(_module.D18.create_DC47((_this).f27, p1))), !(_2015_v3).equals(_2015_v3), (_2031_v15).Merge(_2031_v15), globalState)).length);
        (globalState).f0 = ((p0) ? (_module.__default.safeDivisionInt(_2030_i3, _this.f34)) : (new BigNumber(861)));
        let _2037_v21;
        let _nw346 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _2037_v21 = _nw346;
        let _2038_v22;
        _2038_v22 = _dafny.Seq.of((_this).f33, p3, p0);
        let _index250 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_2037_v21).length));
        (_2037_v21)[_index250] = (_2030_i3).plus((_this).fm4(_dafny.Set.fromElements(p2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(566)), ((_2039_p1) => function (_2040_i4) {
          return _2039_p1;
        })(p1))), (_this).f33, _2038_v22, globalState));
        let _2041_v23;
        _2041_v23 = _dafny.Seq.UnicodeFromString("jbd");
        let _index251 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_2037_v21).length));
        let _rhs368 = (_this).f32;
        let _rhs369 = _2033_v18;
        let _rhs370 = _dafny.Seq.Concat(_2041_v23, _2041_v23);
        let _lhs287 = _2037_v21;
        let _lhs288 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_2037_v21).length));
        _lhs287[_lhs288] = _rhs368;
        _2033_v18 = _rhs369;
        _2041_v23 = _rhs370;
      }
      let _2042_v24;
      _2042_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(76),p3);
      let _2043_v26;
      _2043_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,(_this).f27);
      let _2044_v28;
      _2044_v28 = _module.D1.create_DC1(_2042_v24);
      let _2045_v29;
      _2045_v29 = _dafny.Seq.of((_this).f33, p0);
      let _2046_v30;
      _2046_v30 = _dafny.Seq.of(_2042_v24);
      let _2047_v32;
      _2047_v32 = _dafny.Set.fromElements(((_this).f30)[_module.__default.safeIndex((_this).f32, new BigNumber(((_this).f30).length))], (_dafny.ZERO).minus(new BigNumber((_2015_v3).length)), _this.f34, (_this).f32, _this.f34);
      let _2048_v34;
      let _nw347 = Array((new BigNumber(19)).toNumber());
      _nw347[(_dafny.ZERO).toNumber()] = _2042_v24;
      _nw347[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,(_this).f33);
      _nw347[(new BigNumber(2)).toNumber()] = _2042_v24;
      _nw347[(new BigNumber(3)).toNumber()] = (function () {
        let _coll63 = new _dafny.Map();
        for (const _compr_63 of (_2043_v26).Keys.Elements) {
          let _2049_v25 = _compr_63;
          if ((_2043_v26).contains(_2049_v25)) {
            _coll63.push([(_2049_v25).minus(new BigNumber(((_this).f30).length)),(_this).f33]);
          }
        }
        return _coll63;
      }()).Merge(function () {
        let _coll64 = new _dafny.Map();
        for (const _compr_64 of _dafny.IntegerRange(new BigNumber(432), new BigNumber(-747))) {
          let _2050_v27 = _compr_64;
          if (((new BigNumber(432)).isLessThanOrEqualTo(_2050_v27)) && ((_2050_v27).isLessThan(new BigNumber(-747)))) {
            _coll64.push([_module.__default.safeDivisionInt(_2050_v27, (_this).f32),p3]);
          }
        }
        return _coll64;
      }());
      _nw347[(new BigNumber(4)).toNumber()] = _2042_v24;
      _nw347[(new BigNumber(5)).toNumber()] = ((_2044_v28).dtor_cf1).update(new BigNumber((_2043_v26).length), p3);
      _nw347[(new BigNumber(6)).toNumber()] = _2042_v24;
      _nw347[(new BigNumber(7)).toNumber()] = (_2042_v24).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2045_v29).length),p3)).length),p3));
      _nw347[(new BigNumber(8)).toNumber()] = (_2046_v30)[_module.__default.safeIndex(new BigNumber((_2015_v3).length), new BigNumber((_2046_v30).length))];
      _nw347[(new BigNumber(9)).toNumber()] = _2042_v24;
      _nw347[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f34,false);
      _nw347[(new BigNumber(11)).toNumber()] = _module.__default.fm48(p3, p3, _this.f34, globalState);
      _nw347[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f34,!(p3));
      _nw347[(new BigNumber(13)).toNumber()] = function () {
        let _coll65 = new _dafny.Map();
        for (const _compr_65 of (_2047_v32).Elements) {
          let _2051_v31 = _compr_65;
          if ((_2047_v32).contains(_2051_v31)) {
            _coll65.push([_module.__default.safeDivisionInt(_2051_v31, (_module.D13.create_DC32((_this).f32, (_this).f32, (_this).f30, p3)).dtor_cf57),(_this).f33]);
          }
        }
        return _coll65;
      }();
      _nw347[(new BigNumber(14)).toNumber()] = (_2042_v24).Merge(_module.__default.fm48(p0, p3, (_this).f27, globalState));
      _nw347[(new BigNumber(15)).toNumber()] = _2042_v24;
      _nw347[(new BigNumber(16)).toNumber()] = (_2042_v24).Merge(function () {
        let _coll66 = new _dafny.Map();
        for (const _compr_66 of ((_this).f30).Elements) {
          let _2052_v33 = _compr_66;
          if (_dafny.Seq.contains((_this).f30, _2052_v33)) {
            _coll66.push([(_2052_v33).plus(new BigNumber((_2015_v3).length)),(_this).f33]);
          }
        }
        return _coll66;
      }());
      _nw347[(new BigNumber(17)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f32,p3)).Merge(_2042_v24);
      _nw347[(new BigNumber(18)).toNumber()] = _2042_v24;
      _2048_v34 = _nw347;
      _2048_v34 = _2048_v34;
      return;
    }
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = _dafny.ZERO;
      let r2 = [];
      let _2053_v0;
      _2053_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,(_this).f32);
      let _2054_v1;
      _2054_v1 = _dafny.Seq.of((_this).f33, p3);
      let _2055_v2;
      _2055_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2053_v0,new BigNumber((_2054_v1).length));
      let _2056_v3;
      _2056_v3 = _dafny.Seq.of((_this).f30, (_this).f30, (_this).f30);
      let _2057_v4;
      _2057_v4 = _dafny.Seq.UnicodeFromString("pfggacs");
      let _2058_v5;
      _2058_v5 = new _dafny.CodePoint('p'.codePointAt(0));
      let _2059_v6;
      let _nw348 = new _module.C5();
      _nw348.__ctor(_2055_v2, (_this).f29, (_2056_v3)[_module.__default.safeIndex((_this).f32, new BigNumber((_2056_v3).length))], (new BigNumber((_dafny.Seq.update(_2057_v4, _module.__default.safeIndex(p1, new BigNumber((_2057_v4).length)), _2058_v5)).length)).minus((_this).f27), (_this).f28);
      _2059_v6 = _nw348;
      let _2060_v7;
      let _nw349 = new _module.C3();
      _nw349.__ctor(new BigNumber(90), (_this).f28);
      _2060_v7 = _nw349;
      let _2061_v9;
      _2061_v9 = _dafny.Set.fromElements(_this.f34, _this.f34);
      let _2062_v10;
      _2062_v10 = _module.D11.create_DC23(_dafny.MultiSet.fromElements(function () {
  let _coll67 = new _dafny.Set();
  for (const _compr_67 of _dafny.IntegerRange(new BigNumber(-381), new BigNumber(691))) {
    let _2063_v8 = _compr_67;
    if (((new BigNumber(-381)).isLessThanOrEqualTo(_2063_v8)) && ((_2063_v8).isLessThan(new BigNumber(691)))) {
      _coll67.add(_module.__default.safeModuloInt(_2063_v8, p1));
    }
  }
  return _coll67;
}(), _2061_v9, _2061_v9));
      let _source29 = ((p0) ? (_2062_v10) : (_2062_v10));
      if (_source29.is_DC24) {
        let _2064___mcc_h0 = (_source29).cf43;
        let _2065___mcc_h1 = (_source29).cf44;
        let _2066_cf44 = _2065___mcc_h1;
        let _2067_cf43 = _2064___mcc_h0;
        let _2068_v11;
        let _init42 = function (_2069_i0) {
          return (_2069_i0).multipliedBy((_this).f32);
        };
        let _nw350 = Array((new BigNumber(18)).toNumber());
        for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw350.length); _i0_42++) {
          _nw350[_i0_42] = _init42(new BigNumber(_i0_42));
        }
        _2068_v11 = _nw350;
        _2068_v11 = _2068_v11;
        let _2070_v12;
        _2070_v12 = _dafny.Seq.of(_2054_v1, _dafny.Seq.update(_dafny.Seq.of((_this).f33), _module.__default.safeIndex((_this).f27, new BigNumber((_dafny.Seq.of((_this).f33)).length)), p0));
        let _2071_v13;
        _2071_v13 = _module.D6.create_DC12(false, _2070_v12, _module.__default.fm31(p3, p2, globalState));
        _2071_v13 = _2071_v13;
        let _2072_v14;
        let _init43 = ((_2073_v0) => function (_2074_i1) {
          return _2073_v0;
        })(_2053_v0);
        let _nw351 = Array((new BigNumber(11)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw351.length); _i0_43++) {
          _nw351[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _2072_v14 = _nw351;
        let _source30 = _module.D16.create_DC39(_2072_v14);
        if (_source30.is_DC40) {
          let _2075___mcc_h4 = (_source30).cf74;
          let _2076___mcc_h5 = (_source30).cf75;
          let _2077___mcc_h6 = (_source30).cf76;
          let _2078_cf76 = _2077___mcc_h6;
          let _2079_cf75 = _2076___mcc_h5;
          let _2080_cf74 = _2075___mcc_h4;
          let _2081_v15;
          _2081_v15 = _dafny.Seq.of(_2057_v4);
          _2057_v4 = (_2081_v15)[_module.__default.safeIndex(_2079_cf75, new BigNumber((_2081_v15).length))];
          let _2082_v16;
          _2082_v16 = _dafny.Seq.of(_2067_cf43);
          let _2083_v17;
          _2083_v17 = _dafny.Seq.of(_2053_v0);
          let _rhs371 = (_this).f30;
          let _rhs372 = (((_this).f33) ? (p0) : (_module.__default.fm0(_2078_cf76, globalState)));
          let _rhs373 = _2083_v17;
          let _lhs289 = globalState;
          _2082_v16 = _rhs371;
          _lhs289.f26 = _rhs372;
          _2083_v17 = _rhs373;
          _2080_cf74 = !(((_this).f27).isLessThanOrEqualTo((_2079_cf75).multipliedBy(_this.f34)));
          _2068_v11 = _2068_v11;
        } else if (_source30.is_DC41) {
          let _2084___mcc_h7 = (_source30).cf77;
          let _2085_cf77 = _2084___mcc_h7;
          let _2086_v19;
          _2086_v19 = _dafny.MultiSet.fromElements(_2061_v9, function () {
            let _coll68 = new _dafny.Set();
            for (const _compr_68 of _dafny.IntegerRange(new BigNumber(799), new BigNumber(419))) {
              let _2087_v18 = _compr_68;
              if (((new BigNumber(799)).isLessThanOrEqualTo(_2087_v18)) && ((_2087_v18).isLessThan(new BigNumber(419)))) {
                _coll68.add(_module.__default.safeModuloInt(_2087_v18, _2067_cf43));
              }
            }
            return _coll68;
          }());
          let _2088_v20;
          _2088_v20 = _module.D12.create_DC28(((_this).f32).multipliedBy((_dafny.ZERO).minus(new BigNumber((_2086_v19).cardinality()))), ((_this).f30)[_module.__default.safeIndex(new BigNumber(-352), new BigNumber(((_this).f30).length))], (_this).f33);
          _2088_v20 = _2088_v20;
          let _2089_v21;
          let _nw352 = Array((new BigNumber(4)).toNumber()).fill([]);
          _2089_v21 = _nw352;
          let _2090_v22;
          let _nw353 = new _module.C1();
          _nw353.__ctor(p0, _dafny.MultiSet.fromElements(p2), _dafny.Seq.update((_this).f30, _module.__default.safeIndex((_this).f32, new BigNumber(((_this).f30).length)), new BigNumber((_2054_v1).length)), (_dafny.ZERO).minus((_dafny.ZERO).minus(_2067_cf43)), (_this).f28);
          _2090_v22 = _nw353;
          let _2091_v23;
          let _nw354 = Array((new BigNumber(18)).toNumber());
          _nw354[(_dafny.ZERO).toNumber()] = _2090_v22;
          _nw354[(_dafny.ONE).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(2)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(3)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(4)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(5)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(6)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(7)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(8)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(9)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(10)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(11)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(12)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(13)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(14)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(15)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(16)).toNumber()] = _2090_v22;
          _nw354[(new BigNumber(17)).toNumber()] = _2090_v22;
          _2091_v23 = _nw354;
          let _index252 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_2089_v21).length));
          (_2089_v21)[_index252] = _2091_v23;
          let _index253 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_2089_v21).length));
          (_2089_v21)[_index253] = _2091_v23;
          (globalState).f24 = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(32)), ((_2092_v5) => function (_2093_i2) {
            return _2092_v5;
          })(_2058_v5)), _2057_v4)).length)).plus(new BigNumber((_2054_v1).length));
          (globalState).f21 = _module.__default.safeModuloInt((_this).f32, (_this).f27);
        } else if (_source30.is_DC42) {
          let _2094___mcc_h8 = (_source30).cf78;
          let _2095___mcc_h9 = (_source30).cf79;
          let _2096___mcc_h10 = (_source30).cf80;
          let _2097___mcc_h11 = (_source30).cf81;
          let _2098_cf81 = _2097___mcc_h11;
          let _2099_cf80 = _2096___mcc_h10;
          let _2100_cf79 = _2095___mcc_h9;
          let _2101_cf78 = _2094___mcc_h8;
          (globalState).f3 = (_this).f27;
          let _2102_v24;
          let _nw355 = Array((new BigNumber(4)).toNumber());
          _nw355[(_dafny.ZERO).toNumber()] = _2099_cf80;
          _nw355[(_dafny.ONE).toNumber()] = true;
          _nw355[(new BigNumber(2)).toNumber()] = (_this).f33;
          _nw355[(new BigNumber(3)).toNumber()] = (p1).isLessThan((_this).f32);
          _2102_v24 = _nw355;
          let _index254 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_2102_v24).length));
          (_2102_v24)[_index254] = p0;
          let _index255 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_2102_v24).length));
          (_2102_v24)[_index255] = (_2054_v1)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_2057_v4, _2057_v4)).length), new BigNumber((_2054_v1).length))];
          let _2103_v25;
          let _nw356 = new _module.C2();
          _nw356.__ctor((_this).f29, (_dafny.ZERO).minus((new BigNumber(753)).multipliedBy(p1)));
          _2103_v25 = _nw356;
          let _index256 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_2068_v11).length));
          (_2068_v11)[_index256] = _2067_cf43;
          let _2104_v26;
          let _nw357 = new _module.C3();
          _nw357.__ctor((_this).f27, (_this).f28);
          _2104_v26 = _nw357;
          let _2105_v27;
          _2105_v27 = _module.D16.create_DC40((_2054_v1)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("lpyjyi")).length), new BigNumber((_2054_v1).length))], p1, (_this).f27);
          let _2106_v28;
          _2106_v28 = _dafny.Map.Empty.slice().updateUnsafe(_2104_v26,(_2105_v27).dtor_cf76);
          let _index257 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_2068_v11).length));
          let _rhs374 = _2103_v25;
          let _rhs375 = new BigNumber(446);
          let _rhs376 = (((_this).f33) ? (_dafny.Map.Empty.slice().updateUnsafe(_2104_v26,new BigNumber(((_this).f29).cardinality()))) : ((((_this).f33) ? (_2106_v28) : (_2106_v28))));
          let _lhs290 = _2068_v11;
          let _lhs291 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_2068_v11).length));
          _2103_v25 = _rhs374;
          _lhs290[_lhs291] = _rhs375;
          _2106_v28 = _rhs376;
          (globalState).f21 = (_2068_v11)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_2068_v11).length))];
        } else {
          let _2107___mcc_h12 = (_source30).cf73;
          let _2108_cf73 = _2107___mcc_h12;
          let _rhs377 = _module.__default.safeDivisionInt(new BigNumber(556), p1);
          let _rhs378 = (_2054_v1)[_module.__default.safeIndex(new BigNumber(((_this).f30).length), new BigNumber((_2054_v1).length))];
          let _lhs292 = globalState;
          let _lhs293 = globalState;
          _lhs292.f15 = _rhs377;
          _lhs293.f26 = _rhs378;
          let _2109_v29;
          _2109_v29 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_dafny.Set.fromElements(p3)).length));
          let _2110_v30;
          _2110_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_2109_v29);
          let _rhs379 = p2;
          let _rhs380 = ((!(p3)) ? (_2057_v4) : (_2057_v4));
          let _rhs381 = (new BigNumber((_dafny.Seq.Concat(_2057_v4, _dafny.Seq.UnicodeFromString("nq"))).length)).plus((_this).f32);
          let _rhs382 = ((p0) ? (_this.f34) : ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_2057_v4).length), new BigNumber(760)))));
          let _rhs383 = new BigNumber(((_2110_v30).Merge(_2110_v30)).length);
          let _lhs294 = globalState;
          let _lhs295 = globalState;
          let _lhs296 = globalState;
          let _lhs297 = globalState;
          _lhs294.f26 = _rhs379;
          _2057_v4 = _rhs380;
          _lhs295.f18 = _rhs381;
          _lhs296.f21 = _rhs382;
          _lhs297.f21 = _rhs383;
          let _2111_v31;
          let _nw358 = Array((new BigNumber(2)).toNumber());
          _nw358[(_dafny.ZERO).toNumber()] = _2057_v4;
          _nw358[(_dafny.ONE).toNumber()] = ((p3) ? (_2057_v4) : (_dafny.Seq.UnicodeFromString("yggtvlhp")));
          _2111_v31 = _nw358;
          let _rhs384 = (_2057_v4)[_module.__default.safeIndex(new BigNumber(((_this).f31).cardinality()), new BigNumber((_2057_v4).length))];
          let _rhs385 = (_this).f28;
          let _rhs386 = _module.__default.safeDivisionInt(new BigNumber((_2054_v1).length), _2067_cf43);
          let _lhs298 = globalState;
          _2058_v5 = _rhs384;
          _2111_v31 = _rhs385;
          _lhs298.f24 = _rhs386;
          (globalState).f0 = (_this).f27;
        }
        r0 = _2061_v9;
      } else if (_source29.is_DC23) {
        let _2112___mcc_h2 = (_source29).cf42;
        let _2113_cf42 = _2112___mcc_h2;
        (globalState).f26 = (_this).fm3(((_this).f31).Union(_dafny.MultiSet.fromElements(p2)), globalState);
        (globalState).f26 = (_this).f33;
        let _2114_v32;
        _2114_v32 = _module.D9.create_DC19(p0);
        let _2115_v33;
        _2115_v33 = _dafny.MultiSet.fromElements(_2114_v32);
        let _2116_v34;
        let _nw359 = new _module.C1();
        _nw359.__ctor(p2, (_this).f29, _dafny.Seq.update((_this).f30, _module.__default.safeIndex((((_2115_v33).contains(_2114_v32)) ? ((_2115_v33).get(_2114_v32)) : (_this.f34)), new BigNumber(((_this).f30).length)), (_this).f27), new BigNumber((_module.__default.fm25(p2, globalState)).length), (_this).f28);
        _2116_v34 = _nw359;
        let _2117_v35;
        let _nw360 = Array((new BigNumber(6)).toNumber()).fill(false);
        _2117_v35 = _nw360;
        let _index258 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_2117_v35).length));
        (_2117_v35)[_index258] = ((_this).f33) || ((_this).f33);
        let _2118_v36;
        _2118_v36 = _2116_v34.f35;
        let _2119_v37;
        _2119_v37 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2058_v5);
        let _2120_v38;
        _2120_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2119_v37,(_this).f33);
        let _2121_v39;
        _2121_v39 = _dafny.Map.Empty.slice().updateUnsafe(_2118_v36,(((_2120_v38).contains(_2119_v37)) ? ((_2120_v38).get(_2119_v37)) : (p3)));
        let _2122_v40;
        _2122_v40 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2121_v39);
        let _index259 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_2117_v35).length));
        (_2117_v35)[_index259] = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(140)), ((_2123_v39) => function (_2124_i3) {
          return _2123_v39;
        })(_2121_v39)), (((_2122_v40).contains(p2)) ? ((_2122_v40).get(p2)) : (_2121_v39)));
      } else {
        let _2125___mcc_h3 = (_source29).cf45;
        let _2126_cf45 = _2125___mcc_h3;
        let _2127_v41;
        _2127_v41 = _dafny.MultiSet.fromElements(new BigNumber(284));
        let _2128_v42;
        let _nw361 = new _module.C5();
        _nw361.__ctor((_2059_v6).f38, (_this).f31, (_this).f30, (_this).f32, (_this).f28);
        _2128_v42 = _nw361;
        let _2129_v43;
        _2129_v43 = _dafny.Map.Empty.slice().updateUnsafe(_2128_v42,true);
        let _2130_v44;
        let _nw362 = Array((new BigNumber(20)).toNumber());
        _nw362[(_dafny.ZERO).toNumber()] = p3;
        _nw362[(_dafny.ONE).toNumber()] = (_this).f33;
        _nw362[(new BigNumber(2)).toNumber()] = p3;
        _nw362[(new BigNumber(3)).toNumber()] = true;
        _nw362[(new BigNumber(4)).toNumber()] = p3;
        _nw362[(new BigNumber(5)).toNumber()] = p2;
        _nw362[(new BigNumber(6)).toNumber()] = p0;
        _nw362[(new BigNumber(7)).toNumber()] = !(p0);
        _nw362[(new BigNumber(8)).toNumber()] = p2;
        _nw362[(new BigNumber(9)).toNumber()] = ((false) ? ((_this).f33) : (!(p0)));
        _nw362[(new BigNumber(10)).toNumber()] = !(!(p0)) || (p0);
        _nw362[(new BigNumber(11)).toNumber()] = !(!(p0));
        _nw362[(new BigNumber(12)).toNumber()] = p0;
        _nw362[(new BigNumber(13)).toNumber()] = !(p0);
        _nw362[(new BigNumber(14)).toNumber()] = (_2127_v41).equals((_2127_v41).update(p1, _module.__default.abs(new BigNumber((_2129_v43).length))));
        _nw362[(new BigNumber(15)).toNumber()] = p3;
        _nw362[(new BigNumber(16)).toNumber()] = p2;
        _nw362[(new BigNumber(17)).toNumber()] = p0;
        _nw362[(new BigNumber(18)).toNumber()] = true;
        _nw362[(new BigNumber(19)).toNumber()] = p0;
        _2130_v44 = _nw362;
        let _index260 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_2130_v44).length));
        (_2130_v44)[_index260] = _dafny.Seq.IsPrefixOf(_2054_v1, _2054_v1);
        let _index261 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_2130_v44).length));
        (_2130_v44)[_index261] = _module.__default.fm0((_this).f32, globalState);
        let _index262 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_2130_v44).length));
        (_2130_v44)[_index262] = p0;
        let _2131_v45;
        let _nw363 = new _module.C1();
        _nw363.__ctor(p0, ((_this).f31).Intersect((_this).f29), (_this).f30, (_2128_v42).f27, (_this).f28);
        _2131_v45 = _nw363;
        let _2132_v46;
        _2132_v46 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ogu"), _dafny.Seq.update(_2057_v4, _module.__default.safeIndex(p1, new BigNumber((_2057_v4).length)), _2058_v5));
        let _rhs387 = _module.__default.safeDivisionInt(p1, (_this).f27);
        let _rhs388 = _2131_v45;
        let _rhs389 = _2132_v46;
        let _lhs299 = _this;
        _lhs299.f34 = _rhs387;
        _2131_v45 = _rhs388;
        _2132_v46 = _rhs389;
        let _2133_v47;
        let _init44 = function (_2134_i4) {
          return (_2134_i4).plus(_this.f34);
        };
        let _nw364 = Array((new BigNumber(19)).toNumber());
        for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw364.length); _i0_44++) {
          _nw364[_i0_44] = _init44(new BigNumber(_i0_44));
        }
        _2133_v47 = _nw364;
        let _2135_v48;
        _2135_v48 = _dafny.MultiSet.fromElements(_2133_v47, _2133_v47, _2133_v47, _2133_v47);
        _2135_v48 = _2135_v48;
      }
      let _2136_v49;
      _2136_v49 = _module.D4.create_DC8(_2058_v5, new BigNumber((_dafny.Seq.UnicodeFromString("kybibih")).length), p2);
      let _2137_v50;
      _2137_v50 = _module.D11.create_DC24((_dafny.ZERO).minus((_this).f32), (((_this).f33) ? (_2136_v49) : (_2136_v49)));
      let _source31 = _2137_v50;
      if (_source31.is_DC24) {
        let _2138___mcc_h13 = (_source31).cf43;
        let _2139___mcc_h14 = (_source31).cf44;
        let _2140_cf44 = _2139___mcc_h14;
        let _2141_cf43 = _2138___mcc_h13;
        let _2142_v51;
        _2142_v51 = _dafny.Map.Empty.slice().updateUnsafe(p2,!_dafny.areEqual((_this).f30, _dafny.Seq.of(_this.f34, _2141_cf43)));
        _2142_v51 = (_2142_v51).update(p2, (_this).f33);
        if (p3) {
          let _2143_v52;
          let _init45 = function (_2144_i5) {
            return (_2144_i5).minus((_this).f32);
          };
          let _nw365 = Array((new BigNumber(4)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw365.length); _i0_45++) {
            _nw365[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _2143_v52 = _nw365;
          let _2145_v53;
          let _nw366 = new _module.C2();
          _nw366.__ctor((_this).f29, p1);
          _2145_v53 = _nw366;
          let _2146_v54;
          _2146_v54 = _dafny.Seq.of(_2145_v53);
          let _index263 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_2143_v52).length));
          (_2143_v52)[_index263] = (new BigNumber((_2146_v54).length)).multipliedBy((_this).f27);
          let _index264 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_2143_v52).length));
          (_2143_v52)[_index264] = ((_this).f32).multipliedBy(new BigNumber((_dafny.Seq.Concat(_2057_v4, _2057_v4)).length));
          let _2147_v55;
          _2147_v55 = _dafny.Set.fromElements((_this).f33, false);
          let _2148_v56;
          let _nw367 = Array((new BigNumber(10)).toNumber());
          _nw367[(_dafny.ZERO).toNumber()] = true;
          _nw367[(_dafny.ONE).toNumber()] = (_this).f33;
          _nw367[(new BigNumber(2)).toNumber()] = p2;
          _nw367[(new BigNumber(3)).toNumber()] = false;
          _nw367[(new BigNumber(4)).toNumber()] = (_this).f33;
          _nw367[(new BigNumber(5)).toNumber()] = p2;
          _nw367[(new BigNumber(6)).toNumber()] = p0;
          _nw367[(new BigNumber(7)).toNumber()] = p2;
          _nw367[(new BigNumber(8)).toNumber()] = true;
          _nw367[(new BigNumber(9)).toNumber()] = p2;
          _2148_v56 = _nw367;
          let _2149_v57;
          let _nw368 = Array((new BigNumber(8)).toNumber());
          _nw368[(_dafny.ZERO).toNumber()] = _2148_v56;
          _nw368[(_dafny.ONE).toNumber()] = _2148_v56;
          _nw368[(new BigNumber(2)).toNumber()] = _2148_v56;
          _nw368[(new BigNumber(3)).toNumber()] = _2148_v56;
          _nw368[(new BigNumber(4)).toNumber()] = _2148_v56;
          _nw368[(new BigNumber(5)).toNumber()] = _2148_v56;
          _nw368[(new BigNumber(6)).toNumber()] = _2148_v56;
          _nw368[(new BigNumber(7)).toNumber()] = _2148_v56;
          _2149_v57 = _nw368;
          let _2150_v58;
          _2150_v58 = _module.D2.create_DC4((_2145_v53).f32, new BigNumber(982), _this.f34, _2058_v5, _2057_v4);
          let _2151_v59;
          _2151_v59 = _module.D16.create_DC42(p1, _2149_v57, p3, _module.D2.create_DC5(_2150_v58));
          let _2152_v60;
          _2152_v60 = _dafny.MultiSet.fromElements((new BigNumber(399)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2147_v55,p3)).length)), _this.f34, _this.f34, (_2151_v59).dtor_cf78, (_2145_v53).f32);
          let _2153_v61;
          _2153_v61 = _dafny.Seq.of(_2061_v9);
          let _index265 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_2143_v52).length));
          (_2143_v52)[_index265] = (((_2152_v60).contains((p1).plus((_2143_v52)[_module.__default.safeIndex(new BigNumber(820), new BigNumber((_2143_v52).length))]))) ? ((_2152_v60).get((p1).plus((_2143_v52)[_module.__default.safeIndex(new BigNumber(820), new BigNumber((_2143_v52).length))]))) : (new BigNumber((_dafny.Seq.update(_2153_v61, _module.__default.safeIndex(p1, new BigNumber((_2153_v61).length)), _2061_v9)).length)));
          let _index266 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_2143_v52).length));
          let _rhs390 = p0;
          let _rhs391 = (_2145_v53).f32;
          let _lhs300 = globalState;
          let _lhs301 = _2143_v52;
          let _lhs302 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_2143_v52).length));
          _lhs300.f26 = _rhs390;
          _lhs301[_lhs302] = _rhs391;
          (globalState).f26 = true;
          let _2154_v62;
          _2154_v62 = _module.D12.create_DC27(p0, _2152_v60, (((_2152_v60).contains(new BigNumber((_2061_v9).length))) ? ((_2152_v60).get(new BigNumber((_2061_v9).length))) : (p1)), _2143_v52);
          let _2155_v63;
          _2155_v63 = _dafny.Seq.of(_2154_v62);
          (globalState).f24 = (new BigNumber((_dafny.Seq.Concat(_2155_v63, _dafny.Seq.of(_2154_v62))).length)).minus(new BigNumber(((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_2057_v4, _module.__default.safeIndex(new BigNumber(830), new BigNumber((_2057_v4).length)), new _dafny.CodePoint('k'.codePointAt(0)))).length))).Intersect(_2061_v9)).length));
        } else {
          (globalState).f0 = _module.__default.safeModuloInt(_2141_cf43, (_this).f32);
          let _2156_v64;
          _2156_v64 = _dafny.Seq.of((_this).f27);
          let _2157_v65;
          _2157_v65 = _module.D9.create_DC18(p0);
          let _rhs392 = _module.__default.fm9((_this).f33, globalState);
          let _rhs393 = !(p3);
          let _rhs394 = _module.__default.fm52((_this).f33, _module.__default.fm0((_this).f32, globalState), false, new BigNumber(914), globalState);
          let _rhs395 = _dafny.Seq.Concat(_2054_v1, _2054_v1);
          let _rhs396 = new BigNumber(844);
          let _lhs303 = globalState;
          let _lhs304 = globalState;
          _2156_v64 = _rhs392;
          _lhs303.f26 = _rhs393;
          _2157_v65 = _rhs394;
          _2054_v1 = _rhs395;
          _lhs304.f21 = _rhs396;
          (globalState).f18 = _this.f34;
          (globalState).f3 = _this.f34;
          let _2158_v66;
          let _init46 = ((_2159_cf43, _2160_p1, _2161_v0) => function (_2162_i6) {
            return !(_dafny.Map.Empty.slice().updateUnsafe(_2159_cf43,_dafny.Map.Empty.slice().updateUnsafe(_this.f34,_2160_p1))).equals(_dafny.Map.Empty.slice().updateUnsafe(_2159_cf43,_2161_v0));
          })(_2141_cf43, p1, _2053_v0);
          let _nw369 = Array((new BigNumber(4)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw369.length); _i0_46++) {
            _nw369[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _2158_v66 = _nw369;
          let _index267 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_2158_v66).length));
          (_2158_v66)[_index267] = p3;
          let _2163_v67;
          _2163_v67 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2054_v1).length),_2058_v5);
          let _2164_v68;
          let _nw370 = Array((new BigNumber(5)).toNumber());
          _nw370[(_dafny.ZERO).toNumber()] = (new BigNumber((_2163_v67).length)).multipliedBy((_2060_v7).fm45(globalState));
          _nw370[(_dafny.ONE).toNumber()] = new BigNumber(638);
          _nw370[(new BigNumber(2)).toNumber()] = (p1).plus(_2141_cf43);
          _nw370[(new BigNumber(3)).toNumber()] = (_2141_cf43).multipliedBy(p1);
          _nw370[(new BigNumber(4)).toNumber()] = _2141_cf43;
          _2164_v68 = _nw370;
          let _index268 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_2164_v68).length));
          (_2164_v68)[_index268] = (_this).f32;
          let _2165_v69;
          _2165_v69 = _module.D2.create_DC4((_this).fm7(globalState), (_this).f32, _2141_cf43, _2058_v5, _module.__default.fm37(new BigNumber((_2061_v9).length), _this.f34, globalState));
          let _2166_v70;
          _2166_v70 = _dafny.Map.Empty.slice().updateUnsafe(_2165_v69,p3);
          let _2167_v71;
          _2167_v71 = _dafny.Seq.of(_2166_v70);
          let _index269 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_2158_v66).length));
          let _index270 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_2164_v68).length));
          let _rhs397 = ((new BigNumber(((_module.D8.create_DC15(_2142_v51)).dtor_cf28).length)).plus(p1)).isLessThanOrEqualTo(new BigNumber((_2142_v51).length));
          let _rhs398 = !(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_2167_v71, _2167_v71))).contains(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm36(p3, globalState),p2));
          let _rhs399 = (_this).fm7(globalState);
          let _rhs400 = _module.__default.safeDivisionInt(new BigNumber(870), _2141_cf43);
          let _lhs305 = globalState;
          let _lhs306 = _2158_v66;
          let _lhs307 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_2158_v66).length));
          let _lhs308 = _2164_v68;
          let _lhs309 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_2164_v68).length));
          let _lhs310 = globalState;
          _lhs305.f26 = _rhs397;
          _lhs306[_lhs307] = _rhs398;
          _lhs308[_lhs309] = _rhs399;
          _lhs310.f3 = _rhs400;
        }
        let _2168_v72;
        let _init47 = function (_2169_i7) {
          return (_this).f30;
        };
        let _nw371 = Array((new BigNumber(22)).toNumber());
        for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw371.length); _i0_47++) {
          _nw371[_i0_47] = _init47(new BigNumber(_i0_47));
        }
        _2168_v72 = _nw371;
        let _index271 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_2168_v72).length));
        (_2168_v72)[_index271] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-680)), function (_2170_i8) {
          return (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f32));
        });
        let _2171_v73;
        _2171_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-289),true);
        let _2172_v74;
        let _nw372 = Array((new BigNumber(22)).toNumber());
        _nw372[(_dafny.ZERO).toNumber()] = (_this).f33;
        _nw372[(_dafny.ONE).toNumber()] = p0;
        _nw372[(new BigNumber(2)).toNumber()] = !(!(p0)) || (p0);
        _nw372[(new BigNumber(3)).toNumber()] = (p3) === (p0);
        _nw372[(new BigNumber(4)).toNumber()] = (_this).f33;
        _nw372[(new BigNumber(5)).toNumber()] = p3;
        _nw372[(new BigNumber(6)).toNumber()] = (p0) && (p0);
        _nw372[(new BigNumber(7)).toNumber()] = p3;
        _nw372[(new BigNumber(8)).toNumber()] = p3;
        _nw372[(new BigNumber(9)).toNumber()] = _dafny.areEqual(_dafny.Seq.of(p1, (_this).f27), _module.__default.fm16((_this).fm7(globalState), globalState));
        _nw372[(new BigNumber(10)).toNumber()] = (_this).f33;
        _nw372[(new BigNumber(11)).toNumber()] = (new BigNumber((_2171_v73).length)).isLessThan(_this.f34);
        _nw372[(new BigNumber(12)).toNumber()] = ((_this).f31).IsSubsetOf(_module.__default.fm32(globalState));
        _nw372[(new BigNumber(13)).toNumber()] = !(p2);
        _nw372[(new BigNumber(14)).toNumber()] = (_2061_v9).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber(525)));
        _nw372[(new BigNumber(15)).toNumber()] = p3;
        _nw372[(new BigNumber(16)).toNumber()] = _module.__default.fm0((_this).fm7(globalState), globalState);
        _nw372[(new BigNumber(17)).toNumber()] = p3;
        _nw372[(new BigNumber(18)).toNumber()] = (_this).fm5(globalState);
        _nw372[(new BigNumber(19)).toNumber()] = (p1).isLessThan((_this).f32);
        _nw372[(new BigNumber(20)).toNumber()] = !(p3);
        _nw372[(new BigNumber(21)).toNumber()] = p0;
        _2172_v74 = _nw372;
        let _index272 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_2172_v74).length));
        (_2172_v74)[_index272] = false;
        let _index273 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_2168_v72).length));
        let _index274 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_2172_v74).length));
        let _rhs401 = !(p0) || (p3);
        let _rhs402 = p1;
        let _rhs403 = _dafny.Seq.of(_this.f34);
        let _rhs404 = (_this).f33;
        let _rhs405 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2054_v1, _dafny.Seq.of(p2)), _2054_v1);
        let _lhs311 = globalState;
        let _lhs312 = globalState;
        let _lhs313 = _2168_v72;
        let _lhs314 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_2168_v72).length));
        let _lhs315 = _2172_v74;
        let _lhs316 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_2172_v74).length));
        _lhs311.f26 = _rhs401;
        _lhs312.f24 = _rhs402;
        _lhs313[_lhs314] = _rhs403;
        _lhs315[_lhs316] = _rhs404;
        _2054_v1 = _rhs405;
        (globalState).f3 = _2141_cf43;
      } else if (_source31.is_DC23) {
        let _2173___mcc_h15 = (_source31).cf42;
        let _2174_cf42 = _2173___mcc_h15;
        let _2175_v75;
        _2175_v75 = _module.D15.create_DC36(_2058_v5, (_this).f32);
        let _source32 = _2175_v75;
        if (_source32.is_DC36) {
          let _2176___mcc_h17 = (_source32).cf64;
          let _2177___mcc_h18 = (_source32).cf65;
          let _2178_cf65 = _2177___mcc_h18;
          let _2179_cf64 = _2176___mcc_h17;
          let _2180_v76;
          _2180_v76 = _module.D16.create_DC40(_dafny.Seq.IsProperPrefixOf(_2057_v4, _dafny.Seq.update(_2057_v4, _module.__default.safeIndex(p1, new BigNumber((_2057_v4).length)), new _dafny.CodePoint('g'.codePointAt(0)))), _this.f34, (_dafny.ZERO).minus(((!(p0)) ? ((_dafny.ZERO).minus((_this).f27)) : ((_this).f27))));
          let _2181_v77;
          let _nw373 = Array((new BigNumber(21)).toNumber()).fill(false);
          _2181_v77 = _nw373;
          let _2182_v78;
          _2182_v78 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,_2054_v1));
          let _2183_v79;
          _2183_v79 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2061_v9);
          let _2184_v80;
          _2184_v80 = _dafny.Map.Empty.slice().updateUnsafe((_this).f33,p0);
          let _pat_let_tv46 = _2182_v78;
          let _pat_let_tv47 = _2182_v78;
          let _rhs406 = function (_pat_let62_0) {
            return function (_2185_dt__update__tmp_h0) {
              return function (_pat_let63_0) {
                return function (_2186_dt__update_hcf75_h0) {
                  return _module.D16.create_DC40((_2185_dt__update__tmp_h0).dtor_cf74, _2186_dt__update_hcf75_h0, (_2185_dt__update__tmp_h0).dtor_cf76);
                }(_pat_let63_0);
              }(new BigNumber(((_pat_let_tv46)[_module.__default.safeIndex((_this).f27, new BigNumber((_pat_let_tv47).length))]).length));
            }(_pat_let62_0);
          }(_module.D16.create_DC40(p3, _2178_cf65, p1));
          let _rhs407 = new BigNumber(((((_2183_v79).contains((new BigNumber((_2184_v80).length)).multipliedBy(_this.f34))) ? ((_2183_v79).get((new BigNumber((_2184_v80).length)).multipliedBy(_this.f34))) : (_2061_v9))).length);
          let _rhs408 = (((_this).f33) ? (_2181_v77) : (_2181_v77));
          let _lhs317 = globalState;
          _2180_v76 = _rhs406;
          _lhs317.f15 = _rhs407;
          _2181_v77 = _rhs408;
          (globalState).f0 = new BigNumber((_dafny.Seq.update(_2057_v4, _module.__default.safeIndex(new BigNumber((_2057_v4).length), new BigNumber((_2057_v4).length)), _2058_v5)).length);
          let _2187_v81;
          let _nw374 = new _module.C4();
          _nw374.__ctor(p1, (_this).f28);
          _2187_v81 = _nw374;
          let _2188_v82;
          _2188_v82 = _2057_v4;
          let _2189_v83;
          _2189_v83 = _dafny.MultiSet.fromElements(_2057_v4);
          let _2190_v84;
          _2190_v84 = _module.D7.create_DC14((_this).f27, _2187_v81, _2188_v82, _2189_v83);
          let _2191_v85;
          _2191_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_2190_v84);
          _2191_v85 = (_2191_v85).update(new BigNumber(165), _2190_v84);
          let _2192_v86;
          _2192_v86 = _dafny.Map.Empty.slice().updateUnsafe(_2058_v5,p1);
          _2192_v86 = (_2192_v86).update(new _dafny.CodePoint('n'.codePointAt(0)), (_this).f27);
        } else if (_source32.is_DC37) {
          let _2193___mcc_h19 = (_source32).cf66;
          let _2194___mcc_h20 = (_source32).cf67;
          let _2195___mcc_h21 = (_source32).cf68;
          let _2196___mcc_h22 = (_source32).cf69;
          let _2197_cf69 = _2196___mcc_h22;
          let _2198_cf68 = _2195___mcc_h21;
          let _2199_cf67 = _2194___mcc_h20;
          let _2200_cf66 = _2193___mcc_h19;
          let _2201_v87;
          let _nw375 = Array((new BigNumber(7)).toNumber()).fill(false);
          _2201_v87 = _nw375;
          let _index275 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_2201_v87).length));
          (_2201_v87)[_index275] = _2197_cf69;
          let _index276 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_2201_v87).length));
          (_2201_v87)[_index276] = !(new BigNumber(-729)).isEqualTo(((_this).f30)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(654)), ((_2202_v5) => function (_2203_i9) {
            return _2202_v5;
          })(_2058_v5))).length), new BigNumber(((_this).f30).length))]);
          let _2204_v88;
          _2204_v88 = _dafny.Set.fromElements(_2200_cf66, (_2201_v87)[_module.__default.safeIndex(new BigNumber(143), new BigNumber((_2201_v87).length))]);
          let _2205_v89;
          _2205_v89 = _dafny.MultiSet.fromElements(((_2197_cf69) ? (_2054_v1) : (_2054_v1)), _2054_v1, _dafny.Seq.Concat(_module.__default.fm41(_2199_cf67, new BigNumber((_2204_v88).length), p0, globalState), _2054_v1));
          _2205_v89 = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_module.__default.fm44(globalState), _dafny.Seq.of(_2054_v1)));
          let _2206_v90;
          let _nw376 = Array((new BigNumber(9)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2206_v90 = _nw376;
          let _index277 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_2206_v90).length));
          (_2206_v90)[_index277] = _2058_v5;
          let _index278 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_2206_v90).length));
          (_2206_v90)[_index278] = _2058_v5;
          let _2207_v91;
          _2207_v91 = _dafny.Seq.of((_this).f27);
          _2207_v91 = (_this).f30;
        } else if (_source32.is_DC38) {
          let _2208___mcc_h23 = (_source32).cf70;
          let _2209___mcc_h24 = (_source32).cf71;
          let _2210___mcc_h25 = (_source32).cf72;
          let _2211_cf72 = _2210___mcc_h25;
          let _2212_cf71 = _2209___mcc_h24;
          let _2213_cf70 = _2208___mcc_h23;
          let _2214_v92;
          _2214_v92 = _dafny.Map.Empty.slice().updateUnsafe(true,(_2061_v9).Intersect(_2061_v9));
          (globalState).f3 = new BigNumber(((((_2214_v92).contains(_dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(466)), ((_2217_v5) => function (_2218_i10) {
            return _2217_v5;
          })(_2058_v5)), _2057_v4))) ? ((_2214_v92).get(_dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(466)), ((_2215_v5) => function (_2216_i10) {
            return _2215_v5;
          })(_2058_v5)), _2057_v4))) : (_2061_v9))).length);
          let _2219_v93;
          let _nw377 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
          _2219_v93 = _nw377;
          _2219_v93 = _2219_v93;
          (globalState).f15 = p1;
          let _2220_v94;
          let _nw378 = Array((new BigNumber(3)).toNumber());
          _nw378[(_dafny.ZERO).toNumber()] = _2211_cf72;
          _nw378[(_dafny.ONE).toNumber()] = _2211_cf72;
          _nw378[(new BigNumber(2)).toNumber()] = (_this).f32;
          _2220_v94 = _nw378;
          let _2221_v95;
          _2221_v95 = _dafny.Map.Empty.slice().updateUnsafe(_2220_v94,(_this).f29);
          let _2222_v96;
          let _nw379 = new _module.C5();
          _nw379.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_2053_v0,new BigNumber(((_this).f30).length)), (_this).f29, _dafny.Seq.update((_this).f30, _module.__default.safeIndex(new BigNumber((((((_2221_v95).contains(_2220_v94)) ? ((_2221_v95).get(_2220_v94)) : ((_this).f29))).update(false, _module.__default.abs(new BigNumber((_dafny.Seq.update((_this).f30, _module.__default.safeIndex(_this.f34, new BigNumber(((_this).f30).length)), _2211_cf72)).length)))).cardinality()), new BigNumber(((_this).f30).length)), _this.f34), ((p2) ? (p1) : (((((_this).f29).contains(p2)) ? (((_this).f29).get(p2)) : (p1)))), (_this).f28);
          _2222_v96 = _nw379;
        } else {
          let _2223___mcc_h26 = (_source32).cf63;
          let _2224_cf63 = _2223___mcc_h26;
          let _2225_v97;
          _2225_v97 = _module.D8.create_DC16(_this.f34, (_this).f27);
          let _2226_v98;
          _2226_v98 = _dafny.Map.Empty.slice().updateUnsafe(_2225_v97,new BigNumber((_2053_v0).length));
          (globalState).f0 = new BigNumber((_2226_v98).length);
          r0 = ((_2061_v9).Union(_2061_v9)).Union(_2061_v9);
          let _2227_v99;
          let _nw380 = new _module.C2();
          _nw380.__ctor(_dafny.MultiSet.fromElements(p0), _this.f34);
          _2227_v99 = _nw380;
          let _2228_v100;
          _2228_v100 = _dafny.Map.Empty.slice().updateUnsafe(_2054_v1,false);
          _2228_v100 = (((p3) ? (_dafny.Map.Empty.slice().updateUnsafe(_2054_v1,p0)) : (_2228_v100))).Merge(_module.__default.fm53(_2058_v5, globalState));
        }
        let _2229_v101;
        let _nw381 = Array((new BigNumber(4)).toNumber()).fill(_dafny.MultiSet.Empty);
        _2229_v101 = _nw381;
        let _index279 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_2229_v101).length));
        (_2229_v101)[_index279] = (_this).f31;
        let _index280 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_2229_v101).length));
        (_2229_v101)[_index280] = ((_this).f29).Union((_this).f29);
        (globalState).f26 = (_this).f33;
        let _2230_v102;
        _2230_v102 = _dafny.MultiSet.fromElements((_this).f32);
        let _2231_v103;
        let _nw382 = new _module.C2();
        _nw382.__ctor(_dafny.MultiSet.fromElements(p0, (_this).f33, (_this).f33, true, p3), new BigNumber((_2230_v102).cardinality()));
        _2231_v103 = _nw382;
      } else {
        let _2232___mcc_h16 = (_source31).cf45;
        let _2233_cf45 = _2232___mcc_h16;
        let _2234_v104;
        let _nw383 = new _module.C4();
        _nw383.__ctor(_this.f34, (_this).f28);
        _2234_v104 = _nw383;
        let _2235_v106;
        _2235_v106 = _dafny.MultiSet.fromElements((_this).f27, p1, p1);
        let _2236_v107;
        let _nw384 = Array((new BigNumber(27)).toNumber());
        _nw384[(_dafny.ZERO).toNumber()] = !((_this).f33);
        _nw384[(_dafny.ONE).toNumber()] = p0;
        _nw384[(new BigNumber(2)).toNumber()] = p0;
        _nw384[(new BigNumber(3)).toNumber()] = p0;
        _nw384[(new BigNumber(4)).toNumber()] = p0;
        _nw384[(new BigNumber(5)).toNumber()] = !_dafny.areEqual(_2054_v1, _2054_v1);
        _nw384[(new BigNumber(6)).toNumber()] = p2;
        _nw384[(new BigNumber(7)).toNumber()] = (new BigNumber((_2057_v4).length)).isLessThan((_dafny.ZERO).minus((_this).f27));
        _nw384[(new BigNumber(8)).toNumber()] = !(false) || (p3);
        _nw384[(new BigNumber(9)).toNumber()] = (_this.f34).isLessThanOrEqualTo(p1);
        _nw384[(new BigNumber(10)).toNumber()] = !((_this).f33) || (p2);
        _nw384[(new BigNumber(11)).toNumber()] = !(p0) || (p3);
        _nw384[(new BigNumber(12)).toNumber()] = (function () {
          let _coll69 = new _dafny.Map();
          for (const _compr_69 of _dafny.IntegerRange(new BigNumber(371), new BigNumber(979))) {
            let _2237_v105 = _compr_69;
            if (((new BigNumber(371)).isLessThanOrEqualTo(_2237_v105)) && ((_2237_v105).isLessThan(new BigNumber(979)))) {
              _coll69.push([(_2237_v105).minus(new BigNumber(566)),(_this).f32]);
            }
          }
          return _coll69;
        }()).contains((_this).f32);
        _nw384[(new BigNumber(13)).toNumber()] = (_this).f33;
        _nw384[(new BigNumber(14)).toNumber()] = (_2235_v106).IsSubsetOf(_2235_v106);
        _nw384[(new BigNumber(15)).toNumber()] = p2;
        _nw384[(new BigNumber(16)).toNumber()] = p3;
        _nw384[(new BigNumber(17)).toNumber()] = !((_2059_v6).fm3((_this).f31, globalState));
        _nw384[(new BigNumber(18)).toNumber()] = (_2054_v1)[_module.__default.safeIndex((_this).f32, new BigNumber((_2054_v1).length))];
        _nw384[(new BigNumber(19)).toNumber()] = p0;
        _nw384[(new BigNumber(20)).toNumber()] = (_this).f33;
        _nw384[(new BigNumber(21)).toNumber()] = p2;
        _nw384[(new BigNumber(22)).toNumber()] = p0;
        _nw384[(new BigNumber(23)).toNumber()] = p3;
        _nw384[(new BigNumber(24)).toNumber()] = p0;
        _nw384[(new BigNumber(25)).toNumber()] = p2;
        _nw384[(new BigNumber(26)).toNumber()] = p3;
        _2236_v107 = _nw384;
        let _index281 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_2236_v107).length));
        (_2236_v107)[_index281] = (_this).f33;
        let _2238_v108;
        _2238_v108 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f27).isLessThanOrEqualTo(p1),!((_2234_v104).fm35(_2058_v5, (_this).f27, globalState)));
        let _2239_v109;
        _2239_v109 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,_2057_v4);
        let _index282 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_2236_v107).length));
        let _rhs409 = !_dafny.Seq.contains(_dafny.Seq.Concat((((_2239_v109).contains(_this.f34)) ? ((_2239_v109).get(_this.f34)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(228)), ((_2240_v5) => function (_2241_i11) {
          return _2240_v5;
        })(_2058_v5)))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(7)), ((_2242_v5) => function (_2243_i12) {
          return _2242_v5;
        })(_2058_v5))), _2058_v5);
        let _rhs410 = _2238_v108;
        let _lhs318 = _2236_v107;
        let _lhs319 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_2236_v107).length));
        _lhs318[_lhs319] = _rhs409;
        _2238_v108 = _rhs410;
        let _2244_v110;
        let _nw385 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _2244_v110 = _nw385;
        let _index283 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_2244_v110).length));
        (_2244_v110)[_index283] = (_this).f32;
        let _index284 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_2244_v110).length));
        (_2244_v110)[_index284] = _module.__default.fm1(globalState);
        _2057_v4 = _module.__default.fm37((((_2053_v0).contains(new BigNumber(333))) ? ((_2053_v0).get(new BigNumber(333))) : ((_this).f32)), new BigNumber(474), globalState);
      }
      let _hi12 = (_this).f27;
      for (let _2245_i13 = (new BigNumber((_dafny.Seq.UnicodeFromString("e")).length)).plus(_this.f34); _2245_i13.isLessThan(_hi12); _2245_i13 = _2245_i13.plus(_dafny.ONE)) {
        r1 = (_2245_i13).multipliedBy(_2245_i13);
        let _2246_v111;
        let _nw386 = new _module.C3();
        _nw386.__ctor((_this).f32, (_this).f28);
        _2246_v111 = _nw386;
        (globalState).f26 = !(p3) || ((_this).f33);
        let _2247_v112;
        let _init48 = ((_2248_p0) => function (_2249_i14) {
          return _dafny.Map.Empty.slice().updateUnsafe(_2248_p0,true);
        })(p0);
        let _nw387 = Array((new BigNumber(13)).toNumber());
        for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw387.length); _i0_48++) {
          _nw387[_i0_48] = _init48(new BigNumber(_i0_48));
        }
        _2247_v112 = _nw387;
        let _2250_v113;
        _2250_v113 = _dafny.Set.fromElements(_2057_v4, _module.__default.fm8(globalState));
        let _rhs411 = _module.__default.safeDivisionInt((new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(831), (_this).f27, (_this).f27, _2245_i13))).cardinality())).minus(new BigNumber((_2057_v4).length)), p1);
        let _rhs412 = _module.__default.fm15((_this).fm4(_2250_v113, !(p3), _2054_v1, globalState), _2245_i13, globalState);
        let _rhs413 = _2247_v112;
        let _rhs414 = p1;
        let _lhs320 = globalState;
        let _lhs321 = globalState;
        _lhs320.f18 = _rhs411;
        _2053_v0 = _rhs412;
        _2247_v112 = _rhs413;
        _lhs321.f15 = _rhs414;
      }
      let _2251_v114;
      _2251_v114 = _dafny.Set.fromElements((_this).f31);
      (globalState).f21 = ((_this.f34).plus(new BigNumber((_dafny.Seq.UnicodeFromString("hgatgwca")).length))).multipliedBy(new BigNumber((_2251_v114).length));
      let _2252_v115;
      _2252_v115 = _dafny.MultiSet.fromElements(_this.f34, p1);
      let _2253_v116;
      _2253_v116 = _module.D4.create_DC9(_this.f34, p3);
      let _2254_v117;
      _2254_v117 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2252_v115).cardinality()),_2253_v116);
      r0 = function () {
        let _coll70 = new _dafny.Set();
        for (const _compr_70 of (_2254_v117).Keys.Elements) {
          let _2255_v118 = _compr_70;
          if ((_2254_v117).contains(_2255_v118)) {
            _coll70.add((_2255_v118).multipliedBy(_module.__default.safeDivisionInt(new BigNumber(-421), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(280))).cardinality()))).length))));
          }
        }
        return _coll70;
      }();
      r1 = new BigNumber(245);
      let _nw388 = Array((new BigNumber(15)).toNumber()).fill([]);
      r2 = _nw388;
      return [r0, r1, r2];
    }
    get f33() {
      let _this = this;
      return _this._f33;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
