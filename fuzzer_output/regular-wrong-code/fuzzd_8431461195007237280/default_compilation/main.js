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
    static fm0(globalState) {
      return (_dafny.ZERO).minus((_dafny.ZERO).minus((_module.__default.safeModuloInt(new BigNumber(4), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),false)).length))).plus(new BigNumber(592))));
    };
    static fm1(p0, p1, p2, globalState) {
      return !(!((((!(true)) ? (_dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)))) : (_dafny.Set.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)))))).IsProperSubsetOf(_dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0))))));
    };
    static fm2(globalState) {
      if (!(!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("anivmhual"), new _dafny.CodePoint('x'.codePointAt(0))))) {
        return _dafny.Set.fromElements(false);
      } else {
        return _dafny.Set.fromElements(true, !(true), true);
      }
    };
    static fm5(p0, globalState) {
      if (_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("ujvkkcrk"), new _dafny.CodePoint('w'.codePointAt(0)))) {
        return _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(16)), function (_0_i0) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("rpe")).length)))).length))).length));
      } else {
        return _dafny.Set.fromElements(new BigNumber(233));
      }
    };
    static fm7(p0, p1, globalState) {
      return ((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (function () {
          let _coll1 = new _dafny.Set();
          for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(24))).length))).cardinality()),true)).length))).Keys.Elements) {
            let _1_v1 = _compr_1;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(24))).length))).cardinality()),true)).length))).contains(_1_v1)) {
              _coll1.add(_1_v1);
            }
          }
          return _coll1;
        }()).Elements) {
          let _2_v0 = _compr_0;
          if ((function () {
            let _coll2 = new _dafny.Set();
            for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(24))).length))).cardinality()),true)).length))).Keys.Elements) {
              let _3_v1 = _compr_2;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(24))).length))).cardinality()),true)).length))).contains(_3_v1)) {
                _coll2.add(_3_v1);
              }
            }
            return _coll2;
          }()).contains(_2_v0)) {
            _coll0.push([_2_v0,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('v'.codePointAt(0)))).length)]);
          }
        }
        return _coll0;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(508)), function (_4_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })).length)))).Merge(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (function () {
          let _coll4 = new _dafny.Map();
          for (const _compr_4 of (_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)))).Elements) {
            let _5_v3 = _compr_4;
            if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0))), _5_v3)) {
              _coll4.push([_5_v3,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(true)), _dafny.MultiSet.FromArray(_dafny.Seq.of(true)), _dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(true))).length), new BigNumber((_dafny.Seq.of(new BigNumber(629), new BigNumber(834))).length))).length)]);
            }
          }
          return _coll4;
        }()).Keys.Elements) {
          let _6_v2 = _compr_3;
          if ((function () {
            let _coll5 = new _dafny.Map();
            for (const _compr_5 of (_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)))).Elements) {
              let _5_v3 = _compr_5;
              if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0))), _5_v3)) {
                _coll5.push([_5_v3,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(true)), _dafny.MultiSet.FromArray(_dafny.Seq.of(true)), _dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(true))).length), new BigNumber((_dafny.Seq.of(new BigNumber(629), new BigNumber(834))).length))).length)]);
              }
            }
            return _coll5;
          }()).contains(_6_v2)) {
            _coll3.push([_6_v2,new BigNumber((function () {
              let _coll6 = new _dafny.Map();
              for (const _compr_6 of _dafny.IntegerRange(new BigNumber(309), new BigNumber(772))) {
                let _7_v4 = _compr_6;
                if (((new BigNumber(309)).isLessThanOrEqualTo(_7_v4)) && ((_7_v4).isLessThan(new BigNumber(772)))) {
                  _coll6.push([_module.__default.safeDivisionInt(_7_v4, new BigNumber(-942)),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("wg")).length))]);
                }
              }
              return _coll6;
            }()).length)]);
          }
        }
        return _coll3;
      }());
    };
    static fm8(globalState) {
      return _dafny.Seq.UnicodeFromString("kh");
    };
    static fm9(globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Seq.of(false))).Union((_dafny.MultiSet.fromElements(_dafny.Seq.of(true))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(722)), function (_8_i0) {
        return _dafny.Seq.of(false);
      }))));
    };
    static fm10(globalState) {
      return new _dafny.CodePoint('a'.codePointAt(0));
    };
    static fm11(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(35)), function (_9_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("nsov"),false))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("k"),!(true))).Merge(function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("b"), _dafny.Seq.UnicodeFromString("v"))).Elements) {
          let _10_v0 = _compr_7;
          if ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("b"), _dafny.Seq.UnicodeFromString("v"))).contains(_10_v0)) {
            _coll7.push([_10_v0,true]);
          }
        }
        return _coll7;
      }()));
    };
    static fm14(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("pess")).length)), new BigNumber((_dafny.Seq.UnicodeFromString("pj")).length), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(501),true)).length))).length)).multipliedBy(new BigNumber(143)));
    };
    static fm15(p0, p1, p2, globalState) {
      if ((_module.D1.create_DC6(!(false), false)).dtor_cf8) {
        return ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("u")))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("ddchqe"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(788)), function (_11_i0) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        }))));
      } else {
        return (_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(597)), function (_12_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        })));
      }
    };
    static fm16(globalState) {
      return _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true, false, false, true, true)).length),new BigNumber(698)));
    };
    static fm17(globalState) {
      return _dafny.Set.fromElements((new BigNumber(472)).minus(new BigNumber(-750)));
    };
    static fm18(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe((_module.D12.create_DC31(true, _dafny.Seq.UnicodeFromString("wyhk"), new BigNumber(716), true)).dtor_cf47,(_module.D5.create_DC14(!(false), new BigNumber(560))).dtor_cf17)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("qgxmchpc"),new BigNumber(411)));
    };
    static fm19(p0, p1, globalState) {
      return function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(-378), new BigNumber(-241))) {
          let _13_v0 = _compr_8;
          if (((new BigNumber(-378)).isLessThanOrEqualTo(_13_v0)) && ((_13_v0).isLessThan(new BigNumber(-241)))) {
            _coll8.push([(_13_v0).multipliedBy(new BigNumber(-530)),(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(249)), function (_14_i0) {
              return new BigNumber(690);
            })).length)).minus(new BigNumber((_dafny.Seq.of(false, true)).length))]);
          }
        }
        return _coll8;
      }();
    };
    static fm20(p0, globalState) {
      let _source0 = _module.D3.create_DC10(!(false));
      if (_source0.is_DC9) {
        let _15___mcc_h0 = (_source0).cf11;
        let _16___mcc_h1 = (_source0).cf12;
        let _17_cf12 = _16___mcc_h1;
        let _18_cf11 = _15___mcc_h0;
        return new _dafny.CodePoint('u'.codePointAt(0));
      } else if (_source0.is_DC10) {
        let _19___mcc_h2 = (_source0).cf13;
        let _20_cf13 = _19___mcc_h2;
        return new _dafny.CodePoint('e'.codePointAt(0));
      } else {
        let _21___mcc_h3 = (_source0).cf10;
        let _22_cf10 = _21___mcc_h3;
        return _22_cf10;
      }
    };
    static fm21(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ihvn"), _dafny.Seq.UnicodeFromString("nvbkgws")), _dafny.Seq.UnicodeFromString("qjxpor"));
    };
    static fm22(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(282)), _dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(266)), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-107))).length)), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(571)))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber(-604)))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(453)))));
    };
    static fm23(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),true)));
    };
    static fm24(globalState) {
      return _module.D3.create_DC10(!(true));
    };
    static fm25(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC1();
    };
    static fm26(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(-304))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(597)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-128)));
    };
    static fm27(globalState) {
      return _dafny.Seq.Concat(((true) ? (_dafny.Seq.of(false, true)) : (_dafny.Seq.of(false, true, true))), _dafny.Seq.of(false));
    };
    static fm28(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-633), new BigNumber((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(-876), new BigNumber(398))) {
          let _23_v0 = _compr_9;
          if (((new BigNumber(-876)).isLessThanOrEqualTo(_23_v0)) && ((_23_v0).isLessThan(new BigNumber(398)))) {
            _coll9.add(_module.__default.safeModuloInt(_23_v0, new BigNumber(378)));
          }
        }
        return _coll9;
      }()).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(80)), function (_24_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })).length))))).Difference(_dafny.MultiSet.fromElements(new BigNumber(52)))).Union(((false) ? (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(466)), function (_25_i1) {
        return _dafny.Map.Empty.slice().updateUnsafe(true,true);
      })).length),_dafny.Set.fromElements(new BigNumber(168)))).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0)))).length)),false)).length),new BigNumber(569))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))) : (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(976))).length))))));
    };
    static fm29(globalState) {
      return _module.D1.create_DC4(new BigNumber(295), _module.__default.safeModuloInt(new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality()), new BigNumber(420)), (new BigNumber(151)).isLessThanOrEqualTo(new BigNumber(-521)));
    };
    static fm30(p0, globalState) {
      return _dafny.Seq.of(new BigNumber(-734));
    };
    static fm31(p0, p1, globalState) {
      let _source1 = _module.D12.create_DC29(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(595))).length), new BigNumber(463)));
      if (_source1.is_DC30) {
        let _26___mcc_h0 = (_source1).cf44;
        let _27___mcc_h1 = (_source1).cf45;
        let _28_cf45 = _27___mcc_h1;
        let _29_cf44 = _26___mcc_h0;
        return _module.D1.create_DC5(new BigNumber(993));
      } else if (_source1.is_DC31) {
        let _30___mcc_h2 = (_source1).cf46;
        let _31___mcc_h3 = (_source1).cf47;
        let _32___mcc_h4 = (_source1).cf48;
        let _33___mcc_h5 = (_source1).cf49;
        let _34_cf49 = _33___mcc_h5;
        let _35_cf48 = _32___mcc_h4;
        let _36_cf47 = _31___mcc_h3;
        let _37_cf46 = _30___mcc_h2;
        return _module.D1.create_DC5((_dafny.ZERO).minus(_35_cf48));
      } else if (_source1.is_DC29) {
        let _38___mcc_h6 = (_source1).cf43;
        let _39_cf43 = _38___mcc_h6;
        return _module.D1.create_DC5(new BigNumber(215));
      } else {
        let _40___mcc_h7 = (_source1).cf50;
        let _41_cf50 = _40___mcc_h7;
        if (false) {
          return _module.D1.create_DC5(new BigNumber(-264));
        } else {
          return _module.D1.create_DC5(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(new BigNumber(784))).length))).length));
        }
      }
    };
    static m0(p0, p1, p2, globalState) {
      let r0 = [];
      let r1 = false;
      let r2 = _dafny.ZERO;
      if (!(p1)) {
        let _42_v0;
        _42_v0 = new _dafny.CodePoint('d'.codePointAt(0));
        let _43_v1;
        _43_v1 = _dafny.Seq.UnicodeFromString("srfvaskfl");
        let _44_v2;
        _44_v2 = _module.D0.create_DC0(p2);
        let _45_v3;
        _45_v3 = _dafny.Map.Empty.slice().updateUnsafe(_43_v1,(_44_v2).dtor_cf0);
        let _46_v4;
        let _nw0 = new _module.C3();
        _nw0.__ctor(_42_v0, _45_v3);
        _46_v4 = _nw0;
        let _47_v5;
        _47_v5 = _dafny.Set.fromElements(_46_v4, _46_v4);
        let _48_v6;
        _48_v6 = _dafny.Seq.of(_47_v5);
        let _49_v7;
        _49_v7 = _module.D1.create_DC6(p1, p1);
        let _50_v8;
        _50_v8 = _dafny.Seq.of(p1);
        let _51_v9;
        _51_v9 = _dafny.MultiSet.fromElements(p2, p0, (_dafny.ZERO).minus(p2), (_dafny.ZERO).minus(p0), p0);
        let _52_v10;
        let _nw1 = Array((new BigNumber(17)).toNumber());
        _nw1[(_dafny.ZERO).toNumber()] = p1;
        _nw1[(_dafny.ONE).toNumber()] = p1;
        _nw1[(new BigNumber(2)).toNumber()] = p1;
        _nw1[(new BigNumber(3)).toNumber()] = ((_48_v6)[_module.__default.safeIndex(p0, new BigNumber((_48_v6).length))]).IsProperSubsetOf(_47_v5);
        _nw1[(new BigNumber(4)).toNumber()] = (_49_v7).dtor_cf8;
        _nw1[(new BigNumber(5)).toNumber()] = p1;
        _nw1[(new BigNumber(6)).toNumber()] = !(((_dafny.ZERO).minus(p0)).isLessThan(new BigNumber((_module.__default.fm21(globalState)).length)));
        _nw1[(new BigNumber(7)).toNumber()] = p1;
        _nw1[(new BigNumber(8)).toNumber()] = !(false);
        _nw1[(new BigNumber(9)).toNumber()] = (_50_v8)[_module.__default.safeIndex(new BigNumber((_51_v9).cardinality()), new BigNumber((_50_v8).length))];
        _nw1[(new BigNumber(10)).toNumber()] = p1;
        _nw1[(new BigNumber(11)).toNumber()] = p1;
        _nw1[(new BigNumber(12)).toNumber()] = p1;
        _nw1[(new BigNumber(13)).toNumber()] = p1;
        _nw1[(new BigNumber(14)).toNumber()] = (p0).isLessThanOrEqualTo(p0);
        _nw1[(new BigNumber(15)).toNumber()] = p1;
        _nw1[(new BigNumber(16)).toNumber()] = p1;
        _52_v10 = _nw1;
        let _index0 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_52_v10).length));
        (_52_v10)[_index0] = !(p1);
        let _index1 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_52_v10).length));
        (_52_v10)[_index1] = p1;
        let _53_v11;
        _53_v11 = _dafny.Map.Empty.slice().updateUnsafe((_52_v10)[_module.__default.safeIndex(new BigNumber(517), new BigNumber((_52_v10).length))],_42_v0);
        _53_v11 = (_53_v11).update(p1, _42_v0);
        let _index2 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_52_v10).length));
        (_52_v10)[_index2] = true;
        (globalState).f6 = new BigNumber((_43_v1).length);
        let _index3 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_52_v10).length));
        (_52_v10)[_index3] = !((_52_v10)[_module.__default.safeIndex(new BigNumber(517), new BigNumber((_52_v10).length))]) || (p1);
      } else {
        let _54_v12;
        _54_v12 = new _dafny.CodePoint('j'.codePointAt(0));
        let _55_v13;
        _55_v13 = _dafny.Seq.UnicodeFromString("vqatmsp");
        let _56_v14;
        let _nw2 = Array((new BigNumber(6)).toNumber());
        _nw2[(_dafny.ZERO).toNumber()] = p0;
        _nw2[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(986)), function (_57_i0) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        }), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(986)), function (_58_i0) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        })).length)), _54_v12)).length);
        _nw2[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(p0, (_dafny.ZERO).minus(new BigNumber((_55_v13).length)));
        _nw2[(new BigNumber(3)).toNumber()] = p0;
        _nw2[(new BigNumber(4)).toNumber()] = new BigNumber(-689);
        _nw2[(new BigNumber(5)).toNumber()] = p2;
        _56_v14 = _nw2;
        let _index4 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_56_v14).length));
        (_56_v14)[_index4] = p0;
        let _59_v15;
        _59_v15 = _dafny.MultiSet.fromElements(p0);
        let _60_v16;
        let _nw3 = Array((new BigNumber(4)).toNumber());
        _nw3[(_dafny.ZERO).toNumber()] = _59_v15;
        _nw3[(_dafny.ONE).toNumber()] = (_59_v15).Difference(_59_v15);
        _nw3[(new BigNumber(2)).toNumber()] = _59_v15;
        _nw3[(new BigNumber(3)).toNumber()] = (_59_v15).Intersect(_59_v15);
        _60_v16 = _nw3;
        let _index5 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_60_v16).length));
        (_60_v16)[_index5] = _module.__default.fm28(p1, p2, p1, globalState);
        let _61_v17;
        _61_v17 = _dafny.MultiSet.fromElements(p1, p1, p1, !(!(p1)));
        let _62_v18;
        let _nw4 = Array((new BigNumber(17)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = p1;
        _nw4[(_dafny.ONE).toNumber()] = p1;
        _nw4[(new BigNumber(2)).toNumber()] = p1;
        _nw4[(new BigNumber(3)).toNumber()] = true;
        _nw4[(new BigNumber(4)).toNumber()] = p1;
        _nw4[(new BigNumber(5)).toNumber()] = p1;
        _nw4[(new BigNumber(6)).toNumber()] = p1;
        _nw4[(new BigNumber(7)).toNumber()] = p1;
        _nw4[(new BigNumber(8)).toNumber()] = p1;
        _nw4[(new BigNumber(9)).toNumber()] = ((p1) ? (p1) : (p1));
        _nw4[(new BigNumber(10)).toNumber()] = p1;
        _nw4[(new BigNumber(11)).toNumber()] = true;
        _nw4[(new BigNumber(12)).toNumber()] = _module.__default.fm1(p1, _54_v12, new BigNumber(245), globalState);
        _nw4[(new BigNumber(13)).toNumber()] = !(!(p1));
        _nw4[(new BigNumber(14)).toNumber()] = (_61_v17).IsDisjointFrom(_61_v17);
        _nw4[(new BigNumber(15)).toNumber()] = _module.__default.fm1(true, _54_v12, p2, globalState);
        _nw4[(new BigNumber(16)).toNumber()] = p1;
        _62_v18 = _nw4;
        let _63_v19;
        _63_v19 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("fjijcedg"),p2);
        let _64_v20;
        let _nw5 = new _module.C2();
        _nw5.__ctor(_54_v12, _63_v19);
        _64_v20 = _nw5;
        let _65_v21;
        _65_v21 = _dafny.Set.fromElements(_64_v20);
        let _index6 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_56_v14).length));
        let _index7 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_60_v16).length));
        let _rhs0 = _module.__default.safeDivisionInt(p2, p2);
        let _rhs1 = (_59_v15).Difference(_59_v15);
        let _rhs2 = _62_v18;
        let _rhs3 = new BigNumber(664);
        let _rhs4 = !(!(_65_v21).contains(_64_v20));
        let _lhs0 = _56_v14;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_56_v14).length));
        let _lhs2 = _60_v16;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_60_v16).length));
        let _lhs4 = globalState;
        _lhs0[_lhs1] = _rhs0;
        _lhs2[_lhs3] = _rhs1;
        r0 = _rhs2;
        _lhs4.f6 = _rhs3;
        r1 = _rhs4;
        let _66_v22;
        let _nw6 = new _module.C3();
        _nw6.__ctor(_54_v12, _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("umuxo"),(_56_v14)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((_56_v14).length))]));
        _66_v22 = _nw6;
        let _67_v23;
        let _init0 = ((_68_v13) => function (_69_i1) {
          return _68_v13;
        })(_55_v13);
        let _nw7 = Array((new BigNumber(7)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw7.length); _i0_0++) {
          _nw7[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _67_v23 = _nw7;
        let _index8 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_67_v23).length));
        (_67_v23)[_index8] = _55_v13;
        let _index9 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_67_v23).length));
        let _rhs5 = p1;
        let _rhs6 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(38)), ((_70_v12) => function (_71_i2) {
          return _70_v12;
        })(_54_v12)), _55_v13), _dafny.Seq.update(_dafny.Seq.update(_55_v13, _module.__default.safeIndex((_56_v14)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((_56_v14).length))], new BigNumber((_55_v13).length)), _54_v12), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.update(_55_v13, _module.__default.safeIndex((_56_v14)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((_56_v14).length))], new BigNumber((_55_v13).length)), _54_v12)).length)), new _dafny.CodePoint('q'.codePointAt(0))));
        let _rhs7 = _66_v22;
        let _rhs8 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("estdt"), _dafny.Seq.UnicodeFromString("mfn"));
        let _lhs5 = _67_v23;
        let _lhs6 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_67_v23).length));
        r1 = _rhs5;
        _55_v13 = _rhs6;
        _66_v22 = _rhs7;
        _lhs5[_lhs6] = _rhs8;
        let _72_v24;
        _72_v24 = _module.D0.create_DC2(_module.D0.create_DC0(p2));
        let _73_v25;
        _73_v25 = _module.D11.create_DC27(_63_v19);
        let _74_v26;
        let _nw8 = new _module.C1();
        _nw8.__ctor(new BigNumber(687), p1, _module.__default.fm20(_module.__default.fm1(p1, _54_v12, new BigNumber((_59_v15).cardinality()), globalState), globalState), ((_73_v25).dtor_cf39).Merge((_63_v19).update((_67_v23)[_module.__default.safeIndex(new BigNumber(671), new BigNumber((_67_v23).length))], (_56_v14)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((_56_v14).length))])));
        _74_v26 = _nw8;
        let _75_v27;
        _75_v27 = _dafny.Map.Empty.slice().updateUnsafe(_54_v12,p0);
        let _76_v28;
        let _nw9 = new _module.C0();
        _nw9.__ctor((_67_v23)[_module.__default.safeIndex(new BigNumber(671), new BigNumber((_67_v23).length))], _75_v27, _54_v12, _dafny.Map.Empty.slice().updateUnsafe(_55_v13,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),(_56_v14)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((_56_v14).length))])).length)));
        _76_v28 = _nw9;
        let _77_v29;
        _77_v29 = _dafny.Map.Empty.slice().updateUnsafe(_76_v28,p2);
        let _rhs9 = _62_v18;
        let _rhs10 = _72_v24;
        let _rhs11 = _74_v26;
        let _rhs12 = (((_77_v29).contains(_76_v28)) ? ((_77_v29).get(_76_v28)) : (((_56_v14)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((_56_v14).length))]).minus(p2)));
        let _rhs13 = _62_v18;
        _62_v18 = _rhs9;
        _72_v24 = _rhs10;
        _74_v26 = _rhs11;
        r2 = _rhs12;
        _62_v18 = _rhs13;
        (_76_v28).f18 = _54_v12;
        let _index10 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_62_v18).length));
        (_62_v18)[_index10] = (_74_v26).f23;
        let _index11 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_62_v18).length));
        (_62_v18)[_index11] = _dafny.Seq.contains(_dafny.Seq.of(_74_v26.f22), new BigNumber(316));
      }
      let _78_v30;
      _78_v30 = _dafny.Seq.UnicodeFromString("lgcj");
      let _hi0 = new BigNumber((_dafny.Seq.Concat(_78_v30, _module.__default.fm21(globalState))).length);
      for (let _79_i3 = p0; _79_i3.isLessThan(_hi0); _79_i3 = _79_i3.plus(_dafny.ONE)) {
        let _80_v31;
        _80_v31 = _dafny.Map.Empty.slice().updateUnsafe(false,_79_i3);
        (globalState).f6 = (((_80_v31).contains(false)) ? ((_80_v31).get(false)) : (_79_i3));
        let _81_v32;
        _81_v32 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),_dafny.Seq.Concat(_78_v30, _78_v30));
        let _82_v33;
        _82_v33 = new _dafny.CodePoint('y'.codePointAt(0));
        _81_v32 = (_81_v32).update(_module.__default.fm1(p1, _82_v33, new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(538)), ((_83_v33) => function (_84_i4) {
          return _83_v33;
        })(_82_v33)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(538)), ((_85_v33) => function (_86_i4) {
          return _85_v33;
        })(_82_v33))).length)), _module.__default.fm20(_module.__default.fm1(p1, _82_v33, p2, globalState), globalState))).length), globalState), _78_v30);
        let _87_v34;
        _87_v34 = _dafny.Map.Empty.slice().updateUnsafe(_82_v33,p2);
        let _88_v35;
        _88_v35 = _dafny.Map.Empty.slice().updateUnsafe(_78_v30,p0);
        let _89_v36;
        let _nw10 = new _module.C0();
        _nw10.__ctor(_dafny.Seq.update(_78_v30, _module.__default.safeIndex(_79_i3, new BigNumber((_78_v30).length)), _82_v33), (_87_v34).Merge(_87_v34), new _dafny.CodePoint('r'.codePointAt(0)), (_88_v35).Merge(_88_v35));
        _89_v36 = _nw10;
        let _90_v37;
        let _nw11 = Array((new BigNumber(3)).toNumber()).fill(_dafny.MultiSet.Empty);
        _90_v37 = _nw11;
        let _91_v38;
        _91_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,_90_v37);
        let _92_v39;
        let _nw12 = Array((new BigNumber(24)).toNumber());
        _nw12[(_dafny.ZERO).toNumber()] = _90_v37;
        _nw12[(_dafny.ONE).toNumber()] = _90_v37;
        _nw12[(new BigNumber(2)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(3)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(4)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(5)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(6)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(7)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(8)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(9)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(10)).toNumber()] = (((_91_v38).contains(p0)) ? ((_91_v38).get(p0)) : (_90_v37));
        _nw12[(new BigNumber(11)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(12)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(13)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(14)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(15)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(16)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(17)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(18)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(19)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(20)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(21)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(22)).toNumber()] = _90_v37;
        _nw12[(new BigNumber(23)).toNumber()] = _90_v37;
        _92_v39 = _nw12;
        let _index12 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_92_v39).length));
        (_92_v39)[_index12] = _90_v37;
        let _index13 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_92_v39).length));
        (_92_v39)[_index13] = _90_v37;
      }
      let _93_v40;
      let _nw13 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
      _93_v40 = _nw13;
      let _index14 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length));
      (_93_v40)[_index14] = _module.__default.safeDivisionInt(p2, p0);
      let _94_v41;
      let _nw14 = Array((new BigNumber(6)).toNumber()).fill(false);
      _94_v41 = _nw14;
      let _index15 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
      (_94_v41)[_index15] = ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(p2)).length))).isLessThanOrEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(p2)));
      let _95_v42;
      _95_v42 = new _dafny.CodePoint('j'.codePointAt(0));
      let _index16 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length));
      let _index17 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
      let _rhs14 = new BigNumber(715);
      let _rhs15 = !(p1);
      let _rhs16 = _dafny.Seq.Concat(_dafny.Seq.of(_95_v42, _module.__default.fm10(globalState)), _module.__default.fm8(globalState));
      let _lhs7 = _93_v40;
      let _lhs8 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length));
      let _lhs9 = _94_v41;
      let _lhs10 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
      _lhs7[_lhs8] = _rhs14;
      _lhs9[_lhs10] = _rhs15;
      _78_v30 = _rhs16;
      let _96_v43;
      _96_v43 = _dafny.Map.Empty.slice().updateUnsafe(_95_v42,(_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]);
      let _97_v45;
      _97_v45 = _dafny.Map.Empty.slice().updateUnsafe(_95_v42,new BigNumber(979));
      _96_v43 = (function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (_97_v45).Keys.Elements) {
          let _98_v44 = _compr_10;
          if ((_97_v45).contains(_98_v44)) {
            _coll10.push([_98_v44,(_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]]);
          }
        }
        return _coll10;
      }()).Merge((_96_v43).Merge(_dafny.Map.Empty.slice().updateUnsafe(_95_v42,(_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))])));
      let _99_v46;
      _99_v46 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),p1);
      if (((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))], new BigNumber((_99_v46).length)))).isLessThan(p2)) {
        let _100_v47;
        _100_v47 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
        let _101_v48;
        _101_v48 = _dafny.Seq.of(_100_v47);
        if (!_dafny.areEqual(_101_v48, _dafny.Seq.Create(_module.__default.abs(new BigNumber(472)), ((_102_p2, _103_v41) => function (_104_i5) {
          return function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of _dafny.IntegerRange(new BigNumber(-992), new BigNumber(954))) {
              let _105_v49 = _compr_11;
              if (((new BigNumber(-992)).isLessThanOrEqualTo(_105_v49)) && ((_105_v49).isLessThan(new BigNumber(954)))) {
                _coll11.push([_module.__default.safeDivisionInt(_105_v49, _102_p2),(_103_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_103_v41).length))]]);
              }
            }
            return _coll11;
          }();
        })(p2, _94_v41)))) {
          let _106_v50;
          _106_v50 = _module.D9.create_DC23(_dafny.Seq.Create(_module.__default.abs(new BigNumber(865)), ((_107_p0) => function (_108_i6) {
  return _dafny.Set.fromElements(_107_p0);
})(p0)), p2, p1, p2);
          let _index18 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          let _index19 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length));
          let _index20 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          let _rhs17 = _dafny.Seq.update(_dafny.Seq.of((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))]), _module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_dafny.Seq.of((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))])).length)), (_106_v50).dtor_cf31);
          let _rhs18 = ((p1) ? ((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]) : (!(p1)));
          let _rhs19 = _module.__default.safeModuloInt((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))], p0);
          let _rhs20 = (_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))];
          let _rhs21 = (_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))];
          let _lhs11 = globalState;
          let _lhs12 = _94_v41;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          let _lhs14 = _93_v40;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length));
          let _lhs16 = _94_v41;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          _lhs11.f1 = _rhs17;
          _lhs12[_lhs13] = _rhs18;
          _lhs14[_lhs15] = _rhs19;
          _lhs16[_lhs17] = _rhs20;
          r1 = _rhs21;
          let _109_v51;
          _109_v51 = _dafny.Map.Empty.slice().updateUnsafe(_78_v30,p0);
          let _110_v52;
          _110_v52 = _dafny.Set.fromElements(true, true, !(p1));
          let _111_v53;
          _111_v53 = _module.D0.create_DC0(p0);
          _109_v51 = (_109_v51).update(_dafny.Seq.update(_78_v30, _module.__default.safeIndex((_module.__default.fm31(_110_v52, p0, globalState)).dtor_cf6, new BigNumber((_78_v30).length)), _95_v42), (_111_v53).dtor_cf0);
          let _index21 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          (_94_v41)[_index21] = p1;
          let _112_v54;
          _112_v54 = _dafny.Map.Empty.slice().updateUnsafe((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))],p2);
          _112_v54 = (_112_v54).update((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))], (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(p2), new BigNumber(119))));
          let _index22 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length));
          (_93_v40)[_index22] = new BigNumber(-852);
        } else {
          let _113_v55;
          _113_v55 = _module.D10.create_DC25(_93_v40);
          let _114_v56;
          _114_v56 = _dafny.Seq.of((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]);
          let _pat_let_tv0 = _93_v40;
          let _115_v57;
          _115_v57 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let0_0) {
            return function (_116_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_117_dt__update_hcf33_h0) {
                  return _module.D10.create_DC25(_117_dt__update_hcf33_h0);
                }(_pat_let1_0);
              }(_pat_let_tv0);
            }(_pat_let0_0);
          }(_113_v55)).dtor_cf33,_114_v56);
          _115_v57 = _115_v57;
          let _118_v58;
          _118_v58 = _dafny.Seq.of(p0, p0);
          let _index23 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length));
          let _index24 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          let _rhs22 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(79)), ((_119_v42) => function (_120_i7) {
            return _119_v42;
          })(_95_v42))).length);
          let _rhs23 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))], (_dafny.ZERO).minus(p2)), p2);
          let _rhs24 = _118_v58;
          let _rhs25 = false;
          let _lhs18 = globalState;
          let _lhs19 = _93_v40;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length));
          let _lhs21 = globalState;
          let _lhs22 = _94_v41;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          _lhs18.f6 = _rhs22;
          _lhs19[_lhs20] = _rhs23;
          _lhs21.f1 = _rhs24;
          _lhs22[_lhs23] = _rhs25;
          r1 = (_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))];
          _95_v42 = _95_v42;
          let _index25 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          (_94_v41)[_index25] = !(p1);
        }
        if (true) {
          let _121_v59;
          _121_v59 = _dafny.MultiSet.fromElements(p1);
          let _index26 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          (_94_v41)[_index26] = (((_121_v59).IsProperSubsetOf(_121_v59)) ? (p1) : ((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]));
          let _122_v60;
          _122_v60 = _dafny.Seq.of((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]);
          let _123_v61;
          _123_v61 = _dafny.Map.Empty.slice().updateUnsafe(_78_v30,new BigNumber((_122_v60).length));
          let _124_v62;
          let _nw15 = new _module.C1();
          _nw15.__ctor(p2, (_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))], _95_v42, (_123_v61).update(_78_v30, p0));
          _124_v62 = _nw15;
          let _125_v63;
          _125_v63 = _dafny.MultiSet.fromElements(_124_v62, _124_v62, _124_v62);
          let _126_v64;
          _126_v64 = _dafny.Map.Empty.slice().updateUnsafe((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))],_125_v63);
          let _127_v65;
          _127_v65 = _dafny.MultiSet.fromElements(new BigNumber((_126_v64).length));
          let _index27 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          (_94_v41)[_index27] = !((_127_v65).contains((_dafny.ZERO).minus((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))])));
          let _128_v66;
          _128_v66 = _dafny.Map.Empty.slice().updateUnsafe(p0,_93_v40);
          let _129_v67;
          _129_v67 = _dafny.Set.fromElements(p0);
          let _130_v68;
          _130_v68 = _dafny.MultiSet.fromElements(_93_v40, _93_v40, (((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]) ? ((((_128_v66).contains(new BigNumber((_129_v67).length))) ? ((_128_v66).get(new BigNumber((_129_v67).length))) : (_93_v40))) : (_93_v40)));
          _130_v68 = _130_v68;
          let _index28 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length));
          (_93_v40)[_index28] = _module.__default.safeModuloInt((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))], (p0).multipliedBy((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))]));
          let _131_v69;
          let _nw16 = new _module.C0();
          _nw16.__ctor(_78_v30, _module.__default.fm14((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))], globalState), _95_v42, _123_v61);
          _131_v69 = _nw16;
        } else {
          let _132_v70;
          _132_v70 = _dafny.Set.fromElements(p2);
          let _133_v71;
          _133_v71 = _dafny.Seq.of((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]);
          let _134_v72;
          _134_v72 = _dafny.MultiSet.fromElements(_78_v30);
          let _135_v73;
          _135_v73 = _dafny.Seq.of((_133_v71)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_78_v30, _module.__default.safeIndex(new BigNumber((_134_v72).cardinality()), new BigNumber((_78_v30).length)), _95_v42)).length),(((_100_v47).contains(p2)) ? ((_100_v47).get(p2)) : (p1)))).length), new BigNumber((_133_v71).length))]);
          let _index29 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          let _rhs26 = ((_132_v70).Intersect(_132_v70)).IsDisjointFrom((_132_v70).Union(_132_v70));
          let _rhs27 = (_135_v73)[_module.__default.safeIndex((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))], new BigNumber((_135_v73).length))];
          let _lhs24 = _94_v41;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          r1 = _rhs26;
          _lhs24[_lhs25] = _rhs27;
          let _index30 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length));
          (_93_v40)[_index30] = _module.__default.safeModuloInt(new BigNumber((_78_v30).length), p0);
          let _index31 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length));
          (_93_v40)[_index31] = p2;
          _95_v42 = _95_v42;
          let _136_v74;
          _136_v74 = _module.D1.create_DC3(p1);
          let _137_v75;
          _137_v75 = _module.D0.create_DC1();
          let _138_v76;
          let _init1 = ((_139_v71) => function (_140_i8) {
            return _139_v71;
          })(_133_v71);
          let _nw17 = Array((new BigNumber(25)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw17.length); _i0_1++) {
            _nw17[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _138_v76 = _nw17;
          let _index32 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_138_v76).length));
          (_138_v76)[_index32] = _133_v71;
          let _141_v77;
          let _nw18 = Array((new BigNumber(4)).toNumber()).fill([]);
          _141_v77 = _nw18;
          let _142_v78;
          let _nw19 = Array((new BigNumber(17)).toNumber()).fill(_module.D1.Default());
          _142_v78 = _nw19;
          let _index33 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_141_v77).length));
          (_141_v77)[_index33] = _142_v78;
          let _index34 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_138_v76).length));
          let _index35 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_141_v77).length));
          let _rhs28 = _136_v74;
          let _rhs29 = new BigNumber(-121);
          let _rhs30 = ((p1) ? (_137_v75) : (_137_v75));
          let _rhs31 = _133_v71;
          let _rhs32 = _142_v78;
          let _lhs26 = globalState;
          let _lhs27 = _138_v76;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_138_v76).length));
          let _lhs29 = _141_v77;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_141_v77).length));
          _136_v74 = _rhs28;
          _lhs26.f6 = _rhs29;
          _137_v75 = _rhs30;
          _lhs27[_lhs28] = _rhs31;
          _lhs29[_lhs30] = _rhs32;
        }
        let _143_v79;
        _143_v79 = _module.D6.create_DC15(_97_v45);
        let _source2 = _143_v79;
        if (_source2.is_DC16) {
          let _144___mcc_h0 = (_source2).cf19;
          let _145___mcc_h1 = (_source2).cf20;
          let _146_cf20 = _145___mcc_h1;
          let _147_cf19 = _144___mcc_h0;
          let _148_v80;
          _148_v80 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(613),_78_v30);
          _148_v80 = ((_148_v80).Merge(_148_v80)).Merge((_148_v80).Merge(_148_v80));
          let _149_v82;
          _149_v82 = _module.D9.create_DC22(_78_v30);
          let _150_v83;
          _150_v83 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("qribx"), _module.__default.safeIndex(new BigNumber(73), new BigNumber((_dafny.Seq.UnicodeFromString("qribx")).length)), _95_v42), (_149_v82).dtor_cf27);
          let _151_v84;
          let _nw20 = new _module.C2();
          _nw20.__ctor(_95_v42, function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of (_150_v83).Elements) {
              let _152_v81 = _compr_12;
              if (_dafny.Seq.contains(_150_v83, _152_v81)) {
                _coll12.push([_152_v81,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(p2, p0)).length),p1)).length)]);
              }
            }
            return _coll12;
          }());
          _151_v84 = _nw20;
          let _153_v85;
          _153_v85 = _dafny.Set.fromElements(_100_v47);
          let _154_v87;
          _154_v87 = _dafny.MultiSet.fromElements(function () {
            let _coll13 = new _dafny.Map();
            for (const _compr_13 of _dafny.IntegerRange(new BigNumber(763), new BigNumber(-498))) {
              let _155_v86 = _compr_13;
              if (((new BigNumber(763)).isLessThanOrEqualTo(_155_v86)) && ((_155_v86).isLessThan(new BigNumber(-498)))) {
                _coll13.push([(_155_v86).multipliedBy(new BigNumber((_100_v47).length)),(_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]]);
              }
            }
            return _coll13;
          }());
          _153_v85 = ((_153_v85).Difference(function () {
            let _coll14 = new _dafny.Set();
            for (const _compr_14 of (_154_v87).Elements) {
              let _156_v88 = _compr_14;
              if ((_154_v87).contains(_156_v88)) {
                _coll14.add(_156_v88);
              }
            }
            return _coll14;
          }())).Difference(_153_v85);
          let _157_v89;
          _157_v89 = _dafny.Map.Empty.slice().updateUnsafe(_78_v30,p2);
          let _158_v90;
          let _nw21 = new _module.C1();
          _nw21.__ctor(p2, _dafny.Seq.IsPrefixOf(_module.__default.fm30(p1, globalState), _dafny.Seq.of(p2, (_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))])), new _dafny.CodePoint('u'.codePointAt(0)), _157_v89);
          _158_v90 = _nw21;
          let _159_v91;
          let _nw22 = new _module.C3();
          _nw22.__ctor(_95_v42, (_157_v89).Merge(_157_v89));
          _159_v91 = _nw22;
          let _160_v92;
          let _nw23 = Array((new BigNumber(16)).toNumber());
          _160_v92 = _nw23;
          let _index36 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_160_v92).length));
          (_160_v92)[_index36] = _159_v91;
          let _161_v93;
          _161_v93 = _dafny.MultiSet.fromElements((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))], true);
          let _pat_let_tv1 = _93_v40;
          let _pat_let_tv2 = _93_v40;
          let _pat_let_tv3 = p1;
          let _index37 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length));
          let _index38 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_160_v92).length));
          let _rhs33 = _158_v90;
          let _rhs34 = _159_v91;
          let _rhs35 = ((new BigNumber((_78_v30).length)).multipliedBy((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))])).plus((p2).minus(new BigNumber((_161_v93).cardinality())));
          let _rhs36 = (function (_pat_let2_0) {
            return function (_162_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_163_dt__update_hcf38_h0) {
                  return function (_pat_let4_0) {
                    return function (_164_dt__update_hcf37_h0) {
                      return function (_pat_let5_0) {
                        return function (_165_dt__update_hcf34_h0) {
                          return _module.D10.create_DC26(_165_dt__update_hcf34_h0, (_162_dt__update__tmp_h1).dtor_cf35, (_162_dt__update__tmp_h1).dtor_cf36, _164_dt__update_hcf37_h0, _163_dt__update_hcf38_h0);
                        }(_pat_let5_0);
                      }(new BigNumber(-203));
                    }(_pat_let4_0);
                  }(!(_pat_let_tv3));
                }(_pat_let3_0);
              }((_pat_let_tv2)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_pat_let_tv1).length))]);
            }(_pat_let2_0);
          }(_module.D10.create_DC26(p2, _94_v41, (_158_v90).f23, (_158_v90).f23, p2))).dtor_cf34;
          let _rhs37 = _159_v91;
          let _lhs31 = globalState;
          let _lhs32 = _93_v40;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length));
          let _lhs34 = _160_v92;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_160_v92).length));
          _158_v90 = _rhs33;
          _159_v91 = _rhs34;
          _lhs31.f6 = _rhs35;
          _lhs32[_lhs33] = _rhs36;
          _lhs34[_lhs35] = _rhs37;
        } else if (_source2.is_DC15) {
          let _166___mcc_h2 = (_source2).cf18;
          let _167_cf18 = _166___mcc_h2;
          let _index39 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          (_94_v41)[_index39] = (_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))];
          _78_v30 = _dafny.Seq.Concat(_dafny.Seq.Concat(_78_v30, _78_v30), _78_v30);
          let _168_v94;
          let _nw24 = Array((new BigNumber(26)).toNumber());
          _168_v94 = _nw24;
          let _169_v95;
          let _nw25 = new _module.C3();
          _nw25.__ctor(_95_v42, _dafny.Map.Empty.slice().updateUnsafe(_78_v30,p2));
          _169_v95 = _nw25;
          let _index40 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_168_v94).length));
          (_168_v94)[_index40] = _169_v95;
          let _170_v97;
          _170_v97 = _dafny.Map.Empty.slice().updateUnsafe(_78_v30,new BigNumber(((_module.D12.create_DC29(_dafny.Set.fromElements(p2, p2))).dtor_cf43).length));
          let _171_v98;
          let _nw26 = new _module.C0();
          _nw26.__ctor(_78_v30, function () {
            let _coll15 = new _dafny.Map();
            for (const _compr_15 of (_78_v30).Elements) {
              let _172_v96 = _compr_15;
              if (_dafny.Seq.contains(_78_v30, _172_v96)) {
                _coll15.push([_172_v96,p0]);
              }
            }
            return _coll15;
          }(), _95_v42, (_170_v97).Merge(_170_v97));
          _171_v98 = _nw26;
          let _173_v99;
          let _init2 = ((_174_v47) => function (_175_i9) {
            return _174_v47;
          })(_100_v47);
          let _nw27 = Array((new BigNumber(24)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw27.length); _i0_2++) {
            _nw27[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _173_v99 = _nw27;
          let _index41 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_168_v94).length));
          let _rhs38 = _169_v95;
          let _rhs39 = _171_v98;
          let _rhs40 = _173_v99;
          let _lhs36 = _168_v94;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_168_v94).length));
          _lhs36[_lhs37] = _rhs38;
          _171_v98 = _rhs39;
          _173_v99 = _rhs40;
          let _176_v100;
          _176_v100 = _dafny.Set.fromElements(p1);
          let _177_v101;
          _177_v101 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _178_v102;
          _178_v102 = _dafny.Map.Empty.slice().updateUnsafe((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))],_177_v101);
          (globalState).f6 = new BigNumber(((((_176_v100).IsSubsetOf(_176_v100)) ? (_177_v101) : (((((_178_v102).contains(p1)) ? ((_178_v102).get(p1)) : (_177_v101))).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,p0))))).length);
        } else {
          let _179___mcc_h3 = (_source2).cf21;
          let _180_cf21 = _179___mcc_h3;
          let _181_v103;
          _181_v103 = _dafny.Map.Empty.slice().updateUnsafe(((true) ? (p1) : (p1)),(_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))]);
          _181_v103 = _181_v103;
          let _index42 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          (_94_v41)[_index42] = (_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))];
          let _index43 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          (_94_v41)[_index43] = (p2).isLessThan(p2);
          (globalState).f6 = p0;
        }
        let _rhs41 = (((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]) ? (_module.__default.safeModuloInt(p0, new BigNumber(362))) : ((p0).plus(p2)));
        let _rhs42 = _94_v41;
        let _lhs38 = globalState;
        _lhs38.f6 = _rhs41;
        _94_v41 = _rhs42;
        let _182_v104;
        let _nw28 = Array((_dafny.ONE).toNumber()).fill([]);
        _182_v104 = _nw28;
        let _183_v105;
        _183_v105 = _module.D6.create_DC16(p1, _182_v104);
        let _184_v106;
        _184_v106 = _module.D6.create_DC17(_183_v105);
        let _185_v107;
        _185_v107 = _module.D6.create_DC17(((p1) ? (_183_v105) : (_184_v106)));
        _185_v107 = _185_v107;
      } else {
        let _186_v108;
        _186_v108 = _dafny.Map.Empty.slice().updateUnsafe(_78_v30,(_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))]);
        let _187_v109;
        let _nw29 = new _module.C2();
        _nw29.__ctor(new _dafny.CodePoint('f'.codePointAt(0)), _186_v108);
        _187_v109 = _nw29;
        let _188_v110;
        _188_v110 = _dafny.Seq.of(_187_v109);
        let _189_v111;
        _189_v111 = _dafny.Set.fromElements(new BigNumber((_188_v110).length), (_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))]);
        if ((_189_v111).IsDisjointFrom(_189_v111)) {
          let _190_v112;
          _190_v112 = _dafny.Map.Empty.slice().updateUnsafe(((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]) === ((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]),p0);
          let _191_v113;
          let _nw30 = new _module.C0();
          _nw30.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(525)), ((_192_v42) => function (_193_i10) {
            return _192_v42;
          })(_95_v42)), _97_v45, _95_v42, _module.__default.fm18((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))], globalState));
          _191_v113 = _nw30;
          let _194_v114;
          _194_v114 = _dafny.Seq.of(_191_v113);
          _190_v112 = (_190_v112).update(!_dafny.areEqual(_194_v114, _194_v114), _module.__default.safeDivisionInt((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))], p0));
          let _195_v115;
          _195_v115 = _dafny.Seq.of(p1, p1);
          let _196_v116;
          _196_v116 = _dafny.Seq.of(p1, (_195_v115)[_module.__default.safeIndex(new BigNumber(-484), new BigNumber((_195_v115).length))], p1);
          let _197_v117;
          _197_v117 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_195_v115).length),p2);
          _197_v117 = _197_v117;
          let _198_v118;
          let _nw31 = new _module.C3();
          _nw31.__ctor(new _dafny.CodePoint('p'.codePointAt(0)), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("jfumgxqqe"),p0));
          _198_v118 = _nw31;
          let _199_v119;
          _199_v119 = _module.D9.create_DC24(false);
          let _200_v120;
          _200_v120 = _dafny.Map.Empty.slice().updateUnsafe(_198_v118,(p1) === (!((_199_v119).dtor_cf32)));
          _200_v120 = (_200_v120).update(_198_v118, p1);
          let _index44 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          let _rhs43 = (_187_v109).fm3(p1, new BigNumber((_dafny.Seq.Concat(_78_v30, _78_v30)).length), globalState);
          let _rhs44 = _95_v42;
          let _rhs45 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(p2, new BigNumber((_dafny.Seq.of(_94_v41, _94_v41)).length)), (_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))]);
          let _lhs39 = _94_v41;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length));
          let _lhs41 = _191_v113;
          let _lhs42 = globalState;
          _lhs39[_lhs40] = _rhs43;
          _lhs41.f18 = _rhs44;
          _lhs42.f6 = _rhs45;
          let _201_v122;
          _201_v122 = _dafny.Seq.of(_dafny.Set.fromElements(p0), (function () {
            let _coll16 = new _dafny.Set();
            for (const _compr_16 of _dafny.IntegerRange(new BigNumber(-685), new BigNumber(204))) {
              let _202_v121 = _compr_16;
              if (((new BigNumber(-685)).isLessThanOrEqualTo(_202_v121)) && ((_202_v121).isLessThan(new BigNumber(204)))) {
                _coll16.add(_module.__default.safeModuloInt(_202_v121, new BigNumber((_dafny.Set.fromElements((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))])).length)));
              }
            }
            return _coll16;
          }()).Intersect(_189_v111));
          _201_v122 = _dafny.Seq.Concat(_dafny.Seq.of(_189_v111), _dafny.Seq.of(_189_v111, _189_v111));
        } else {
          r2 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(((_dafny.ZERO).minus(p0)).minus(p2), (_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))]));
          let _203_v123;
          let _init3 = ((_204_v30, _205_p2) => function (_206_i11) {
            return (_204_v30)[_module.__default.safeIndex(_205_p2, new BigNumber((_204_v30).length))];
          })(_78_v30, p2);
          let _nw32 = Array((new BigNumber(7)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw32.length); _i0_3++) {
            _nw32[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _203_v123 = _nw32;
          let _index45 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_203_v123).length));
          (_203_v123)[_index45] = _95_v42;
          let _index46 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_203_v123).length));
          (_203_v123)[_index46] = (_module.D3.create_DC8(_95_v42)).dtor_cf10;
          let _207_v124;
          _207_v124 = _dafny.Map.Empty.slice().updateUnsafe(_189_v111,_dafny.Seq.of(p1));
          let _208_v125;
          _208_v125 = _dafny.Seq.of(_189_v111);
          let _209_v126;
          _209_v126 = _module.D9.create_DC23(_208_v125, p2, p1, (_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))]);
          let _210_v127;
          _210_v127 = _dafny.Seq.of((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))], (_209_v126).dtor_cf30);
          let _211_v128;
          _211_v128 = _dafny.Seq.of(_210_v127);
          let _212_v129;
          _212_v129 = _dafny.Map.Empty.slice().updateUnsafe(_210_v127,(_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]);
          let _213_v130;
          _213_v130 = _dafny.Map.Empty.slice().updateUnsafe(_212_v129,_210_v127);
          let _214_v132;
          let _nw33 = Array((new BigNumber(6)).toNumber());
          _nw33[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((((_207_v124).contains(_189_v111)) ? ((_207_v124).get(_189_v111)) : (_210_v127)), _210_v127);
          _nw33[(_dafny.ONE).toNumber()] = _210_v127;
          _nw33[(new BigNumber(2)).toNumber()] = _210_v127;
          _nw33[(new BigNumber(3)).toNumber()] = (_211_v128)[_module.__default.safeIndex((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))], new BigNumber((_211_v128).length))];
          _nw33[(new BigNumber(4)).toNumber()] = _210_v127;
          _nw33[(new BigNumber(5)).toNumber()] = (((_213_v130).contains(function () {
            let _coll18 = new _dafny.Map();
            for (const _compr_18 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-229)), ((_219_v127) => function (_220_i12) {
              return _219_v127;
            })(_210_v127))).Elements) {
              let _221_v131 = _compr_18;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-229)), ((_222_v127) => function (_220_i12) {
                return _222_v127;
              })(_210_v127)), _221_v131)) {
                _coll18.push([_221_v131,(_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]]);
              }
            }
            return _coll18;
          }())) ? ((_213_v130).get(function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-229)), ((_215_v127) => function (_216_i12) {
              return _215_v127;
            })(_210_v127))).Elements) {
              let _217_v131 = _compr_17;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-229)), ((_218_v127) => function (_216_i12) {
                return _218_v127;
              })(_210_v127)), _217_v131)) {
                _coll17.push([_217_v131,(_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]]);
              }
            }
            return _coll17;
          }())) : (_dafny.Seq.of(p1, (_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))])));
          _214_v132 = _nw33;
          _214_v132 = _214_v132;
          let _223_v133;
          _223_v133 = _dafny.Set.fromElements(p1);
          let _224_v134;
          _224_v134 = _dafny.Seq.of(new BigNumber((_223_v133).length));
          let _225_v135;
          _225_v135 = _dafny.Seq.of(_186_v108);
          let _226_v136;
          let _nw34 = new _module.C1();
          _nw34.__ctor((_dafny.ZERO).minus((_224_v134)[_module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_224_v134).length))]), p1, (_203_v123)[_module.__default.safeIndex(new BigNumber(967), new BigNumber((_203_v123).length))], (_225_v135)[_module.__default.safeIndex(new BigNumber((_210_v127).length), new BigNumber((_225_v135).length))]);
          _226_v136 = _nw34;
          r2 = p2;
        }
        r1 = (function () {
          let _coll19 = new _dafny.Set();
          for (const _compr_19 of (_189_v111).Elements) {
            let _227_v137 = _compr_19;
            if ((_189_v111).contains(_227_v137)) {
              _coll19.add(_module.__default.safeModuloInt(_227_v137, new BigNumber(36)));
            }
          }
          return _coll19;
        }()).IsDisjointFrom(_189_v111);
        let _228_v138;
        let _nw35 = Array((new BigNumber(11)).toNumber()).fill(_dafny.MultiSet.Empty);
        _228_v138 = _nw35;
        let _229_v139;
        _229_v139 = _dafny.MultiSet.fromElements(_94_v41);
        let _230_v140;
        _230_v140 = _dafny.Map.Empty.slice().updateUnsafe(true,_229_v139);
        let _index47 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_228_v138).length));
        (_228_v138)[_index47] = (((_230_v140).contains(p1)) ? ((_230_v140).get(p1)) : (_229_v139));
        let _index48 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_228_v138).length));
        (_228_v138)[_index48] = _229_v139;
        let _231_v141;
        _231_v141 = _dafny.Seq.of((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))], _module.__default.fm0(globalState));
        (globalState).f1 = _dafny.Seq.Concat((((_187_v109).fm3(p1, p0, globalState)) ? (_231_v141) : (_231_v141)), _231_v141);
        let _232_v142;
        _232_v142 = _dafny.Set.fromElements(p1);
        let _233_v143;
        _233_v143 = _module.D12.create_DC31(p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(758)), ((_234_v42) => function (_235_i13) {
  return _234_v42;
})(_95_v42)), new BigNumber((_module.__default.fm30((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))], globalState)).length), p1);
        let _rhs46 = ((!((_232_v142).IsDisjointFrom(_232_v142))) ? ((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))]) : (new BigNumber((_dafny.Seq.Concat((_233_v143).dtor_cf47, _78_v30)).length)));
        let _rhs47 = ((p1) ? (((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))]).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))]),_94_v41)).length))) : (p0));
        let _lhs43 = globalState;
        _lhs43.f6 = _rhs46;
        r2 = _rhs47;
      }
      let _236_v144;
      _236_v144 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.update(_78_v30, _module.__default.safeIndex(p0, new BigNumber((_78_v30).length)), _95_v42), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.update(_78_v30, _module.__default.safeIndex(p0, new BigNumber((_78_v30).length)), _95_v42)).length)), _95_v42),p2);
      let _237_v145;
      let _nw36 = new _module.C2();
      _nw36.__ctor(_95_v42, _236_v144);
      _237_v145 = _nw36;
      r0 = _94_v41;
      r1 = ((p1) ? ((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))]) : (p1));
      let _238_v146;
      _238_v146 = _dafny.Set.fromElements((_93_v40)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_93_v40).length))]);
      r2 = (_237_v145).fm4((_94_v41)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_94_v41).length))], _95_v42, _module.__default.fm10(globalState), (_238_v146).IsProperSubsetOf(_238_v146), globalState);
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _239_v0;
      _239_v0 = new BigNumber(140);
      let _240_v1;
      _240_v1 = _dafny.Seq.of((_dafny.ZERO).minus(_239_v0));
      let _241_v2;
      _241_v2 = true;
      let _242_v3;
      _242_v3 = _dafny.Map.Empty.slice().updateUnsafe(!(true),_241_v2);
      let _243_v4;
      _243_v4 = new _dafny.CodePoint('q'.codePointAt(0));
      let _244_v5;
      _244_v5 = _dafny.Map.Empty.slice().updateUnsafe(_243_v4,(_dafny.ZERO).minus(_239_v0));
      let _245_v7;
      let _init4 = ((_246_v0) => function (_247_i0) {
        return function () {
          let _coll20 = new _dafny.Map();
          for (const _compr_20 of _dafny.IntegerRange(new BigNumber(399), new BigNumber(-703))) {
            let _248_v6 = _compr_20;
            if (((new BigNumber(399)).isLessThanOrEqualTo(_248_v6)) && ((_248_v6).isLessThan(new BigNumber(-703)))) {
              _coll20.push([_module.__default.safeModuloInt(_248_v6, new BigNumber(668)),_246_v0]);
            }
          }
          return _coll20;
        }();
      })(_239_v0);
      let _nw37 = Array((new BigNumber(10)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw37.length); _i0_4++) {
        _nw37[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _245_v7 = _nw37;
      let _249_v8;
      _249_v8 = _dafny.Set.fromElements(_241_v2);
      let _250_v9;
      let _nw38 = Array((new BigNumber(12)).toNumber());
      _nw38[(_dafny.ZERO).toNumber()] = _239_v0;
      _nw38[(_dafny.ONE).toNumber()] = _239_v0;
      _nw38[(new BigNumber(2)).toNumber()] = _239_v0;
      _nw38[(new BigNumber(3)).toNumber()] = _239_v0;
      _nw38[(new BigNumber(4)).toNumber()] = _239_v0;
      _nw38[(new BigNumber(5)).toNumber()] = new BigNumber(44);
      _nw38[(new BigNumber(6)).toNumber()] = _239_v0;
      _nw38[(new BigNumber(7)).toNumber()] = _239_v0;
      _nw38[(new BigNumber(8)).toNumber()] = _239_v0;
      _nw38[(new BigNumber(9)).toNumber()] = _239_v0;
      _nw38[(new BigNumber(10)).toNumber()] = _239_v0;
      _nw38[(new BigNumber(11)).toNumber()] = new BigNumber((_249_v8).length);
      _250_v9 = _nw38;
      let _251_v10;
      let _nw39 = Array((new BigNumber(5)).toNumber()).fill(false);
      _251_v10 = _nw39;
      let _252_v11;
      _252_v11 = _dafny.Map.Empty.slice().updateUnsafe(_250_v9,_251_v10);
      let _253_v12;
      _253_v12 = _dafny.Seq.UnicodeFromString("oalh");
      let _254_v13;
      _254_v13 = _dafny.Seq.of(_250_v9);
      let _255_v14;
      _255_v14 = _dafny.Seq.of(false, _241_v2);
      let _256_v15;
      _256_v15 = _dafny.Map.Empty.slice().updateUnsafe(_239_v0,_241_v2);
      let _257_v16;
      _257_v16 = _dafny.Map.Empty.slice().updateUnsafe(_239_v0,(((_256_v15).contains(new BigNumber((_dafny.Seq.UnicodeFromString("qgoah")).length))) ? ((_256_v15).get(new BigNumber((_dafny.Seq.UnicodeFromString("qgoah")).length))) : (_241_v2)));
      let _258_v17;
      _258_v17 = _dafny.Seq.of(_257_v16);
      let _259_globalState;
      let _nw40 = new _module.GlobalState();
      _nw40.__ctor(new BigNumber(515), _240_v1, _240_v1, _242_v3, new BigNumber(523), _244_v5, new BigNumber(914), _245_v7, _252_v11, _253_v12, _dafny.Map.Empty.slice().updateUnsafe(_251_v10,_241_v2), _dafny.Seq.Concat(_dafny.Seq.of(_250_v9, _250_v9, _250_v9, _250_v9, _250_v9), _254_v13), false, _255_v14, false, new BigNumber(682), _dafny.Seq.Concat(_dafny.Seq.of(_257_v16), _258_v17), true);
      _259_globalState = _nw40;
      let _hi1 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(902), _239_v0));
      for (let _260_i1 = _239_v0; _260_i1.isLessThan(_hi1); _260_i1 = _260_i1.plus(_dafny.ONE)) {
        _241_v2 = (((_241_v2) ? (_260_i1) : ((_dafny.ZERO).minus(_260_i1)))).isLessThan(((_dafny.ZERO).minus(_239_v0)).multipliedBy(_239_v0));
        let _261_v18;
        _261_v18 = _dafny.Seq.of(_253_v12, _253_v12);
        _261_v18 = _dafny.Seq.Concat(_dafny.Seq.of(_253_v12), _261_v18);
        let _262_v19;
        _262_v19 = _dafny.Set.fromElements(_250_v9);
        let _rhs48 = _241_v2;
        let _rhs49 = _241_v2;
        let _rhs50 = _240_v1;
        let _rhs51 = _module.__default.safeDivisionInt(new BigNumber(((_262_v19).Intersect(_262_v19)).length), new BigNumber((_240_v1).length));
        let _lhs44 = _259_globalState;
        _241_v2 = _rhs48;
        _241_v2 = _rhs49;
        _lhs44.f1 = _rhs50;
        _239_v0 = _rhs51;
        let _index49 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_251_v10).length));
        (_251_v10)[_index49] = _241_v2;
        let _263_v20;
        _263_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_259_globalState),((_241_v2) ? (_243_v4) : (_243_v4)));
        let _264_v21;
        _264_v21 = _module.D0.create_DC0(_module.__default.fm0(_259_globalState));
        let _index50 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_251_v10).length));
        let _rhs52 = (((_263_v20).contains(new BigNumber(288))) ? ((_263_v20).get(new BigNumber(288))) : (_243_v4));
        let _rhs53 = ((_264_v21).dtor_cf0).minus(_239_v0);
        let _rhs54 = _241_v2;
        let _rhs55 = ((false) ? (false) : (_241_v2));
        let _rhs56 = (_264_v21).dtor_cf0;
        let _lhs45 = _259_globalState;
        let _lhs46 = _251_v10;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_251_v10).length));
        _243_v4 = _rhs52;
        _lhs45.f6 = _rhs53;
        _lhs46[_lhs47] = _rhs54;
        _241_v2 = _rhs55;
        _239_v0 = _rhs56;
      }
      let _265_i2;
      _265_i2 = _dafny.ZERO;
      L0: {
        while (_241_v2) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_265_i2)) {
              break L0;
            }
            _265_i2 = (_265_i2).plus(_dafny.ONE);
            _242_v3 = (_242_v3).update(!((new BigNumber(904)).isLessThanOrEqualTo(_239_v0)), _241_v2);
            let _rhs57 = _243_v4;
            let _rhs58 = _239_v0;
            let _lhs48 = _259_globalState;
            _243_v4 = _rhs57;
            _lhs48.f6 = _rhs58;
            let _266_v22;
            _266_v22 = _dafny.Map.Empty.slice().updateUnsafe(_251_v10,true);
            (_259_globalState).f6 = ((_241_v2) ? (_239_v0) : (((_dafny.ZERO).minus(new BigNumber(-533))).multipliedBy(new BigNumber((_266_v22).length))));
            let _index51 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_250_v9).length));
            (_250_v9)[_index51] = _239_v0;
            let _index52 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_250_v9).length));
            (_250_v9)[_index52] = (new BigNumber(-29)).minus(_239_v0);
          }
        }
      }
      let _index53 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length));
      (_250_v9)[_index53] = _239_v0;
      let _index54 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length));
      (_250_v9)[_index54] = _module.__default.safeModuloInt(_239_v0, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("nxertem")).length)));
      let _267_v23;
      let _268_v24;
      let _269_v25;
      let _out0;
      let _out1;
      let _out2;
      let _outcollector0 = _module.__default.m0((_dafny.ZERO).minus((_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))]), _module.__default.fm1((((_256_v15).contains(new BigNumber((_dafny.Seq.UnicodeFromString("mjgtu")).length))) ? ((_256_v15).get(new BigNumber((_dafny.Seq.UnicodeFromString("mjgtu")).length))) : (_241_v2)), _243_v4, new BigNumber(-261), _259_globalState), (new BigNumber((_dafny.Seq.update(_253_v12, _module.__default.safeIndex(_239_v0, new BigNumber((_253_v12).length)), _243_v4)).length)).multipliedBy((_dafny.ZERO).minus(_239_v0)), _259_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _out2 = _outcollector0[2];
      _267_v23 = _out0;
      _268_v24 = _out1;
      _269_v25 = _out2;
      _253_v12 = _253_v12;
      let _270_v26;
      _270_v26 = _dafny.Map.Empty.slice().updateUnsafe(_268_v24,_255_v14);
      let _index55 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length));
      let _rhs59 = _module.__default.fm0(_259_globalState);
      let _rhs60 = _dafny.Seq.Concat((((_270_v26).contains(_268_v24)) ? ((_270_v26).get(_268_v24)) : (_255_v14)), _255_v14);
      let _lhs49 = _250_v9;
      let _lhs50 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length));
      _lhs49[_lhs50] = _rhs59;
      _255_v14 = _rhs60;
      _268_v24 = _241_v2;
      let _271_v27;
      let _init5 = ((_272_v4) => function (_273_i3) {
        return _272_v4;
      })(_243_v4);
      let _nw41 = Array((_dafny.ONE).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw41.length); _i0_5++) {
        _nw41[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _271_v27 = _nw41;
      let _274_v28;
      _274_v28 = _dafny.Map.Empty.slice().updateUnsafe((_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))],_271_v27);
      if ((_dafny.Set.fromElements(_module.__default.fm1(_241_v2, _243_v4, new BigNumber((_274_v28).length), _259_globalState), _268_v24, (((_242_v3).contains(_268_v24)) ? ((_242_v3).get(_268_v24)) : (_268_v24)))).IsSubsetOf(_module.__default.fm2(_259_globalState))) {
        let _275_v29;
        _275_v29 = _module.D0.create_DC1();
        _275_v29 = _275_v29;
        let _276_v30;
        _276_v30 = _dafny.Map.Empty.slice().updateUnsafe(_241_v2,new BigNumber(531));
        let _277_v31;
        let _nw42 = new _module.C3();
        _nw42.__ctor(_module.__default.fm20(false, _259_globalState), _dafny.Map.Empty.slice().updateUnsafe(_253_v12,new BigNumber((_276_v30).length)));
        _277_v31 = _nw42;
        (_259_globalState).f6 = (new BigNumber((function () {
          let _coll21 = new _dafny.Map();
          for (const _compr_21 of _dafny.IntegerRange(new BigNumber(584), new BigNumber(334))) {
            let _278_v32 = _compr_21;
            if (((new BigNumber(584)).isLessThanOrEqualTo(_278_v32)) && ((_278_v32).isLessThan(new BigNumber(334)))) {
              _coll21.push([(_278_v32).minus(_239_v0),_module.D5.create_DC12(_dafny.MultiSet.fromElements(_241_v2))]);
            }
          }
          return _coll21;
        }()).length)).multipliedBy((_269_v25).plus((_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))]));
        let _index56 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_267_v23).length));
        (_267_v23)[_index56] = !(((_268_v24) ? (_268_v24) : (_241_v2)));
        let _279_v33;
        _279_v33 = _dafny.Set.fromElements(_269_v25);
        let _280_v34;
        _280_v34 = _dafny.Seq.of(_279_v33);
        let _281_v35;
        _281_v35 = _dafny.Map.Empty.slice().updateUnsafe(_239_v0,_280_v34);
        let _index57 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_267_v23).length));
        let _rhs61 = !(_268_v24);
        let _rhs62 = _241_v2;
        let _rhs63 = _module.__default.fm25(_242_v3, _module.D9.create_DC23((((_281_v35).contains(_239_v0)) ? ((_281_v35).get(_239_v0)) : (_280_v34)), (_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))], _241_v2, _239_v0), _module.__default.fm1(false, _243_v4, _239_v0, _259_globalState), (_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))], _259_globalState);
        let _rhs64 = !(_241_v2) || ((_269_v25).isLessThan(new BigNumber((_module.__default.fm26(_259_globalState)).length)));
        let _rhs65 = !(_268_v24);
        let _lhs51 = _267_v23;
        let _lhs52 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_267_v23).length));
        _lhs51[_lhs52] = _rhs61;
        _241_v2 = _rhs62;
        _275_v29 = _rhs63;
        _241_v2 = _rhs64;
        _268_v24 = _rhs65;
        _241_v2 = !(_269_v25).isEqualTo(_269_v25);
      } else {
        let _282_v36;
        let _nw43 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _282_v36 = _nw43;
        let _283_v37;
        _283_v37 = _dafny.Map.Empty.slice().updateUnsafe(_282_v36,_251_v10);
        _283_v37 = (_283_v37).update(_282_v36, _267_v23);
        let _284_v38;
        let _init6 = ((_285_v14) => function (_286_i4) {
          return _dafny.Seq.Concat(_285_v14, _285_v14);
        })(_255_v14);
        let _nw44 = Array((new BigNumber(16)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw44.length); _i0_6++) {
          _nw44[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _284_v38 = _nw44;
        _284_v38 = _284_v38;
        let _287_v39;
        let _nw45 = Array((new BigNumber(19)).toNumber()).fill([]);
        _287_v39 = _nw45;
        let _288_v40;
        _288_v40 = _module.D6.create_DC16(_241_v2, _287_v39);
        let _source3 = _288_v40;
        if (_source3.is_DC16) {
          let _289___mcc_h0 = (_source3).cf19;
          let _290___mcc_h1 = (_source3).cf20;
          let _291_cf20 = _290___mcc_h1;
          let _292_cf19 = _289___mcc_h0;
          let _293_v41;
          _293_v41 = _module.D7.create_DC18(_249_v8);
          let _index58 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length));
          (_250_v9)[_index58] = (_269_v25).minus((new BigNumber(((_293_v41).dtor_cf22).length)).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_292_cf19,_269_v25)).length)));
          _292_cf19 = _268_v24;
          let _294_v42;
          _294_v42 = _dafny.MultiSet.fromElements((_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))]);
          let _index59 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_251_v10).length));
          (_251_v10)[_index59] = !(!(_294_v42).contains(new BigNumber((_253_v12).length)));
          let _index60 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_251_v10).length));
          (_251_v10)[_index60] = _241_v2;
          _239_v0 = _module.__default.fm0(_259_globalState);
        } else if (_source3.is_DC15) {
          let _295___mcc_h2 = (_source3).cf18;
          let _296_cf18 = _295___mcc_h2;
          let _297_v43;
          let _298_v44;
          let _299_v45;
          let _out3;
          let _out4;
          let _out5;
          let _outcollector1 = _module.__default.m0(new BigNumber((_253_v12).length), _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(158)), function (_300_i5) {
            return new BigNumber(629);
          }), _module.__default.fm0(_259_globalState)), new BigNumber((_253_v12).length), _259_globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _out5 = _outcollector1[2];
          _297_v43 = _out3;
          _298_v44 = _out4;
          _299_v45 = _out5;
          let _301_v46;
          let _nw46 = new _module.C3();
          _nw46.__ctor(_243_v4, _dafny.Map.Empty.slice().updateUnsafe(_253_v12,_239_v0));
          _301_v46 = _nw46;
          _301_v46 = _301_v46;
          let _302_v47;
          _302_v47 = _dafny.MultiSet.fromElements(_298_v44, true);
          let _303_v48;
          _303_v48 = _module.D3.create_DC10((_dafny.MultiSet.FromArray(_dafny.Seq.of(_241_v2))).IsProperSubsetOf(_302_v47));
          _303_v48 = _303_v48;
          let _304_v49;
          let _nw47 = Array((new BigNumber(10)).toNumber()).fill([]);
          _304_v49 = _nw47;
          let _305_v50;
          _305_v50 = _dafny.Map.Empty.slice().updateUnsafe(_241_v2,_282_v36);
          let _index61 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_304_v49).length));
          (_304_v49)[_index61] = (((_305_v50).contains(_298_v44)) ? ((_305_v50).get(_298_v44)) : (_282_v36));
          let _index62 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_304_v49).length));
          (_304_v49)[_index62] = _282_v36;
        } else {
          let _306___mcc_h3 = (_source3).cf21;
          let _307_cf21 = _306___mcc_h3;
          let _308_v51;
          _308_v51 = _dafny.Map.Empty.slice().updateUnsafe(_253_v12,new BigNumber(((_module.D9.create_DC22(_dafny.Seq.Create(_module.__default.abs(new BigNumber(932)), ((_309_v4) => function (_310_i6) {
  return _309_v4;
})(_243_v4)))).dtor_cf27).length));
          let _311_v52;
          let _nw48 = new _module.C3();
          _nw48.__ctor(_243_v4, _308_v51);
          _311_v52 = _nw48;
          _311_v52 = _311_v52;
          let _312_v53;
          _312_v53 = _dafny.Map.Empty.slice().updateUnsafe(_268_v24,((_241_v2) ? (_250_v9) : ((_254_v13)[_module.__default.safeIndex(_239_v0, new BigNumber((_254_v13).length))])));
          _250_v9 = (((_312_v53).contains((_311_v52).fm3(_241_v2, (_dafny.ZERO).minus(_239_v0), _259_globalState))) ? ((_312_v53).get((_311_v52).fm3(_241_v2, (_dafny.ZERO).minus(_239_v0), _259_globalState))) : (_250_v9));
          let _313_v54;
          let _nw49 = new _module.C3();
          _nw49.__ctor(_243_v4, (_308_v51).Merge(_308_v51));
          _313_v54 = _nw49;
          let _314_v55;
          let _315_v56;
          let _316_v57;
          let _317_v58;
          let _out6;
          let _out7;
          let _out8;
          let _out9;
          let _outcollector2 = (_311_v52).m1(_module.__default.fm1(_241_v2, new _dafny.CodePoint('m'.codePointAt(0)), new BigNumber((_312_v53).length), _259_globalState), _259_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _out8 = _outcollector2[2];
          _out9 = _outcollector2[3];
          _314_v55 = _out6;
          _315_v56 = _out7;
          _316_v57 = _out8;
          _317_v58 = _out9;
        }
        let _318_v59;
        _318_v59 = _dafny.Map.Empty.slice().updateUnsafe(_253_v12,_239_v0);
        let _319_v60;
        let _nw50 = new _module.C1();
        _nw50.__ctor((_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))], _268_v24, _243_v4, _318_v59);
        _319_v60 = _nw50;
        let _320_v61;
        _320_v61 = _dafny.Map.Empty.slice().updateUnsafe(_319_v60,(_dafny.ZERO).minus(_module.__default.safeModuloInt(_239_v0, _269_v25)));
        _320_v61 = (_320_v61).update(_319_v60, _269_v25);
        let _321_v62;
        _321_v62 = _dafny.MultiSet.fromElements(_268_v24, true);
        let _322_v63;
        _322_v63 = _dafny.Map.Empty.slice().updateUnsafe(_243_v4,(_dafny.MultiSet.FromArray(_dafny.Seq.of((_319_v60).f23))).Intersect(_321_v62));
        _322_v63 = (_322_v63).update(_243_v4, _321_v62);
      }
      let _323_v64;
      let _nw51 = Array((new BigNumber(18)).toNumber()).fill(_module.D7.Default());
      _323_v64 = _nw51;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_323_v64).length))) {
        let _324_i7 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_324_i7)) && ((_324_i7).isLessThan(new BigNumber((_323_v64).length))))) {
          (_323_v64)[(_324_i7)] = _module.D7.create_DC20(_241_v2, _241_v2);
        }
      }
      let _325_v65;
      _325_v65 = _dafny.Map.Empty.slice().updateUnsafe(_269_v25,new BigNumber(458));
      let _326_i8;
      _326_i8 = _dafny.ZERO;
      L1: {
        while ((_269_v25).isLessThanOrEqualTo(((_dafny.ZERO).minus(new BigNumber((_325_v65).length))).multipliedBy(new BigNumber((_253_v12).length)))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_326_i8)) {
              break L1;
            }
            _326_i8 = (_326_i8).plus(_dafny.ONE);
            _269_v25 = new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_240_v1, _module.__default.safeIndex(new BigNumber(104), new BigNumber((_240_v1).length)), (_dafny.ZERO).minus(_269_v25)), _module.__default.safeIndex(_module.__default.safeDivisionInt(_269_v25, _239_v0), new BigNumber((_dafny.Seq.update(_240_v1, _module.__default.safeIndex(new BigNumber(104), new BigNumber((_240_v1).length)), (_dafny.ZERO).minus(_269_v25))).length)), new BigNumber(123))).length);
            let _327_v66;
            _327_v66 = _dafny.Set.fromElements((_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))], (_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))]);
            let _328_v67;
            _328_v67 = _dafny.Seq.of(_327_v66, _327_v66);
            _328_v67 = ((_241_v2) ? (_dafny.Seq.of(_327_v66, _327_v66, _327_v66, _327_v66, _dafny.Set.fromElements(_239_v0))) : (_dafny.Seq.of(function () {
              let _coll22 = new _dafny.Set();
              for (const _compr_22 of _dafny.IntegerRange(new BigNumber(-240), new BigNumber(336))) {
                let _329_v68 = _compr_22;
                if (((new BigNumber(-240)).isLessThanOrEqualTo(_329_v68)) && ((_329_v68).isLessThan(new BigNumber(336)))) {
                  _coll22.add((_329_v68).minus(new BigNumber((_256_v15).length)));
                }
              }
              return _coll22;
            }(), _dafny.Set.fromElements(new BigNumber(90)), _327_v66, _327_v66, _327_v66)));
            let _330_v70;
            _330_v70 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), ((_331_v4) => function (_332_i9) {
              return _331_v4;
            })(_243_v4)));
            let _333_v71;
            let _nw52 = new _module.C2();
            _nw52.__ctor(_243_v4, function () {
              let _coll23 = new _dafny.Map();
              for (const _compr_23 of (_330_v70).Elements) {
                let _334_v69 = _compr_23;
                if (_dafny.Seq.contains(_330_v70, _334_v69)) {
                  _coll23.push([_334_v69,(_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))]]);
                }
              }
              return _coll23;
            }());
            _333_v71 = _nw52;
            _325_v65 = (_325_v65).update((_dafny.ZERO).minus((_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))]), _module.__default.safeDivisionInt(new BigNumber(378), _269_v25));
          }
        }
      }
      let _335_v72;
      let _nw53 = Array((new BigNumber(27)).toNumber());
      _335_v72 = _nw53;
      let _336_v73;
      _336_v73 = _dafny.Map.Empty.slice().updateUnsafe(_335_v72,_240_v1);
      _336_v73 = (_336_v73).update(_335_v72, _240_v1);
      let _337_v74;
      _337_v74 = _module.D1.create_DC6(_dafny.Seq.IsPrefixOf(_253_v12, _253_v12), _268_v24);
      let _source4 = _337_v74;
      if (_source4.is_DC4) {
        let _338___mcc_h4 = (_source4).cf3;
        let _339___mcc_h5 = (_source4).cf4;
        let _340___mcc_h6 = (_source4).cf5;
        let _341_cf5 = _340___mcc_h6;
        let _342_cf4 = _339___mcc_h5;
        let _343_cf3 = _338___mcc_h4;
        let _344_v75;
        _344_v75 = _dafny.Map.Empty.slice().updateUnsafe(_253_v12,_342_cf4);
        let _345_v76;
        let _nw54 = new _module.C0();
        _nw54.__ctor(_dafny.Seq.UnicodeFromString("ikrmgew"), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),(_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))]), _243_v4, _344_v75);
        _345_v76 = _nw54;
        let _346_v77;
        _346_v77 = _dafny.MultiSet.fromElements(_345_v76);
        let _347_v78;
        _347_v78 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_345_v76), _dafny.MultiSet.fromElements(_345_v76), (_346_v77).update(_345_v76, _module.__default.abs(_module.__default.fm0(_259_globalState))), _dafny.MultiSet.fromElements(_345_v76, _345_v76));
        _347_v78 = _347_v78;
        let _348_v79;
        _348_v79 = _dafny.Set.fromElements(new BigNumber(788), _269_v25, new BigNumber(979), _343_cf3, (_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))]);
        let _349_v80;
        _349_v80 = _dafny.Seq.of(_348_v79, _348_v79);
        let _350_v81;
        let _351_v82;
        let _352_v83;
        let _out10;
        let _out11;
        let _out12;
        let _outcollector3 = _module.__default.m0((_dafny.ZERO).minus((_module.D9.create_DC23(_349_v80, _269_v25, _241_v2, new BigNumber(277))).dtor_cf31), _241_v2, _239_v0, _259_globalState);
        _out10 = _outcollector3[0];
        _out11 = _outcollector3[1];
        _out12 = _outcollector3[2];
        _350_v81 = _out10;
        _351_v82 = _out11;
        _352_v83 = _out12;
        if (_241_v2) {
          _351_v82 = !((_341_cf5) === (_module.__default.fm1(_268_v24, _345_v76.f18, _239_v0, _259_globalState)));
          (_259_globalState).f6 = _239_v0;
          let _353_v86;
          _353_v86 = _dafny.Seq.of(_251_v10);
          let _354_v87;
          let _nw55 = Array((new BigNumber(29)).toNumber());
          _nw55[(_dafny.ZERO).toNumber()] = _350_v81;
          _nw55[(_dafny.ONE).toNumber()] = _350_v81;
          _nw55[(new BigNumber(2)).toNumber()] = _267_v23;
          _nw55[(new BigNumber(3)).toNumber()] = _251_v10;
          _nw55[(new BigNumber(4)).toNumber()] = _350_v81;
          _nw55[(new BigNumber(5)).toNumber()] = _350_v81;
          _nw55[(new BigNumber(6)).toNumber()] = _251_v10;
          _nw55[(new BigNumber(7)).toNumber()] = _251_v10;
          _nw55[(new BigNumber(8)).toNumber()] = _350_v81;
          _nw55[(new BigNumber(9)).toNumber()] = _350_v81;
          _nw55[(new BigNumber(10)).toNumber()] = _350_v81;
          _nw55[(new BigNumber(11)).toNumber()] = _251_v10;
          _nw55[(new BigNumber(12)).toNumber()] = _267_v23;
          _nw55[(new BigNumber(13)).toNumber()] = _267_v23;
          _nw55[(new BigNumber(14)).toNumber()] = _267_v23;
          _nw55[(new BigNumber(15)).toNumber()] = _350_v81;
          _nw55[(new BigNumber(16)).toNumber()] = _350_v81;
          _nw55[(new BigNumber(17)).toNumber()] = _251_v10;
          _nw55[(new BigNumber(18)).toNumber()] = _251_v10;
          _nw55[(new BigNumber(19)).toNumber()] = _267_v23;
          _nw55[(new BigNumber(20)).toNumber()] = _267_v23;
          _nw55[(new BigNumber(21)).toNumber()] = _251_v10;
          _nw55[(new BigNumber(22)).toNumber()] = _267_v23;
          _nw55[(new BigNumber(23)).toNumber()] = _267_v23;
          _nw55[(new BigNumber(24)).toNumber()] = _251_v10;
          _nw55[(new BigNumber(25)).toNumber()] = _267_v23;
          _nw55[(new BigNumber(26)).toNumber()] = _267_v23;
          _nw55[(new BigNumber(27)).toNumber()] = _350_v81;
          _nw55[(new BigNumber(28)).toNumber()] = (_353_v86)[_module.__default.safeIndex(new BigNumber((_249_v8).length), new BigNumber((_353_v86).length))];
          _354_v87 = _nw55;
          let _355_v88;
          _355_v88 = _module.D6.create_DC16(_341_cf5, _354_v87);
          let _356_v89;
          _356_v89 = _module.D6.create_DC17(_355_v88);
          let _357_v90;
          _357_v90 = _dafny.Map.Empty.slice().updateUnsafe(_356_v89,_341_cf5);
          let _358_v93;
          let _nw56 = Array((new BigNumber(18)).toNumber());
          _nw56[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_253_v12).length)),_241_v2);
          _nw56[(_dafny.ONE).toNumber()] = _256_v15;
          _nw56[(new BigNumber(2)).toNumber()] = _257_v16;
          _nw56[(new BigNumber(3)).toNumber()] = _257_v16;
          _nw56[(new BigNumber(4)).toNumber()] = _257_v16;
          _nw56[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_249_v8).length),_module.__default.fm1(true, _243_v4, _352_v83, _259_globalState));
          _nw56[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_342_cf4,true);
          _nw56[(new BigNumber(7)).toNumber()] = function () {
            let _coll24 = new _dafny.Map();
            for (const _compr_24 of _dafny.IntegerRange(new BigNumber(919), new BigNumber(570))) {
              let _359_v84 = _compr_24;
              if (((new BigNumber(919)).isLessThanOrEqualTo(_359_v84)) && ((_359_v84).isLessThan(new BigNumber(570)))) {
                _coll24.push([_module.__default.safeModuloInt(_359_v84, _239_v0),_341_cf5]);
              }
            }
            return _coll24;
          }();
          _nw56[(new BigNumber(8)).toNumber()] = _257_v16;
          _nw56[(new BigNumber(9)).toNumber()] = _256_v15;
          _nw56[(new BigNumber(10)).toNumber()] = _257_v16;
          _nw56[(new BigNumber(11)).toNumber()] = (function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of _dafny.IntegerRange(new BigNumber(699), new BigNumber(764))) {
              let _360_v85 = _compr_25;
              if (((new BigNumber(699)).isLessThanOrEqualTo(_360_v85)) && ((_360_v85).isLessThan(new BigNumber(764)))) {
                _coll25.push([(_360_v85).plus(_269_v25),_341_cf5]);
              }
            }
            return _coll25;
          }()).Merge(_256_v15);
          _nw56[(new BigNumber(12)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_239_v0,!((((_357_v90).contains(_356_v89)) ? ((_357_v90).get(_356_v89)) : (_241_v2))))).Merge(_257_v16);
          _nw56[(new BigNumber(13)).toNumber()] = function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of _dafny.IntegerRange(new BigNumber(58), new BigNumber(424))) {
              let _361_v91 = _compr_26;
              if (((new BigNumber(58)).isLessThanOrEqualTo(_361_v91)) && ((_361_v91).isLessThan(new BigNumber(424)))) {
                _coll26.push([(_361_v91).minus(new BigNumber(335)),true]);
              }
            }
            return _coll26;
          }();
          _nw56[(new BigNumber(14)).toNumber()] = _256_v15;
          _nw56[(new BigNumber(15)).toNumber()] = _257_v16;
          _nw56[(new BigNumber(16)).toNumber()] = function () {
            let _coll27 = new _dafny.Map();
            for (const _compr_27 of _dafny.IntegerRange(new BigNumber(230), new BigNumber(64))) {
              let _362_v92 = _compr_27;
              if (((new BigNumber(230)).isLessThanOrEqualTo(_362_v92)) && ((_362_v92).isLessThan(new BigNumber(64)))) {
                _coll27.push([(_362_v92).minus(_239_v0),_268_v24]);
              }
            }
            return _coll27;
          }();
          _nw56[(new BigNumber(17)).toNumber()] = _257_v16;
          _358_v93 = _nw56;
          let _index63 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_358_v93).length));
          (_358_v93)[_index63] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_253_v12).length),!(_268_v24));
          let _index64 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_358_v93).length));
          (_358_v93)[_index64] = ((_module.__default.fm1(_351_v82, _345_v76.f18, (_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))], _259_globalState)) ? (_256_v15) : (_257_v16));
          let _363_v94;
          _363_v94 = _module.D10.create_DC25(_250_v9);
          let _364_v95;
          let _nw57 = Array((new BigNumber(9)).toNumber());
          _nw57[(_dafny.ZERO).toNumber()] = _250_v9;
          _nw57[(_dafny.ONE).toNumber()] = _250_v9;
          _nw57[(new BigNumber(2)).toNumber()] = _250_v9;
          _nw57[(new BigNumber(3)).toNumber()] = (_363_v94).dtor_cf33;
          _nw57[(new BigNumber(4)).toNumber()] = _250_v9;
          _nw57[(new BigNumber(5)).toNumber()] = _250_v9;
          _nw57[(new BigNumber(6)).toNumber()] = _250_v9;
          _nw57[(new BigNumber(7)).toNumber()] = (_254_v13)[_module.__default.safeIndex(_352_v83, new BigNumber((_254_v13).length))];
          _nw57[(new BigNumber(8)).toNumber()] = ((_268_v24) ? (_250_v9) : (_250_v9));
          _364_v95 = _nw57;
          let _index65 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_364_v95).length));
          (_364_v95)[_index65] = _250_v9;
          let _index66 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_364_v95).length));
          let _index67 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length));
          let _rhs66 = _240_v1;
          let _rhs67 = _250_v9;
          let _rhs68 = _269_v25;
          let _lhs53 = _259_globalState;
          let _lhs54 = _364_v95;
          let _lhs55 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_364_v95).length));
          let _lhs56 = _250_v9;
          let _lhs57 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length));
          _lhs53.f1 = _rhs66;
          _lhs54[_lhs55] = _rhs67;
          _lhs56[_lhs57] = _rhs68;
          _268_v24 = _351_v82;
        } else {
          let _365_v96;
          let _366_v97;
          let _367_v98;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector4 = _module.__default.m0(_343_cf3, (_269_v25).isEqualTo(new BigNumber((_module.__default.fm2(_259_globalState)).length)), new BigNumber((_348_v79).length), _259_globalState);
          _out13 = _outcollector4[0];
          _out14 = _outcollector4[1];
          _out15 = _outcollector4[2];
          _365_v96 = _out13;
          _366_v97 = _out14;
          _367_v98 = _out15;
          _268_v24 = false;
          let _368_v99;
          let _nw58 = new _module.C3();
          _nw58.__ctor(_345_v76.f18, _dafny.Map.Empty.slice().updateUnsafe(_253_v12,_269_v25));
          _368_v99 = _nw58;
          let _369_v100;
          let _nw59 = new _module.C2();
          _nw59.__ctor(_345_v76.f18, _344_v75);
          _369_v100 = _nw59;
          let _370_v101;
          _370_v101 = _dafny.Map.Empty.slice().updateUnsafe(_368_v99,_369_v100);
          let _371_v102;
          _371_v102 = _dafny.MultiSet.fromElements((_370_v101).Merge(_370_v101));
          let _372_v103;
          _372_v103 = _dafny.Seq.of(_370_v101);
          _371_v102 = (_dafny.MultiSet.FromArray(_372_v103)).Difference(_371_v102);
          _352_v83 = _239_v0;
          _239_v0 = _367_v98;
        }
        _243_v4 = _243_v4;
      } else if (_source4.is_DC5) {
        let _373___mcc_h7 = (_source4).cf6;
        let _374_cf6 = _373___mcc_h7;
        _249_v8 = (((!(_268_v24)) && (_241_v2)) ? (((_241_v2) ? (_module.__default.fm2(_259_globalState)) : (_249_v8))) : (_249_v8));
        let _375_v104;
        let _nw60 = new _module.C0();
        _nw60.__ctor(_253_v12, _244_v5, _243_v4, _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vf"),(_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))]));
        _375_v104 = _nw60;
        _375_v104 = _375_v104;
        let _376_v105;
        let _nw61 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
        _376_v105 = _nw61;
        _376_v105 = _376_v105;
        let _377_v106;
        _377_v106 = _dafny.MultiSet.fromElements((((_242_v3).contains(false)) ? ((_242_v3).get(false)) : (_241_v2)), !(_241_v2));
        let _378_v107;
        _378_v107 = _module.D1.create_DC4((_dafny.ZERO).minus(new BigNumber((_377_v106).cardinality())), (_dafny.ZERO).minus(_269_v25), _241_v2);
        let _379_v108;
        _379_v108 = _dafny.Seq.of(_249_v8);
        let _index68 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length));
        let _rhs69 = (_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))];
        let _rhs70 = _module.D1.create_DC4(_239_v0, _269_v25, true);
        let _rhs71 = _379_v108;
        let _rhs72 = _module.__default.fm27(_259_globalState);
        let _lhs58 = _250_v9;
        let _lhs59 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length));
        let _lhs60 = _259_globalState;
        _lhs58[_lhs59] = _rhs69;
        _378_v107 = _rhs70;
        _379_v108 = _rhs71;
        _lhs60.f13 = _rhs72;
      } else if (_source4.is_DC6) {
        let _380___mcc_h8 = (_source4).cf7;
        let _381___mcc_h9 = (_source4).cf8;
        let _382_cf8 = _381___mcc_h9;
        let _383_cf7 = _380___mcc_h8;
        let _384_v109;
        let _385_v110;
        let _386_v111;
        let _out16;
        let _out17;
        let _out18;
        let _outcollector5 = _module.__default.m0(_239_v0, (_383_cf7) && (_241_v2), _239_v0, _259_globalState);
        _out16 = _outcollector5[0];
        _out17 = _outcollector5[1];
        _out18 = _outcollector5[2];
        _384_v109 = _out16;
        _385_v110 = _out17;
        _386_v111 = _out18;
        let _387_v112;
        _387_v112 = _dafny.Map.Empty.slice().updateUnsafe(_253_v12,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_269_v25,_383_cf7)).length)));
        let _388_v113;
        _388_v113 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_269_v25),_387_v112);
        let _389_v114;
        let _nw62 = new _module.C1();
        _nw62.__ctor((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_269_v25, _386_v111)), true, _243_v4, (((_388_v113).contains(_239_v0)) ? ((_388_v113).get(_239_v0)) : (_387_v112)));
        _389_v114 = _nw62;
        let _390_v115;
        let _out19;
        _out19 = (_389_v114).m2(!(_382_cf8), _389_v114.f22, (_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))], _386_v111, _259_globalState);
        _390_v115 = _out19;
        let _391_v116;
        let _nw63 = new _module.C0();
        _nw63.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(102)), ((_392_v4) => function (_393_i10) {
          return _392_v4;
        })(_243_v4)), _244_v5, _243_v4, _387_v112);
        _391_v116 = _nw63;
      } else {
        let _394___mcc_h10 = (_source4).cf2;
        let _395_cf2 = _394___mcc_h10;
        let _396_v117;
        _396_v117 = _dafny.Map.Empty.slice().updateUnsafe(_253_v12,(_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))]);
        let _397_v118;
        _397_v118 = _dafny.MultiSet.fromElements(_243_v4);
        let _398_v119;
        let _nw64 = new _module.C0();
        _nw64.__ctor(_dafny.Seq.Concat(_253_v12, _253_v12), _244_v5, _243_v4, (_396_v117).Merge((_dafny.Map.Empty.slice().updateUnsafe(_253_v12,new BigNumber(((_397_v118).update(_243_v4, _module.__default.abs(_269_v25))).cardinality()))).update(_dafny.Seq.UnicodeFromString("rtoijma"), (_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))])));
        _398_v119 = _nw64;
        let _399_v120;
        _399_v120 = _module.D5.create_DC14(_268_v24, (_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))]);
        let _400_v121;
        _400_v121 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_399_v120,_241_v2),_269_v25);
        let _401_v122;
        _401_v122 = _dafny.Map.Empty.slice().updateUnsafe(_399_v120,_268_v24);
        _400_v121 = _dafny.Map.Empty.slice().updateUnsafe(_401_v122,(((_325_v65).contains((_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))])) ? ((_325_v65).get((_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))])) : (_269_v25)));
        let _rhs73 = _243_v4;
        let _rhs74 = (((_242_v3).contains((_255_v14)[_module.__default.safeIndex(_269_v25, new BigNumber((_255_v14).length))])) ? ((_242_v3).get((_255_v14)[_module.__default.safeIndex(_269_v25, new BigNumber((_255_v14).length))])) : (_268_v24));
        _243_v4 = _rhs73;
        _395_cf2 = _rhs74;
        let _402_v123;
        let _nw65 = Array((new BigNumber(12)).toNumber());
        _nw65[(_dafny.ZERO).toNumber()] = _253_v12;
        _nw65[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(416)), ((_403_v4) => function (_404_i11) {
          return _403_v4;
        })(_243_v4));
        _nw65[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("p");
        _nw65[(new BigNumber(3)).toNumber()] = _253_v12;
        _nw65[(new BigNumber(4)).toNumber()] = _module.__default.fm8(_259_globalState);
        _nw65[(new BigNumber(5)).toNumber()] = (_398_v119).f20;
        _nw65[(new BigNumber(6)).toNumber()] = ((!(_268_v24)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(49)), ((_405_v4) => function (_406_i12) {
          return _405_v4;
        })(_243_v4))) : (_dafny.Seq.UnicodeFromString("f")));
        _nw65[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(466)), ((_407_v4) => function (_408_i13) {
          return _407_v4;
        })(_243_v4));
        _nw65[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("l"), (_398_v119).f20);
        _nw65[(new BigNumber(9)).toNumber()] = _253_v12;
        _nw65[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("art");
        _nw65[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("fhfctskoe");
        _402_v123 = _nw65;
        let _index69 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_402_v123).length));
        (_402_v123)[_index69] = _dafny.Seq.Concat(_253_v12, _dafny.Seq.update(_253_v12, _module.__default.safeIndex(_239_v0, new BigNumber((_253_v12).length)), _243_v4));
        let _index70 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_402_v123).length));
        (_402_v123)[_index70] = _253_v12;
      }
      let _409_i14;
      _409_i14 = _dafny.ZERO;
      L2: {
        while (false) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_409_i14)) {
              break L2;
            }
            _409_i14 = (_409_i14).plus(_dafny.ONE);
            let _410_v124;
            let _nw66 = new _module.C0();
            _nw66.__ctor(_253_v12, _244_v5, _243_v4, _dafny.Map.Empty.slice().updateUnsafe(_253_v12,new BigNumber((_dafny.Set.fromElements(new BigNumber((_255_v14).length))).length)));
            _410_v124 = _nw66;
            _410_v124 = _410_v124;
            let _411_v125;
            _411_v125 = _dafny.Map.Empty.slice().updateUnsafe(_253_v12,new BigNumber((_253_v12).length));
            let _412_v126;
            let _nw67 = new _module.C0();
            _nw67.__ctor(_253_v12, (_410_v124).f21, ((_268_v24) ? (_243_v4) : (_243_v4)), (_411_v125).Merge(_411_v125));
            _412_v126 = _nw67;
            let _413_v127;
            let _nw68 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.of());
            _413_v127 = _nw68;
            let _index71 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_413_v127).length));
            (_413_v127)[_index71] = _dafny.Seq.Concat(_dafny.Seq.of(_268_v24), _255_v14);
            let _index72 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_413_v127).length));
            (_413_v127)[_index72] = _255_v14;
            let _414_v128;
            _414_v128 = _dafny.Map.Empty.slice().updateUnsafe(_250_v9,_411_v125);
            let _415_v129;
            let _nw69 = new _module.C3();
            _nw69.__ctor(_243_v4, ((((_414_v128).contains(_250_v9)) ? ((_414_v128).get(_250_v9)) : (_411_v125))).Merge(_411_v125));
            _415_v129 = _nw69;
          }
        }
      }
      let _416_v130;
      _416_v130 = _dafny.Map.Empty.slice().updateUnsafe(_268_v24,_239_v0);
      let _hi2 = _module.__default.fm0(_259_globalState);
      for (let _417_i15 = new BigNumber((_416_v130).length); _417_i15.isLessThan(_hi2); _417_i15 = _417_i15.plus(_dafny.ONE)) {
        _241_v2 = (((_417_i15).isLessThanOrEqualTo(_239_v0)) ? (!(_241_v2)) : (_268_v24));
        let _418_v131;
        _418_v131 = _dafny.Map.Empty.slice().updateUnsafe(_417_i15,_dafny.Map.Empty.slice().updateUnsafe((((_416_v130).contains(_241_v2)) ? ((_416_v130).get(_241_v2)) : (_417_i15)),new BigNumber((_249_v8).length)));
        _418_v131 = _418_v131;
        _242_v3 = _242_v3;
        _250_v9 = _250_v9;
      }
      _269_v25 = _239_v0;
      if (_268_v24) {
        let _419_v132;
        _419_v132 = _module.D3.create_DC8(_module.__default.fm20(_268_v24, _259_globalState));
        let _420_v134;
        _420_v134 = _dafny.Seq.of(_253_v12);
        let _421_v135;
        let _nw70 = new _module.C3();
        _nw70.__ctor((_419_v132).dtor_cf10, function () {
          let _coll28 = new _dafny.Map();
          for (const _compr_28 of (_420_v134).Elements) {
            let _422_v133 = _compr_28;
            if (_dafny.Seq.contains(_420_v134, _422_v133)) {
              _coll28.push([_422_v133,_269_v25]);
            }
          }
          return _coll28;
        }());
        _421_v135 = _nw70;
        let _423_v136;
        _423_v136 = _dafny.Map.Empty.slice().updateUnsafe(_421_v135,((_241_v2) ? (_241_v2) : (!((_421_v135).fm3(_268_v24, new BigNumber((_253_v12).length), _259_globalState)))));
        let _424_v137;
        _424_v137 = _dafny.Seq.of(_module.__default.fm28(_241_v2, (_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))], _268_v24, _259_globalState));
        _423_v136 = (_423_v136).update(_421_v135, !(new BigNumber(((_424_v137)[_module.__default.safeIndex(_269_v25, new BigNumber((_424_v137).length))]).cardinality())).isEqualTo(_239_v0));
        (_259_globalState).f6 = (_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))];
        let _425_v138;
        _425_v138 = _dafny.Map.Empty.slice().updateUnsafe(_253_v12,_module.__default.fm0(_259_globalState));
        let _426_v139;
        let _nw71 = new _module.C1();
        _nw71.__ctor(_269_v25, _241_v2, _243_v4, _425_v138);
        _426_v139 = _nw71;
        let _427_v140;
        let _nw72 = Array((new BigNumber(23)).toNumber()).fill([]);
        _427_v140 = _nw72;
        let _index73 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_427_v140).length));
        (_427_v140)[_index73] = _250_v9;
        let _index74 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_427_v140).length));
        (_427_v140)[_index74] = _250_v9;
        _253_v12 = _module.__default.fm8(_259_globalState);
      } else {
        (_259_globalState).f6 = _269_v25;
        let _428_v141;
        let _nw73 = Array((new BigNumber(10)).toNumber());
        _nw73[(_dafny.ZERO).toNumber()] = _250_v9;
        _nw73[(_dafny.ONE).toNumber()] = _250_v9;
        _nw73[(new BigNumber(2)).toNumber()] = _250_v9;
        _nw73[(new BigNumber(3)).toNumber()] = _250_v9;
        _nw73[(new BigNumber(4)).toNumber()] = _250_v9;
        _nw73[(new BigNumber(5)).toNumber()] = _250_v9;
        _nw73[(new BigNumber(6)).toNumber()] = _250_v9;
        _nw73[(new BigNumber(7)).toNumber()] = _250_v9;
        _nw73[(new BigNumber(8)).toNumber()] = _250_v9;
        _nw73[(new BigNumber(9)).toNumber()] = _250_v9;
        _428_v141 = _nw73;
        let _index75 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_428_v141).length));
        (_428_v141)[_index75] = _250_v9;
        let _index76 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_428_v141).length));
        (_428_v141)[_index76] = _250_v9;
        let _429_v142;
        _429_v142 = _module.D5.create_DC14(false, new BigNumber(51));
        let _rhs75 = (_module.__default.fm29(_259_globalState)).dtor_cf5;
        let _rhs76 = (((_dafny.ZERO).minus((_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))])).minus(new BigNumber((_253_v12).length))).minus((_429_v142).dtor_cf17);
        let _rhs77 = ((_250_v9)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_250_v9).length))]).isLessThan((_dafny.ZERO).minus(_269_v25));
        _268_v24 = _rhs75;
        _239_v0 = _rhs76;
        _241_v2 = _rhs77;
        let _430_v143;
        _430_v143 = _dafny.MultiSet.fromElements(_268_v24, _268_v24, _241_v2);
        let _431_v144;
        let _432_v145;
        let _433_v146;
        let _out20;
        let _out21;
        let _out22;
        let _outcollector6 = _module.__default.m0(new BigNumber((_430_v143).cardinality()), _268_v24, (_dafny.ZERO).minus((_269_v25).multipliedBy(new BigNumber((_module.__default.fm30(_241_v2, _259_globalState)).length))), _259_globalState);
        _out20 = _outcollector6[0];
        _out21 = _outcollector6[1];
        _out22 = _outcollector6[2];
        _431_v144 = _out20;
        _432_v145 = _out21;
        _433_v146 = _out22;
        _432_v145 = _268_v24;
      }
      process.stdout.write(_dafny.toString(_239_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_240_v1, _dafny.Seq.of(new BigNumber(-140)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_241_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(false,true).updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_243_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new BigNumber(-140)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_245_v7)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_245_v7)[_dafny.ONE]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_245_v7)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_245_v7)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_245_v7)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_245_v7)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_245_v7)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_245_v7)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_245_v7)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_245_v7)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_249_v8).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v9)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v9)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v9)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v9)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_252_v11).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_253_v12).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_254_v13).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_255_v14, _dafny.Seq.of(false, true, false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v15).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(140),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_257_v16).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(140),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_258_v17, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(140),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_259_globalState.f1, _dafny.Seq.of(new BigNumber(-140)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_259_globalState).f2, _dafny.Seq.of(new BigNumber(-140)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_globalState).f3).equals(_dafny.Map.Empty.slice().updateUnsafe(false,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_globalState).f5).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new BigNumber(-140)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_259_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_259_globalState).f7)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_259_globalState).f7)[_dafny.ONE]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_259_globalState).f7)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_259_globalState).f7)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_259_globalState).f7)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_259_globalState).f7)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_259_globalState).f7)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_259_globalState).f7)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_259_globalState).f7)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_259_globalState).f7)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_259_globalState).f8).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_259_globalState).f9).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_259_globalState).f10).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_259_globalState).f11).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_259_globalState.f13, _dafny.Seq.of(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_259_globalState).f16, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(140),true), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(140),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_265_i2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_267_v23)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_267_v23)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_268_v24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_269_v25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v26).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_271_v27)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_274_v28).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[_dafny.ZERO]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[_dafny.ZERO]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[_dafny.ONE]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[_dafny.ONE]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(2)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(2)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(3)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(3)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(4)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(4)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(5)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(5)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(6)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(6)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(7)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(7)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(8)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(8)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(9)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(9)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(10)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(10)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(11)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(11)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(12)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(12)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(13)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(13)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(14)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(14)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(15)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(15)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(16)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(16)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(17)]).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_323_v64)[new BigNumber(17)]).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_325_v65).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-976),new BigNumber(458)).updateUnsafe(new BigNumber(-592),new BigNumber(378)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_326_i8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_336_v73).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_337_v74).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_337_v74).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_409_i14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_416_v130).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(140)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1() {
      let $dt = new D0(0);
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC2(cf1) {
      let $dt = new D0(2);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf1) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf1, other.cf1);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1();
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
    static create_DC4(cf3, cf4, cf5) {
      let $dt = new D1(0);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC5(cf6) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC6(cf7, cf8) {
      let $dt = new D1(2);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC3(cf2) {
      let $dt = new D1(3);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf7 === other.cf7 && this.cf8 === other.cf8;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf2 === other.cf2;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC7(cf9) {
      let $dt = new D2(0);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9);
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC9(cf11, cf12) {
      let $dt = new D3(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC10(cf13) {
      let $dt = new D3(1);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC8(cf10) {
      let $dt = new D3(2);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf11 === other.cf11 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf13 === other.cf13;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(false, _dafny.ZERO);
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
    static create_DC11(cf14) {
      let $dt = new D4(0);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf14 === other.cf14;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
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
    static create_DC13() {
      let $dt = new D5(0);
      return $dt;
    }
    static create_DC14(cf16, cf17) {
      let $dt = new D5(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC12(cf15) {
      let $dt = new D5(2);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf15) + ")";
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
        return other.$tag === 1 && this.cf16 === other.cf16 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13();
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
    static create_DC16(cf19, cf20) {
      let $dt = new D6(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC15(cf18) {
      let $dt = new D6(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC17(cf21) {
      let $dt = new D6(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf19 === other.cf19 && this.cf20 === other.cf20;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC16(false, []);
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
    static create_DC19(cf23) {
      let $dt = new D7(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC20(cf24, cf25) {
      let $dt = new D7(1);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC18(cf22) {
      let $dt = new D7(2);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf23 === other.cf23;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf24 === other.cf24 && this.cf25 === other.cf25;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC19(false);
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
    static create_DC21(cf26) {
      let $dt = new D8(0);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf26, other.cf26);
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC23(cf28, cf29, cf30, cf31) {
      let $dt = new D9(0);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC24(cf32) {
      let $dt = new D9(1);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC22(cf27) {
      let $dt = new D9(2);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC22" + "(" + this.cf27.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf28, other.cf28) && _dafny.areEqual(this.cf29, other.cf29) && this.cf30 === other.cf30 && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf32 === other.cf32;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC23(_dafny.Seq.of(), _dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC26(cf34, cf35, cf36, cf37, cf38) {
      let $dt = new D10(0);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC25(cf33) {
      let $dt = new D10(1);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC25" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf34, other.cf34) && this.cf35 === other.cf35 && this.cf36 === other.cf36 && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf33 === other.cf33;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC26(_dafny.ZERO, [], false, false, _dafny.ZERO);
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
    static create_DC28(cf40, cf41, cf42) {
      let $dt = new D11(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC27(cf39) {
      let $dt = new D11(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC28" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf40 === other.cf40 && _dafny.areEqual(this.cf41, other.cf41) && this.cf42 === other.cf42;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC28(null, _dafny.ZERO, false);
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
    static create_DC30(cf44, cf45) {
      let $dt = new D12(0);
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC31(cf46, cf47, cf48, cf49) {
      let $dt = new D12(1);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC29(cf43) {
      let $dt = new D12(2);
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC32(cf50) {
      let $dt = new D12(3);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get is_DC32() { return this.$tag === 3; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC30" + "(" + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf46) + ", " + this.cf47.toVerbatimString(true) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC29" + "(" + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf44 === other.cf44 && this.cf45 === other.cf45;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47) && _dafny.areEqual(this.cf48, other.cf48) && this.cf49 === other.cf49;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf50, other.cf50);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC30(false, false);
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
    static create_DC33(cf51) {
      let $dt = new D13(0);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC33" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf51, other.cf51);
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
          return D13.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = _dafny.Seq.of();
      this.f6 = _dafny.ZERO;
      this.f13 = _dafny.Seq.of();
      this._f0 = _dafny.ZERO;
      this._f2 = _dafny.Seq.of();
      this._f3 = _dafny.Map.Empty;
      this._f4 = _dafny.ZERO;
      this._f5 = _dafny.Map.Empty;
      this._f7 = [];
      this._f8 = _dafny.Map.Empty;
      this._f9 = _dafny.Seq.UnicodeFromString("");
      this._f10 = _dafny.Map.Empty;
      this._f11 = _dafny.Seq.of();
      this._f12 = false;
      this._f14 = false;
      this._f15 = _dafny.ZERO;
      this._f16 = _dafny.Seq.of();
      this._f17 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
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
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
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
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f18 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f19 = _dafny.Map.Empty;
      this._f20 = _dafny.Seq.UnicodeFromString("");
      this._f21 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    set f18(value) {
      let _this = this;
      _this._f18 = value;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    __ctor(f20, f21, f18, f19) {
      let _this = this;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      return;
    }
    fm6(p0, p1, globalState) {
      let _this = this;
      return !((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)).minus(new BigNumber(((_this).f20).length))).isEqualTo((new BigNumber(-574)).plus(new BigNumber(924)));
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f18 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f19 = _dafny.Map.Empty;
      this.f22 = _dafny.ZERO;
      this._f23 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    set f18(value) {
      let _this = this;
      _this._f18 = value;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    __ctor(f22, f23, f18, f19) {
      let _this = this;
      (_this).f22 = f22;
      (_this)._f23 = f23;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      return;
    }
    fm3(p0, p1, globalState) {
      let _this = this;
      return (_this).f23;
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (new BigNumber(390)).multipliedBy((_dafny.ZERO).minus((_dafny.ZERO).minus(((!((_this).f23)) ? (new BigNumber(851)) : (new BigNumber(366))))));
    };
    fm12(globalState) {
      let _this = this;
      return (_this).f23;
    };
    fm13(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (!(!((_this).f23))) {
        return _this.f18;
      } else {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = false;
      let _434_v0;
      _434_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_this.f22);
      let _435_v1;
      _435_v1 = _dafny.Set.fromElements(_434_v0, _434_v0);
      r3 = ((_dafny.Set.fromElements(_434_v0)).Union(_435_v1)).IsSubsetOf(_435_v1);
      (globalState).f6 = (_this.f22).plus(_this.f22);
      let _436_v2;
      _436_v2 = _dafny.MultiSet.fromElements(_this.f22, _this.f22, _this.f22);
      let _437_i0;
      _437_i0 = _dafny.ZERO;
      L3: {
        while (((_436_v2).update(_this.f22, _module.__default.abs(_this.f22))).IsSubsetOf(((p0) ? (_436_v2) : (_436_v2)))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_437_i0)) {
              break L3;
            }
            _437_i0 = (_437_i0).plus(_dafny.ONE);
            let _438_v3;
            _438_v3 = _module.D3.create_DC10(p0);
            let _source5 = _438_v3;
            if (_source5.is_DC9) {
              let _439___mcc_h0 = (_source5).cf11;
              let _440___mcc_h1 = (_source5).cf12;
              let _441_cf12 = _440___mcc_h1;
              let _442_cf11 = _439___mcc_h0;
              let _443_v4;
              _443_v4 = _dafny.Seq.of(!(_442_cf11));
              let _444_v5;
              _444_v5 = _dafny.MultiSet.fromElements((_this).f23, (_this).f23, p0);
              let _445_v6;
              _445_v6 = _dafny.Map.Empty.slice().updateUnsafe(_441_cf12,new BigNumber((_444_v5).cardinality()));
              let _446_v7;
              _446_v7 = _module.D3.create_DC8(_this.f18);
              let _447_v8;
              _447_v8 = _dafny.Seq.of(_this.f22);
              let _448_v9;
              _448_v9 = _dafny.Set.fromElements(_442_cf11);
              let _449_v10;
              _449_v10 = _dafny.Set.fromElements(_441_cf12);
              let _450_v11;
              _450_v11 = _dafny.Map.Empty.slice().updateUnsafe(_442_cf11,_447_v8);
              let _451_v12;
              let _nw74 = Array((new BigNumber(21)).toNumber());
              _nw74[(_dafny.ZERO).toNumber()] = p0;
              _nw74[(_dafny.ONE).toNumber()] = _module.__default.fm1(p0, _this.f18, _this.f22, globalState);
              _nw74[(new BigNumber(2)).toNumber()] = _442_cf11;
              _nw74[(new BigNumber(3)).toNumber()] = false;
              _nw74[(new BigNumber(4)).toNumber()] = _442_cf11;
              _nw74[(new BigNumber(5)).toNumber()] = !((_this).f23);
              _nw74[(new BigNumber(6)).toNumber()] = (_this).f23;
              _nw74[(new BigNumber(7)).toNumber()] = _module.__default.fm1((_this).f23, _this.f18, _this.f22, globalState);
              _nw74[(new BigNumber(8)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_443_v4, _dafny.Seq.of((_this).f23, p0, p0, _442_cf11));
              _nw74[(new BigNumber(9)).toNumber()] = (_441_cf12).isLessThanOrEqualTo(new BigNumber((_445_v6).length));
              _nw74[(new BigNumber(10)).toNumber()] = !(_this.f22).isEqualTo(_this.f22);
              _nw74[(new BigNumber(11)).toNumber()] = _442_cf11;
              _nw74[(new BigNumber(12)).toNumber()] = _module.__default.fm1((_this).f23, (_446_v7).dtor_cf10, (_447_v8)[_module.__default.safeIndex(_441_cf12, new BigNumber((_447_v8).length))], globalState);
              _nw74[(new BigNumber(13)).toNumber()] = (_module.__default.fm2(globalState)).IsSubsetOf(_448_v9);
              _nw74[(new BigNumber(14)).toNumber()] = !(false);
              _nw74[(new BigNumber(15)).toNumber()] = p0;
              _nw74[(new BigNumber(16)).toNumber()] = _dafny.Seq.contains((((_450_v11).contains((_this).f23)) ? ((_450_v11).get((_this).f23)) : (_447_v8)), new BigNumber((_449_v10).length));
              _nw74[(new BigNumber(17)).toNumber()] = (_441_cf12).isEqualTo(_441_cf12);
              _nw74[(new BigNumber(18)).toNumber()] = (_this).f23;
              _nw74[(new BigNumber(19)).toNumber()] = _442_cf11;
              _nw74[(new BigNumber(20)).toNumber()] = _module.__default.fm1(false, _this.f18, new BigNumber(-279), globalState);
              _451_v12 = _nw74;
              _451_v12 = _451_v12;
              let _452_v13;
              _452_v13 = _module.D0.create_DC2(_module.D0.create_DC1());
              let _453_v14;
              _453_v14 = _module.D0.create_DC2(_452_v13);
              _453_v14 = _453_v14;
              let _454_v15;
              _454_v15 = _dafny.Seq.UnicodeFromString("hkdt");
              let _455_v16;
              _455_v16 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ylwbg"), _454_v15);
              _455_v16 = (_455_v16).Union(_dafny.Set.fromElements(_454_v15, _454_v15, _454_v15));
              let _456_v17;
              _456_v17 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,_this.f22);
              let _457_v19;
              _457_v19 = _dafny.Seq.of(_454_v15, _454_v15, _454_v15);
              let _458_v20;
              _458_v20 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(170)), function (_459_i1) {
                return _this.f18;
              }),_441_cf12), (_this).f19, function () {
                let _coll29 = new _dafny.Map();
                for (const _compr_29 of (_457_v19).Elements) {
                  let _460_v18 = _compr_29;
                  if (_dafny.Seq.contains(_457_v19, _460_v18)) {
                    _coll29.push([_460_v18,_this.f22]);
                  }
                }
                return _coll29;
              }(), (_this).f19);
              let _461_v21;
              let _nw75 = new _module.C0();
              _nw75.__ctor(_454_v15, _456_v17, new _dafny.CodePoint('x'.codePointAt(0)), (_458_v20)[_module.__default.safeIndex(_this.f22, new BigNumber((_458_v20).length))]);
              _461_v21 = _nw75;
            } else if (_source5.is_DC10) {
              let _462___mcc_h2 = (_source5).cf13;
              let _463_cf13 = _462___mcc_h2;
              _463_cf13 = ((_module.D3.create_DC9(_463_cf13, _module.__default.fm0(globalState))).dtor_cf11) === (_463_cf13);
              let _464_v22;
              let _nw76 = Array((new BigNumber(13)).toNumber()).fill(false);
              _464_v22 = _nw76;
              let _465_v23;
              _465_v23 = _dafny.Map.Empty.slice().updateUnsafe(_464_v22,_this.f22);
              let _466_v24;
              _466_v24 = _dafny.MultiSet.fromElements((_this).f23);
              let _467_v25;
              _467_v25 = _dafny.Set.fromElements(false);
              let _468_v26;
              let _nw77 = Array((new BigNumber(8)).toNumber());
              _nw77[(_dafny.ZERO).toNumber()] = (_this.f22).plus(_this.f22);
              _nw77[(_dafny.ONE).toNumber()] = ((((_465_v23).contains(_464_v22)) ? ((_465_v23).get(_464_v22)) : (new BigNumber((_466_v24).cardinality())))).multipliedBy(_this.f22);
              _nw77[(new BigNumber(2)).toNumber()] = _this.f22;
              _nw77[(new BigNumber(3)).toNumber()] = _this.f22;
              _nw77[(new BigNumber(4)).toNumber()] = _this.f22;
              _nw77[(new BigNumber(5)).toNumber()] = new BigNumber((_467_v25).length);
              _nw77[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_this.f22);
              _nw77[(new BigNumber(7)).toNumber()] = (new BigNumber(251)).minus(_this.f22);
              _468_v26 = _nw77;
              let _index77 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_468_v26).length));
              (_468_v26)[_index77] = _module.__default.safeDivisionInt(_this.f22, new BigNumber((_module.__default.fm14(_module.__default.fm0(globalState), globalState)).length));
              let _469_v27;
              _469_v27 = _dafny.Seq.of(_module.__default.fm0(globalState));
              let _index78 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_468_v26).length));
              (_468_v26)[_index78] = _module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.FromArray(_469_v27)).cardinality()), _this.f22);
              let _470_v28;
              let _nw78 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _470_v28 = _nw78;
              let _index79 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_470_v28).length));
              (_470_v28)[_index79] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(132)), function (_471_i2) {
                return _this.f18;
              });
              let _472_v29;
              _472_v29 = _dafny.Seq.UnicodeFromString("vphuxrsd");
              let _index80 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_470_v28).length));
              (_470_v28)[_index80] = _dafny.Seq.Concat(_472_v29, _dafny.Seq.UnicodeFromString("f"));
              let _index81 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_468_v26).length));
              (_468_v26)[_index81] = (new BigNumber(249)).multipliedBy((_dafny.ZERO).minus(_this.f22));
            } else {
              let _473___mcc_h3 = (_source5).cf10;
              let _474_cf10 = _473___mcc_h3;
              let _475_v30;
              _475_v30 = _dafny.MultiSet.fromElements((_this).f23, p0, p0, (_this).f23, (_this).f23);
              let _476_v31;
              _476_v31 = _module.D5.create_DC12(_475_v30);
              let _477_v32;
              _477_v32 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,(_476_v31).dtor_cf15);
              let _478_v33;
              _478_v33 = _dafny.Seq.of(p0);
              let _479_v34;
              let _nw79 = Array((new BigNumber(17)).toNumber());
              _nw79[(_dafny.ZERO).toNumber()] = _475_v30;
              _nw79[(_dafny.ONE).toNumber()] = _475_v30;
              _nw79[(new BigNumber(2)).toNumber()] = (((_477_v32).contains(new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality()))) ? ((_477_v32).get(new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality()))) : (_dafny.MultiSet.fromElements((_this).f23)));
              _nw79[(new BigNumber(3)).toNumber()] = _475_v30;
              _nw79[(new BigNumber(4)).toNumber()] = _475_v30;
              _nw79[(new BigNumber(5)).toNumber()] = _475_v30;
              _nw79[(new BigNumber(6)).toNumber()] = _475_v30;
              _nw79[(new BigNumber(7)).toNumber()] = (_475_v30).update((_this).f23, _module.__default.abs(_this.f22));
              _nw79[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.FromArray(_478_v33);
              _nw79[(new BigNumber(9)).toNumber()] = (_475_v30).Intersect(_475_v30);
              _nw79[(new BigNumber(10)).toNumber()] = _475_v30;
              _nw79[(new BigNumber(11)).toNumber()] = _475_v30;
              _nw79[(new BigNumber(12)).toNumber()] = _475_v30;
              _nw79[(new BigNumber(13)).toNumber()] = _475_v30;
              _nw79[(new BigNumber(14)).toNumber()] = _475_v30;
              _nw79[(new BigNumber(15)).toNumber()] = _475_v30;
              _nw79[(new BigNumber(16)).toNumber()] = (_475_v30).Difference(_dafny.MultiSet.FromArray(_478_v33));
              _479_v34 = _nw79;
              let _index82 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_479_v34).length));
              (_479_v34)[_index82] = (_475_v30).Union(_475_v30);
              let _480_v35;
              _480_v35 = _dafny.Seq.UnicodeFromString("hoka");
              let _481_v36;
              _481_v36 = _dafny.MultiSet.fromElements(((p0) ? (_480_v35) : (_480_v35)), _dafny.Seq.Concat(_480_v35, _480_v35), _dafny.Seq.Create(_module.__default.abs(new BigNumber(874)), function (_482_i3) {
                return _this.f18;
              }), _480_v35, _dafny.Seq.update(_480_v35, _module.__default.safeIndex(_this.f22, new BigNumber((_480_v35).length)), _474_cf10));
              let _483_v37;
              _483_v37 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,new BigNumber((_475_v30).cardinality()));
              let _index83 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_479_v34).length));
              let _rhs78 = p0;
              let _rhs79 = (_475_v30).Union(_475_v30);
              let _rhs80 = _module.__default.fm15(_this.f18, (_dafny.ZERO).minus(_this.f22), _438_v3, globalState);
              let _rhs81 = (_dafny.ZERO).minus(((_this.f22).minus(_this.f22)).multipliedBy((((_483_v37).contains(new BigNumber((_434_v0).length))) ? ((_483_v37).get(new BigNumber((_434_v0).length))) : (_this.f22))));
              let _lhs61 = _479_v34;
              let _lhs62 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_479_v34).length));
              let _lhs63 = globalState;
              r3 = _rhs78;
              _lhs61[_lhs62] = _rhs79;
              _481_v36 = _rhs80;
              _lhs63.f6 = _rhs81;
              let _484_v38;
              _484_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f23);
              let _485_v39;
              _485_v39 = _dafny.Set.fromElements(_this.f22, _this.f22, new BigNumber(295));
              let _486_v40;
              _486_v40 = _dafny.Seq.of(_485_v39);
              _484_v38 = (_484_v38).update((_485_v39).IsDisjointFrom((_486_v40)[_module.__default.safeIndex(new BigNumber(-737), new BigNumber((_486_v40).length))]), p0);
              let _487_v41;
              let _init7 = function (_488_i4) {
                return (_488_i4).multipliedBy(_this.f22);
              };
              let _nw80 = Array((new BigNumber(21)).toNumber());
              for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw80.length); _i0_7++) {
                _nw80[_i0_7] = _init7(new BigNumber(_i0_7));
              }
              _487_v41 = _nw80;
              let _index84 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_487_v41).length));
              (_487_v41)[_index84] = _this.f22;
              let _index85 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_487_v41).length));
              (_487_v41)[_index85] = (_dafny.ZERO).minus(_this.f22);
              let _489_v42;
              _489_v42 = _module.D5.create_DC14(((p0) ? (p0) : ((_this).f23)), (_487_v41)[_module.__default.safeIndex(new BigNumber(162), new BigNumber((_487_v41).length))]);
              let _490_v43;
              _490_v43 = _dafny.Seq.of(new BigNumber((_480_v35).length));
              let _pat_let_tv4 = _487_v41;
              let _pat_let_tv5 = _487_v41;
              let _pat_let_tv6 = _487_v41;
              let _pat_let_tv7 = _487_v41;
              let _pat_let_tv8 = _490_v43;
              let _pat_let_tv9 = _490_v43;
              _489_v42 = function (_pat_let6_0) {
                return function (_495_dt__update__tmp_h2) {
                  return function (_pat_let11_0) {
                    return function (_496_dt__update_hcf17_h2) {
                      return _module.D5.create_DC14((_495_dt__update__tmp_h2).dtor_cf16, _496_dt__update_hcf17_h2);
                    }(_pat_let11_0);
                  }(new BigNumber((_dafny.Seq.Concat(_pat_let_tv8, _pat_let_tv9)).length));
                }(_pat_let6_0);
              }(function (_pat_let7_0) {
                return function (_493_dt__update__tmp_h1) {
                  return function (_pat_let10_0) {
                    return function (_494_dt__update_hcf17_h1) {
                      return _module.D5.create_DC14((_493_dt__update__tmp_h1).dtor_cf16, _494_dt__update_hcf17_h1);
                    }(_pat_let10_0);
                  }((_pat_let_tv7)[_module.__default.safeIndex(new BigNumber(162), new BigNumber((_pat_let_tv6).length))]);
                }(_pat_let7_0);
              }(function (_pat_let8_0) {
                return function (_491_dt__update__tmp_h0) {
                  return function (_pat_let9_0) {
                    return function (_492_dt__update_hcf17_h0) {
                      return _module.D5.create_DC14((_491_dt__update__tmp_h0).dtor_cf16, _492_dt__update_hcf17_h0);
                    }(_pat_let9_0);
                  }((_pat_let_tv5)[_module.__default.safeIndex(new BigNumber(162), new BigNumber((_pat_let_tv4).length))]);
                }(_pat_let8_0);
              }(_489_v42)));
            }
            let _497_v44;
            _497_v44 = _dafny.Seq.of((_this).f23);
            let _498_v45;
            _498_v45 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,_this.f22);
            let _499_v46;
            _499_v46 = _dafny.Seq.of(_498_v45, _498_v45);
            let _500_v47;
            _500_v47 = _dafny.Seq.UnicodeFromString("gr");
            let _501_v48;
            _501_v48 = _497_v44;
            let _502_v49;
            let _nw81 = Array((new BigNumber(15)).toNumber());
            _nw81[(_dafny.ZERO).toNumber()] = false;
            _nw81[(_dafny.ONE).toNumber()] = (_this).f23;
            _nw81[(new BigNumber(2)).toNumber()] = (_497_v44)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((_497_v44).length))];
            _nw81[(new BigNumber(3)).toNumber()] = (_this).f23;
            _nw81[(new BigNumber(4)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_499_v46, _module.__default.fm16(globalState));
            _nw81[(new BigNumber(5)).toNumber()] = (_this).f23;
            _nw81[(new BigNumber(6)).toNumber()] = p0;
            _nw81[(new BigNumber(7)).toNumber()] = _dafny.Seq.contains(_500_v47, _this.f18);
            _nw81[(new BigNumber(8)).toNumber()] = _dafny.areEqual((_501_v48), _497_v44);
            _nw81[(new BigNumber(9)).toNumber()] = p0;
            _nw81[(new BigNumber(10)).toNumber()] = p0;
            _nw81[(new BigNumber(11)).toNumber()] = !((_this).f23) || (p0);
            _nw81[(new BigNumber(12)).toNumber()] = !(p0);
            _nw81[(new BigNumber(13)).toNumber()] = _module.__default.fm1(p0, _this.f18, _this.f22, globalState);
            _nw81[(new BigNumber(14)).toNumber()] = p0;
            _502_v49 = _nw81;
            let _index86 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_502_v49).length));
            (_502_v49)[_index86] = p0;
            let _index87 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_502_v49).length));
            (_502_v49)[_index87] = _module.__default.fm1(p0, ((!(p0)) ? (_this.f18) : (_this.f18)), _module.__default.fm0(globalState), globalState);
            let _503_v50;
            let _nw82 = new _module.C0();
            _nw82.__ctor(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("s"), _500_v47), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),_this.f22), _this.f18, ((_this).f19).Merge((_this).f19));
            _503_v50 = _nw82;
            let _504_v51;
            let _nw83 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
            _504_v51 = _nw83;
            let _index88 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_504_v51).length));
            (_504_v51)[_index88] = new BigNumber(((_503_v50).f20).length);
            let _index89 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_504_v51).length));
            (_504_v51)[_index89] = _this.f22;
          }
        }
      }
      let _hi3 = _this.f22;
      for (let _505_i5 = _this.f22; _505_i5.isLessThan(_hi3); _505_i5 = _505_i5.plus(_dafny.ONE)) {
        let _506_v52;
        _506_v52 = _dafny.Seq.of(_505_i5, new BigNumber(-684));
        let _507_v53;
        _507_v53 = _dafny.Seq.of(_dafny.Seq.of(_505_i5), _506_v52);
        (globalState).f6 = (_505_i5).multipliedBy((new BigNumber((_507_v53).length)).multipliedBy(_505_i5));
        let _508_v54;
        _508_v54 = _dafny.Seq.UnicodeFromString("isnmbr");
        let _rhs82 = (_this).f23;
        let _rhs83 = (_dafny.ZERO).minus(new BigNumber((_508_v54).length));
        r2 = _rhs82;
        r1 = _rhs83;
        r3 = p0;
        r2 = !(p0);
      }
      let _509_v55;
      _509_v55 = _dafny.Seq.UnicodeFromString("sk");
      let _510_i6;
      _510_i6 = _dafny.ZERO;
      L4: {
        while ((((_this).f23) ? (!_dafny.areEqual(_509_v55, _509_v55)) : ((_this).f23))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_510_i6)) {
              break L4;
            }
            _510_i6 = (_510_i6).plus(_dafny.ONE);
            let _511_v56;
            _511_v56 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,_this.f22);
            let _512_v57;
            let _nw84 = new _module.C0();
            _nw84.__ctor(_509_v55, _511_v56, (((_this).f23) ? (_this.f18) : (_this.f18)), ((_this).f19).Merge((_this).f19));
            _512_v57 = _nw84;
            let _513_v58;
            _513_v58 = _dafny.Seq.of(p0);
            let _514_v59;
            _514_v59 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,_513_v58);
            let _515_v60;
            _515_v60 = _dafny.Seq.Concat((((_514_v59).contains(_this.f22)) ? ((_514_v59).get(_this.f22)) : (_dafny.Seq.of((_this).f23))), _513_v58);
            let _516_v61;
            _516_v61 = _dafny.Set.fromElements(p0, _module.__default.fm1(!(!(p0)), _this.f18, new BigNumber(169), globalState), (_this).f23);
            let _517_v62;
            _517_v62 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,(_this).f23);
            let _518_v63;
            let _nw85 = Array((new BigNumber(4)).toNumber());
            _nw85[(_dafny.ZERO).toNumber()] = true;
            _nw85[(_dafny.ONE).toNumber()] = (_516_v61).IsDisjointFrom(_dafny.Set.fromElements((_this).f23));
            _nw85[(new BigNumber(2)).toNumber()] = (_this).f23;
            _nw85[(new BigNumber(3)).toNumber()] = (((_517_v62).contains(_this.f22)) ? ((_517_v62).get(_this.f22)) : ((_this).f23));
            _518_v63 = _nw85;
            let _index90 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_518_v63).length));
            (_518_v63)[_index90] = p0;
            let _519_v64;
            _519_v64 = _dafny.Seq.of(_517_v62, _517_v62);
            let _index91 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_518_v63).length));
            let _rhs84 = _this.f22;
            let _rhs85 = _515_v60;
            let _rhs86 = true;
            let _rhs87 = _this.f22;
            let _rhs88 = (_517_v62).equals((_519_v64)[_module.__default.safeIndex(new BigNumber(((_512_v57).f20).length), new BigNumber((_519_v64).length))]);
            let _lhs64 = globalState;
            let _lhs65 = globalState;
            let _lhs66 = _518_v63;
            let _lhs67 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_518_v63).length));
            _lhs64.f6 = _rhs84;
            _515_v60 = _rhs85;
            r3 = _rhs86;
            _lhs65.f6 = _rhs87;
            _lhs66[_lhs67] = _rhs88;
            let _index92 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_518_v63).length));
            (_518_v63)[_index92] = (_513_v58)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_513_v58).length))];
            let _520_v65;
            _520_v65 = _module.D3.create_DC8(((false) ? (_this.f18) : (_this.f18)));
            let _source6 = _520_v65;
            if (_source6.is_DC9) {
              let _521___mcc_h4 = (_source6).cf11;
              let _522___mcc_h5 = (_source6).cf12;
              let _523_cf12 = _522___mcc_h5;
              let _524_cf11 = _521___mcc_h4;
              let _525_v66;
              let _nw86 = new _module.C0();
              _nw86.__ctor(_dafny.Seq.UnicodeFromString("isrngtcpi"), ((_524_cf11) ? ((_512_v57).f21) : (_dafny.Map.Empty.slice().updateUnsafe(_this.f18,new BigNumber((_dafny.Seq.of(_this.f22, _523_cf12)).length)))), _this.f18, (_this).f19);
              _525_v66 = _nw86;
              let _526_v67;
              let _nw87 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
              _526_v67 = _nw87;
              let _index93 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_526_v67).length));
              (_526_v67)[_index93] = (_this.f22).plus(_523_cf12);
              let _index94 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_526_v67).length));
              (_526_v67)[_index94] = _523_cf12;
              let _527_v68;
              let _nw88 = Array((new BigNumber(11)).toNumber());
              _nw88[(_dafny.ZERO).toNumber()] = _this.f18;
              _nw88[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('w'.codePointAt(0));
              _nw88[(new BigNumber(2)).toNumber()] = (_this).fm13(_523_cf12, _this.f22, new BigNumber((_dafny.MultiSet.fromElements(!(p0))).cardinality()), (_this).f23, globalState);
              _nw88[(new BigNumber(3)).toNumber()] = _this.f18;
              _nw88[(new BigNumber(4)).toNumber()] = _this.f18;
              _nw88[(new BigNumber(5)).toNumber()] = ((_524_cf11) ? (_this.f18) : (_this.f18));
              _nw88[(new BigNumber(6)).toNumber()] = (_this).fm13((_526_v67)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_526_v67).length))], _523_cf12, _this.f22, !(false), globalState);
              _nw88[(new BigNumber(7)).toNumber()] = _this.f18;
              _nw88[(new BigNumber(8)).toNumber()] = _this.f18;
              _nw88[(new BigNumber(9)).toNumber()] = _this.f18;
              _nw88[(new BigNumber(10)).toNumber()] = _this.f18;
              _527_v68 = _nw88;
              _527_v68 = _527_v68;
              r3 = _524_cf11;
            } else if (_source6.is_DC10) {
              let _528___mcc_h6 = (_source6).cf13;
              let _529_cf13 = _528___mcc_h6;
              let _530_v69;
              _530_v69 = _module.D5.create_DC14(((_518_v63)[_module.__default.safeIndex(new BigNumber(342), new BigNumber((_518_v63).length))]) === (_529_cf13), (_this.f22).minus(_this.f22));
              _530_v69 = _530_v69;
              r1 = (_dafny.ZERO).minus(_this.f22);
              let _531_v70;
              let _nw89 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _531_v70 = _nw89;
              let _index95 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_531_v70).length));
              (_531_v70)[_index95] = (_512_v57).f20;
              let _index96 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_531_v70).length));
              (_531_v70)[_index96] = (_512_v57).f20;
              let _532_v71;
              let _nw90 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
              _532_v71 = _nw90;
              let _init8 = ((_533_v60) => function (_534_i7) {
                return _533_v60;
              })(_515_v60);
              let _nw91 = Array((new BigNumber(9)).toNumber());
              for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw91.length); _i0_8++) {
                _nw91[_i0_8] = _init8(new BigNumber(_i0_8));
              }
              _532_v71 = _nw91;
            } else {
              let _535___mcc_h7 = (_source6).cf10;
              let _536_cf10 = _535___mcc_h7;
              _536_cf10 = _536_cf10;
              (globalState).f6 = _this.f22;
              let _537_v72;
              _537_v72 = _dafny.Set.fromElements(_this.f22, new BigNumber(513));
              let _538_v73;
              _538_v73 = _dafny.Seq.of(new BigNumber((_517_v62).length), _this.f22);
              _537_v72 = (_dafny.Set.fromElements((_538_v73)[_module.__default.safeIndex(_this.f22, new BigNumber((_538_v73).length))])).Intersect(_537_v72);
              let _539_v74;
              _539_v74 = _dafny.Seq.of(_436_v2);
              let _540_v75;
              _540_v75 = _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(836), _this.f22), _dafny.Seq.Create(_module.__default.abs(new BigNumber(264)), function (_541_i8) {
                return _this.f22;
              })), _dafny.Seq.of(new BigNumber((_539_v74).length), new BigNumber((_538_v73).length)), _538_v73);
              _540_v75 = (function () {
                let _coll30 = new _dafny.Set();
                for (const _compr_30 of (_dafny.Seq.of(_538_v73)).Elements) {
                  let _542_v76 = _compr_30;
                  if (_dafny.Seq.contains(_dafny.Seq.of(_538_v73), _542_v76)) {
                    _coll30.add(_542_v76);
                  }
                }
                return _coll30;
              }()).Union(_540_v75);
            }
          }
        }
      }
      let _543_v77;
      let _nw92 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
      _543_v77 = _nw92;
      let _544_v78;
      _544_v78 = _module.D3.create_DC8(new _dafny.CodePoint('r'.codePointAt(0)));
      let _545_v79;
      _545_v79 = _dafny.MultiSet.fromElements(_509_v55, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-325)), function (_546_i9) {
        return _this.f18;
      }), _dafny.Seq.UnicodeFromString("xdamu"));
      let _547_v80;
      let _nw93 = Array((new BigNumber(13)).toNumber()).fill(false);
      _547_v80 = _nw93;
      let _548_v81;
      _548_v81 = _dafny.Seq.of(_this.f22, _this.f22, _this.f22, _this.f22, _this.f22);
      let _549_v82;
      _549_v82 = _dafny.Map.Empty.slice().updateUnsafe(_547_v80,(_548_v81)[_module.__default.safeIndex(_this.f22, new BigNumber((_548_v81).length))]);
      let _550_v83;
      _550_v83 = _dafny.Seq.of((_this).f23);
      let _rhs89 = _543_v77;
      let _rhs90 = _544_v78;
      let _rhs91 = _545_v79;
      let _rhs92 = _dafny.Map.Empty.slice().updateUnsafe(_547_v80,_this.f22);
      let _rhs93 = _550_v83;
      let _lhs68 = globalState;
      _543_v77 = _rhs89;
      _544_v78 = _rhs90;
      _545_v79 = _rhs91;
      _549_v82 = _rhs92;
      _lhs68.f13 = _rhs93;
      r0 = _module.D0.create_DC0(_this.f22);
      r1 = _this.f22;
      r2 = !(!_dafny.areEqual(_dafny.Seq.UnicodeFromString("yjptbhmhv"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(423)), function (_551_i10) {
        return _this.f18;
      })));
      r3 = p0;
      return [r0, r1, r2, r3];
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _552_v0;
      let _nw94 = Array((new BigNumber(15)).toNumber());
      _nw94[(_dafny.ZERO).toNumber()] = _this.f18;
      _nw94[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
      _nw94[(new BigNumber(2)).toNumber()] = _this.f18;
      _nw94[(new BigNumber(3)).toNumber()] = _this.f18;
      _nw94[(new BigNumber(4)).toNumber()] = (_this).fm13(p3, p2, new BigNumber(968), p0, globalState);
      _nw94[(new BigNumber(5)).toNumber()] = _this.f18;
      _nw94[(new BigNumber(6)).toNumber()] = _this.f18;
      _nw94[(new BigNumber(7)).toNumber()] = _this.f18;
      _nw94[(new BigNumber(8)).toNumber()] = _this.f18;
      _nw94[(new BigNumber(9)).toNumber()] = _this.f18;
      _nw94[(new BigNumber(10)).toNumber()] = _this.f18;
      _nw94[(new BigNumber(11)).toNumber()] = _this.f18;
      _nw94[(new BigNumber(12)).toNumber()] = _this.f18;
      _nw94[(new BigNumber(13)).toNumber()] = _this.f18;
      _nw94[(new BigNumber(14)).toNumber()] = _this.f18;
      _552_v0 = _nw94;
      let _553_v1;
      _553_v1 = _dafny.Set.fromElements(_552_v0, _552_v0);
      (globalState).f6 = (_this).fm4(!((_this).f23), (((_this).f23) ? (_this.f18) : (_this.f18)), _this.f18, (_553_v1).contains(_552_v0), globalState);
      let _554_v2;
      _554_v2 = _dafny.Seq.UnicodeFromString("jdgvd");
      let _555_v3;
      _555_v3 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_554_v2, _554_v2), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Concat(_554_v2, _554_v2)).length)), _this.f18),!(_module.__default.fm1((_this).f23, _this.f18, new BigNumber((_554_v2).length), globalState)));
      let _556_v4;
      _556_v4 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,_554_v2);
      _555_v3 = (_555_v3).update(_dafny.Seq.Concat((((_556_v4).contains(p1)) ? ((_556_v4).get(p1)) : (_554_v2)), _554_v2), false);
      let _557_i0;
      _557_i0 = _dafny.ZERO;
      L5: {
        while ((p2).isLessThanOrEqualTo(p1)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_557_i0)) {
              break L5;
            }
            _557_i0 = (_557_i0).plus(_dafny.ONE);
            (globalState).f6 = (p2).minus(p1);
            let _558_v5;
            _558_v5 = _dafny.Map.Empty.slice().updateUnsafe(p3,_this.f18);
            let _559_v6;
            _559_v6 = _module.D1.create_DC4(p1, p2, (_this).f23);
            r0 = (_this).fm4((_this).fm12(globalState), _this.f18, (((_558_v5).contains((_559_v6).dtor_cf4)) ? ((_558_v5).get((_559_v6).dtor_cf4)) : (new _dafny.CodePoint('y'.codePointAt(0)))), (_this).f23, globalState);
            (_this).f22 = p1;
            let _560_v7;
            let _nw95 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
            _560_v7 = _nw95;
            let _index97 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_560_v7).length));
            (_560_v7)[_index97] = _module.__default.safeModuloInt(p2, new BigNumber((_558_v5).length));
            let _index98 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_560_v7).length));
            (_560_v7)[_index98] = new BigNumber(365);
          }
        }
      }
      let _561_v8;
      _561_v8 = _dafny.Seq.of(p0, false);
      if ((_module.__default.safeModuloInt(p1, new BigNumber((_561_v8).length))).isLessThanOrEqualTo(p1)) {
        let _562_v9;
        _562_v9 = _dafny.MultiSet.fromElements(false, (_this).f23);
        _562_v9 = (_562_v9).Union(_562_v9);
        _561_v8 = _dafny.Seq.Concat(_561_v8, _561_v8);
        let _563_v10;
        _563_v10 = _dafny.Set.fromElements(new BigNumber(299));
        _563_v10 = _module.__default.fm17(globalState);
        let _564_v11;
        _564_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,p3);
        let _565_v12;
        _565_v12 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),new BigNumber((_564_v11).length));
        let _566_v14;
        _566_v14 = _dafny.Seq.of(_554_v2, _554_v2);
        let _567_v15;
        let _nw96 = new _module.C0();
        _nw96.__ctor(_554_v2, _565_v12, _this.f18, function () {
          let _coll31 = new _dafny.Map();
          for (const _compr_31 of (_dafny.Seq.update(_566_v14, _module.__default.safeIndex(p3, new BigNumber((_566_v14).length)), _554_v2)).Elements) {
            let _568_v13 = _compr_31;
            if (_dafny.Seq.contains(_dafny.Seq.update(_566_v14, _module.__default.safeIndex(p3, new BigNumber((_566_v14).length)), _554_v2), _568_v13)) {
              _coll31.push([_568_v13,new BigNumber((_dafny.MultiSet.fromElements(_this.f18, _this.f18, new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), _this.f18)).cardinality())]);
            }
          }
          return _coll31;
        }());
        _567_v15 = _nw96;
        let _569_v16;
        _569_v16 = _dafny.Map.Empty.slice().updateUnsafe(_567_v15,_567_v15);
        let _570_v17;
        _570_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(((_569_v16).contains(_567_v15)) ? ((_569_v16).get(_567_v15)) : (_567_v15)));
        _570_v17 = (_570_v17).update(!_dafny.areEqual(_561_v8, _561_v8), _567_v15);
        let _571_v18;
        _571_v18 = _module.D6.create_DC15((_567_v15).f21);
        let _572_v19;
        _572_v19 = _dafny.MultiSet.fromElements((_565_v12).Merge(_565_v12), (_571_v18).dtor_cf18);
        _572_v19 = _572_v19;
      } else {
        r0 = p1;
        if ((_this).f23) {
          let _573_v20;
          _573_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_this.f18);
          _573_v20 = (_573_v20).update(!((_this).f23), _this.f18);
          let _574_v21;
          _574_v21 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC6(!(!(false)), true),p0);
          let _575_v22;
          _575_v22 = _module.D1.create_DC6((_this).f23, p0);
          _574_v21 = (_574_v21).update(_575_v22, (_this).f23);
          (_this).f22 = p2;
          (globalState).f6 = (_dafny.ZERO).minus(_this.f22);
          let _576_v23;
          _576_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,p0);
          let _577_v24;
          let _nw97 = Array((new BigNumber(11)).toNumber());
          _nw97[(_dafny.ZERO).toNumber()] = (new BigNumber((_576_v23).length)).minus(new BigNumber((_561_v8).length));
          _nw97[(_dafny.ONE).toNumber()] = (new BigNumber(744)).plus(p1);
          _nw97[(new BigNumber(2)).toNumber()] = p1;
          _nw97[(new BigNumber(3)).toNumber()] = _this.f22;
          _nw97[(new BigNumber(4)).toNumber()] = (_this).fm4((_this).f23, _this.f18, _this.f18, p0, globalState);
          _nw97[(new BigNumber(5)).toNumber()] = _module.__default.fm0(globalState);
          _nw97[(new BigNumber(6)).toNumber()] = p1;
          _nw97[(new BigNumber(7)).toNumber()] = p3;
          _nw97[(new BigNumber(8)).toNumber()] = new BigNumber(-798);
          _nw97[(new BigNumber(9)).toNumber()] = p2;
          _nw97[(new BigNumber(10)).toNumber()] = p2;
          _577_v24 = _nw97;
          let _578_v25;
          _578_v25 = _dafny.Seq.of(new BigNumber((_561_v8).length));
          let _579_v26;
          _579_v26 = _dafny.MultiSet.fromElements((_this).f23, p0);
          let _index99 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_577_v24).length));
          (_577_v24)[_index99] = _module.__default.safeModuloInt((_578_v25)[_module.__default.safeIndex(p2, new BigNumber((_578_v25).length))], new BigNumber((_579_v26).cardinality()));
          let _index100 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_577_v24).length));
          (_577_v24)[_index100] = p3;
        } else {
          let _580_v27;
          let _nw98 = new _module.C0();
          _nw98.__ctor(_554_v2, _module.__default.fm14(new BigNumber(356), globalState), _this.f18, (_this).f19);
          _580_v27 = _nw98;
          let _581_v28;
          _581_v28 = _dafny.Seq.of(_580_v27);
          let _582_v29;
          _582_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_dafny.Seq.Concat(_581_v28, _581_v28));
          let _583_v30;
          let _nw99 = Array((new BigNumber(25)).toNumber()).fill(false);
          _583_v30 = _nw99;
          let _584_v31;
          _584_v31 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),_this.f22);
          let _585_v32;
          let _nw100 = new _module.C0();
          _nw100.__ctor(_dafny.Seq.UnicodeFromString("v"), _584_v31, _this.f18, (_580_v27).f19);
          _585_v32 = _nw100;
          let _586_v33;
          _586_v33 = _dafny.Map.Empty.slice().updateUnsafe(_585_v32,!((_this).f23));
          let _index101 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_583_v30).length));
          (_583_v30)[_index101] = (((_586_v33).contains(_585_v32)) ? ((_586_v33).get(_585_v32)) : (p0));
          let _index102 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_583_v30).length));
          let _rhs94 = (_582_v29).update((_this).f23, _581_v28);
          let _rhs95 = (_this).f23;
          let _lhs69 = _583_v30;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_583_v30).length));
          _582_v29 = _rhs94;
          _lhs69[_lhs70] = _rhs95;
          let _587_v34;
          _587_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,((_this).f23) === ((_this).f23));
          _587_v34 = (_587_v34).update(_dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(42)), ((_588_v27) => function (_589_i1) {
            return _588_v27.f18;
          })(_580_v27)), _554_v2), (_this).f23);
          let _590_v35;
          _590_v35 = _dafny.Seq.of(_this.f22, p1, p2);
          let _591_v36;
          _591_v36 = _dafny.Set.fromElements(new BigNumber(304), new BigNumber((_590_v35).length));
          let _index103 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_583_v30).length));
          (_583_v30)[_index103] = _module.__default.fm1(p0, _580_v27.f18, (new BigNumber((_591_v36).length)).minus((_590_v35)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-813)), ((_592_v27) => function (_593_i2) {
            return _592_v27.f18;
          })(_580_v27))).length), new BigNumber((_590_v35).length))]), globalState);
          let _594_v37;
          _594_v37 = _dafny.Seq.of((_585_v32).f20, _554_v2, _dafny.Seq.Concat(_554_v2, _554_v2));
          let _rhs96 = (_590_v35)[_module.__default.safeIndex(new BigNumber((_555_v3).length), new BigNumber((_590_v35).length))];
          let _rhs97 = _594_v37;
          let _rhs98 = _module.__default.safeModuloInt((_module.D5.create_DC14(!(false), p1)).dtor_cf17, _module.__default.safeDivisionInt(new BigNumber(-221), new BigNumber(274)));
          let _lhs71 = globalState;
          let _lhs72 = globalState;
          _lhs71.f6 = _rhs96;
          _594_v37 = _rhs97;
          _lhs72.f6 = _rhs98;
          _554_v2 = _dafny.Seq.update((_585_v32).f20, _module.__default.safeIndex(p1, new BigNumber(((_585_v32).f20).length)), _580_v27.f18);
        }
        if ((_this).fm3(!((p0) && ((_this).f23)), (new BigNumber((_554_v2).length)).plus(new BigNumber((_554_v2).length)), globalState)) {
          (_this).f18 = _this.f18;
          let _595_v38;
          _595_v38 = false;
          _595_v38 = !(p0);
          (_this).f22 = p2;
          let _596_v39;
          _596_v39 = _module.D1.create_DC4(p3, p1, _595_v38);
          let _597_v40;
          _597_v40 = _dafny.Map.Empty.slice().updateUnsafe((_596_v39).dtor_cf5,(_this).f23);
          _597_v40 = (_597_v40).update((_this).f23, (_this).f23);
          (_this).f22 = (_dafny.ZERO).minus(p2);
        } else {
          let _598_v41;
          _598_v41 = _dafny.MultiSet.fromElements(p1, p1, p1, p1);
          let _599_v42;
          _599_v42 = _module.D3.create_DC10((_598_v41).IsSubsetOf(_598_v41));
          _599_v42 = (((_this).f23) ? (_module.D3.create_DC10((_this).f23)) : (_599_v42));
          let _600_v43;
          let _init9 = ((_601_p3) => function (_602_i3) {
            return _module.__default.safeModuloInt(_602_i3, _601_p3);
          })(p3);
          let _nw101 = Array((new BigNumber(10)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw101.length); _i0_9++) {
            _nw101[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _600_v43 = _nw101;
          let _index104 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_600_v43).length));
          (_600_v43)[_index104] = p3;
          let _603_v44;
          _603_v44 = _dafny.Set.fromElements(p2, p3);
          let _index105 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_600_v43).length));
          (_600_v43)[_index105] = (p2).plus((new BigNumber(693)).multipliedBy(new BigNumber((_603_v44).length)));
          let _604_v45;
          _604_v45 = _dafny.Seq.of(_552_v0, _552_v0, _552_v0, _552_v0, _552_v0);
          let _605_v46;
          _605_v46 = _dafny.Seq.of(_604_v45);
          let _606_v47;
          _606_v47 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("tqhnwiyw"), _554_v2, _554_v2);
          let _607_v48;
          _607_v48 = _dafny.Map.Empty.slice().updateUnsafe(_554_v2,(_605_v46)[_module.__default.safeIndex(new BigNumber((_606_v47).length), new BigNumber((_605_v46).length))]);
          _607_v48 = (_607_v48).update(_dafny.Seq.Concat(_554_v2, _dafny.Seq.UnicodeFromString("jrv")), _604_v45);
          let _608_v49;
          _608_v49 = true;
          let _609_v50;
          _609_v50 = _dafny.MultiSet.fromElements((_554_v2)[_module.__default.safeIndex(p2, new BigNumber((_554_v2).length))], _this.f18);
          _608_v49 = ((_608_v49) ? ((_561_v8)[_module.__default.safeIndex((((_609_v50).contains(_this.f18)) ? ((_609_v50).get(_this.f18)) : (new BigNumber((_598_v41).cardinality()))), new BigNumber((_561_v8).length))]) : ((new BigNumber((_609_v50).cardinality())).isLessThanOrEqualTo((_600_v43)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_600_v43).length))])));
          let _610_v51;
          _610_v51 = _dafny.Map.Empty.slice().updateUnsafe(_608_v49,_this.f22);
          _610_v51 = (_610_v51).update(p0, (new BigNumber(843)).minus(p2));
        }
        if ((_561_v8)[_module.__default.safeIndex(_this.f22, new BigNumber((_561_v8).length))]) {
          let _611_v52;
          let _nw102 = Array((new BigNumber(24)).toNumber()).fill([]);
          _611_v52 = _nw102;
          let _612_v53;
          _612_v53 = _module.D6.create_DC16(!((_this).f23), _611_v52);
          let _613_v54;
          _613_v54 = _module.D6.create_DC17(_612_v53);
          let _614_v55;
          _614_v55 = _dafny.Map.Empty.slice().updateUnsafe(_613_v54,new _dafny.CodePoint('y'.codePointAt(0)));
          _614_v55 = (_614_v55).update(_613_v54, _this.f18);
          (_this).f22 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), ((_615_p3, _616_p0) => function (_617_i4) {
            return _module.__default.safeDivisionInt(_615_p3, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_616_p0)).length));
          })(p3, p0))).length);
          let _618_v56;
          _618_v56 = _dafny.MultiSet.fromElements(new BigNumber(527));
          let _619_v57;
          _619_v57 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p2,_this.f18));
          let _620_v58;
          _620_v58 = _dafny.Set.fromElements((_this).f23);
          let _621_v59;
          let _nw103 = Array((new BigNumber(27)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = p0;
          _nw103[(_dafny.ONE).toNumber()] = (_this).f23;
          _nw103[(new BigNumber(2)).toNumber()] = (_this).f23;
          _nw103[(new BigNumber(3)).toNumber()] = !(true) || (!(_module.__default.fm1(p0, _this.f18, new BigNumber((_618_v56).cardinality()), globalState)));
          _nw103[(new BigNumber(4)).toNumber()] = !((_this).f23);
          _nw103[(new BigNumber(5)).toNumber()] = p0;
          _nw103[(new BigNumber(6)).toNumber()] = true;
          _nw103[(new BigNumber(7)).toNumber()] = (_this).fm3(false, new BigNumber((_dafny.Seq.of(_this.f22)).length), globalState);
          _nw103[(new BigNumber(8)).toNumber()] = p0;
          _nw103[(new BigNumber(9)).toNumber()] = !((_this).f23);
          _nw103[(new BigNumber(10)).toNumber()] = (_this).f23;
          _nw103[(new BigNumber(11)).toNumber()] = (_this).f23;
          _nw103[(new BigNumber(12)).toNumber()] = true;
          _nw103[(new BigNumber(13)).toNumber()] = (_561_v8)[_module.__default.safeIndex(p2, new BigNumber((_561_v8).length))];
          _nw103[(new BigNumber(14)).toNumber()] = p0;
          _nw103[(new BigNumber(15)).toNumber()] = (p2).isLessThanOrEqualTo(_this.f22);
          _nw103[(new BigNumber(16)).toNumber()] = (_619_v57).IsProperSubsetOf(_619_v57);
          _nw103[(new BigNumber(17)).toNumber()] = p0;
          _nw103[(new BigNumber(18)).toNumber()] = true;
          _nw103[(new BigNumber(19)).toNumber()] = p0;
          _nw103[(new BigNumber(20)).toNumber()] = !((_this).f23);
          _nw103[(new BigNumber(21)).toNumber()] = (_this).f23;
          _nw103[(new BigNumber(22)).toNumber()] = (_this).f23;
          _nw103[(new BigNumber(23)).toNumber()] = (_620_v58).IsDisjointFrom(_620_v58);
          _nw103[(new BigNumber(24)).toNumber()] = p0;
          _nw103[(new BigNumber(25)).toNumber()] = p0;
          _nw103[(new BigNumber(26)).toNumber()] = (_this).f23;
          _621_v59 = _nw103;
          let _index106 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_621_v59).length));
          (_621_v59)[_index106] = p0;
          let _index107 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_621_v59).length));
          (_621_v59)[_index107] = (((p1).isLessThanOrEqualTo(new BigNumber((_554_v2).length))) ? (p0) : (!(p2).isEqualTo(p2)));
          let _622_v60;
          _622_v60 = _module.D1.create_DC4(p2, new BigNumber(192), (_this).f23);
          let _623_v61;
          _623_v61 = _dafny.Seq.of((_this).fm4(_module.__default.fm1((_this).f23, _this.f18, p3, globalState), _this.f18, new _dafny.CodePoint('d'.codePointAt(0)), !(p0), globalState));
          let _624_v62;
          _624_v62 = _dafny.Set.fromElements(p3, _this.f22, new BigNumber(241), p2, new BigNumber((_618_v56).cardinality()));
          let _625_v63;
          let _nw104 = Array((new BigNumber(24)).toNumber());
          _nw104[(_dafny.ZERO).toNumber()] = _622_v60;
          _nw104[(_dafny.ONE).toNumber()] = _622_v60;
          _nw104[(new BigNumber(2)).toNumber()] = _module.D1.create_DC4(p1, p1, true);
          _nw104[(new BigNumber(3)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(4)).toNumber()] = _module.D1.create_DC4(p1, p3, (_this).f23);
          _nw104[(new BigNumber(5)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(6)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(7)).toNumber()] = _module.D1.create_DC4(new BigNumber((_623_v61).length), new BigNumber((_624_v62).length), false);
          _nw104[(new BigNumber(8)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(9)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(10)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(11)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(12)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(13)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(14)).toNumber()] = _module.D1.create_DC4(p3, p3, (_this).f23);
          _nw104[(new BigNumber(15)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(16)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(17)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(18)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(19)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(20)).toNumber()] = _module.D1.create_DC4(_this.f22, p2, p0);
          _nw104[(new BigNumber(21)).toNumber()] = _622_v60;
          _nw104[(new BigNumber(22)).toNumber()] = _module.D1.create_DC4(new BigNumber((_561_v8).length), _this.f22, (_561_v8)[_module.__default.safeIndex(p3, new BigNumber((_561_v8).length))]);
          _nw104[(new BigNumber(23)).toNumber()] = _622_v60;
          _625_v63 = _nw104;
          let _626_v64;
          _626_v64 = _dafny.MultiSet.fromElements(_625_v63, _625_v63, _625_v63);
          _626_v64 = _626_v64;
          let _627_v65;
          _627_v65 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,p3);
          let _628_v66;
          let _nw105 = new _module.C0();
          _nw105.__ctor(_dafny.Seq.UnicodeFromString("scd"), _627_v65, _this.f18, (_this).f19);
          _628_v66 = _nw105;
        } else {
          (globalState).f6 = p3;
          let _629_v67;
          _629_v67 = _dafny.MultiSet.fromElements(new BigNumber((_561_v8).length), p3, new BigNumber(542));
          let _rhs99 = p2;
          let _rhs100 = _629_v67;
          let _lhs73 = _this;
          _lhs73.f22 = _rhs99;
          _629_v67 = _rhs100;
          let _630_v68;
          _630_v68 = _module.D3.create_DC8((_this).fm13((_dafny.ZERO).minus(p2), p1, p2, (_this).f23, globalState));
          let _631_v69;
          _631_v69 = _dafny.Map.Empty.slice().updateUnsafe(_630_v68,(_this).f23);
          let _632_v70;
          _632_v70 = _dafny.Set.fromElements(p0, !(p0), p0);
          _631_v69 = (_631_v69).update(function (_pat_let12_0) {
            return function (_633_dt__update__tmp_h0) {
              return function (_pat_let13_0) {
                return function (_634_dt__update_hcf10_h0) {
                  return _module.D3.create_DC8(_634_dt__update_hcf10_h0);
                }(_pat_let13_0);
              }(new _dafny.CodePoint('a'.codePointAt(0)));
            }(_pat_let12_0);
          }(_630_v68), (_dafny.Set.fromElements(p0)).IsProperSubsetOf(_632_v70));
          let _635_v71;
          _635_v71 = true;
          _635_v71 = _635_v71;
          let _636_v72;
          let _nw106 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _636_v72 = _nw106;
          let _index108 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_636_v72).length));
          (_636_v72)[_index108] = _module.__default.safeModuloInt(p3, p3);
          let _637_v73;
          _637_v73 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,new BigNumber(931));
          let _index109 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_636_v72).length));
          let _rhs101 = _this.f22;
          let _rhs102 = _module.__default.safeDivisionInt((((_629_v67).contains(p1)) ? ((_629_v67).get(p1)) : ((((_637_v73).contains(_this.f22)) ? ((_637_v73).get(_this.f22)) : (new BigNumber(698))))), (_dafny.ZERO).minus(((_635_v71) ? (p1) : (_this.f22))));
          let _rhs103 = (p0) && (false);
          let _lhs74 = _636_v72;
          let _lhs75 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_636_v72).length));
          let _lhs76 = globalState;
          _lhs74[_lhs75] = _rhs101;
          _lhs76.f6 = _rhs102;
          _635_v71 = _rhs103;
        }
        let _638_v74;
        _638_v74 = true;
        _638_v74 = _module.__default.fm1((_553_v1).IsSubsetOf(_553_v1), new _dafny.CodePoint('y'.codePointAt(0)), p3, globalState);
      }
      (globalState).f6 = new BigNumber(-101);
      let _639_i5;
      _639_i5 = _dafny.ZERO;
      L6: {
        while ((_561_v8)[_module.__default.safeIndex(new BigNumber(297), new BigNumber((_561_v8).length))]) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_639_i5)) {
              break L6;
            }
            _639_i5 = (_639_i5).plus(_dafny.ONE);
            let _640_v75;
            let _init10 = ((_641_p3) => function (_642_i6) {
              return (_642_i6).plus(_641_p3);
            })(p3);
            let _nw107 = Array((new BigNumber(4)).toNumber());
            for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw107.length); _i0_10++) {
              _nw107[_i0_10] = _init10(new BigNumber(_i0_10));
            }
            _640_v75 = _nw107;
            let _643_v76;
            _643_v76 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,false);
            let _644_v77;
            _644_v77 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus(new BigNumber((_643_v76).length)));
            let _index110 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_640_v75).length));
            (_640_v75)[_index110] = _module.__default.safeDivisionInt(p1, new BigNumber((_644_v77).length));
            let _index111 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_640_v75).length));
            (_640_v75)[_index111] = p1;
            let _645_v78;
            _645_v78 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(p0, (_this).f23)).length), p1);
            let _646_v79;
            _646_v79 = _module.D3.create_DC9(p0, new BigNumber((_645_v78).cardinality()));
            let _647_v80;
            _647_v80 = _module.D1.create_DC4((_646_v79).dtor_cf12, _this.f22, (_this).f23);
            _647_v80 = _647_v80;
            (globalState).f6 = _this.f22;
            _640_v75 = _640_v75;
          }
        }
      }
      r0 = (p1).plus(_module.__default.safeModuloInt(new BigNumber(599), _this.f22));
      return r0;
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f18 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f19 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    set f18(value) {
      let _this = this;
      _this._f18 = value;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    __ctor(f18, f19) {
      let _this = this;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      return;
    }
    fm3(p0, p1, globalState) {
      let _this = this;
      if ((new BigNumber((_dafny.Seq.UnicodeFromString("fywgn")).length)).isLessThanOrEqualTo(new BigNumber(896))) {
        return (new BigNumber((_dafny.Seq.UnicodeFromString("bnebdr")).length)).isLessThanOrEqualTo(new BigNumber(329));
      } else {
        return true;
      }
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(false, false, !(true), false))).length)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(971)), function (_648_i0) {
        return new BigNumber((_dafny.Seq.of(true, !(!(false)), true, !(!(true)), !(false))).length);
      })).length)));
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = false;
      let _649_v0;
      let _init11 = function (_650_i0) {
        return (_650_i0).minus(new BigNumber((_dafny.Seq.UnicodeFromString("fugltux")).length));
      };
      let _nw108 = Array((new BigNumber(15)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw108.length); _i0_11++) {
        _nw108[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _649_v0 = _nw108;
      let _651_v1;
      _651_v1 = new BigNumber(-605);
      let _index112 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length));
      (_649_v0)[_index112] = _651_v1;
      let _652_v2;
      let _init12 = ((_653_p0, _654_v1) => function (_655_i1) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("yxgl"),_653_p0)).length),_module.D0.create_DC0(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-553)), ((_656_v1) => function (_657_i2) {
  return _656_v1;
})(_654_v1))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_654_v1,_module.D0.create_DC0(_654_v1)));
      })(p0, _651_v1);
      let _nw109 = Array((new BigNumber(28)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw109.length); _i0_12++) {
        _nw109[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _652_v2 = _nw109;
      let _index113 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_649_v0).length));
      (_649_v0)[_index113] = _651_v1;
      let _658_v3;
      _658_v3 = _dafny.Seq.UnicodeFromString("uuruhxs");
      let _659_v4;
      _659_v4 = _dafny.Map.Empty.slice().updateUnsafe(_651_v1,_658_v3);
      let _660_v5;
      let _nw110 = Array((new BigNumber(5)).toNumber()).fill(false);
      _660_v5 = _nw110;
      let _661_v6;
      _661_v6 = _dafny.MultiSet.fromElements(_660_v5, _660_v5);
      let _662_v7;
      _662_v7 = _dafny.Seq.of(!(p0));
      let _663_v8;
      _663_v8 = _662_v7;
      let _pat_let_tv10 = _651_v1;
      let _pat_let_tv11 = _651_v1;
      let _index114 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length));
      let _index115 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_649_v0).length));
      let _rhs104 = (new BigNumber((_659_v4).length)).isLessThanOrEqualTo((new BigNumber(579)).multipliedBy(new BigNumber((_661_v6).cardinality())));
      let _rhs105 = _this.f18;
      let _rhs106 = _651_v1;
      let _rhs107 = _652_v2;
      let _rhs108 = function (_source7) {
        let _664___mcc_h0 = _source7;
        let _665_cf9 = _664___mcc_h0;
        return (_pat_let_tv10).multipliedBy(_pat_let_tv11);
      }(_663_v8);
      let _lhs77 = _this;
      let _lhs78 = _649_v0;
      let _lhs79 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length));
      let _lhs80 = _649_v0;
      let _lhs81 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_649_v0).length));
      r2 = _rhs104;
      _lhs77.f18 = _rhs105;
      _lhs78[_lhs79] = _rhs106;
      _652_v2 = _rhs107;
      _lhs80[_lhs81] = _rhs108;
      let _666_v9;
      _666_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,(_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))]);
      let _667_v10;
      let _nw111 = new _module.C0();
      _nw111.__ctor(_658_v3, _666_v9, _this.f18, (_this).f19);
      _667_v10 = _nw111;
      _651_v1 = (_651_v1).multipliedBy(new BigNumber(988));
      let _668_v11;
      _668_v11 = _dafny.Map.Empty.slice().updateUnsafe((_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))],(_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))]);
      let _669_v12;
      _669_v12 = _dafny.Seq.of((new BigNumber((_668_v11).length)).plus(_651_v1), (_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))], (_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))]);
      r1 = (_669_v12)[_module.__default.safeIndex(_651_v1, new BigNumber((_669_v12).length))];
      let _hi4 = _651_v1;
      for (let _670_i3 = (_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))]; _670_i3.isLessThan(_hi4); _670_i3 = _670_i3.plus(_dafny.ONE)) {
        let _671_v13;
        let _out23;
        _out23 = (_this).m4((_670_i3).multipliedBy(_651_v1), _module.__default.fm0(globalState), _660_v5, globalState);
        _671_v13 = _out23;
        r3 = _671_v13;
        let _672_v14;
        _672_v14 = _module.D0.create_DC1();
        let _source8 = _672_v14;
        if (_source8.is_DC1) {
          let _673_v15;
          _673_v15 = _dafny.MultiSet.fromElements(_671_v13);
          let _index116 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length));
          (_649_v0)[_index116] = new BigNumber((_673_v15).cardinality());
          (globalState).f13 = _dafny.Seq.Concat(_dafny.Seq.Concat(_662_v7, _662_v7), _dafny.Seq.of(_671_v13));
          let _index117 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_660_v5).length));
          (_660_v5)[_index117] = false;
          let _index118 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_660_v5).length));
          (_660_v5)[_index118] = _671_v13;
          let _674_v16;
          _674_v16 = _dafny.MultiSet.fromElements(_649_v0, _649_v0, _649_v0);
          let _675_v17;
          _675_v17 = _dafny.Seq.of(_674_v16, _674_v16);
          let _index119 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_660_v5).length));
          let _index120 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_660_v5).length));
          let _rhs109 = _671_v13;
          let _rhs110 = _module.__default.fm0(globalState);
          let _rhs111 = !((_675_v17)[_module.__default.safeIndex(_651_v1, new BigNumber((_675_v17).length))]).equals((_674_v16).update(_649_v0, _module.__default.abs(new BigNumber(((_667_v10).f20).length))));
          let _rhs112 = (_module.__default.safeModuloInt((_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))], _651_v1)).isLessThanOrEqualTo(_670_i3);
          let _rhs113 = p0;
          let _lhs82 = globalState;
          let _lhs83 = _660_v5;
          let _lhs84 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_660_v5).length));
          let _lhs85 = _660_v5;
          let _lhs86 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_660_v5).length));
          _671_v13 = _rhs109;
          _lhs82.f6 = _rhs110;
          _671_v13 = _rhs111;
          _lhs83[_lhs84] = _rhs112;
          _lhs85[_lhs86] = _rhs113;
          _671_v13 = _671_v13;
        } else if (_source8.is_DC0) {
          let _676___mcc_h1 = (_source8).cf0;
          let _677_cf0 = _676___mcc_h1;
          let _678_v18;
          _678_v18 = _dafny.Set.fromElements(_671_v13);
          _678_v18 = _678_v18;
          (globalState).f6 = (_651_v1).multipliedBy(_651_v1);
          let _679_v19;
          let _nw112 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.of());
          _679_v19 = _nw112;
          let _index121 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_679_v19).length));
          (_679_v19)[_index121] = _dafny.Seq.of(_671_v13);
          let _index122 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_679_v19).length));
          (_679_v19)[_index122] = _663_v8;
          let _680_v20;
          _680_v20 = _dafny.Map.Empty.slice().updateUnsafe(_658_v3,!(p0));
          let _681_v21;
          _681_v21 = _dafny.Map.Empty.slice().updateUnsafe(((_671_v13) ? (new BigNumber(-196)) : (_670_i3)),(_dafny.Map.Empty.slice().updateUnsafe((_667_v10).f20,_671_v13)).Merge(_680_v20));
          _681_v21 = (_681_v21).update((_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))], ((_671_v13) ? (_module.__default.fm11(globalState)) : ((_680_v20).update((_667_v10).f20, true))));
        } else {
          let _682___mcc_h2 = (_source8).cf1;
          let _683_cf1 = _682___mcc_h2;
          let _684_v22;
          let _nw113 = new _module.C0();
          _nw113.__ctor(_dafny.Seq.UnicodeFromString("qfwpi"), _dafny.Map.Empty.slice().updateUnsafe(_this.f18,_651_v1), new _dafny.CodePoint('m'.codePointAt(0)), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update((_667_v10).f20, _module.__default.safeIndex((_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))], new BigNumber(((_667_v10).f20).length)), _this.f18),(_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))]));
          _684_v22 = _nw113;
          let _685_v23;
          _685_v23 = _684_v22;
          let _686_v24;
          _686_v24 = (_685_v23);
          _684_v22 = (_684_v22);
          let _687_v26;
          _687_v26 = _dafny.MultiSet.fromElements(_658_v3);
          let _688_v27;
          let _nw114 = new _module.C0();
          _nw114.__ctor((_667_v10).f20, _666_v9, new _dafny.CodePoint('l'.codePointAt(0)), function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of (_687_v26).Elements) {
              let _689_v25 = _compr_32;
              if ((_687_v26).contains(_689_v25)) {
                _coll32.push([_689_v25,new BigNumber(((_667_v10).f20).length)]);
              }
            }
            return _coll32;
          }());
          _688_v27 = _nw114;
          let _690_v28;
          _690_v28 = _dafny.Map.Empty.slice().updateUnsafe(_667_v10,(_dafny.ZERO).minus(_651_v1));
          let _rhs114 = ((((_690_v28).contains(_688_v27)) ? ((_690_v28).get(_688_v27)) : ((_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))]))).isLessThan(_651_v1);
          let _rhs115 = new BigNumber(6);
          _671_v13 = _rhs114;
          _651_v1 = _rhs115;
          let _691_v29;
          let _nw115 = new _module.C1();
          _nw115.__ctor(_651_v1, _671_v13, _this.f18, _module.__default.fm18(new BigNumber(((_667_v10).f20).length), globalState));
          _691_v29 = _nw115;
          let _692_v30;
          _692_v30 = _dafny.Map.Empty.slice().updateUnsafe(_662_v7,_691_v29);
          let _693_v31;
          _693_v31 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(_662_v7,_691_v29));
          let _rhs116 = ((((_693_v31).contains(p0)) ? ((_693_v31).get(p0)) : (_692_v30))).Merge(_692_v30);
          let _rhs117 = (_651_v1).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(357)), ((_694_v29) => function (_695_i4) {
            return _694_v29.f18;
          })(_691_v29))).length));
          _692_v30 = _rhs116;
          _651_v1 = _rhs117;
        }
        r2 = !(p0);
      }
      let _index123 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_660_v5).length));
      (_660_v5)[_index123] = p0;
      let _index124 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length));
      let _index125 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_660_v5).length));
      let _rhs118 = (_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))];
      let _rhs119 = _dafny.areEqual((_667_v10).f20, (_667_v10).f20);
      let _rhs120 = _this.f18;
      let _rhs121 = _651_v1;
      let _lhs87 = _649_v0;
      let _lhs88 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length));
      let _lhs89 = _660_v5;
      let _lhs90 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_660_v5).length));
      let _lhs91 = _this;
      _lhs87[_lhs88] = _rhs118;
      _lhs89[_lhs90] = _rhs119;
      _lhs91.f18 = _rhs120;
      _651_v1 = _rhs121;
      r0 = _module.D0.create_DC0(((_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))]).minus((_this).fm4(p0, _this.f18, new _dafny.CodePoint('r'.codePointAt(0)), (_660_v5)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_660_v5).length))], globalState)));
      r1 = (((_668_v11).contains((_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))])) ? ((_668_v11).get((_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))])) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(380)), ((_696_v11) => function (_697_i5) {
        return _696_v11;
      })(_668_v11))).length)));
      r2 = (((_660_v5)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_660_v5).length))]) ? (p0) : (((_649_v0)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_649_v0).length))]).isLessThan(_651_v1)));
      let _698_v32;
      _698_v32 = _dafny.Set.fromElements(!(false));
      r3 = (_698_v32).IsSubsetOf(_698_v32);
      return [r0, r1, r2, r3];
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      r0 = p2;
      if (p0) {
        r0 = p3;
        r0 = p3;
        let _699_v0;
        let _nw116 = Array((new BigNumber(20)).toNumber()).fill(false);
        _699_v0 = _nw116;
        _699_v0 = _699_v0;
        let _index126 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_699_v0).length));
        (_699_v0)[_index126] = p0;
        let _index127 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_699_v0).length));
        (_699_v0)[_index127] = ((_dafny.ZERO).minus(p3)).isLessThan(p3);
        let _index128 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_699_v0).length));
        (_699_v0)[_index128] = (true) === ((p3).isLessThan(p2));
      } else {
        let _700_v1;
        _700_v1 = false;
        let _701_v2;
        _701_v2 = _dafny.Set.fromElements(_700_v1);
        _700_v1 = (_701_v2).IsSubsetOf(_701_v2);
        let _702_v3;
        _702_v3 = _dafny.Seq.UnicodeFromString("wrx");
        let _703_v4;
        _703_v4 = _dafny.Map.Empty.slice().updateUnsafe(_700_v1,!(_700_v1));
        let _704_v5;
        _704_v5 = _dafny.Map.Empty.slice().updateUnsafe((_702_v3)[_module.__default.safeIndex(new BigNumber((_703_v4).length), new BigNumber((_702_v3).length))],p2);
        let _705_v6;
        _705_v6 = _module.D6.create_DC15((_704_v5).Merge(_704_v5));
        let _706_v7;
        let _nw117 = Array((new BigNumber(8)).toNumber());
        _nw117[(_dafny.ZERO).toNumber()] = _702_v3;
        _nw117[(_dafny.ONE).toNumber()] = _702_v3;
        _nw117[(new BigNumber(2)).toNumber()] = _702_v3;
        _nw117[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_702_v3, _module.__default.safeIndex(p2, new BigNumber((_702_v3).length)), new _dafny.CodePoint('p'.codePointAt(0)));
        _nw117[(new BigNumber(4)).toNumber()] = _702_v3;
        _nw117[(new BigNumber(5)).toNumber()] = _702_v3;
        _nw117[(new BigNumber(6)).toNumber()] = _702_v3;
        _nw117[(new BigNumber(7)).toNumber()] = _702_v3;
        _706_v7 = _nw117;
        let _707_v8;
        let _nw118 = Array((new BigNumber(17)).toNumber()).fill(false);
        _707_v8 = _nw118;
        let _708_v9;
        _708_v9 = _dafny.MultiSet.fromElements(_700_v1, !(false), _700_v1, _module.__default.fm1(_700_v1, new _dafny.CodePoint('r'.codePointAt(0)), p2, globalState));
        let _index129 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_707_v8).length));
        (_707_v8)[_index129] = (_708_v9).IsSubsetOf(_708_v9);
        let _709_v10;
        _709_v10 = _module.D0.create_DC0((_dafny.ZERO).minus(p1));
        let _index130 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_707_v8).length));
        let _rhs122 = _module.D6.create_DC15((_704_v5).Merge(_704_v5));
        let _rhs123 = _706_v7;
        let _rhs124 = _module.__default.fm1(_700_v1, _this.f18, new BigNumber(-405), globalState);
        let _rhs125 = _module.D0.create_DC0(p2);
        let _rhs126 = _this.f18;
        let _lhs92 = _707_v8;
        let _lhs93 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_707_v8).length));
        let _lhs94 = _this;
        _705_v6 = _rhs122;
        _706_v7 = _rhs123;
        _lhs92[_lhs93] = _rhs124;
        _709_v10 = _rhs125;
        _lhs94.f18 = _rhs126;
        let _710_v11;
        _710_v11 = _module.D1.create_DC6(true, p0);
        let _index131 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_707_v8).length));
        (_707_v8)[_index131] = (_710_v11).dtor_cf7;
        let _711_v12;
        _711_v12 = _dafny.Map.Empty.slice().updateUnsafe(p1,!((_707_v8)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_707_v8).length))]));
        let _712_v13;
        _712_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_702_v3).length),(((_711_v12).contains(p3)) ? ((_711_v12).get(p3)) : (p0)));
        _712_v13 = (_712_v13).update(p1, _700_v1);
        (globalState).f6 = _module.__default.fm0(globalState);
      }
      let _713_v14;
      _713_v14 = true;
      let _714_v15;
      _714_v15 = _dafny.Set.fromElements(p0);
      let _715_v16;
      _715_v16 = _module.D7.create_DC18(_714_v15);
      _713_v14 = (_714_v15).equals((_715_v16).dtor_cf22);
      let _716_v17;
      _716_v17 = _dafny.Map.Empty.slice().updateUnsafe(_713_v14,p0);
      let _717_v18;
      _717_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((((_716_v17).update(p0, _713_v14)).Merge(_716_v17)).length),p2);
      let _718_v19;
      _718_v19 = _dafny.Seq.UnicodeFromString("oiyohkc");
      let _719_v20;
      let _nw119 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
      _719_v20 = _nw119;
      let _index132 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_719_v20).length));
      (_719_v20)[_index132] = p1;
      let _720_v21;
      _720_v21 = _dafny.MultiSet.fromElements(_713_v14, _713_v14);
      let _721_v22;
      _721_v22 = _module.D5.create_DC12(_720_v21);
      let _pat_let_tv12 = p2;
      let _pat_let_tv13 = p1;
      let _index133 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_719_v20).length));
      let _rhs127 = (_module.__default.fm19(p0, p0, globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p2,p2));
      let _rhs128 = _718_v19;
      let _rhs129 = function (_source9) {
        if (_source9.is_DC13) {
          return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("k")).length)), (_dafny.Seq.of(_pat_let_tv12))))).cardinality());
        } else if (_source9.is_DC14) {
          let _722___mcc_h0 = (_source9).cf16;
          let _723___mcc_h1 = (_source9).cf17;
          let _724_cf17 = _723___mcc_h1;
          let _725_cf16 = _722___mcc_h0;
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(378)), function (_726_i0) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          })).length));
        } else {
          let _727___mcc_h2 = (_source9).cf15;
          let _728_cf15 = _727___mcc_h2;
          return _pat_let_tv13;
        }
      }(_721_v22);
      let _lhs95 = _719_v20;
      let _lhs96 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_719_v20).length));
      _717_v18 = _rhs127;
      _718_v19 = _rhs128;
      _lhs95[_lhs96] = _rhs129;
      let _729_v23;
      _729_v23 = _dafny.Map.Empty.slice().updateUnsafe(_713_v14,p3);
      _729_v23 = (_729_v23).update(p0, p2);
      let _730_v24;
      _730_v24 = _dafny.MultiSet.fromElements(new BigNumber(296));
      let _731_v25;
      let _nw120 = new _module.C1();
      _nw120.__ctor(p3, (_730_v24).contains((_719_v20)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_719_v20).length))]), new _dafny.CodePoint('y'.codePointAt(0)), ((_this).f19).Merge((_this).f19));
      _731_v25 = _nw120;
      let _732_v26;
      _732_v26 = _dafny.Seq.of(_731_v25);
      _731_v25 = (_732_v26)[_module.__default.safeIndex(p3, new BigNumber((_732_v26).length))];
      r0 = p1;
      return r0;
    }
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _hi5 = p1;
      for (let _733_i0 = (p1).plus(p1); _733_i0.isLessThan(_hi5); _733_i0 = _733_i0.plus(_dafny.ONE)) {
        let _734_v0;
        _734_v0 = false;
        let _735_v1;
        _735_v1 = _dafny.Map.Empty.slice().updateUnsafe(_734_v0,p0);
        let _736_v2;
        _736_v2 = _dafny.MultiSet.fromElements(_734_v0);
        let _737_v3;
        _737_v3 = _dafny.Set.fromElements(new BigNumber((_735_v1).length), new BigNumber((_736_v2).cardinality()));
        let _738_v4;
        _738_v4 = _dafny.Map.Empty.slice().updateUnsafe(_737_v3,(p1).multipliedBy(p0));
        _738_v4 = (_738_v4).update(_737_v3, _module.__default.fm0(globalState));
        let _739_v5;
        _739_v5 = _dafny.Seq.UnicodeFromString("gtyc");
        (globalState).f6 = (p1).minus(new BigNumber((_dafny.Seq.of(_739_v5, _739_v5)).length));
        _734_v0 = true;
        let _740_v6;
        _740_v6 = _dafny.Map.Empty.slice().updateUnsafe(_739_v5,_734_v0);
        let _index134 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((p2).length));
        (p2)[_index134] = (((_740_v6).contains(_739_v5)) ? ((_740_v6).get(_739_v5)) : (_734_v0));
        let _741_v7;
        _741_v7 = _module.D5.create_DC13();
        let _742_v8;
        _742_v8 = _dafny.Map.Empty.slice().updateUnsafe(_741_v7,_733_i0);
        let _index135 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((p2).length));
        let _rhs130 = !(_module.__default.fm2(globalState)).contains(_734_v0);
        let _rhs131 = _742_v8;
        let _rhs132 = _module.__default.fm1((_734_v0) || (_734_v0), new _dafny.CodePoint('v'.codePointAt(0)), (_dafny.ZERO).minus(p1), globalState);
        let _lhs97 = p2;
        let _lhs98 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((p2).length));
        _lhs97[_lhs98] = _rhs130;
        _742_v8 = _rhs131;
        _734_v0 = _rhs132;
      }
      let _743_v9;
      _743_v9 = false;
      let _744_i1;
      _744_i1 = _dafny.ZERO;
      L7: {
        while (_743_v9) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_744_i1)) {
              break L7;
            }
            _744_i1 = (_744_i1).plus(_dafny.ONE);
            (globalState).f6 = (p0).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(p1)));
            let _745_v10;
            _745_v10 = _dafny.Seq.UnicodeFromString("yesur");
            let _746_v11;
            _746_v11 = _dafny.Map.Empty.slice().updateUnsafe(_743_v9,(_745_v10)[_module.__default.safeIndex(p1, new BigNumber((_745_v10).length))]);
            let _747_v12;
            _747_v12 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_743_v9,new _dafny.CodePoint('c'.codePointAt(0))));
            let _748_v13;
            _748_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,_745_v10);
            _746_v11 = (_747_v12)[_module.__default.safeIndex(new BigNumber(((((_748_v13).contains(new BigNumber(-768))) ? ((_748_v13).get(new BigNumber(-768))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(175)), function (_749_i2) {
              return _this.f18;
            })))).length), new BigNumber((_747_v12).length))];
            r0 = _module.__default.fm1(_743_v9, _this.f18, p1, globalState);
            let _750_v14;
            _750_v14 = _dafny.Map.Empty.slice().updateUnsafe(!(_743_v9),_743_v9);
            let _751_v15;
            _751_v15 = _dafny.Seq.of(_743_v9);
            let _752_v16;
            _752_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-623),_this.f18);
            let _753_v17;
            _753_v17 = _dafny.Map.Empty.slice().updateUnsafe(_751_v15,_752_v16);
            let _754_v19;
            _754_v19 = _dafny.Set.fromElements(_751_v15, _751_v15, _751_v15);
            let _755_v20;
            let _nw121 = Array((new BigNumber(21)).toNumber());
            _nw121[(_dafny.ZERO).toNumber()] = _743_v9;
            _nw121[(_dafny.ONE).toNumber()] = _743_v9;
            _nw121[(new BigNumber(2)).toNumber()] = ((_743_v9) ? (true) : (!((((_750_v14).contains(_743_v9)) ? ((_750_v14).get(_743_v9)) : (_743_v9)))));
            _nw121[(new BigNumber(3)).toNumber()] = _743_v9;
            _nw121[(new BigNumber(4)).toNumber()] = _743_v9;
            _nw121[(new BigNumber(5)).toNumber()] = _dafny.Seq.contains(_751_v15, _743_v9);
            _nw121[(new BigNumber(6)).toNumber()] = _743_v9;
            _nw121[(new BigNumber(7)).toNumber()] = true;
            _nw121[(new BigNumber(8)).toNumber()] = true;
            _nw121[(new BigNumber(9)).toNumber()] = !(((_743_v9) ? (_743_v9) : (true)));
            _nw121[(new BigNumber(10)).toNumber()] = _743_v9;
            _nw121[(new BigNumber(11)).toNumber()] = (_751_v15)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,true)).length), new BigNumber((_751_v15).length))];
            _nw121[(new BigNumber(12)).toNumber()] = _743_v9;
            _nw121[(new BigNumber(13)).toNumber()] = _743_v9;
            _nw121[(new BigNumber(14)).toNumber()] = !(_743_v9) || (_743_v9);
            _nw121[(new BigNumber(15)).toNumber()] = _743_v9;
            _nw121[(new BigNumber(16)).toNumber()] = (function () {
              let _coll33 = new _dafny.Set();
              for (const _compr_33 of (_753_v17).Keys.Elements) {
                let _756_v18 = _compr_33;
                if ((_753_v17).contains(_756_v18)) {
                  _coll33.add(_756_v18);
                }
              }
              return _coll33;
            }()).IsDisjointFrom(_754_v19);
            _nw121[(new BigNumber(17)).toNumber()] = _743_v9;
            _nw121[(new BigNumber(18)).toNumber()] = _743_v9;
            _nw121[(new BigNumber(19)).toNumber()] = _743_v9;
            _nw121[(new BigNumber(20)).toNumber()] = _743_v9;
            _755_v20 = _nw121;
            let _757_v21;
            _757_v21 = _dafny.Map.Empty.slice().updateUnsafe((_745_v10)[_module.__default.safeIndex(new BigNumber(316), new BigNumber((_745_v10).length))],p1);
            let _758_v22;
            let _nw122 = new _module.C0();
            _nw122.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(214)), function (_759_i3) {
              return _this.f18;
            }), _757_v21, _this.f18, (_this).f19);
            _758_v22 = _nw122;
            let _760_v23;
            _760_v23 = _dafny.Map.Empty.slice().updateUnsafe(_743_v9,_755_v20);
            let _rhs133 = p2;
            let _rhs134 = new BigNumber((((_743_v9) ? (_760_v23) : ((_760_v23).Merge(_760_v23)))).length);
            let _rhs135 = ((_743_v9) ? (_758_v22) : (_758_v22));
            let _rhs136 = _module.__default.fm20(_743_v9, globalState);
            let _lhs99 = globalState;
            let _lhs100 = _this;
            _755_v20 = _rhs133;
            _lhs99.f6 = _rhs134;
            _758_v22 = _rhs135;
            _lhs100.f18 = _rhs136;
          }
        }
      }
      let _761_v24;
      let _nw123 = Array((new BigNumber(28)).toNumber()).fill([]);
      _761_v24 = _nw123;
      let _index136 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_761_v24).length));
      (_761_v24)[_index136] = p2;
      let _index137 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_761_v24).length));
      (_761_v24)[_index137] = p2;
      let _762_v25;
      _762_v25 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(-174), p1));
      let _763_v26;
      _763_v26 = _dafny.Map.Empty.slice().updateUnsafe(_743_v9,p0);
      let _764_v27;
      _764_v27 = _dafny.Seq.UnicodeFromString("jcyiogjr");
      let _765_v28;
      _765_v28 = _dafny.Set.fromElements(new BigNumber((_763_v26).length), new BigNumber((_764_v27).length));
      let _766_v29;
      _766_v29 = _dafny.Seq.of(p0);
      _762_v25 = (_dafny.MultiSet.fromElements(_765_v28)).update((_765_v28).Union(function () {
        let _coll34 = new _dafny.Set();
        for (const _compr_34 of (_766_v29).Elements) {
          let _767_v30 = _compr_34;
          if (_dafny.Seq.contains(_766_v29, _767_v30)) {
            _coll34.add(_module.__default.safeDivisionInt(_767_v30, new BigNumber((function () {
              let _coll35 = new _dafny.Set();
              for (const _compr_35 of (function () {
                let _coll36 = new _dafny.Map();
                for (const _compr_36 of _dafny.IntegerRange(new BigNumber(814), new BigNumber(478))) {
                  let _768_v31 = _compr_36;
                  if (((new BigNumber(814)).isLessThanOrEqualTo(_768_v31)) && ((_768_v31).isLessThan(new BigNumber(478)))) {
                    _coll36.push([(_768_v31).plus(new BigNumber(-657)),new BigNumber(846)]);
                  }
                }
                return _coll36;
              }()).Keys.Elements) {
                let _769_v32 = _compr_35;
                if ((function () {
                  let _coll37 = new _dafny.Map();
                  for (const _compr_37 of _dafny.IntegerRange(new BigNumber(814), new BigNumber(478))) {
                    let _768_v31 = _compr_37;
                    if (((new BigNumber(814)).isLessThanOrEqualTo(_768_v31)) && ((_768_v31).isLessThan(new BigNumber(478)))) {
                      _coll37.push([(_768_v31).plus(new BigNumber(-657)),new BigNumber(846)]);
                    }
                  }
                  return _coll37;
                }()).contains(_769_v32)) {
                  _coll35.add((_769_v32).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(883)), function (_770_i4) {
                    return new _dafny.CodePoint('t'.codePointAt(0));
                  })).length)));
                }
              }
              return _coll35;
            }()).length)));
          }
        }
        return _coll34;
      }()), _module.__default.abs(_module.__default.fm0(globalState)));
      let _771_i5;
      _771_i5 = _dafny.ZERO;
      L8: {
        while ((_765_v28).IsDisjointFrom(_765_v28)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_771_i5)) {
              break L8;
            }
            _771_i5 = (_771_i5).plus(_dafny.ONE);
            (globalState).f6 = p0;
            (globalState).f6 = p0;
            let _772_v33;
            _772_v33 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,p0);
            let _773_v34;
            let _nw124 = new _module.C0();
            _nw124.__ctor(_764_v27, _772_v33, _this.f18, (_this).f19);
            _773_v34 = _nw124;
            let _774_v35;
            _774_v35 = _773_v34;
            let _source10 = _774_v35;
            let _775___mcc_h0 = _source10;
            let _776_cf14 = _775___mcc_h0;
            let _777_v36;
            _777_v36 = _dafny.Seq.of(_743_v9);
            _764_v27 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_module.__default.fm21(globalState), _module.__default.safeIndex(p1, new BigNumber((_module.__default.fm21(globalState)).length)), _this.f18), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_module.__default.fm21(globalState), _module.__default.safeIndex(p1, new BigNumber((_module.__default.fm21(globalState)).length)), _this.f18)).length)), _776_cf14.f18), _dafny.Seq.update(_dafny.Seq.Concat(_764_v27, _764_v27), _module.__default.safeIndex(new BigNumber((_777_v36).length), new BigNumber((_dafny.Seq.Concat(_764_v27, _764_v27)).length)), _this.f18));
            let _778_v37;
            _778_v37 = _dafny.Seq.of(_763_v26);
            let _779_v38;
            _779_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm22(globalState));
            let _780_v39;
            _780_v39 = _dafny.Seq.of((((_779_v38).contains(p1)) ? ((_779_v38).get(p1)) : (_778_v37)));
            let _pat_let_tv14 = _743_v9;
            let _pat_let_tv15 = p0;
            _778_v37 = (_780_v39)[_module.__default.safeIndex((function (_pat_let14_0) {
              return function (_783_dt__update__tmp_h1) {
                return function (_pat_let17_0) {
                  return function (_784_dt__update_hcf17_h0) {
                    return _module.D5.create_DC14((_783_dt__update__tmp_h1).dtor_cf16, _784_dt__update_hcf17_h0);
                  }(_pat_let17_0);
                }(_pat_let_tv15);
              }(_pat_let14_0);
            }(function (_pat_let15_0) {
              return function (_781_dt__update__tmp_h0) {
                return function (_pat_let16_0) {
                  return function (_782_dt__update_hcf16_h0) {
                    return _module.D5.create_DC14(_782_dt__update_hcf16_h0, (_781_dt__update__tmp_h0).dtor_cf17);
                  }(_pat_let16_0);
                }(!(_pat_let_tv14));
              }(_pat_let15_0);
            }(_module.D5.create_DC14(_743_v9, p1)))).dtor_cf17, new BigNumber((_780_v39).length))];
            let _785_v40;
            _785_v40 = _dafny.Map.Empty.slice().updateUnsafe(_764_v27,_743_v9);
            _785_v40 = (_785_v40).update(_764_v27, _743_v9);
            (globalState).f6 = (_dafny.ZERO).minus(_module.__default.fm0(globalState));
            let _786_v41;
            _786_v41 = _766_v29;
            _786_v41 = _786_v41;
          }
        }
      }
      let _787_v42;
      let _nw125 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
      _787_v42 = _nw125;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_787_v42).length))) {
        let _788_i6 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_788_i6)) && ((_788_i6).isLessThan(new BigNumber((_787_v42).length))))) {
          (_787_v42)[(_788_i6)] = (_788_i6).plus(_module.__default.safeModuloInt(p0, p1));
        }
      }
      r0 = !(_743_v9);
      return r0;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f18 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f19 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    set f18(value) {
      let _this = this;
      _this._f18 = value;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    __ctor(f18, f19) {
      let _this = this;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      return;
    }
    fm3(p0, p1, globalState) {
      let _this = this;
      return !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(514)), function (_789_i0) {
        return new BigNumber((_dafny.Seq.of(new BigNumber(937), new BigNumber(678))).length);
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(880)), function (_790_i1) {
        return new BigNumber(307);
      })), _dafny.Seq.Create(_module.__default.abs(new BigNumber(647)), function (_791_i2) {
        return new BigNumber(-125);
      }));
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber(265))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(373), new BigNumber(858)))).cardinality()), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hkdhj"), _dafny.Seq.UnicodeFromString("wtgkshg"))).length));
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = false;
      let _792_i0;
      _792_i0 = _dafny.ZERO;
      L9: {
        while (p0) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_792_i0)) {
              break L9;
            }
            _792_i0 = (_792_i0).plus(_dafny.ONE);
            let _793_v0;
            _793_v0 = new BigNumber(652);
            let _794_v1;
            _794_v1 = _dafny.MultiSet.fromElements(_793_v0);
            let _795_v2;
            _795_v2 = _dafny.Seq.of(_this.f18);
            let _796_v3;
            _796_v3 = _dafny.Set.fromElements((_dafny.ZERO).minus(_793_v0), _793_v0);
            let _797_v4;
            _797_v4 = _dafny.Seq.of(false, p0);
            let _798_v5;
            _798_v5 = _797_v4;
            let _799_v6;
            _799_v6 = _dafny.Map.Empty.slice().updateUnsafe((_798_v5),_793_v0);
            let _800_v7;
            _800_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_799_v6).length),p0);
            let _801_v8;
            _801_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_800_v7).length),p0);
            let _802_v9;
            let _init13 = ((_803_v0) => function (_804_i1) {
              return _module.__default.safeModuloInt(_804_i1, _803_v0);
            })(_793_v0);
            let _nw126 = Array((new BigNumber(21)).toNumber());
            for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw126.length); _i0_13++) {
              _nw126[_i0_13] = _init13(new BigNumber(_i0_13));
            }
            _802_v9 = _nw126;
            let _805_v10;
            _805_v10 = _dafny.Map.Empty.slice().updateUnsafe(_802_v9,(_this).fm3(p0, _793_v0, globalState));
            let _806_v11;
            let _nw127 = Array((new BigNumber(23)).toNumber());
            _nw127[(_dafny.ZERO).toNumber()] = (p0) || (p0);
            _nw127[(_dafny.ONE).toNumber()] = (_794_v1).IsSubsetOf(_794_v1);
            _nw127[(new BigNumber(2)).toNumber()] = false;
            _nw127[(new BigNumber(3)).toNumber()] = (_module.D1.create_DC4(_793_v0, _793_v0, true)).dtor_cf5;
            _nw127[(new BigNumber(4)).toNumber()] = _dafny.areEqual(_795_v2, _795_v2);
            _nw127[(new BigNumber(5)).toNumber()] = false;
            _nw127[(new BigNumber(6)).toNumber()] = p0;
            _nw127[(new BigNumber(7)).toNumber()] = p0;
            _nw127[(new BigNumber(8)).toNumber()] = (p0) || (true);
            _nw127[(new BigNumber(9)).toNumber()] = ((p0) ? (p0) : (p0));
            _nw127[(new BigNumber(10)).toNumber()] = p0;
            _nw127[(new BigNumber(11)).toNumber()] = p0;
            _nw127[(new BigNumber(12)).toNumber()] = !(p0);
            _nw127[(new BigNumber(13)).toNumber()] = !((_module.__default.fm5(_793_v0, globalState)).IsDisjointFrom(_796_v3));
            _nw127[(new BigNumber(14)).toNumber()] = (new BigNumber((_800_v7).length)).isLessThan(new BigNumber((_795_v2).length));
            _nw127[(new BigNumber(15)).toNumber()] = (_793_v0).isLessThan(_793_v0);
            _nw127[(new BigNumber(16)).toNumber()] = p0;
            _nw127[(new BigNumber(17)).toNumber()] = p0;
            _nw127[(new BigNumber(18)).toNumber()] = (((_805_v10).contains(_802_v9)) ? ((_805_v10).get(_802_v9)) : (p0));
            _nw127[(new BigNumber(19)).toNumber()] = (_793_v0).isLessThan(_793_v0);
            _nw127[(new BigNumber(20)).toNumber()] = (_794_v1).IsDisjointFrom(_794_v1);
            _nw127[(new BigNumber(21)).toNumber()] = p0;
            _nw127[(new BigNumber(22)).toNumber()] = (p0) || (p0);
            _806_v11 = _nw127;
            let _index138 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_806_v11).length));
            (_806_v11)[_index138] = p0;
            let _index139 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_802_v9).length));
            (_802_v9)[_index139] = _793_v0;
            let _index140 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_806_v11).length));
            let _index141 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_802_v9).length));
            let _rhs137 = (_793_v0).isLessThan(_793_v0);
            let _rhs138 = (_793_v0).multipliedBy(_793_v0);
            let _rhs139 = p0;
            let _rhs140 = new BigNumber(620);
            let _lhs101 = globalState;
            let _lhs102 = _806_v11;
            let _lhs103 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_806_v11).length));
            let _lhs104 = _802_v9;
            let _lhs105 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_802_v9).length));
            r3 = _rhs137;
            _lhs101.f6 = _rhs138;
            _lhs102[_lhs103] = _rhs139;
            _lhs104[_lhs105] = _rhs140;
            let _807_v12;
            _807_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_795_v2).length),_793_v0);
            _800_v7 = (_dafny.Map.Empty.slice().updateUnsafe(_793_v0,p0)).update((_dafny.ZERO).minus(((_802_v9)[_module.__default.safeIndex(new BigNumber(977), new BigNumber((_802_v9).length))]).multipliedBy(new BigNumber((_807_v12).length))), (_806_v11)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_806_v11).length))]);
            let _808_v13;
            let _nw128 = new _module.C0();
            _nw128.__ctor(_dafny.Seq.update((((_module.D1.create_DC3(p0)).dtor_cf2) ? (_795_v2) : (_dafny.Seq.UnicodeFromString("as"))), _module.__default.safeIndex((_802_v9)[_module.__default.safeIndex(new BigNumber(977), new BigNumber((_802_v9).length))], new BigNumber(((((_module.D1.create_DC3(p0)).dtor_cf2) ? (_795_v2) : (_dafny.Seq.UnicodeFromString("as")))).length)), _this.f18), _module.__default.fm7((_806_v11)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_806_v11).length))], _793_v0, globalState), _this.f18, ((_this).f19).Merge((_this).f19));
            _808_v13 = _nw128;
            let _809_v14;
            let _nw129 = new _module.C0();
            _nw129.__ctor((_808_v13).f20, (_808_v13).f21, _this.f18, (_this).f19);
            _809_v14 = _nw129;
          }
        }
      }
      let _810_v15;
      _810_v15 = new BigNumber(854);
      let _811_v16;
      _811_v16 = _dafny.MultiSet.fromElements(_810_v15);
      let _812_v17;
      let _nw130 = Array((new BigNumber(21)).toNumber());
      _nw130[(_dafny.ZERO).toNumber()] = _810_v15;
      _nw130[(_dafny.ONE).toNumber()] = _810_v15;
      _nw130[(new BigNumber(2)).toNumber()] = _810_v15;
      _nw130[(new BigNumber(3)).toNumber()] = _810_v15;
      _nw130[(new BigNumber(4)).toNumber()] = new BigNumber(200);
      _nw130[(new BigNumber(5)).toNumber()] = _810_v15;
      _nw130[(new BigNumber(6)).toNumber()] = _810_v15;
      _nw130[(new BigNumber(7)).toNumber()] = _810_v15;
      _nw130[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_810_v15);
      _nw130[(new BigNumber(9)).toNumber()] = new BigNumber(497);
      _nw130[(new BigNumber(10)).toNumber()] = _810_v15;
      _nw130[(new BigNumber(11)).toNumber()] = _810_v15;
      _nw130[(new BigNumber(12)).toNumber()] = _810_v15;
      _nw130[(new BigNumber(13)).toNumber()] = new BigNumber((_811_v16).cardinality());
      _nw130[(new BigNumber(14)).toNumber()] = _810_v15;
      _nw130[(new BigNumber(15)).toNumber()] = _810_v15;
      _nw130[(new BigNumber(16)).toNumber()] = _810_v15;
      _nw130[(new BigNumber(17)).toNumber()] = _810_v15;
      _nw130[(new BigNumber(18)).toNumber()] = _810_v15;
      _nw130[(new BigNumber(19)).toNumber()] = new BigNumber(192);
      _nw130[(new BigNumber(20)).toNumber()] = new BigNumber(-38);
      _812_v17 = _nw130;
      let _813_v18;
      _813_v18 = _dafny.Map.Empty.slice().updateUnsafe(_812_v17,_this.f18);
      let _814_v19;
      _814_v19 = _dafny.Seq.UnicodeFromString("kbgkabcen");
      let _hi6 = (new BigNumber((_814_v19).length)).minus(_810_v15);
      for (let _815_i2 = new BigNumber((_813_v18).length); _815_i2.isLessThan(_hi6); _815_i2 = _815_i2.plus(_dafny.ONE)) {
        let _816_v20;
        _816_v20 = _module.D3.create_DC8(_this.f18);
        let _817_v21;
        _817_v21 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _818_v22;
        _818_v22 = _dafny.Seq.of(p0);
        let _819_v23;
        _819_v23 = _dafny.Map.Empty.slice().updateUnsafe(_817_v21,new BigNumber((_818_v22).length));
        let _820_v24;
        _820_v24 = _dafny.Seq.of(p0, !(p0) || (_module.__default.fm1(p0, (_816_v20).dtor_cf10, (((_819_v23).contains(_817_v21)) ? ((_819_v23).get(_817_v21)) : (_815_i2)), globalState)), p0);
        r1 = new BigNumber((_818_v22).length);
        (globalState).f6 = _810_v15;
        (globalState).f6 = _810_v15;
        r1 = ((!_dafny.areEqual(_dafny.Seq.UnicodeFromString("pqpigsw"), _dafny.Seq.UnicodeFromString("mlp"))) ? (_810_v15) : (_810_v15));
      }
      let _hi7 = _810_v15;
      for (let _821_i3 = _810_v15; _821_i3.isLessThan(_hi7); _821_i3 = _821_i3.plus(_dafny.ONE)) {
        let _822_v25;
        let _init14 = ((_823_v19) => function (_824_i4) {
          return _823_v19;
        })(_814_v19);
        let _nw131 = Array((new BigNumber(17)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw131.length); _i0_14++) {
          _nw131[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _822_v25 = _nw131;
        let _index142 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_822_v25).length));
        (_822_v25)[_index142] = _814_v19;
        let _index143 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_822_v25).length));
        (_822_v25)[_index143] = _dafny.Seq.Concat(_814_v19, _814_v19);
        _812_v17 = ((false) ? (_812_v17) : (_812_v17));
        let _825_v26;
        _825_v26 = _dafny.Map.Empty.slice().updateUnsafe(_821_i3,p0);
        let _826_v27;
        _826_v27 = _dafny.Set.fromElements(p0);
        let _827_v28;
        _827_v28 = _dafny.Seq.of((((_825_v26).contains(new BigNumber((_811_v16).cardinality()))) ? ((_825_v26).get(new BigNumber((_811_v16).cardinality()))) : (p0)), p0, p0, _module.__default.fm1(p0, _this.f18, new BigNumber((_826_v27).length), globalState), p0);
        let _828_v29;
        _828_v29 = _dafny.Seq.of(true, (_827_v28)[_module.__default.safeIndex(_810_v15, new BigNumber((_827_v28).length))]);
        let _829_v30;
        _829_v30 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),new BigNumber((_814_v19).length));
        let _830_v31;
        let _nw132 = new _module.C0();
        _nw132.__ctor((_822_v25)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_822_v25).length))], _829_v30, _this.f18, (_this).f19);
        _830_v31 = _nw132;
        let _831_v32;
        _831_v32 = _dafny.Set.fromElements(_830_v31, _830_v31, _830_v31);
        let _rhs141 = (((_827_v28)[_module.__default.safeIndex(_810_v15, new BigNumber((_827_v28).length))]) ? (_810_v15) : (_821_i3));
        let _rhs142 = !(((_831_v32).Union(_831_v32)).IsProperSubsetOf(_831_v32));
        let _lhs106 = globalState;
        _lhs106.f6 = _rhs141;
        r3 = _rhs142;
        (globalState).f6 = _821_i3;
      }
      let _832_v33;
      let _init15 = ((_833_p0) => function (_834_i5) {
        return _833_p0;
      })(p0);
      let _nw133 = Array((new BigNumber(9)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw133.length); _i0_15++) {
        _nw133[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _832_v33 = _nw133;
      _832_v33 = _832_v33;
      let _835_v34;
      _835_v34 = _module.D1.create_DC3(false);
      r2 = (!((_835_v34).dtor_cf2)) && (p0);
      let _836_v35;
      let _nw134 = new _module.C0();
      _nw134.__ctor(((p0) ? (_814_v19) : (_814_v19)), _module.__default.fm7(p0, _810_v15, globalState), new _dafny.CodePoint('g'.codePointAt(0)), (_this).f19);
      _836_v35 = _nw134;
      let _837_v36;
      _837_v36 = _module.D0.create_DC0((_dafny.ZERO).minus(_810_v15));
      r0 = _837_v36;
      let _838_v37;
      _838_v37 = _dafny.Seq.of(_810_v15, _810_v15, _810_v15);
      r1 = (_810_v15).multipliedBy((new BigNumber((_838_v37).length)).minus(_810_v15));
      r2 = p0;
      r3 = p0;
      return [r0, r1, r2, r3];
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _839_v0;
      let _init16 = ((_840_p3) => function (_841_i0) {
        return (_841_i0).plus(_840_p3);
      })(p3);
      let _nw135 = Array((new BigNumber(8)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw135.length); _i0_16++) {
        _nw135[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _839_v0 = _nw135;
      let _index144 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length));
      (_839_v0)[_index144] = p3;
      let _842_v1;
      _842_v1 = true;
      let _843_v2;
      _843_v2 = _dafny.Set.fromElements(_842_v1);
      let _844_v3;
      _844_v3 = _dafny.Map.Empty.slice().updateUnsafe(p3,_843_v2);
      let _845_v4;
      _845_v4 = _module.D3.create_DC9(_842_v1, _module.__default.fm0(globalState));
      let _pat_let_tv16 = p2;
      let _pat_let_tv17 = p2;
      let _pat_let_tv18 = p3;
      let _index145 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length));
      let _rhs143 = (_dafny.ZERO).minus(function (_source11) {
        if (_source11.is_DC9) {
          let _846___mcc_h0 = (_source11).cf11;
          let _847___mcc_h1 = (_source11).cf12;
          let _848_cf12 = _847___mcc_h1;
          let _849_cf11 = _846___mcc_h0;
          return _848_cf12;
        } else if (_source11.is_DC10) {
          let _850___mcc_h2 = (_source11).cf13;
          let _851_cf13 = _850___mcc_h2;
          return _module.__default.safeModuloInt(_pat_let_tv16, _pat_let_tv17);
        } else {
          let _852___mcc_h3 = (_source11).cf10;
          let _853_cf10 = _852___mcc_h3;
          return _pat_let_tv18;
        }
      }(_845_v4));
      let _rhs144 = false;
      let _rhs145 = _844_v3;
      let _lhs107 = _839_v0;
      let _lhs108 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length));
      _lhs107[_lhs108] = _rhs143;
      _842_v1 = _rhs144;
      _844_v3 = _rhs145;
      let _854_i1;
      _854_i1 = _dafny.ZERO;
      L10: {
        while (_842_v1) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_854_i1)) {
              break L10;
            }
            _854_i1 = (_854_i1).plus(_dafny.ONE);
            (globalState).f6 = ((_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))]).multipliedBy(p3);
            let _855_v5;
            _855_v5 = _dafny.Seq.UnicodeFromString("qxtvrh");
            let _856_v6;
            _856_v6 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,new BigNumber(308));
            let _857_v7;
            let _nw136 = new _module.C0();
            _nw136.__ctor(_855_v5, _856_v6, _this.f18, ((_this).f19).update(_855_v5, (_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))]));
            _857_v7 = _nw136;
            let _index146 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length));
            (_839_v0)[_index146] = _module.__default.fm0(globalState);
            if (((p3).minus(p2)).isLessThanOrEqualTo((_dafny.ZERO).minus((_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))]))) {
              let _858_v8;
              _858_v8 = _dafny.Seq.of((_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))]);
              (globalState).f1 = _dafny.Seq.Concat(_858_v8, _dafny.Seq.of(p2));
              let _859_v9;
              _859_v9 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(p2, p1),_dafny.Map.Empty.slice().updateUnsafe(p1,p0));
              _859_v9 = (_859_v9).update((_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))], _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),p0));
              let _860_v10;
              _860_v10 = _module.D1.create_DC3(true);
              let _861_v11;
              _861_v11 = _dafny.Map.Empty.slice().updateUnsafe(p3,_860_v10);
              let _862_v12;
              let _nw137 = new _module.C0();
              _nw137.__ctor(_module.__default.fm8(globalState), (_857_v7).f21, _this.f18, (_dafny.Map.Empty.slice().updateUnsafe((_857_v7).f20,new BigNumber((_861_v11).length))).update(_dafny.Seq.UnicodeFromString("vnfqve"), (_dafny.ZERO).minus(_module.__default.fm0(globalState))));
              _862_v12 = _nw137;
              let _863_v13;
              let _init17 = ((_864_p0) => function (_865_i2) {
                return _864_p0;
              })(p0);
              let _nw138 = Array((new BigNumber(23)).toNumber());
              for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw138.length); _i0_17++) {
                _nw138[_i0_17] = _init17(new BigNumber(_i0_17));
              }
              _863_v13 = _nw138;
              let _866_v14;
              _866_v14 = _dafny.Map.Empty.slice().updateUnsafe(_862_v12,_863_v13);
              _866_v14 = (_866_v14).update(_862_v12, _863_v13);
              let _867_v15;
              _867_v15 = _dafny.Seq.of(true, _842_v1);
              let _868_v16;
              _868_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
              let _869_v17;
              _869_v17 = _dafny.Seq.update(_867_v15, _module.__default.safeIndex(new BigNumber((_868_v16).length), new BigNumber((_867_v15).length)), p0);
              let _870_v18;
              _870_v18 = _dafny.MultiSet.fromElements(_869_v17);
              let _871_v19;
              _871_v19 = _dafny.Seq.of(_867_v15);
              let _872_v20;
              _872_v20 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_871_v19));
              let _873_v21;
              let _nw139 = Array((new BigNumber(11)).toNumber());
              _nw139[(_dafny.ZERO).toNumber()] = (_870_v18).Intersect(_870_v18);
              _nw139[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements(_869_v17, _869_v17, _869_v17, _867_v15, _dafny.Seq.of(p0, p0, false, _842_v1, p0));
              _nw139[(new BigNumber(2)).toNumber()] = _870_v18;
              _nw139[(new BigNumber(3)).toNumber()] = (_870_v18).Difference(_870_v18);
              _nw139[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.fromElements(_869_v17, _869_v17);
              _nw139[(new BigNumber(5)).toNumber()] = _module.__default.fm9(globalState);
              _nw139[(new BigNumber(6)).toNumber()] = _870_v18;
              _nw139[(new BigNumber(7)).toNumber()] = (_872_v20)[_module.__default.safeIndex((_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))], new BigNumber((_872_v20).length))];
              _nw139[(new BigNumber(8)).toNumber()] = (_872_v20)[_module.__default.safeIndex((_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))], new BigNumber((_872_v20).length))];
              _nw139[(new BigNumber(9)).toNumber()] = _870_v18;
              _nw139[(new BigNumber(10)).toNumber()] = (_dafny.MultiSet.fromElements(_869_v17)).update(_869_v17, _module.__default.abs((_this).fm4(_842_v1, _this.f18, new _dafny.CodePoint('u'.codePointAt(0)), _842_v1, globalState)));
              _873_v21 = _nw139;
              let _index147 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_873_v21).length));
              (_873_v21)[_index147] = _870_v18;
              let _874_v22;
              _874_v22 = _dafny.MultiSet.fromElements(_842_v1);
              let _index148 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_873_v21).length));
              (_873_v21)[_index148] = ((_870_v18).update(_869_v17, _module.__default.abs(new BigNumber((_874_v22).cardinality())))).Intersect(_870_v18);
              let _875_v23;
              _875_v23 = _dafny.Set.fromElements(new BigNumber(-192), (_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))], (_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))], (_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))]);
              _842_v1 = (_875_v23).IsSubsetOf((_875_v23).Union(_875_v23));
            } else {
              let _876_v24;
              _876_v24 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_module.__default.fm8(globalState), (_857_v7).f20),_842_v1);
              _842_v1 = (((_876_v24).contains(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fvox"), (_857_v7).f20))) ? ((_876_v24).get(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fvox"), (_857_v7).f20))) : (_842_v1));
              let _877_v25;
              let _nw140 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
              _877_v25 = _nw140;
              let _878_v26;
              _878_v26 = _dafny.Set.fromElements(p3);
              let _879_v27;
              _879_v27 = _dafny.Seq.of(p1, (_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))], p1, new BigNumber((_878_v26).length), p1);
              let _index149 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_877_v25).length));
              (_877_v25)[_index149] = _879_v27;
              let _880_v28;
              let _init18 = ((_881_v2) => function (_882_i3) {
                return _881_v2;
              })(_843_v2);
              let _nw141 = Array((new BigNumber(22)).toNumber());
              for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw141.length); _i0_18++) {
                _nw141[_i0_18] = _init18(new BigNumber(_i0_18));
              }
              _880_v28 = _nw141;
              let _index150 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_880_v28).length));
              (_880_v28)[_index150] = _dafny.Set.fromElements(_842_v1, _842_v1, _842_v1, p0, !(_842_v1));
              let _883_v29;
              let _nw142 = Array((new BigNumber(19)).toNumber()).fill(false);
              _883_v29 = _nw142;
              let _index151 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_883_v29).length));
              (_883_v29)[_index151] = _842_v1;
              let _884_v30;
              _884_v30 = _dafny.Seq.of(!(_842_v1));
              let _index152 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_877_v25).length));
              let _index153 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_880_v28).length));
              let _index154 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_883_v29).length));
              let _rhs146 = (_884_v30)[_module.__default.safeIndex(p1, new BigNumber((_884_v30).length))];
              let _rhs147 = _module.D3.create_DC9(p0, p3);
              let _rhs148 = _dafny.Seq.Concat(_dafny.Seq.Concat(_879_v27, _879_v27), _879_v27);
              let _rhs149 = ((_843_v2).Union(_843_v2)).Intersect(_dafny.Set.fromElements(p0));
              let _rhs150 = _dafny.Seq.IsProperPrefixOf(((p0) ? (_879_v27) : (_879_v27)), _dafny.Seq.Concat(_879_v27, _879_v27));
              let _lhs109 = _877_v25;
              let _lhs110 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_877_v25).length));
              let _lhs111 = _880_v28;
              let _lhs112 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_880_v28).length));
              let _lhs113 = _883_v29;
              let _lhs114 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_883_v29).length));
              _842_v1 = _rhs146;
              _845_v4 = _rhs147;
              _lhs109[_lhs110] = _rhs148;
              _lhs111[_lhs112] = _rhs149;
              _lhs113[_lhs114] = _rhs150;
              let _index155 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length));
              (_839_v0)[_index155] = new BigNumber(600);
              let _index156 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_883_v29).length));
              (_883_v29)[_index156] = (_884_v30)[_module.__default.safeIndex(p1, new BigNumber((_884_v30).length))];
              _883_v29 = _883_v29;
            }
          }
        }
      }
      let _885_v31;
      _885_v31 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
      let _886_v32;
      _886_v32 = _dafny.Seq.of(_885_v31, _885_v31, _885_v31);
      let _index157 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length));
      (_839_v0)[_index157] = (_dafny.ZERO).minus(((!(p0) || (!(_842_v1))) ? ((_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))]) : (new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-147)), ((_887_v1) => function (_888_i4) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_this.f18, _this.f18, _this.f18)).cardinality()),_887_v1);
      })(_842_v1)), _886_v32))).cardinality()))));
      let _889_v33;
      _889_v33 = _dafny.Map.Empty.slice().updateUnsafe((_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))],new BigNumber(135));
      let _hi8 = new BigNumber((_889_v33).length);
      for (let _890_i5 = p3; _890_i5.isLessThan(_hi8); _890_i5 = _890_i5.plus(_dafny.ONE)) {
        let _891_v34;
        _891_v34 = _module.D1.create_DC5(p1);
        let _892_v35;
        let _nw143 = Array((new BigNumber(6)).toNumber());
        _nw143[(_dafny.ZERO).toNumber()] = _891_v34;
        _nw143[(_dafny.ONE).toNumber()] = ((_842_v1) ? (_891_v34) : (_module.D1.create_DC5(new BigNumber(717))));
        _nw143[(new BigNumber(2)).toNumber()] = _891_v34;
        _nw143[(new BigNumber(3)).toNumber()] = _891_v34;
        _nw143[(new BigNumber(4)).toNumber()] = _891_v34;
        _nw143[(new BigNumber(5)).toNumber()] = _891_v34;
        _892_v35 = _nw143;
        let _index158 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_892_v35).length));
        (_892_v35)[_index158] = _891_v34;
        let _index159 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_892_v35).length));
        (_892_v35)[_index159] = _module.D1.create_DC5(p2);
        _842_v1 = (p0) && (_842_v1);
        let _893_v36;
        let _nw144 = Array((new BigNumber(10)).toNumber()).fill(false);
        _893_v36 = _nw144;
        let _index160 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_893_v36).length));
        (_893_v36)[_index160] = !(p0);
        let _894_v37;
        _894_v37 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p1), p3, p1);
        let _895_v38;
        _895_v38 = _dafny.Map.Empty.slice().updateUnsafe(_842_v1,p0);
        let _896_v39;
        let _nw145 = new _module.C2();
        _nw145.__ctor(_this.f18, (_this).f19);
        _896_v39 = _nw145;
        let _897_v40;
        _897_v40 = _dafny.Map.Empty.slice().updateUnsafe(_896_v39,_842_v1);
        let _898_v41;
        _898_v41 = _module.D1.create_DC4(p1, new BigNumber((_897_v40).length), _842_v1);
        let _index161 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_893_v36).length));
        let _rhs151 = ((_this).fm3(_842_v1, (((_894_v37).contains((_dafny.ZERO).minus(p3))) ? ((_894_v37).get((_dafny.ZERO).minus(p3))) : (new BigNumber((_895_v38).length))), globalState)) && (((p0) ? (_module.__default.fm1(p0, _module.__default.fm10(globalState), _890_i5, globalState)) : (_842_v1)));
        let _rhs152 = _843_v2;
        let _rhs153 = (_898_v41).dtor_cf5;
        let _lhs115 = _893_v36;
        let _lhs116 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_893_v36).length));
        _lhs115[_lhs116] = _rhs151;
        _843_v2 = _rhs152;
        _842_v1 = _rhs153;
        let _899_v42;
        let _nw146 = Array((new BigNumber(15)).toNumber()).fill([]);
        _899_v42 = _nw146;
        let _900_v43;
        let _nw147 = Array((new BigNumber(14)).toNumber()).fill([]);
        _900_v43 = _nw147;
        let _index162 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_899_v42).length));
        (_899_v42)[_index162] = _900_v43;
        let _901_v44;
        _901_v44 = _dafny.Seq.of(_842_v1);
        let _902_v45;
        _902_v45 = _dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.Concat(_901_v44, _dafny.Seq.of((_893_v36)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_893_v36).length))], _842_v1, _842_v1, true, !(p0))), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.Concat(_901_v44, _dafny.Seq.of((_893_v36)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_893_v36).length))], _842_v1, _842_v1, true, !(p0)))).length)), _842_v1));
        let _index163 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_899_v42).length));
        let _rhs154 = (_893_v36)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_893_v36).length))];
        let _rhs155 = _900_v43;
        let _rhs156 = _902_v45;
        let _lhs117 = _899_v42;
        let _lhs118 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_899_v42).length));
        _842_v1 = _rhs154;
        _lhs117[_lhs118] = _rhs155;
        _902_v45 = _rhs156;
      }
      let _903_v46;
      let _nw148 = new _module.C2();
      _nw148.__ctor(new _dafny.CodePoint('q'.codePointAt(0)), (_this).f19);
      _903_v46 = _nw148;
      let _904_v47;
      _904_v47 = _dafny.Map.Empty.slice().updateUnsafe(_903_v46,p0);
      let _905_v48;
      _905_v48 = _dafny.Set.fromElements(_904_v47, _904_v47, _904_v47, _904_v47);
      let _906_v49;
      _906_v49 = _dafny.Seq.of(_905_v48);
      let _907_v50;
      _907_v50 = _dafny.MultiSet.fromElements(p0);
      if (!(!((_906_v49)[_module.__default.safeIndex(new BigNumber((_907_v50).cardinality()), new BigNumber((_906_v49).length))]).equals((_905_v48).Difference(_905_v48)))) {
        let _908_v51;
        _908_v51 = _dafny.Seq.of(p0);
        let _909_v52;
        let _nw149 = Array((new BigNumber(29)).toNumber());
        _nw149[(_dafny.ZERO).toNumber()] = p0;
        _nw149[(_dafny.ONE).toNumber()] = p0;
        _nw149[(new BigNumber(2)).toNumber()] = true;
        _nw149[(new BigNumber(3)).toNumber()] = p0;
        _nw149[(new BigNumber(4)).toNumber()] = _842_v1;
        _nw149[(new BigNumber(5)).toNumber()] = _842_v1;
        _nw149[(new BigNumber(6)).toNumber()] = _842_v1;
        _nw149[(new BigNumber(7)).toNumber()] = (_908_v51)[_module.__default.safeIndex(p1, new BigNumber((_908_v51).length))];
        _nw149[(new BigNumber(8)).toNumber()] = _842_v1;
        _nw149[(new BigNumber(9)).toNumber()] = p0;
        _nw149[(new BigNumber(10)).toNumber()] = _842_v1;
        _nw149[(new BigNumber(11)).toNumber()] = p0;
        _nw149[(new BigNumber(12)).toNumber()] = _842_v1;
        _nw149[(new BigNumber(13)).toNumber()] = _842_v1;
        _nw149[(new BigNumber(14)).toNumber()] = _842_v1;
        _nw149[(new BigNumber(15)).toNumber()] = _842_v1;
        _nw149[(new BigNumber(16)).toNumber()] = p0;
        _nw149[(new BigNumber(17)).toNumber()] = p0;
        _nw149[(new BigNumber(18)).toNumber()] = _842_v1;
        _nw149[(new BigNumber(19)).toNumber()] = true;
        _nw149[(new BigNumber(20)).toNumber()] = _842_v1;
        _nw149[(new BigNumber(21)).toNumber()] = p0;
        _nw149[(new BigNumber(22)).toNumber()] = _842_v1;
        _nw149[(new BigNumber(23)).toNumber()] = p0;
        _nw149[(new BigNumber(24)).toNumber()] = p0;
        _nw149[(new BigNumber(25)).toNumber()] = _842_v1;
        _nw149[(new BigNumber(26)).toNumber()] = (_908_v51)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_842_v1,_839_v0)).length)), new BigNumber((_908_v51).length))];
        _nw149[(new BigNumber(27)).toNumber()] = p0;
        _nw149[(new BigNumber(28)).toNumber()] = p0;
        _909_v52 = _nw149;
        let _910_v53;
        _910_v53 = _dafny.Map.Empty.slice().updateUnsafe(p0,_909_v52);
        let _911_v54;
        let _nw150 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _911_v54 = _nw150;
        let _912_v55;
        _912_v55 = _dafny.Map.Empty.slice().updateUnsafe((((_910_v53).contains(_842_v1)) ? ((_910_v53).get(_842_v1)) : (_909_v52)),_911_v54);
        let _913_v56;
        _913_v56 = _dafny.Map.Empty.slice().updateUnsafe((((_912_v55).contains(_909_v52)) ? ((_912_v55).get(_909_v52)) : (_911_v54)),p0);
        _913_v56 = (_913_v56).update(_911_v54, p0);
        let _914_v57;
        _914_v57 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _915_v58;
        _915_v58 = _dafny.Seq.of(_914_v57);
        let _916_v59;
        _916_v59 = _dafny.MultiSet.fromElements(_914_v57, _914_v57, _914_v57, (_914_v57).Merge((_915_v58)[_module.__default.safeIndex(p3, new BigNumber((_915_v58).length))]), _module.__default.fm23(p1, _842_v1, globalState));
        _916_v59 = _916_v59;
        let _917_v60;
        let _nw151 = new _module.C1();
        _nw151.__ctor((p2).plus(p2), p0, _this.f18, _module.__default.fm18((_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))], globalState));
        _917_v60 = _nw151;
        let _918_v61;
        _918_v61 = _dafny.Seq.of(p2);
        let _919_v62;
        _919_v62 = _dafny.Map.Empty.slice().updateUnsafe(_918_v61,_842_v1);
        _842_v1 = !(_842_v1) || ((true) || ((((_919_v62).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-491)), function (_921_i6) {
          return new BigNumber(131);
        }))) ? ((_919_v62).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-491)), function (_920_i6) {
          return new BigNumber(131);
        }))) : (false))));
        let _index164 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length));
        (_839_v0)[_index164] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_908_v51, _module.__default.safeIndex(new BigNumber((_843_v2).length), new BigNumber((_908_v51).length)), !(_842_v1))).length), new BigNumber((_885_v31).length));
      } else {
        let _rhs157 = p2;
        let _rhs158 = p2;
        let _rhs159 = (((_843_v2).IsDisjointFrom(_843_v2)) ? (true) : (!(false)));
        let _lhs119 = globalState;
        _lhs119.f6 = _rhs157;
        r0 = _rhs158;
        _842_v1 = _rhs159;
        let _922_v63;
        _922_v63 = _dafny.Seq.of(_839_v0, _839_v0, _839_v0, _839_v0, _839_v0);
        let _923_v64;
        _923_v64 = _module.D1.create_DC4((_dafny.ZERO).minus((_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))]), (_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))], _842_v1);
        let _924_v65;
        _924_v65 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))]);
        let _925_v66;
        _925_v66 = _dafny.MultiSet.fromElements(p1);
        let _926_v67;
        _926_v67 = _dafny.Seq.UnicodeFromString("wyt");
        let _927_v68;
        let _nw152 = Array((new BigNumber(28)).toNumber());
        _nw152[(_dafny.ZERO).toNumber()] = p0;
        _nw152[(_dafny.ONE).toNumber()] = _842_v1;
        _nw152[(new BigNumber(2)).toNumber()] = _842_v1;
        _nw152[(new BigNumber(3)).toNumber()] = _842_v1;
        _nw152[(new BigNumber(4)).toNumber()] = !(true);
        _nw152[(new BigNumber(5)).toNumber()] = p0;
        _nw152[(new BigNumber(6)).toNumber()] = _842_v1;
        _nw152[(new BigNumber(7)).toNumber()] = !_dafny.Seq.contains(_922_v63, _839_v0);
        _nw152[(new BigNumber(8)).toNumber()] = _842_v1;
        _nw152[(new BigNumber(9)).toNumber()] = !(p0);
        _nw152[(new BigNumber(10)).toNumber()] = (((_885_v31).contains((_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))])) ? ((_885_v31).get((_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))])) : ((_923_v64).dtor_cf5));
        _nw152[(new BigNumber(11)).toNumber()] = (_924_v65).contains(p0);
        _nw152[(new BigNumber(12)).toNumber()] = _842_v1;
        _nw152[(new BigNumber(13)).toNumber()] = (_925_v66).IsSubsetOf(_925_v66);
        _nw152[(new BigNumber(14)).toNumber()] = _842_v1;
        _nw152[(new BigNumber(15)).toNumber()] = ((_845_v4).dtor_cf11) === (p0);
        _nw152[(new BigNumber(16)).toNumber()] = (((_885_v31).contains(p2)) ? ((_885_v31).get(p2)) : (!(p0)));
        _nw152[(new BigNumber(17)).toNumber()] = true;
        _nw152[(new BigNumber(18)).toNumber()] = true;
        _nw152[(new BigNumber(19)).toNumber()] = !(_842_v1) || (p0);
        _nw152[(new BigNumber(20)).toNumber()] = !(p0);
        _nw152[(new BigNumber(21)).toNumber()] = !_dafny.areEqual(_926_v67, _926_v67);
        _nw152[(new BigNumber(22)).toNumber()] = p0;
        _nw152[(new BigNumber(23)).toNumber()] = ((_842_v1) ? (false) : (_842_v1));
        _nw152[(new BigNumber(24)).toNumber()] = _842_v1;
        _nw152[(new BigNumber(25)).toNumber()] = !(true);
        _nw152[(new BigNumber(26)).toNumber()] = ((!(_842_v1)) ? (p0) : (_842_v1));
        _nw152[(new BigNumber(27)).toNumber()] = p0;
        _927_v68 = _nw152;
        let _index165 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length));
        (_927_v68)[_index165] = _842_v1;
        let _928_v69;
        _928_v69 = _dafny.Map.Empty.slice().updateUnsafe(_907_v50,p2);
        let _929_v71;
        _929_v71 = _dafny.Set.fromElements(_907_v50, _907_v50);
        let _930_v72;
        _930_v72 = _dafny.Map.Empty.slice().updateUnsafe(_907_v50,p0);
        let _index166 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length));
        let _rhs160 = (function () {
          let _coll38 = new _dafny.Set();
          for (const _compr_38 of ((_930_v72).update(_907_v50, p0)).Keys.Elements) {
            let _931_v73 = _compr_38;
            if (((_930_v72).update(_907_v50, p0)).contains(_931_v73)) {
              _coll38.add(_931_v73);
            }
          }
          return _coll38;
        }()).IsProperSubsetOf((function () {
          let _coll39 = new _dafny.Set();
          for (const _compr_39 of (_928_v69).Keys.Elements) {
            let _932_v70 = _compr_39;
            if ((_928_v69).contains(_932_v70)) {
              _coll39.add(_932_v70);
            }
          }
          return _coll39;
        }()).Union(_929_v71));
        let _rhs161 = _842_v1;
        let _rhs162 = _926_v67;
        let _lhs120 = _927_v68;
        let _lhs121 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length));
        _lhs120[_lhs121] = _rhs160;
        _842_v1 = _rhs161;
        _926_v67 = _rhs162;
        if (true) {
          _842_v1 = !((_927_v68)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length))]) || (((false) ? (false) : (p0)));
          _889_v33 = _889_v33;
          let _933_v75;
          _933_v75 = _dafny.Seq.of((_927_v68)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length))]);
          let _934_v76;
          _934_v76 = _933_v75;
          let _935_v77;
          _935_v77 = _dafny.Seq.of(_934_v76, _933_v75);
          let _936_v78;
          _936_v78 = _dafny.Seq.of(_935_v77);
          let _937_v79;
          _937_v79 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll40 = new _dafny.Map();
            for (const _compr_40 of _dafny.IntegerRange(new BigNumber(304), new BigNumber(-945))) {
              let _938_v74 = _compr_40;
              if (((new BigNumber(304)).isLessThanOrEqualTo(_938_v74)) && ((_938_v74).isLessThan(new BigNumber(-945)))) {
                _coll40.push([(_938_v74).minus(p1),p0]);
              }
            }
            return _coll40;
          }(),(_936_v78)[_module.__default.safeIndex(p3, new BigNumber((_936_v78).length))]);
          let _939_v80;
          _939_v80 = _dafny.Map.Empty.slice().updateUnsafe(_843_v2,_937_v79);
          _939_v80 = (_939_v80).update((((_927_v68)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length))]) ? (_843_v2) : (_dafny.Set.fromElements((_927_v68)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length))], p0, !((_927_v68)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length))])))), _937_v79);
          let _index167 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length));
          (_927_v68)[_index167] = p0;
          r0 = p3;
        } else {
          let _940_v81;
          _940_v81 = _module.D3.create_DC8(new _dafny.CodePoint('d'.codePointAt(0)));
          let _941_v82;
          _941_v82 = _dafny.Map.Empty.slice().updateUnsafe((_940_v81).dtor_cf10,p3);
          let _942_v83;
          let _nw153 = new _module.C0();
          _nw153.__ctor(_926_v67, _941_v82, _this.f18, ((_this).f19).Merge((_this).f19));
          _942_v83 = _nw153;
          let _943_v84;
          let _nw154 = new _module.C0();
          _nw154.__ctor(_module.__default.fm8(globalState), (_942_v83).f21, _this.f18, (_this).f19);
          _943_v84 = _nw154;
          _943_v84 = _942_v83;
          let _944_v85;
          _944_v85 = _module.D1.create_DC6(_842_v1, (_927_v68)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length))]);
          let _945_v86;
          _945_v86 = _dafny.Map.Empty.slice().updateUnsafe(p3,_944_v85);
          _945_v86 = (_945_v86).update((_839_v0)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_839_v0).length))], _944_v85);
          let _946_v87;
          let _nw155 = new _module.C0();
          _nw155.__ctor(((!(!(true))) ? ((_943_v84).f20) : (_926_v67)), (_941_v82).Merge(_941_v82), _this.f18, (_this).f19);
          _946_v87 = _nw155;
          (globalState).f6 = _module.__default.safeModuloInt((p3).minus(p1), new BigNumber(259));
        }
        let _947_v88;
        _947_v88 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(_842_v1, p0));
        let _948_v89;
        _948_v89 = _module.D3.create_DC10(!(p0));
        let _949_v90;
        _949_v90 = _dafny.Map.Empty.slice().updateUnsafe(_947_v88,_948_v89);
        _949_v90 = (_949_v90).update(_947_v88, _module.__default.fm24(globalState));
        let _950_v92;
        _950_v92 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,_842_v1);
        let _951_v93;
        _951_v93 = _module.D6.create_DC15(function () {
  let _coll41 = new _dafny.Map();
  for (const _compr_41 of (_950_v92).Keys.Elements) {
    let _952_v91 = _compr_41;
    if ((_950_v92).contains(_952_v91)) {
      _coll41.push([_952_v91,new BigNumber((_dafny.Seq.UnicodeFromString("y")).length)]);
    }
  }
  return _coll41;
}());
        let _953_v94;
        _953_v94 = _module.D6.create_DC17(_951_v93);
        let _954_v95;
        _954_v95 = _dafny.Map.Empty.slice().updateUnsafe(_953_v94,(_927_v68)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length))]);
        let _955_v96;
        _955_v96 = _dafny.Seq.of(_954_v95);
        let _956_v97;
        let _nw156 = Array((new BigNumber(29)).toNumber());
        _nw156[(_dafny.ZERO).toNumber()] = _954_v95;
        _nw156[(_dafny.ONE).toNumber()] = (_954_v95).Merge((_954_v95).update(_953_v94, p0));
        _nw156[(new BigNumber(2)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(3)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(4)).toNumber()] = (_954_v95).Merge(_954_v95);
        _nw156[(new BigNumber(5)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(6)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(7)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(8)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(9)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(10)).toNumber()] = (_955_v96)[_module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_955_v96).length))];
        _nw156[(new BigNumber(11)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(12)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(13)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(14)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(15)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(16)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(17)).toNumber()] = ((_842_v1) ? (_954_v95) : (_954_v95));
        _nw156[(new BigNumber(18)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(19)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(20)).toNumber()] = (_954_v95).update(_953_v94, (_this).fm3(!((_927_v68)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length))]), (((_907_v50).contains((_927_v68)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length))])) ? ((_907_v50).get((_927_v68)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length))])) : (new BigNumber((_926_v67).length))), globalState));
        _nw156[(new BigNumber(21)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(22)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(23)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(24)).toNumber()] = (_954_v95).Merge((_954_v95).update(_953_v94, (_927_v68)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_927_v68).length))]));
        _nw156[(new BigNumber(25)).toNumber()] = (_954_v95).Merge(_954_v95);
        _nw156[(new BigNumber(26)).toNumber()] = _954_v95;
        _nw156[(new BigNumber(27)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC17(_module.D6.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(_this.f18,p3))),_842_v1);
        _nw156[(new BigNumber(28)).toNumber()] = _954_v95;
        _956_v97 = _nw156;
        _956_v97 = _956_v97;
      }
      let _957_v98;
      _957_v98 = _dafny.Seq.UnicodeFromString("dv");
      let _958_v99;
      _958_v99 = _module.D9.create_DC22(_957_v98);
      let _959_v100;
      let _nw157 = Array((new BigNumber(16)).toNumber());
      _nw157[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(670)), function (_960_i8) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      });
      _nw157[(_dafny.ONE).toNumber()] = _957_v98;
      _nw157[(new BigNumber(2)).toNumber()] = (_958_v99).dtor_cf27;
      _nw157[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("ude");
      _nw157[(new BigNumber(4)).toNumber()] = _957_v98;
      _nw157[(new BigNumber(5)).toNumber()] = _957_v98;
      _nw157[(new BigNumber(6)).toNumber()] = _957_v98;
      _nw157[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_957_v98, _957_v98);
      _nw157[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(168)), function (_961_i9) {
        return _this.f18;
      });
      _nw157[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_957_v98, _957_v98);
      _nw157[(new BigNumber(10)).toNumber()] = _957_v98;
      _nw157[(new BigNumber(11)).toNumber()] = (_958_v99).dtor_cf27;
      _nw157[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_957_v98, _957_v98);
      _nw157[(new BigNumber(13)).toNumber()] = _957_v98;
      _nw157[(new BigNumber(14)).toNumber()] = _957_v98;
      _nw157[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("o"), _957_v98);
      _959_v100 = _nw157;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_959_v100).length))) {
        let _962_i7 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_962_i7)) && ((_962_i7).isLessThan(new BigNumber((_959_v100).length))))) {
          (_959_v100)[(_962_i7)] = _957_v98;
        }
      }
      r0 = p3;
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _963_v0;
      _963_v0 = _dafny.Seq.UnicodeFromString("bfbxyfb");
      let _964_v1;
      _964_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p2).cardinality()),p1);
      let _965_v2;
      _965_v2 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of(p1, p1)).length), p0, p0, p0, new BigNumber((_963_v0).length));
      let _hi9 = (new BigNumber((_964_v1).length)).minus((_965_v2)[_module.__default.safeIndex(new BigNumber(764), new BigNumber((_965_v2).length))]);
      for (let _966_i0 = (_dafny.ZERO).minus(new BigNumber((_963_v0).length)); _966_i0.isLessThan(_hi9); _966_i0 = _966_i0.plus(_dafny.ONE)) {
        _963_v0 = _963_v0;
        let _967_v3;
        let _nw158 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
        _967_v3 = _nw158;
        let _968_v4;
        _968_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((p2).cardinality()));
        let _index168 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_967_v3).length));
        (_967_v3)[_index168] = _968_v4;
        let _index169 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_967_v3).length));
        (_967_v3)[_index169] = (((p1) && (p1)) ? (_968_v4) : (_968_v4));
        let _969_v5;
        let _970_v6;
        let _971_v7;
        let _out24;
        let _out25;
        let _out26;
        let _outcollector7 = _module.__default.m0((_this).fm4(p1, new _dafny.CodePoint('e'.codePointAt(0)), _this.f18, p1, globalState), p1, _966_i0, globalState);
        _out24 = _outcollector7[0];
        _out25 = _outcollector7[1];
        _out26 = _outcollector7[2];
        _969_v5 = _out24;
        _970_v6 = _out25;
        _971_v7 = _out26;
        let _972_v8;
        _972_v8 = _dafny.MultiSet.fromElements(p2, p2, _dafny.MultiSet.fromElements(_971_v7, _module.__default.fm0(globalState), (_dafny.ZERO).minus(_971_v7), p0, new BigNumber((_963_v0).length)), (_dafny.MultiSet.fromElements(_966_i0)).Union(p2));
        let _973_v9;
        _973_v9 = _dafny.Set.fromElements(_module.__default.fm0(globalState), new BigNumber((_963_v0).length), new BigNumber((_963_v0).length));
        let _rhs163 = _966_i0;
        let _rhs164 = (((_972_v8).contains(p2)) ? ((_972_v8).get(p2)) : (_971_v7));
        let _rhs165 = _module.__default.safeDivisionInt(new BigNumber((_973_v9).length), _module.__default.fm0(globalState));
        let _rhs166 = _module.__default.fm8(globalState);
        let _lhs122 = globalState;
        let _lhs123 = globalState;
        r0 = _rhs163;
        _lhs122.f6 = _rhs164;
        _lhs123.f6 = _rhs165;
        _963_v0 = _rhs166;
      }
      let _974_v10;
      let _nw159 = new _module.C2();
      _nw159.__ctor(_this.f18, (_this).f19);
      _974_v10 = _nw159;
      let _975_v11;
      _975_v11 = _dafny.Seq.of(_974_v10);
      let _976_v12;
      _976_v12 = _dafny.Map.Empty.slice().updateUnsafe(p1,((p1) ? (_975_v11) : (_975_v11)));
      _976_v12 = ((_976_v12).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_975_v11))).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,_975_v11));
      let _977_v13;
      _977_v13 = _dafny.MultiSet.fromElements(false, false, !(p1), p1, _dafny.Seq.IsPrefixOf(_963_v0, _963_v0));
      _977_v13 = _977_v13;
      let _978_v14;
      _978_v14 = true;
      _978_v14 = !(_978_v14);
      _978_v14 = !(p1);
      let _979_i1;
      _979_i1 = _dafny.ZERO;
      L11: {
        while (_978_v14) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_979_i1)) {
              break L11;
            }
            _979_i1 = (_979_i1).plus(_dafny.ONE);
            _978_v14 = p1;
            (_this).f18 = (_963_v0)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_963_v0, _963_v0)).length), new BigNumber((_963_v0).length))];
            _978_v14 = _978_v14;
            let _980_v15;
            _980_v15 = _dafny.Set.fromElements((_this).fm4(p1, (_963_v0)[_module.__default.safeIndex(p0, new BigNumber((_963_v0).length))], _this.f18, _978_v14, globalState));
            _978_v14 = (_980_v15).IsSubsetOf(_980_v15);
          }
        }
      }
      r0 = p0;
      return r0;
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
